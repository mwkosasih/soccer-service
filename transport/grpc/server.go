package server

import (
	"context"
	"github.com/mwkosasih/soccer-service/action"
	"github.com/mwkosasih/soccer-service/builder"
	pb "github.com/mwkosasih/soccer-service/transport/grpc/proto/soccer"
	log "github.com/sirupsen/logrus"
)

type Server struct {
}

func (s Server) CreateTeam(ctx context.Context, req *pb.CreateTeamRequest) (*pb.NoResponse, error) {
	var resp pb.NoResponse

	err := action.NewCreateTeam().Handle(ctx, req.GetName())
	if err != nil {
		log.Errorf("Server.CreateTeam Failed: %v", err)
		return nil, err
	}

	return &resp, nil
}

func (s Server) CreatePlayer(ctx context.Context, req *pb.CreatePlayerRequest) (*pb.NoResponse, error) {
	var resp pb.NoResponse

	createPlayerReq := builder.BuildCreatePlayerRequest(req)
	err := action.NewCreatePlayer().Handle(ctx, createPlayerReq)
	if err != nil {
		log.Errorf("Server.CreatePlayer Failed: %v", err)
		return nil, err
	}

	return &resp, nil
}

func (s Server) GetTeam(ctx context.Context, req *pb.GetTeamRequest) (*pb.GetTeamResponse, error) {
	var resp pb.GetTeamResponse

	teamResp, playerResp, err := action.NewGetTeam().Handle(ctx, req.GetId())
	if err != nil {
		log.Errorf("Server.GetTeam Failed: %v", err)
		return nil, err
	}

	resp = builder.BuildGetTeamResponse(teamResp, playerResp)
	return &resp, nil
}

func (s Server) GetPlayer(ctx context.Context, req *pb.GetPlayerRequest) (*pb.GetPlayerResponse, error) {
	var resp pb.GetPlayerResponse

	response, err := action.NewGetPlayer().Handle(ctx, req.GetId())
	if err != nil {
		log.Errorf("Server.GetPlayer Failed: %v", err)
		return nil, err
	}

	resp.Player = builder.BuildGetPlayerResponse(response)
	return &resp, nil
}

// NewServer create new Server
func NewServer() *Server {
	return &Server{}
}
