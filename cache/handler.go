package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/emeey-lanr/email_service/external"
	"github.com/redis/go-redis/v9"
)

func CacheEmailTemplate (template_code string, rds *redis.Client)(string, error) {
    

    ctx := context.Background()
	cachedData, err := rds.Get(ctx,  "welcome_email").Result()

	if err == redis.Nil{
		// Make a request to the template db if is not cached
		template, err := external.GetTemplate(template_code)
		if err != nil{
          return "", err
		}

	  tempJsonData, _ := json.Marshal(template)
	
	//   cache the response
		if err := rds.Set(ctx, "welcome_email", tempJsonData, 30*time.Minute); err != nil{
			  log.Println("Unable to cache template")

			return "", fmt.Errorf("Unable to cache template")
       
		}

		
	}else if err != nil{
		log.Println("Unable to get cached data")
		return "", err
	}
	
	
	return cachedData, nil
	


}
