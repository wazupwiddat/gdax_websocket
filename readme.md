# gdax_websocket

This go program outputs the BTC/USD orderbook from [GDAX](https://www.gdax.com/) to AWS Kinesis Data Stream.

Need to have a AWS Account (Free Tier should work fine).

To setup AWS account:

1. Setup your AWS account [AWS](https://aws.amazon.com/)
2. Create Access Key [Security Credentials](https://console.aws.amazon.com/iam/home?region=us-east-1#/security_credential)
3. Note credentials profile

The websocket library is implemented via [https://github.com/gorilla/websocket](https://github.com/gorilla/websocket).

To run:

1. Clone this repository into your local go/src directory.
2. Switch to the gdax_websocket directory
3. Make files executable: `chmod u+x *.go`
4. Grab dependencies: `go get ./...`
5. Build: `go build`
6. Run:	`go run main.go gdax.go subscribe.go gdaxkinesis.go`

