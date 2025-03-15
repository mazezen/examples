package service

import (
	"context"
	"github.com/mazezen/itools"
	"github.com/mazezen/zenchat/internel/handler/in"
	"go.uber.org/zap"
)

type ChatService struct{}

func NewChatService() *ChatService {
	return &ChatService{}
}

func (chat *ChatService) ChatMsg(payload *in.ChatMsgPayload) []string {
	var key string
	userIdStr, _ := itools.ToString(payload.UserId)
	targetIdStr, _ := itools.ToString(payload.TargetId)
	if payload.UserId > payload.TargetId {
		key = "msg_" + targetIdStr + "_" + userIdStr
	} else {
		key = "msg_" + userIdStr + "_" + targetIdStr
	}

	var reLs []string
	var err error
	if payload.IsRev {
		reLs, err = itools.Rc.ZRange(context.Background(), key, payload.Start, payload.End).Result()
	} else {
		reLs, err = itools.Rc.ZRevRange(context.Background(), key, payload.Start, payload.End).Result()
	}
	if err != nil {
		itools.AppLog.Error("获取聊天记录失败", zap.Error(err))
		return reLs
	}
	return reLs
}
