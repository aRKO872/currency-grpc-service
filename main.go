package main

import (
	protos "github.com/aRKO872/currency-dummy-grpc-service/protos/currency"
	"github.com/aRKO872/currency-dummy-grpc-service/server"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main () {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrency(log)
	
	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Error("error listening to port", "closing", err.Error)
		os.Exit(1)
	}

	if err := gs.Serve(l); err != nil {
		log.Error("error running server", "closing", err.Error)
		os.Exit(1)
	}
}