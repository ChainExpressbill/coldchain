package account

type LoginParams struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type JoinParams struct {
	Id           string `json:"id,omitempty"`
	Password     string `json:"password,omitempty"`
	Name         string `json:"name,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
}
