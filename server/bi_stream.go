package main

import (
	"io"
	"log"
	"time"

	pb "github.com/surajgoraicse/go-grpc-basic/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		// receiving the stream from the client side
		req, err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil {
			return err
		}
		// first we receive a stream from the client then we send a stream to the client 
		log.Println("Got request with name : ", req.Name , time.Now())
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		// sending the stream to the client
		if err := stream.Send(res); err != nil {
			return err
		}

	}
}