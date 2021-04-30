package builder

import (
	"github.com/mwkosasih/soccer-service/entity"
	pb "github.com/mwkosasih/soccer-service/transport/grpc/proto/soccer"
)

func BuildCreatePlayerRequest(req *pb.CreatePlayerRequest) (resp entity.Player) {
	resp = entity.Player{
		TeamID:   req.TeamId,
		Name:     req.Name,
		Position: req.Position,
		Height:   req.Height,
		Weight:   req.Weight,
	}

	return
}

func BuildGetPlayerResponse(req entity.Player) (resp *pb.Player) {
	resp = &pb.Player{
		Id:        req.ID,
		TeamId:    req.TeamID,
		Name:      req.Name,
		Position:  req.Position,
		Height:    req.Height,
		Weight:    req.Weight,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
	}

	return
}

func BuildGetTeamResponse(teamResp entity.Team, playerResp []entity.Player) (resp pb.GetTeamResponse) {
	var players []*pb.Player
	for _, v := range playerResp {
		players = append(players, &pb.Player{
			Id:        v.ID,
			TeamId:    v.TeamID,
			Name:      v.Name,
			Position:  v.Position,
			Height:    v.Height,
			Weight:    v.Weight,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	resp = pb.GetTeamResponse{
		Id:        teamResp.ID,
		Name:      teamResp.Name,
		CreatedAt: teamResp.CreatedAt,
		UpdatedAt: teamResp.UpdatedAt,
		Players:   players,
	}

	return
}
