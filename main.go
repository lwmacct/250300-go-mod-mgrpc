package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/lwmacct/250300-go-mod-mgrpc/pkg/proto/v10/pb"
	"github.com/lwmacct/250300-go-mod-mlog/pkg/mlog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type ts struct {
	conn   *grpc.ClientConn
	client pb.AskClient
}

type askServer struct {
	pb.UnimplementedAskServer
}

func (s *askServer) Int64AskString(ctx context.Context, in *pb.KvInt64) (*pb.KvString, error) {
	out := &pb.KvString{
		Key: "res",
		Val: fmt.Sprintf("%d", in.Val),
	}
	mlog.Info(mlog.H{
		"in":  in,
		"out": out,
	})
	return out, nil
}

func (t *ts) server() error {
	mlog.Info(mlog.H{"start server": "start server"})

	server := grpc.NewServer()
	pb.RegisterAskServer(server, &askServer{})
	reflection.Register(server)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	return server.Serve(lis)
}

func (t *ts) Init() *ts {
	var err error
	t.conn, err = grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		mlog.Error(mlog.H{"err": err})
	}
	t.client = pb.NewAskClient(t.conn)
	go t.server()
	return t
}

func main() {
	defer mlog.Close()
	mlog.Info(mlog.H{"Hello, World!": "Hello, World!"})
	t := new(ts).Init()

	t.client.Int64AskBool(context.Background(), &pb.KvInt64{
		Val: 1,
	})

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewAskClient(conn)

	{
		res, err := client.Int64AskString(context.Background(), &pb.KvInt64{
			Key: "test",
			Val: 99,
		})
		if err != nil {
			mlog.Error(mlog.H{"err": err})
			return
		}

		fmt.Printf("res: %v\n", res)
		_ = pb.MapStringString{
			Key: "test",
			Val: map[string]string{
				"test": "test",
			},
		}

	}

}
