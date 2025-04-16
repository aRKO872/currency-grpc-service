package server

import (
	"context"
	"github.com/aRKO872/currency-dummy-grpc-service/protos/currency"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log hclog.Logger
	currency.UnimplementedCurrencyServer
}

func NewCurrency(log hclog.Logger) *Currency {
	return &Currency{
		log: log,
	}
}

// To test this do this below : 
// grpcurl --plaintext -d '{"Base":"INR", "Destination":"USD"}' localhost:8081 Currency.GetRate
// Here the key and values for Rate Request are given.
// The server host is given with the port and also the the Service with it's RPC is mentioned

// Also a few commands to play with for GRPC Reqs : 
// - grpcurl --plaintext localhost:8081 describe Currency.GetRate
// - grpcurl --plaintext localhost:8081 list
// - grpcurl --plaintext localhost:8081 list Currency

func (c *Currency) GetRate(ctx context.Context, protoReq *currency.RateRequest) (*currency.RateResponse, error) {
	c.log.Info("info logged", "base", protoReq.GetBase(), "dest", protoReq.GetDestination())
	return &currency.RateResponse{
		Rate: 0.5,
	}, nil
}