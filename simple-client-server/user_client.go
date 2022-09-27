package simple_client_server

import (
	"sync"

	"google.golang.org/grpc"

	userpb "github.com/mshero7/go-grpc/protos/v1/user"
)

var (
	once sync.Once // sync.Once는 gRPC server 내에서 싱글톤으로 초기에 한번만 client를 생성하고, 한번만 생성되고 나서는 각 rpc내에서는 같은 client를 계속 사용하기 위함이다.
	userClient userpb.UserClient
)

func GetUserClient(serviceHost string) userpb.UserClient {
	once.Do(func() {
		conn, _ := grpc.Dial(serviceHost,
			grpc.WithInsecure(),
			grpc.WithBlock(),
		)

		userClient = userpb.NewUserClient(conn)
	})

	return userClient
}