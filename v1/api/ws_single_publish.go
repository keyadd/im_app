package v1

import (
	"encoding/json"
	"github.com/pion/webrtc/v2"
	"go.uber.org/zap"
	"im_app/global"
	"im_app/initialize/wsmanage"
	"im_app/middleware"
	"im_app/utils"
	"im_app/v1/model/request"
	"log"
	"time"
)

type Publish struct {
	*wsmanage.BaseRouter
}

// PreHandle 前置拦截器 处理token
func (p Publish) PreHandle(r wsmanage.Request) error {
	err := middleware.WsJWTMiddleware(r)
	return err
}

func (p Publish) Handle(r wsmanage.Request) {
	data := r.GetData()
	c := r.GetConnection()
	userId, is := c.Get(global.GVA_CONFIG.JWT.UserIdName)
	if is != true {
		global.GVA_LOG.Error("获取不到map中user_id")
		utils.ResponseSuccess(r, &c, "token失效")
		return
	}
	msg := request.Publish{}
	res, err := utils.MapToStruct(data, &msg)
	if err != nil {
		global.GVA_LOG.Error("utils.MapToStruct(data, &joinDate)", zap.Error(err))
		return
	}
	message := res.(*request.Publish)

	c.AddWebRTCPeer(userId.(string), true)

	jsep := webrtc.SessionDescription{
		Type: webrtc.SDPTypeOffer,
		SDP:  message.Jsep.Sdp.(string),
	}

	answer, err := Answer(c, userId.(string), "", jsep, true)
	if err != nil {
		log.Print("创建Answer失败")
		return
	}

	resp := make(map[string]interface{})
	resp["jsep"] = answer
	resp["userId"] = userId
	respByte, err := json.Marshal(resp)
	if err != nil {
		return
	}
	respStr := string(respByte)

	if respStr != "" {
		mgr := c.Server.GetConnMgr()

		iconn, err := mgr.GetConn(userId.(int64))
		if err != nil {
			global.GVA_LOG.Error("conn, err := mgr.GetConn(msg.FriendId)", zap.Error(err))
			return
		}
		utils.ResponseSuccess(r, iconn, respStr)

		conn, err := mgr.GetConn(message.FriendId)
		if err != nil {
			global.GVA_LOG.Error("conn, err := mgr.GetConn(msg.FriendId)", zap.Error(err))
			return
		}

		utils.ResponseSuccess(r, conn, respStr)

	}
	//用户id 连接信息添加
	//mgr := c.Server.GetConnMgr()
	//for _, v := range ids {
	//
	//	conn, err := mgr.GetConn(v)
	//
	//	if err != nil {
	//		global.GVA_LOG.Error("conn, err := mgr.GetConn(msg.FriendId)", zap.Error(err))
	//		return
	//	}
	//	utils.ResponseSuccess(r, conn, message.Content)
	//}

}

func Answer(c wsmanage.Connection, id string, pubid string, offer webrtc.SessionDescription, sender bool) (webrtc.SessionDescription, error) {

	p := c.GetWebRTCPeer(id, sender)

	var err error
	var answer webrtc.SessionDescription
	if sender {
		answer, err = p.AnswerSender(offer)
	} else {
		c.PubPeerLock.RLock()

		pub := c.PubPeers[pubid]
		c.PubPeerLock.RUnlock()
		ticker := time.NewTicker(time.Millisecond * 2000)
		for {
			select {
			case <-ticker.C:
				goto ENDWAIT
			default:
				if pub.VideoTrack == nil || pub.AudioTrack == nil {
					time.Sleep(time.Millisecond * 100)
				} else {
					goto ENDWAIT
				}
			}
		}

	ENDWAIT:
		answer, err = p.AnswerReceiver(offer, &pub.VideoTrack, &pub.AudioTrack)

	}
	return answer, err

}
