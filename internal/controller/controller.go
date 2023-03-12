package controller

import (
	"gateway/internal/client"
	"gateway/internal/message"
	"gateway/pkg/core"
	"gateway/pkg/log"
	"gateway/pkg/network"
	protoGen "gateway/proto"

	"google.golang.org/protobuf/proto"
)

var gameClient *client.ConnInnerClientContext

func Init() {
	core.RegisterMethod(int32(protoGen.ProtoCode_LOGIN_REQUEST), &protoGen.LoginRequest{}, login)
	context, err := network.Dial("tcp", "127.0.0.1:9007")
	if err != nil {
		log.Error(err)
		return
	}
	gameClient = client.NewInnerClientContext(context)
	//add  msg  to game server to add me
	header := &protoGen.InnerHead{
		FromServerUid:    message.BuildServerUid(message.TypeGateway, 35),
		ToServerUid:      0,
		ReceiveServerUid: 0,
		Id:               0,
		SendType:         0,
		ProtoCode:        message.INNER_PROTO_ADD_SERVER,
		CallbackId:       0,
	}

	innerMessage := &client.InnerMessage{
		InnerHeader: header,
		Body:        nil,
	}
	gameClient.Send(innerMessage)
}

func login(ctx network.ChannelContext, request proto.Message) {
	context := ctx.Context().(*client.ConnClientContext)
	loginRequest := request.(*protoGen.LoginRequest)
	log.Infof("login request = %s", loginRequest.LoginToken)
	innerLoginReq := &protoGen.InnerLoginRequest{
		SessionId: context.Sid,
		AccountId: loginRequest.AccountId,
		RoleId:    loginRequest.RoleId,
	}
	msgHeader := &protoGen.InnerHead{
		FromServerUid:    message.BuildServerUid(message.TypeGateway, 35),
		ToServerUid:      message.BuildServerUid(message.TypeGame, 35),
		ReceiveServerUid: 0,
		Id:               loginRequest.RoleId,
		SendType:         0,
		ProtoCode:        message.INNER_PROTO_LOGIN_REQUEST,
		CallbackId:       0,
	}

	innerMsg := &client.InnerMessage{
		InnerHeader: msgHeader,
		Body:        innerLoginReq,
	}
	gameClient.Send(innerMsg)
}
