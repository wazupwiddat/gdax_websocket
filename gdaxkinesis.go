package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/kinesis"
)

// GdaxKinesis provides stuff
type GdaxKinesis struct {
	streamName string
	session    *session.Session
	stream     *kinesis.Kinesis
}

// NewKinesisStream creates a new instance of the GDAX Kinesis stream.
func NewKinesisStream(name string, region string) *GdaxKinesis {
	//s := session.New(&aws.Config{Region: aws.String(region)})
	//s := session.Must(session.NewSession())
	s := session.Must(session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(region)},
		Profile: "jdub",
	}))

	gk := &GdaxKinesis{
		streamName: name,
		session:    s,
		stream:     kinesis.New(s),
	}

	if !gk.streamExists(name) {
		gk.createStream(name, 1)
	}
	return gk
}

func (gk *GdaxKinesis) createStream(name string, shards int64) {
	out, err := gk.stream.CreateStream(&kinesis.CreateStreamInput{
		ShardCount: aws.Int64(shards),
		StreamName: aws.String(name),
	})
	if err != nil {
		panic(err)
	}
	log.Printf("%v\n", out)

	if err := gk.stream.WaitUntilStreamExists(&kinesis.DescribeStreamInput{StreamName: &name}); err != nil {
		panic(err)
	}

	streams, err := gk.stream.DescribeStream(&kinesis.DescribeStreamInput{StreamName: &name})
	if err != nil {
		panic(err)
	}
	log.Printf("%v\n", streams)
}

func (gk GdaxKinesis) streamExists(name string) bool {
	streams, err := gk.stream.DescribeStream(&kinesis.DescribeStreamInput{StreamName: &name})
	if err != nil {
		return false
	}
	return streams != nil
}

func (gk GdaxKinesis) writeMessage(message []byte) error {
	//log.Println(string(message))
	putOutput, err := gk.stream.PutRecord(&kinesis.PutRecordInput{
		Data:         message,
		StreamName:   &gk.streamName,
		PartitionKey: aws.String("key1"),
	})
	if err != nil {
		log.Println("writemessage:", err)
	}
	log.Printf("%v\n", putOutput)
	return nil
}
