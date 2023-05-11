package rpx

import (
	"errors"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 简单的RPC服务示例

// Service.Method
type DemoService struct {
}

type Args struct {
	A, B int
}

// 定义一个RPC服务参数是args，返回值是result
func (d DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("参数错误")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}

// 启动一个JSON服务
func ServeRpc(host string, service interface{}) error {
	// 注册RPC服务
	rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Printf("RPC服务启动失败:%s", err)
	}
	for {
		connect, err := listener.Accept()
		// 连接异常
		if err != nil {
			log.Printf("accept error: %s", err)
			continue
		}
		// 将连接交给JSONRPC去处理
		jsonrpc.ServeConn(connect)
	}
}

// 创建JSONRPC客户端
func NewClient(host string) (*rpc.Client, error) {
	connect, err := net.Dial("tpc", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(connect), nil

}
