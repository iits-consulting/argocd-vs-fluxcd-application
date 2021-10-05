package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"encoding/json"
	"github.com/go-redis/redis"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Redis struct {
	Url string
}

func main() {
	var r Redis
	r.Url, _ = os.LookupEnv("REDIS_URL")

	client := redis.NewClient(&redis.Options{
		Addr:     r.Url,
		Password: "",
		DB:       0,
	})
	fmt.Println("INFO: Checking Redis availability. Ping <-> " + client.Ping().Val())
	randomize := rand.New(rand.NewSource(time.Now().UnixNano()))
	json, err := json.Marshal(Author{Name: "Elliot", Age: randomize.Intn(99)})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("INFO: Author successfully added")
	val, err := client.Get("id1234").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("INFO: Value from Redis: " + val)
}
