package rpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/rcrowley/go-metrics"
	cserver "github.com/rpcxio/rpcx-consul/serverplugin"
	"github.com/smallnest/rpcx/server"
	"log"
	"meadmin/library/context/api"
	"meadmin/library/rpc/config"
	"meadmin/library/rpc/protocol"
	"sync"
	"time"
)

type RPCServer struct {
	Server   *server.Server
	Config   config.Config
	Handlers map[string]api.ResultHandleFunc
}

var once sync.Once
var rpcServer *RPCServer

func GetRpcServer() *RPCServer {
	if rpcServer != nil {
		return rpcServer
	}
	return nil
}

func CreateRpcServer(serviceName string, conf config.Config) *RPCServer {
	once.Do(func() {
		s := server.NewServer()
		rpcServer = &RPCServer{
			Server: s,
			Config: config.Config{
				Network: "tcp",
				Listen:  "127.0.0.1:8007",
			},
			Handlers: map[string]api.ResultHandleFunc{},
		}

		if len(serviceName) > 0 {
			rpcServer.Config.BasePath = "/" + serviceName
			rpcServer.Config.ServiceName = serviceName
		}

		if len(conf.Network) > 0 {
			rpcServer.Config.Network = conf.Network
		}

		if len(conf.Listen) > 0 {
			rpcServer.Config.Listen = conf.Listen
		}

		if len(conf.DiscoveryAddrs) > 0 {
			rpcServer.Config.DiscoveryAddrs = conf.DiscoveryAddrs
		}

		addRegistryPlugin(rpcServer)
	})
	return rpcServer
}

func addRegistryPlugin(s *RPCServer) {
	sadd := fmt.Sprintf("%s@%s", s.Config.Network, s.Config.Listen)
	r := &cserver.ConsulRegisterPlugin{
		ServiceAddress: sadd,
		ConsulServers:  s.Config.DiscoveryAddrs,
		BasePath:       s.Config.BasePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Server.Plugins.Add(r)
}

func (s *RPCServer) Handle(funcName string, fn api.ResultHandleFunc, metadata string) error {
	err := s.Server.RegisterFunctionName(s.Config.ServiceName, funcName, s.CallForFunction, metadata)
	if err == nil {
		s.Handlers[funcName] = fn
	}
	return err
}

func (s *RPCServer) Start() error {
	log.Printf("Listening and serving RPC on :%s", s.Config.Listen)
	err := s.Server.Serve(s.Config.Network, s.Config.Listen)
	return err
}

func (s *RPCServer) CallForFunction(ctx context.Context, args *protocol.Args, replyv *protocol.Reply) error {
	log.Printf("rcp request:" + args.HandlerName)
	if len(args.HandlerName) == 0 {
		return errors.New("No HandlerName args")
	}
	fn, ok := s.Handlers[args.HandlerName]
	if !ok {
		return errors.New("No handle function for" + args.HandlerName)
	}
	apiCtx := api.NewRPCContext(ctx, args)
	result := fn(apiCtx)
	replyv.Data = result
	return nil
}
