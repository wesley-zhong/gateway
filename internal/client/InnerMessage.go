package client

import (
	protoGen "gateway/proto"
	"google.golang.org/protobuf/proto"
)

type InnerMessage struct {
	InnerHeader *protoGen.InnerHead
	Body        proto.Message
}
