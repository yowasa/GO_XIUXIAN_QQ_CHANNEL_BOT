package model

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	gorm.Model
	UserId    string
	UserName  string
	ChannelId string
	MsgId     string
	Msg       string
	EndAt     *time.Time // 执行结束时间
}

func NewEvent(user User, channelId string, msgId string, msg string, time time.Time) {
	var event = Event{
		UserId:    user.UserId,
		UserName:  user.UserName,
		ChannelId: channelId,
		MsgId:     msgId,
		Msg:       msg,
		EndAt:     &time,
	}
	db.Save(&event)
}

func (e *Event) Save() {
	db.Save(&e)
}

func (e *Event) Del() {
	db.Delete(&e)
}

func (e *Event) SelectAll() *[]Event {
	var events *[]Event
	db.Find(&events)
	return events
}

func GetNeedToDeal() *[]Event {
	var events *[]Event
	now := time.Now()
	db.Debug().Find(&events, "end_at <= ?", now)
	return events
}
