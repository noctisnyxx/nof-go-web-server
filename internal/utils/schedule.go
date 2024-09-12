package utils

import (
	"time"
)

type Switch struct {
	Username  string `json:"user_name" bson:"user_name"`
	Device    string `json:"device" bson:"device"`
	DeviceId  int    `json:"device_id" bson:"device_id"`
	CommandId int    `json:"command_id" bson:"command_id"`
}

type PowerMeterCond struct {
	I_A float64 `json:"i_a" bson:"i_a"`
	V_A float64 `json:"v_a" bson:"v_a"`
	KWA float64 `json:"kw_a" bson:"kw_a"`
}

type ScheduleData struct {
	CreatedAt   time.Time       `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" bson:"updated_at"`
	ScheduleId  string          `json:"schedule_id" bson:"schedule_id"`
	Group       string          `json:"group" bson:"group"`
	TestMode    string          `json:"test_mode" bson:"test_mode"`
	Title       string          `json:"title" bson:"title"`
	Status      string          `json:"status" bson:"status"`
	Start       time.Time       `json:"start" bson:"start"`
	End         time.Time       `json:"end" bson:"end"`
	Description string          `json:"description" bson:"description"`
	Switcher    *Switch         `json:"swtchrcond" bson:"swtchrcond"`
	PowerMeter  *PowerMeterCond `json:"pmcond" bson:"pmcond"`
}
