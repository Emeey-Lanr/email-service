package model


type VAribLE struct {
	 Name string `json:"name"`
	 Message string `json:"message"`
}

type DaTA struct {
	Template_code string `json:"template_code"`
	Variable VAribLE `json:"variable"`
}
type QueueResponse struct {
	Correlation_id string `json:"correlation_id"` //for tracking
	Data DaTA `json:"data"` 
	Email string `json:"email"`

}



type EmailData struct {
	Id string `json:"id"`
	Code string `json:"code"`
	Type string `json:"type"`
	 Subject string `json:"subject"`
	 Body string `json:"body"`
	 Language string  `json:"language"`
}
type EmailTemplate struct {
	Success string `json:"success"`
	Data EmailData `json:"data"`
	Message string `json:"message"`
}
