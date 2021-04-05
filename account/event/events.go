package event

type AccountCreated struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       uint   `json:"age"`
	Created   string `json:"created"`
}

type AccountClosed struct {
	UserID  string `json:"user_id"`
	Created string `json:"created"`
}
