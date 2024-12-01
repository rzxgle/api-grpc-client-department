package main

import (
	"apidepartment_client/src/pb/department"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//abrir a conexao
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("error on get connection. error: ", err)
	}
	defer conn.Close()

	//filtrando a consulta
	client := department.NewDepartmentServiceClient(conn)
	stream, err := client.ListPerson(context.Background(), &department.ListPersonRequest{DepartmentId: 1})
	if err != nil {
		log.Fatalln("error on get channel to stream. error: ", err)
	}

	//interando a resposta
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("error on recv. error: ", err)
		}

		fmt.Printf("response: %+v\n", response)
	}
}
