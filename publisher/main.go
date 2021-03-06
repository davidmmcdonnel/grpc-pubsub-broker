package main

import (
	"flag"
	"io/ioutil"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/davidmmcdonnel/grpc-pubsub-broker/protobuf"
	publib "github.com/davidmmcdonnel/grpc-pubsub-broker/publisher/publib"
)

var (
	serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
	topic              = flag.String("topic", "Tennis", "Message topic")
	frequency          = flag.Int("frequency", 1000, "Publishing Frequency")
	file               = flag.String("file", "", "Input file")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()
	client := pb.NewPublisherClient(conn)

	var fileContent []byte
	if *file != "" {
		data, _ := ioutil.ReadFile(*file)
		fileContent = data
	}

	var speed int
	var size int

	go func(size *int, speed *int) {
		for {
			grpclog.Printf("Publishing Message (%f ko/s, %d m/s)", float64(*size)/1000.0, *speed)
			*speed = 0
			*size = 0
			time.Sleep(1 * time.Second)
		}
	}(&size, &speed)

	for {
		if fileContent == nil {
			for _, value := range flag.Args() {
				publib.Publish(client, *topic, &pb.Message{Data: []byte(value)})
				speed++
				size += int(len(value))
				time.Sleep(time.Duration(*frequency))
			}
		} else {
			publib.Publish(client, *topic, &pb.Message{Data: fileContent})
			speed++
			size += int(len(fileContent))
			time.Sleep(time.Duration(*frequency))
		}
	}
}
