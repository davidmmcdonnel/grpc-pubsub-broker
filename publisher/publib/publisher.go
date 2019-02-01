package publisher

import (
	"context"

	"google.golang.org/grpc/grpclog"

	pb "github.com/davidmmcdonnel/grpc-pubsub-broker/protobuf"
)

func Publish(client pb.PublisherClient, key string, msg *pb.Message) {
	request := &pb.PublishRequest{Key: key, Messages: []*pb.Message{msg}}
	_, error := client.Publish(context.Background(), request)
	if error != nil {
		grpclog.Printf("Error publishing")

	}

}
