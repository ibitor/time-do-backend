package core

import (
	"log"
	"timedo/wx"
)

func PushWX(corpID string, agentID int, agentSecret string, content string) {
	client := wx.New(corpID, agentID, agentSecret)
	msg := wx.Message{}
	msg.ToUser = "@all"
	msg.MsgType = "text"

	msg.Text = wx.Content{Content: content}
	err := client.Send(msg)
	if err != nil {
		log.Println(err)
	}
	log.Println("push success")
}
