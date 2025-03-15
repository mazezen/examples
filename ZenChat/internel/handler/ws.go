package handler

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/mazezen/itools"
	"github.com/mazezen/zenchat/internel/dao"
	"github.com/mazezen/zenchat/internel/models"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gopkg.in/fatih/set.v0"
	"net"
	"strconv"
	"sync"
)

type Node struct {
	Conn       *websocket.Conn
	Addr       string
	DataQueue  chan []byte
	GroupsSets set.Interface
}

var clientMap map[uint64]*Node = make(map[uint64]*Node, 0)
var rwLock sync.RWMutex

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				itools.AppLog.Error("发送消息,写入消息失败", zap.Error(err))
				return
			}
			itools.AppLog.Info("发送消息,写入消息成功")
		}
	}
}

func recProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			itools.AppLog.Error("读取消息失败", zap.Error(err))
			return
		}
		broadMsg(data)
	}
}

var upSendChan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	upSendChan <- data
}
func init() {
	go udpSendProc()
	go udpRecProc()
}

func udpSendProc() {
	udpConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
		Zone: "",
	})
	if err != nil {
		itools.AppLog.Error("udp connect err", zap.Error(err))
		return
	}
	defer udpConn.Close()
	for {
		select {
		case data := <-upSendChan:
			_, err = udpConn.Write(data)
			if err != nil {
				itools.AppLog.Error("写入UDP消息失败", zap.Error(err))
				return
			}
		}
	}
}

func udpRecProc() {
	udpConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
		Zone: "",
	})
	if err != nil {
		itools.AppLog.Error("udp connect err", zap.Error(err))
		return
	}
	defer udpConn.Close()
	for {
		var buf [1024]byte
		n, err := udpConn.Read(buf[0:])
		if err != nil {
			itools.AppLog.Error("从UDP中读取消息失败", zap.Error(err))
			return
		}
		dispatch(buf[0:n])
	}
}

func dispatch(data []byte) {
	message := &models.CMessage{}
	err := json.Unmarshal(data, message)
	if err != nil {
		itools.AppLog.Error("json unmarshal消息失败", zap.Error(err))
		return
	}
	switch message.Type {
	case 1: // 单聊
		sendMsgAndSave(message.TargetId, data)
	case 2: // 群聊
		sendCommunityMessage(message.FromId, message.TargetId, data)
	}
}

func sendMsgAndSave(id uint64, data []byte) {
	rwLock.Lock()
	defer rwLock.Unlock()
	targetNode, ok := clientMap[id]

	message := models.CMessage{}
	_ = json.Unmarshal(data, &message)
	ctx := context.Background()
	targetIdStr := strconv.Itoa(int(id))
	userIdStr := strconv.Itoa(int(message.FromId))

	if ok {
		targetNode.DataQueue <- data
	}

	var key string
	if id > message.FromId {
		key = "msg_" + userIdStr + "_" + targetIdStr
	} else {
		key = "msg_" + targetIdStr + "_" + userIdStr
	}

	re, err := itools.Rc.ZRevRange(ctx, key, 0, -1).Result()
	if err != nil {
		itools.AppLog.Error("ZRevRange", zap.Error(err))
		return
	}
	score := float64(cap(re)) + 1
	resS, err := itools.Rc.ZAdd(ctx, key, redis.Z{Score: score, Member: message}).Result()
	if err != nil {
		itools.AppLog.Error("ZAdd", zap.Error(err))
		return
	}
	itools.AppLog.Info("", zap.Int64("resS", resS))
}

func sendCommunityMessage(fromId, targetId uint64, data []byte) (int, error) {
	userIds, err := dao.NewCommunityRelationDao.FindUsers(targetId)
	if err != nil {
		return -1, err
	}
	for _, userId := range *userIds {
		if fromId != userId {
			sendMsgAndSave(userId, data)
		}
	}
	return 0, nil
}
