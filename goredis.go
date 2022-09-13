package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func GoRedis() {
	fmt.Println("############## Go - Redis ##############")
	ctx := context.TODO()
	connectRedis(ctx)
	clearAllFromRedis(ctx)

	post1 := make(map[string]interface{})
	post1["category"] = "Deporte"
	post1["post_id"] = "635216"
	post1["insert_at"] = "2022-09-09"
	post2 := make(map[string]interface{})
	post2["category"] = "Deporte"
	post2["post_id"] = "435216"
	post2["insert_at"] = "2022-09-08"
	post3 := make(map[string]interface{})
	post3["category"] = "Baile"
	post3["post_id"] = "735216"
	post3["insert_at"] = "2022-09-09"
	post4 := make(map[string]interface{})
	post4["category"] = "Baile"
	post4["post_id"] = "535216"
	post4["insert_at"] = "2022-09-07"

	//Add 4 posts
	setHashToRedis(ctx, "Post:1", post1)
	setHashToRedis(ctx, "Post:2", post2)
	setHashToRedis(ctx, "Post:3", post3)
	setHashToRedis(ctx, "Post:4", post4)

	//Maintain insert_at index
	setSetToRedis(ctx, "insert_at:2022-09-09", []string{"1", "3"})
	setSetToRedis(ctx, "insert_at:2022-09-08", []string{"2"})
	setSetToRedis(ctx, "insert_at:2022-09-07", []string{"4"})

	val1 := getSetInterFromRedis(ctx, "insert_at:2022-09-09")
	val2 := getSetInterFromRedis(ctx, "insert_at:2022-09-08")
	val3 := getSetInterFromRedis(ctx, "insert_at:2022-09-07")

	fmt.Println("-------------------")
	fmt.Println("Merge")
	fmt.Println("-------------------")
	for _, v := range val1 {
		resMap := getAllHashFromRedis(ctx, "Post:"+v)
		for k, m := range resMap {
			fmt.Println(k, ":", m)
		}
	}
	fmt.Println("-------------------")
	for _, v := range val2 {
		resMap := getAllHashFromRedis(ctx, "Post:"+v)
		for k, m := range resMap {

			fmt.Println(k, ":", m)
		}
	}
	fmt.Println("-------------------")
	for _, v := range val3 {
		resMap := getAllHashFromRedis(ctx, "Post:"+v)
		for k, m := range resMap {
			fmt.Println(k, ":", m)
		}
	}
}

func connectRedis(ctx context.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)

	redisClient = client
}

func setHashToRedis(ctx context.Context, key string, val map[string]interface{}) {
	err := redisClient.HSet(ctx, key, val).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func setSetToRedis(ctx context.Context, key string, val []string) {
	err := redisClient.SAdd(ctx, key, val).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func getSetInterFromRedis(ctx context.Context, keys ...string) []string {
	_keys := ""
	for i, v := range keys {
		if i < len(keys)-1 {
			_keys += v + ","
		} else if i == len(keys)-1 {
			_keys += v
		}
	}
	val, err := redisClient.SInter(ctx, _keys).Result()
	if err != nil {
		fmt.Println(err)
	}

	return val
}

func getAllHashFromRedis(ctx context.Context, key string) map[string]string {
	val, err := redisClient.HGetAll(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return val
}

func clearAllFromRedis(ctx context.Context) {
	redisClient.FlushAll(ctx)
}
