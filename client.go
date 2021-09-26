package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-linking/protobuf/query"
	"grpc-linking/protobuf/search"
	"log"
)
const SERVER_PORT = "9001"
func main() {

	conn, err := grpc.Dial(":"+SERVER_PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()
	clientSearch := search.NewSearchServiceClient(conn)
	searchResp, err := clientSearch.Search(context.Background(), &search.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}
	log.Printf("resp: %s", searchResp.GetResponse())


	clientQuery := query.NewQueryServiceClient(conn)

	queryRsp, err := clientQuery.Query(context.Background(), &query.QueryRequest{
		Name: "linking",
		Size: "big",
	})
	if err != nil {
		log.Fatalf("client.Query err: %v", err)
	}
	log.Printf("resp: %s,%s", queryRsp.GetDes(),queryRsp.GetPhone())

}