package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)



func Get() {
	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := redisClient.Set(ctx, "test", "value1", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisClient.Get(ctx, "test").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := redisClient.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}