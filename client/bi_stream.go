package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/surajgoraicse/go-grpc-basic/proto"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Println("Bidirectional streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalln("could not send names ", err)
	}

	// receiving stream from the server
	waitc := make(chan struct{})
	go func() { //
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln("Error while streaming : ", err)
			}
			log.Println(message.Message, time.Now())
		}
		close(waitc)
	}()

	// sending stream to the server
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalln("Error while send stream from clent", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Println("Bidirectional streaming finished")

}
