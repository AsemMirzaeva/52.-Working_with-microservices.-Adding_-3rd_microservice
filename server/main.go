package main

import (
	pq "clientstream/pq"
	pb "clientstream/salespb"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type MyServer struct {
	pb.UnimplementedSalesServiceServer
	TransactionSt *pq.SalesTransactionRepo
}

func (s *MyServer) StreamSalesTransactions(stream pb.SalesService_StreamSalesTransactionsServer) error {



	var (
		amount              int
		countOfTransactions int
	)
	for {
		reading, err := stream.Recv()
		if err == io.EOF {

			summary := pb.SalesSummary{
				TotalAmount:       float32(amount),
				TotalTransactions: int32(countOfTransactions),
			}

			err := s.TransactionSt.SaveSummary(&summary)
			if err != nil {
				return err
			}

			return stream.SendAndClose(&summary)
		}

		if err != nil {
			return err
		}

		amount += int(reading.Price)
		countOfTransactions++

		err = s.TransactionSt.SaveTransaction(reading)
		if err != nil {
			return err
		}
	}
}

func main() {

	db, err := pq.ConnectDB()
	if err != nil {
		return 
	}
	defer db.Close()
 
	port := ":9001"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	server := MyServer{TransactionSt: pq.NewSalesTransactionRepo(db)}
	pb.RegisterSalesServiceServer(s, &server)

	log.Printf("Server is listening on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
