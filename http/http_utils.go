package http

import (
	"bytes"
	"fmt"
	"github.com/ccpwcn/kgo"
	"io"
	"net/http"
	"time"
)

type RequestType int
type ResponseType int

const (
	RequestJson      RequestType = 1
	RequestFormData  RequestType = 2
	RequestPlainText RequestType = 3
)

type MyHttpClient struct {
	timeout             time.Duration
	url                 string
	requestContentType  RequestType
	responseContentType ResponseType
	header              map[string]string
	logWriter           io.Writer
}

type ClientOption func(client *MyHttpClient)

func WithTimeout(timeout time.Duration) ClientOption {
	return func(client *MyHttpClient) {
		client.timeout = timeout
	}
}

func WithUrl(url string) ClientOption {
	return func(client *MyHttpClient) {
		client.url = url
	}
}

func WithHeader(header map[string]string) ClientOption {
	return func(client *MyHttpClient) {
		client.header = header
	}
}

func WithRequestContentType(contentType RequestType) ClientOption {
	return func(client *MyHttpClient) {
		client.requestContentType = contentType
	}
}

func WithResponseContentType(contentType ResponseType) ClientOption {
	return func(client *MyHttpClient) {
		client.responseContentType = contentType
	}
}

func NewMyHttpClient(options ...ClientOption) *MyHttpClient {
	client := &MyHttpClient{
		timeout: 15 * time.Second,
	}
	for _, option := range options {
		option(client)
	}
	return client
}

func (hc *MyHttpClient) Post(body []byte) (resBody []byte, err error) {
	client := http.Client{
		Timeout: hc.timeout,
	}
	buffer := bytes.NewBuffer(body)
	req, err := http.NewRequest(http.MethodPost, hc.url, buffer)
	if err != nil {
		return nil, fmt.Errorf("创建HTTP POST请求对象出错：%w", err)
	}
	hc.processHeader(req)
	switch hc.requestContentType {
	case RequestJson:
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		break
	case RequestFormData:
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
		break
	case RequestPlainText:
		req.Header.Set("Content-Type", "text/html; charset=utf-8")
		break
	default:
		break
	}
	// 发起请求
	return hc.sendRequest(client, req)
}

func (hc *MyHttpClient) Get() (resBody []byte, err error) {
	client := http.Client{
		Timeout: hc.timeout,
	}
	req, err := http.NewRequest(http.MethodGet, hc.url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建HTTP POST请求对象出错：%w", err)
	}
	hc.processHeader(req)
	return hc.sendRequest(client, req)
}

func (hc *MyHttpClient) processHeader(req *http.Request) {
	if hc.header != nil && len(hc.header) > 0 {
		for k, v := range hc.header {
			req.Header.Add(k, v)
		}
	}
}

func (hc *MyHttpClient) sendRequest(client http.Client, req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发起POST请求出错：%w", err)
	}
	defer func(Body io.ReadCloser) {
		err1 := Body.Close()
		if err1 != nil {
			if hc.logWriter != nil {
				_, _ = hc.logWriter.Write(kgo.S2B("关闭URL " + hc.url + " 的内容读取器出错：" + err1.Error()))
			}
		}
	}(resp.Body)
	respBytes, err := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求失败，状态码 %d", resp.StatusCode)
	}
	return respBytes, err
}
