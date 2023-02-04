package main

import (
	"encoding/json"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v7"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "123456",
})

func main() {
	m, err := GetModels(1)
	fmt.Printf("model=%+v, err=%v", m, err)
}

type model struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetModels(id int) (model, error) {
	key := "model111111"
	cache, err := client.Get(key).Bytes()
	var res model
	if err != nil {
		fmt.Println("key not exists")
		res, err = GetFromDB(id)
		if err != nil {
			return res, err
		}
		value, err := json.Marshal(res)
		if err != nil {
			return res, err
		}
		client.Set(key, value, 2*time.Minute)
	} else {
		json.Unmarshal(cache, &res)
	}
	return res, nil
}

func GetFromDB(id int) (res model, err error) {
	res = model{1, "13sai"}
	return
}
