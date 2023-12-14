package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	dynamodb "github.com/cohen2804/selfservice/dynamo"
	"github.com/cohen2804/selfservice/generated"
	"google.golang.org/grpc"
)

// should implement the interface myPkgName.InvoicerServer
type prictlGRPCServer struct {
	repo *dynamodb.Repository
	generated.UnimplementedSessionRecordsServiceServer
}

func NewPrictlGRPCServer(repo *dynamodb.Repository) *prictlGRPCServer {
	return &prictlGRPCServer{
		repo: repo,
	}
}

func (s *prictlGRPCServer) ListActions(ctx context.Context, in *generated.ListActionsRq) (*generated.ListActionsRs, error) {
	var err error
	var items []dynamodb.ActionItem
	var pageToken string

	if in.Query.Filter == nil {
		req := dynamodb.QueryRequest{Tenantid: in.TenantID, PageSize: in.Query.PageSize, PageToken: in.Query.PageToken}
		pageToken, items, err = s.repo.GetActionItemsByTenant(req)
	} else {
		req := dynamodb.QueryRequest{Tenantid: in.TenantID, PageSize: in.Query.PageSize, PageToken: in.Query.PageToken, FilterOptions: dynamodb.FilterOption{Field: in.Query.Filter.Field, Value: in.Query.Filter.Value}}
		pageToken, items, err = s.repo.GetFilteredActionItems(req)
	}
	if err != nil {
		fmt.Printf("ListActions: %s", err)
	}
	var res generated.ListActionsRs
	if len(items) > 0 {
		res.PageToken = pageToken

		for _, action := range items {
			res.Records = append(res.Records, &generated.ActionRecord{
				TenantID:     in.TenantID,
				SessionID:    action.Sessionid,
				Action:       generated.ActionType(generated.ActionType_value[action.Actiontype]),
				Status:       generated.Status(generated.ActionType_value[action.Status]),
				User:         &generated.User{UserName: action.Username, Subject: action.Subject},
				StartedDate:  time.Unix(action.StartedTime, 0).Format("02/01/2006 15:04"),
				FinishedDate: time.Unix(action.Finishedtime, 0).Format("02/01/2006 15:04"),
			})
		}
	}
	return &res, nil
}

func (s *prictlGRPCServer) GetActionLog(ctx context.Context, in *generated.GetActionLogRq) (*generated.GetActionLogRs, error) {
	var logs []*generated.LogItem
	logs = append(logs, &generated.LogItem{Message: "I am log"})
	return &generated.GetActionLogRs{LogsList: logs}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	repo, _ := dynamodb.NewRepository()
	prictlServer := NewPrictlGRPCServer(repo)
	generated.RegisterSessionRecordsServiceServer(grpcServer, prictlServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

//go run ./server/server.go