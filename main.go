package main

import (
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func main() {
	Piramid() //excercise 1
	GoRedis() //excercise 2
}
