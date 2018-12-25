package main

import (
	"controller"
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"types"
)

type TeacherControllerServerImpl struct {
}

func (server *TeacherControllerServerImpl) GetTeachers(context.Context, *controller.RequestInput) (*controller.ResponseMessage, error) {
	responseMessage := &controller.ResponseMessage{
		TeacherList: []*types.Teacher{
			{
				Age:  25,
				Id:   1,
				Name: "Milan",
			},
			{
				Age: 26, Id: 2, Name: "Nikul",
			},
		},
	}

	return responseMessage, nil
}

var addr = flag.String("addr", ":1337", "listen addr")

// main start a gRPC server and waits for connection
func main() {
	log.Println(addr)
	flag.Parse()
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	log.Print(err)
	log.Println(lis)

	// create a server instance
	s := TeacherControllerServerImpl{}
	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	controller.RegisterTeacherControllerServer(grpcServer, &s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		log.Print("Worked!")
	}

	log.Println("Server Started")
}
