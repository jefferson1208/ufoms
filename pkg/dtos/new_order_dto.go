package dtos

import "time"

type NewOrderDto struct {
	OrderID      string
	ClordID      string
	TraceID      string
	OrderDate    time.Time
	ScheduleDate time.Time
	ExpireDate   time.Time
	Price        int64
	OrderType    uint8
	TimeInForce  uint8
	Memo         string
	Status       uint8
}
