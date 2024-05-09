package entity

type Mail struct {
	Receiver string `json:"receiver,omitempty"`
	Subject  string `json:"subject,omitempty"`
	Message  string `json:"message,omitempty"`
}
