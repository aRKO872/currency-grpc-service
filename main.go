package main

import (
	"net"
	"os"

	"github.com/aRKO872/currency-grpc-service/data"
	protos "github.com/aRKO872/currency-grpc-service/protos/currency"
	"github.com/aRKO872/currency-grpc-service/server"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main () {
	log := hclog.Default()

	er, err := data.NewRates(log)
	if err != nil {
		log.Error("error creating new rates object", "closing", err.Error)
		os.Exit(1)
	}

	gs := grpc.NewServer()
	cs := server.NewCurrency(er, log)
	
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