package main

import (
	"controller"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"types"
)

func main(){
	serverAddress := "127.0.0.1:7777"

	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())

	if e != nil {
		log.Fatal(e)
	} else {
		defer conn.Close()

		client := controller.NewTeacherControllerClient(conn)

		requestInput := controller.RequestInput{
			Department: &types.Department{
				Name: "Dummy Department",
				Id: 1,
				TeacherList: nil,
			},
		}

		responseMessage, e := client.GetTeachers(context.Background(), &requestInput)

		if e != nil {
			log.Fatal(e)
			return
		}

		log.Println(responseMessage)
	}
}
