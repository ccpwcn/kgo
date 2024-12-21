package kgo

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
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
	instanceMutex     sync.Mutex // 实例锁
	snowflakeInstance *snowflake // 实例
)

// InitSnowflake 初始化雪花算法，只需要在你的程序启动或初始化时调用一次
func InitSnowflake(workerId int64, dataCenterId int64) (err error) {
	if workerId < 0 || workerId > workerIdMax {
		panic(fmt.Sprintf("workId must be between 0 and %d", workerIdMax-1))
	}
	if dataCenterId < 0 || dataCenterId > datacenterIdMax {
		return errors.New(fmt.Sprintf("dataCenterId must be between 0 and %d", datacenterIdMax-1))
	}
	if snowflakeInstance == nil {
		instanceMutex.Lock()
		if snowflakeInstance != nil {
			instanceMutex.Unlock()
			return nil
		}
		defer instanceMutex.Unlock()
		snowflakeInstance = &snowflake{workerId: workerId, dataCenterId: dataCenterId}
		return nil
	} else {
		return nil
	}
}

func SnowflakeId() int64 {
	if snowflakeInstance == nil {
		instanceMutex.Lock()
		if snowflakeInstance != nil {
			instanceMutex.Unlock()
			return snowflakeInstance.nextVal()
		}
		defer instanceMutex.Unlock()
		return snowflakeInstance.nextVal()
	} else {
		return snowflakeInstance.nextVal()
	}
}

func GetSnowflakeId[T string | int64]() (id T) {
	var v interface{}
	if snowflakeInstance == nil {
		instanceMutex.Lock()
		if snowflakeInstance != nil {
			instanceMutex.Unlock()
			v = snowflakeInstance.nextVal()
		}
		defer instanceMutex.Unlock()
		v = snowflakeInstance.nextVal()
	} else {
		v = snowflakeInstance.nextVal()
	}
	t := reflect.TypeOf(id)
	if t.Kind() == reflect.String {
		reflect.ValueOf(&id).Elem().SetString(fmt.Sprintf("%+v", v))
	} else if t.Kind() == reflect.Int64 {
		reflect.ValueOf(&id).Elem().SetInt(v.(int64))
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
		return 0
	}
	s.timestamp = now
	r := t<<timestampShift | (s.dataCenterId << dataCenterIdShift) | (s.workerId << workerIdShift) | (s.sequence)
	return r
}
