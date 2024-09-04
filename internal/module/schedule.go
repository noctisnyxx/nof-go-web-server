package module

import "strconv"

type Switch struct {
	Username  string `json:"user_name"`
	Device    string `json:"device"`
	DeviceId  int    `json:"device_id"`
	CommandId int    `json:"command_id"`
}

type PowerMeterCond struct {
	I_A float64 `json:"i_a"`
	V_A float64 `json:"v_a"`
	KWA float64 `json:"kw_a"`
}

type ScheduleData struct {
	ScheduleId  string          `json:"schedule_id"`
	Group       string          `json:"group"`
	TestMode    string          `json:"test_mode"`
	Title       string          `json:"title"`
	Start       string          `json:"start"`
	Stop        string          `json:"stop"`
	Description string          `json:"description"`
	Switcher    *Switch         `json:"swtchrcond"`
	PowerMeter  *PowerMeterCond `json:"pmcond"`
}

var scheduleIdNum int = 0

func (scheduleData ScheduleData) GenerateScheduleId() string {
	scheduleIdNum++
	return strconv.Itoa(scheduleIdNum)
}
