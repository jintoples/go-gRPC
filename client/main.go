package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/jintoples/go-gRPC/student"
	"google.golang.org/grpc"
)

func getDataStudentByEmail(client pb.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := pb.Student{Email: email}
	student, err := client.FindStudentByEmail(ctx, &s)
	if err != nil {
		log.Fatalln("error when get student", err.Error())
	}

	fmt.Println(student)
}

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":3000", opts...)
	if err != nil {
		log.Fatalln("Error in dial", err.Error())
	}
	defer conn.Close()

	client := pb.NewDataStudentClient(conn)
	getDataStudentByEmail(client, "doe@mail.com")
}
