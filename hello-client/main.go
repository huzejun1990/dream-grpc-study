package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "huzejun_go/dream-grpc-study/hello-server/proto"
	"log"
)

type PerRPCCredentials interface {
	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
	RequireTransportSecurity() bool
}

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "dream",
		"appKey": "123123",
	}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false

}

func main() {
	//creds,_ := credentials.NewClientTLSFromFile("F:\\environment\\GoWorks\\src\\huzejun_go\\dream-grpc-study\\key\\privkey.pem","*.dream.com")

	// 连接到server端，此处禁用安全传输，没有加密和验证
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	conn, err := grpc.Dial("127.0.0.1:9090", opts...)
	//conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 建立连接
	client := pb.NewSayHelloClient(conn)

	// 执行rpc调用 （这个方法在服务器端来实现并返回结果）
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "dream"})

	fmt.Println(resp.GetResponseMsg())
}
