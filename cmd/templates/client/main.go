package main

import (
	templatesV1Pb "ago_grpc_1/pkg/templates/v1"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"os"
	"time"
)

const defaultPort = "9999"
const defaultHost = "0.0.0.0"

func main() {
	port, ok := os.LookupEnv("APP_PORT")
	if !ok {
		port = defaultPort
	}

	host, ok := os.LookupEnv("APP_HOST")
	if !ok {
		host = defaultHost
	}

	if err := execute(net.JoinHostPort(host, port)); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func execute(addr string) (err error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			if err == nil {
				err = cerr
				return
			}
			log.Print(err)
		}
	}()

	client := templatesV1Pb.NewServiceClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	_, err = client.Create(ctx, &templatesV1Pb.CreateRequest{Phone: "+13922233322", Title: "to Jon"})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			log.Print(st.Code())
			log.Print(st.Message())
		}
		return err
	}
	log.Println("Create 'to Jon'")

	response1, err := client.GetList(ctx, &emptypb.Empty{})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			log.Print(st.Code())
			log.Print(st.Message())
		}
		return err
	}
	log.Println("GetList", response1)

	response2, err := client.GetById(ctx, &templatesV1Pb.GetByIdRequest{Id: 2})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			log.Print(st.Code())
			log.Print(st.Message())
		}
		return err
	}
	log.Println("GetById", response2)

	_, err = client.UpdateById(ctx, &templatesV1Pb.UpdateRequest{Id: 1, Title: "best friend"})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			log.Print(st.Code())
			log.Print(st.Message())
		}
		return err
	}
	log.Println("UpdateById")

	_, err = client.DeleteById(ctx, &templatesV1Pb.GetByIdRequest{Id: 3})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			log.Print(st.Code())
			log.Print(st.Message())
		}
		return err
	}
	log.Println("DeleteById")

	response3, err := client.GetList(ctx, &emptypb.Empty{})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			log.Print(st.Code())
			log.Print(st.Message())
		}
		return err
	}
	log.Println("GetList", response3)

	return nil
}
