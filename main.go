package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
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
	p := client.Ping()
	fmt.Println("INFO: Checking Redis availability. Ping <-> " + p.Val())
	if p.Err() != nil {
		log.Fatal(p.Err())
		os.Exit(1)
	}
	for {
		time.Sleep(2 * time.Second)
		writeAndRead(client)
	}
}

func writeAndRead(c *redis.Client) {
	randomize := rand.New(rand.NewSource(time.Now().UnixNano()))
	strconv.Itoa(randomize.Intn(99))
	
	json, err := json.Marshal(Author{Name: "User-"+strconv.Itoa(randomize.Intn(99)), Age: randomize.Intn(99)})
	if err != nil {
		fmt.Println(err)
	}
	id := strconv.Itoa(randomize.Intn(999999))
	err = c.Set(id, json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("INFO: Author successfully added")
	val, err := c.Get(id).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("INFO: Value from Redis: " + val)
}
