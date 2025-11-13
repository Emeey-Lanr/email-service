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
	Subject string `json:"subject"`
	Html_body string `json:"html_body"`
	Text_body string `json:"text_body"`

}



