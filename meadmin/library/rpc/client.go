package rpc

import (
	"context"
	"flag"
	"github.com/mitchellh/mapstructure"
	dclient "github.com/rpcxio/rpcx-consul/client"
	"github.com/smallnest/rpcx/client"
	"meadmin/library/rpc/protocol"
)

var (
	consulAddr = flag.String("consulAddr", "60.164.247.74:8500", "consul address")
	basePath   = flag.String("base", "/rpcx_test", "prefix path")
)

type RemoteService struct {
	ServiceName string
	Discovery   *dclient.ConsulDiscovery
	Client      client.XClient
}

func NewRemoteService(basePath, serviceName string) *RemoteService {
	discovery, _ := dclient.NewConsulDiscovery(basePath, serviceName, []string{*consulAddr}, nil)
	xclient := client.NewXClient(serviceName, client.Failtry, client.RandomSelect, discovery, client.DefaultOption)
	//defer xclient.Close()
	return &RemoteService{
		ServiceName: serviceName,
		Discovery:   discovery,
		Client:      xclient,
	}
}

func (this *RemoteService) Call(funcName string, args any) (*protocol.Reply, error) {
	var remoteArgs protocol.Args
	remoteArgs.Data = args
	remoteArgs.HandlerName = funcName
	reply := &protocol.Reply{}
	err := this.Client.Call(context.Background(), funcName, remoteArgs, reply)
	if err != nil {
		return reply, err
	}
	return reply, nil
}

func CallRemoteService[T any](service *RemoteService, funcName string, args any) (*protocol.RpcResult[T], error) {
	reply, err := service.Call(funcName, args)
	if err != nil {
		return nil, err
	}
	var t protocol.RpcResult[T]
	err = mapstructure.Decode(reply.Data, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
