package api

import (
	"context"
	"fmt"

	pb "github.com/mfmmq/fancy-text/pkg/proto"
)

// Server that generates pretty text
type ServerSimple struct {
	pb.UnimplementedFancyTextGeneratorServer
}
func (s *ServerSimple) GetFancyText(ctx context.Context, incomingText *pb.FancyText) (*pb.FancyText, error) {
	fmt.Printf("Got request: %v\n", incomingText)
	return &pb.FancyText{Text: incomingText.Text}, nil
}