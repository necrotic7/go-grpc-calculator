package main

import (
	"fmt"
	calculatorPB "go-prac/grpc/proto/calculator"
	"log"
	"net"

	"context"

	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
)

type Server struct {
	calculatorPB.UnimplementedCalculatorServiceServer
}

// 實作sum方法
func (s *Server) Sum(ctx context.Context, req *calculatorPB.CalculatorRequest) (*calculatorPB.CalculatorResponse, error) {
	result := decimal.NewFromInt(req.A).Add(decimal.NewFromInt(req.B)).IntPart()
	log.Println(fmt.Sprintf("計算結果 A(%v) + B(%v) = %v", req.A, req.B, result))
	res := &calculatorPB.CalculatorResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	log.Println("starting gRPC server...")
	port := "50051"
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	calculatorPB.RegisterCalculatorServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Panicf("failed to Serve: %v \n", err)
	} else {
		log.Printf("gRPC server listening on %s \n", port)
	}

}
