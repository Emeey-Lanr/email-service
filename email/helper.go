package email

import (
	"encoding/json"
	"strings"

	"github.com/emeey-lanr/email_service/cache"
	"github.com/emeey-lanr/email_service/model"
	"github.com/redis/go-redis/v9"
)

func Decode_template (name, template_code string, rds *redis.Client)(string, string){
  
	var cached model.EmailTemplate

	templateData, _ := cache.CacheEmailTemplate(template_code, rds)

    json.Unmarshal([]byte(templateData), &cached)

   body := cached.Data.Body
   
   subject := strings.Replace(cached.Data.Subject, "{{app_name}}", "Group44", -1)
  bodyMessage := strings.Replace(body, "{{name}}", name, -1)

  return subject, bodyMessage
}