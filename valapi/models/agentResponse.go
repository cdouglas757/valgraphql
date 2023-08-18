package valapi

type AgentResponse struct {
	Status int     `json:"status"`
	Data   []Agent `json:"data"`
}

type Agent struct {
	Uuid string `json:"uuid"`
	Name string `json:"displayName"`
	Role Role   `json:"role"`
}

type Role struct {
	Name string `json:"displayName"`
}
