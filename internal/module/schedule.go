package module

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
	ScheduleId  string          `json:"schedule_id" bson:"schedule_id"`
	Group       string          `json:"group" bson:"group"`
	TestMode    string          `json:"test_mode" bson:"test_mode"`
	Title       string          `json:"title" bson:"title"`
	Start       string          `json:"start" bson:"start"`
	Stop        string          `json:"stop" bson:"stop"`
	Description string          `json:"description" bson:"description"`
	Switcher    *Switch         `json:"swtchrcond" bson:"swtchrcond"`
	PowerMeter  *PowerMeterCond `json:"pmcond" bson:"pmcond"`
}
