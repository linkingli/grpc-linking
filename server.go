package main
import (
	"context"
	"google.golang.org/grpc"
	"grpc-linking/protobuf/query"
	"grpc-linking/protobuf/search"
	"log"
	"net"
)
type SearchService struct{
	search.UnimplementedSearchServiceServer
}
func (s *SearchService) Search(ctx context.Context, r *search.SearchRequest) (*search.SearchResponse, error) {
	return &search.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

type QueryService struct {
	query.UnimplementedQueryServiceServer
}
func (s *QueryService) Query(ctx context.Context, r *query.QueryRequest) (*query.QueryResponse, error) {
	return &query.QueryResponse{Des: "rea",Phone: "11999"}, nil
}


const PORT = "9001"
func main() {
	server := grpc.NewServer()
	search.RegisterSearchServiceServer(server, &SearchService{})
	query.RegisterQueryServiceServer(server, &QueryService{})
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(lis)
}