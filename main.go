package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lwmacct/250300-go-mod-mgrpc/pkg/proto/v10/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	fmt.Println("Hello, World!")

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewAskClient(conn)

	{
		res, err := client.Int64AskBool(context.Background(), &pb.KvInt64{
			Val: 1,
		})

		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		fmt.Printf("res: %v\n", res)
	}

	{
		res, err := client.Int64AskString(context.Background(), &pb.KvInt64{
			Val: 1,
		})

		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		fmt.Printf("res: %v\n", res)
	}
	{

		// bytes
		res, err := client.Int64AskBytes(context.Background(), &pb.KvInt64{
			Val: 1,
		})

		_ = res.Val
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
	}

	{

		res, err := client.Any(context.Background(), &anypb.Any{
			TypeUrl: "type.googleapis.com/google.protobuf.StringValue",
			Value:   []byte(`"Hello, World!"`),
		})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		fmt.Printf("res: %v\n", res)
	}
}
