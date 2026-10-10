package kgo

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

const (
	epoch             = int64(1577808000000)                           // 设置起始时间(时间戳/毫秒)：2020-01-01 00:00:00，有效期69年
	timestampBits41   = uint(41)                                       // 时间戳占用位数
	datacenterIdBits  = uint(2)                                        // 数据中心id所占位数
	workerIdBits      = uint(7)                                        // 机器id所占位数
	sequenceBits      = uint(12)                                       // 序列所占的位数
	timestampMax      = int64(-1 ^ (-1 << timestampBits41))            // 时间戳最大值
	datacenterIdMax   = int64(-1 ^ (-1 << datacenterIdBits))           // 支持的最大数据中心id数量
	workerIdMax       = int64(-1 ^ (-1 << workerIdBits))               // 支持的最大机器id数量
	sequenceMask      = int64(-1 ^ (-1 << sequenceBits))               // 支持的最大序列id数量
	workerIdShift     = sequenceBits                                   // 机器id左移位数
	dataCenterIdShift = sequenceBits + workerIdBits                    // 数据中心id左移位数
	timestampShift    = sequenceBits + workerIdBits + datacenterIdBits // 时间戳左移位数
)

type snowflake struct {
	sync.Mutex         // ID锁
	timestamp    int64 // 时间戳 ，毫秒
	workerId     int64 // 工作节点
	dataCenterId int64 // 数据中心机房id
	sequence     int64 // 序列号
}

var (
	instanceMutex     sync.Mutex   // 实例初始化锁
	snowflakeInstance atomic.Value // 存储 *snowflake 实例，保证并发读写安全（消除 data race）
)

// snowflakeOf 并发安全地读取当前雪花实例，未初始化时返回 nil
func snowflakeOf() *snowflake {
	v := snowflakeInstance.Load()
	if v == nil {
		return nil
	}
	return v.(*snowflake)
}

// InitSnowflake 初始化雪花算法，只需要在你的程序启动或初始化时调用一次。
// 只会生效一次：已初始化后，后续调用（即使参数不同）将被忽略并返回 nil。
// workerId、dataCenterId 越界时返回 error（不再 panic）。
func InitSnowflake(workerId int64, dataCenterId int64) (err error) {
	if workerId < 0 || workerId > workerIdMax {
		return fmt.Errorf("workerId must be between 0 and %d", workerIdMax)
	}
	if dataCenterId < 0 || dataCenterId > datacenterIdMax {
		return fmt.Errorf("dataCenterId must be between 0 and %d", datacenterIdMax)
	}
	instanceMutex.Lock()
	defer instanceMutex.Unlock()
	if snowflakeInstance.Load() != nil {
		return nil
	}
	snowflakeInstance.Store(&snowflake{workerId: workerId, dataCenterId: dataCenterId})
	return nil
}

func SnowflakeId() int64 {
	s := snowflakeOf()
	if s == nil {
		panic("snowflake 尚未初始化，请先调用 InitSnowflake")
	}
	return s.nextVal()
}

func GetSnowflakeId[T string | int64]() (id T) {
	s := snowflakeOf()
	if s == nil {
		panic("snowflake 尚未初始化，请先调用 InitSnowflake")
	}
	v := s.nextVal()
	t := reflect.TypeOf(id)
	if t.Kind() == reflect.String {
		reflect.ValueOf(&id).Elem().SetString(fmt.Sprintf("%+v", v))
	} else if t.Kind() == reflect.Int64 {
		reflect.ValueOf(&id).Elem().SetInt(v)
	}
	return id
}

func (s *snowflake) nextVal() int64 {
	s.Lock()
	defer s.Unlock()
	now := time.Now().UnixNano() / 1000000 // 转毫秒
	if s.timestamp == now {
		// 当同一时间戳（精度：毫秒）下多次生成id会增加序列号
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			// 如果当前序列超出12bit长度，则需要等待下一毫秒
			// 下一毫秒将使用sequence:0
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		// 不同时间戳（精度：毫秒）下直接使用序列号：0
		s.sequence = 0
	}
	t := now - epoch
	if t > timestampMax {
		panic(fmt.Sprintf("epoch must be between 0 and %d", timestampMax-1))
	}
	s.timestamp = now
	r := t<<timestampShift | (s.dataCenterId << dataCenterIdShift) | (s.workerId << workerIdShift) | (s.sequence)
	return r
}
