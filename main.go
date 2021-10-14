package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/djhworld/go-lambda-invoke/golambdainvoke"
)

func main() {
	data := events.CloudwatchLogsEvent{
		AWSLogs: events.CloudwatchLogsRawData{
			Data: "H4sIAAAAAAAAAHWPwQqCQBCGX0Xm7EFtK+smZBEUgXoLCdMhFtKV3akI8d0bLYmibvPPN3wz00CJxmQnTO41whwWQRIctmEcB6sQbFC3CjW3XW8kxpOpP+OC22d1Wml1qZkQGtoMsScxaczKN3plG8zlaHIta5KqWsozoTYw3/djzwhpLwivWFGHGpAFe7DL68JlBUk+l7KSN7tCOEJ4M3/qOI49vMHj+zCKdlFqLaU2ZHV2a4Ct/an0/ivdX8oYc1UVX860fQDQiMdxRQEAAA==",
		},
	}
	response, err := golambdainvoke.Run(golambdainvoke.Input{
		Port:    8001,
		Payload: data,
	})
	if err != nil {
		err = fmt.Errorf("panic with call error(%v)", err)
		panic(err)
	}
	log.Printf("Got resp: %v", response)
}
