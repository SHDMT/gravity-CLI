package cmd

import (
	"fmt"
	gwalletRPC "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
	gravityRPC "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

// rpcURL specifies url for RPC services
//const rpcURL = "10.18.40.1:50052"
const (
	gravityRPCURL = "localhost:50053"
	gwalletRPCURL = "localhost:50052"
)


//设置超时时间
const seconds = 10

type gwalletRPCClient struct {
	gwalletRPC.AccountServiceClient
	conn   *grpc.ClientConn
	cancel context.CancelFunc
}

func (c *gwalletRPCClient) close() {
	if c.AccountServiceClient != nil {
		c.cancel()
		c.conn.Close()
	}
}

func newGwalletRPCClient() (*gwalletRPCClient, context.Context) {
	c := new(gwalletRPCClient)
	var err error
	var ctx context.Context

	c.conn, err = grpc.Dial(gwalletRPCURL, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("grpc client dial failed ")
	}

	c.AccountServiceClient = gwalletRPC.NewAccountServiceClient(c.conn)
	ctx, c.cancel = context.WithTimeout(context.Background(),
		time.Second*time.Duration(seconds))

	return c, ctx

}

type gravityRPCClient struct {
	gravityRPC.ServiceClient
	conn   *grpc.ClientConn
	cancel context.CancelFunc
}

func (c *gravityRPCClient) close() {
	if c.ServiceClient != nil {
		c.cancel()
		c.conn.Close()
	}
}

func newGravityRPCClient() (*gravityRPCClient, context.Context) {
	c := new(gravityRPCClient)
	var err error
	var ctx context.Context

	c.conn, err = grpc.Dial(gravityRPCURL, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("grpc client dial failed ")
	}

	c.ServiceClient = gravityRPC.NewServiceClient(c.conn)
	ctx, c.cancel = context.WithTimeout(context.Background(),
		time.Second*time.Duration(seconds))

	return c, ctx

}
