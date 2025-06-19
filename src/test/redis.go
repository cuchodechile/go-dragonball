package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{Addr: "192.168.16.30:6379"})
	defer rdb.Close()

	var cursor uint64
	var keys []string
	for {
		var err error
		keys, cursor, err = rdb.Scan(ctx, cursor, "*", 1000).Result()
		if err != nil {
			log.Fatalf("scan: %v", err)
		}

		for _, k := range keys {
			fmt.Println(k)
		}
		if cursor == 0 { // fin de la iteración
			break
		}
	}
}