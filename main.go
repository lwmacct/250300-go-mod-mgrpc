package main

import (
	"context"
	"net"

	v10pb "github.com/lwmacct/250300-go-mod-mgrpc/pkg/proto/v10/pb"
	v11pb "github.com/lwmacct/250300-go-mod-mgrpc/pkg/proto/v11/pb"
	"github.com/lwmacct/250300-go-mod-mlog/pkg/mlog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type app struct {
	conn *grpc.ClientConn
	v10c v10pb.RpcClient
	v11c v11pb.RpcClient
}

func (t *app) server() error {
	mlog.Info(mlog.H{"start server": "start server"})

	server := grpc.NewServer()
	v10pb.RegisterRpcServer(server, &v10pb.ImplExamples{})
	v11pb.RegisterRpcServer(server, &v11pb.ImplExamples{})
	reflection.Register(server)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	return server.Serve(lis)
}

func NewApp() (*app, error) {
	var err error
	t := &app{}
	t.conn, err = grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		mlog.Error(mlog.H{"err": err})
		return nil, err
	}
	t.v10c = v10pb.NewRpcClient(t.conn)
	t.v11c = v11pb.NewRpcClient(t.conn)

	go t.server()
	return t, nil
}

func (t *app) Int64AskString() {
	res, err := t.v10c.Int64AskString(context.Background(), &v10pb.KvInt64{
		Key: "test",
		Val: 99,
	})
	mlog.Info(mlog.H{"res": res, "err": err})
}

func main() {
	defer mlog.Close()
	mlog.Info(mlog.H{"Hello, World!": "Hello, World!"})
	app, err := NewApp()
	if err != nil {
		mlog.Error(mlog.H{"err": err})
		return
	}
	_ = app

	app.Int64AskString()

}
