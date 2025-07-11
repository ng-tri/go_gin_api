package auth

import (
	"context"
	"go-gin-api/internal/auth/pb"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
}

var jwtKey = []byte("secret-key")

func (s *Server) VerifyToken(ctx context.Context, req *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	tokenStr := strings.TrimPrefix(req.Token, "Bearer ")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return &pb.VerifyResponse{Valid: false, Message: "Invalid token"}, nil
	}

	return &pb.VerifyResponse{Valid: true, Message: "Token is valid"}, nil
}
