package handler

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/mazezen/itools"
	"github.com/mazezen/zenchat/internel/handler/in"
	"github.com/mazezen/zenchat/internel/service"
	"gopkg.in/fatih/set.v0"
	"net/http"
	"strconv"
)

type ChatHandler struct{}

var NewChatHandler = &ChatHandler{}

var upGrader = &websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Chat 聊天
// 单聊消息体: {"target_id":3,"Type":1,"from_id":2,"Media":1,"Content":"在干嘛"}
func (chat *ChatHandler) Chat(c echo.Context) error {
	id := c.QueryParam("userId")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}

	conn, err := upGrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	node := &Node{
		Conn:       conn,
		DataQueue:  make(chan []byte, 50),
		GroupsSets: set.New(set.ThreadSafe),
	}
	rwLock.Lock()
	clientMap[uint64(userId)] = node
	rwLock.Unlock()

	go sendProc(node)

	go recProc(node)

	return c.JSON(http.StatusOK, itools.Response.ResponseSuccess("", "success send message"))
}

func (chat *ChatHandler) ChatMsg(c echo.Context) error {
	var payload = new(in.ChatMsgPayload)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	if msg := itools.ZhValidateParam(payload); len(msg) > 0 {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, msg))
	}

	chatMsgs := service.NewChatService().ChatMsg(payload)
	return c.JSON(http.StatusOK, itools.Response.ResponseSuccess("", "success", chatMsgs))
}
