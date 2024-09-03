package structs

type HttpResp struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}
