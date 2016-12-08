package redis

import (
	"testing"
	"gopkg.in/redis.v5"
	"fmt"
	"os"
	"time"
)

var client *redis.Client

func TestMain(m *testing.M) {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pong) // PONG
	os.Exit(m.Run())
}

func TestRedis0(t *testing.T) {
	r, err := client.Set("key1", true, time.Second).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("set key1 result: ", r) // set key1 result:  OK

	r, err = client.Get("key1").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("get key1 result: ", r) // get key1 result:  1

	fmt.Println("sleep 1 second")
	time.Sleep(time.Second)
	r, err = client.Get("key1").Result()
	if err != nil {
		fmt.Println(err) // redis: nil
		return
	}
	fmt.Println("get key1 result: ", r)
}
