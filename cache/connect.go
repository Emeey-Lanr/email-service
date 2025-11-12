package cache

import (
	"github.com/redis/go-redis/v9"
	"os"

)



func ConnectToRedis ( )(*redis.Client){
	rds := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
		Password: "",
		DB:0,
	})
 
  return rds
	
}