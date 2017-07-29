package structs

import "time"

type User struct {
	Name         string `storm:"id"`
	MessageCount int
	LastJoin     *time.Time
	LastPart     *time.Time
	LastActive   *time.Time
	FirstSeen    *time.Time
}
