package models

import (
	"github.com/cskr/pubsub"
	"time"
)

var PubSub *pubsub.PubSub

type Event struct {
	Id int

	Uuid            string `sql:"type:uuid;unique"`
	Serial          string
	TimeStamp       string
	NativeTimeStamp string
	EventType       string
	Data            string `sql:"type:jsonb"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Response struct {
	Result string
}

func (e *Event) PublishEvent(Channel string) error {

	PubSub.Pub(e, Channel)

	return nil

}
