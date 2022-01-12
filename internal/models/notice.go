package models

import "time"

type Notice struct {
	User      *User
	NoticeSet *NoticeSet
	Deadline  time.Time
}

type NoticeSet struct {
	Actions []*Action
}
