package api

import (
	"context"
	"fmt"

	pb "github.com/mfmmq/fancy-text/pkg/proto"
)

// Server that generates pretty text
type Server struct {
	pb.UnimplementedFancyTextGeneratorServer
}
func (s *Server) GetFancyText(ctx context.Context, incomingText *pb.FancyText) (*pb.FancyText, error) {
	// Log the request and init some vars
	fmt.Printf("Got request: %v\n", incomingText)


	// Send protobuf back to the user
	return &pb.FancyText{
		Text: incomingText.Text, 
	}, nil
}
