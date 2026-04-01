package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"
	"os"
)

var ctx = context.Background()

func main() {
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "redis"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: redisHost + ":6379",
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		val, err := rdb.Incr(ctx, "counter").Result()
		if err != nil {
			fmt.Fprintf(w, "Hello! Redis error")
			return
		}

		fmt.Fprintf(w, "Hello! Visits: %d\n", val)
	})

	http.ListenAndServe(":8080", nil)
}
