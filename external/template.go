package external

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/emeey-lanr/email_service/model"
)


func GetTemplate (template_code string) ( model.EmailTemplate, error){
	var responseData model.EmailTemplate

	request, err := http.NewRequest("GET",  fmt.Sprintf(`%s/%s`,os.Getenv("TEMPLATE_SERVICE_URL"), template_code), nil)

	if err != nil {
	  return responseData, err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	 defer response.Body.Close()

	 

	 if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil{
		return responseData, err
	 }

   return responseData, nil

}