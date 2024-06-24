package main

import (
	pb "clientstream/salespb"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	clientServer := pb.NewSalesServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	stream, err := clientServer.StreamSalesTransactions(ctx)
	if err != nil {
		log.Fatal(err)
	}

	readings := []pb.SalesTransaction{
		{TransactionId: "1", ProductId: "101", Quantity: 20, Price: 99.99, Timestamp: time.Now().Unix()},
		{TransactionId: "2", ProductId: "540", Quantity: 100, Price: 100.99, Timestamp: time.Now().Unix()},
		{TransactionId: "3", ProductId: "890", Quantity: 35, Price: 200.99, Timestamp: time.Now().Unix()},
		{TransactionId: "4", ProductId: "200", Quantity: 40, Price: 300.99, Timestamp: time.Now().Unix()},
		{TransactionId: "5", ProductId: "999", Quantity: 50, Price: 599.99, Timestamp: time.Now().Unix()},
		{TransactionId: "6", ProductId: "001", Quantity: 80, Price: 699.99, Timestamp: time.Now().Unix()},
	}

	for i := 0; i < len(readings); i++ {

		err := stream.Send(&readings[i])
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Transaction-%d\n", i+1)
		log.Println(&readings[i])
	}

	result, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Total amount: %f\nTotal transaction: %d", result.TotalAmount, result.TotalTransactions)
}
