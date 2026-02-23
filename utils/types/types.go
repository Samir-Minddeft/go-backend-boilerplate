package types

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}
