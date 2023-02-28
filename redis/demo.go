package main

import (
	"context"
	"encoding"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var _ encoding.BinaryMarshaler = new(myStruct2)
var _ encoding.BinaryUnmarshaler = new(myStruct2)

type myStruct2 struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}

func (m *myStruct2) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

func (m *myStruct2) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)

}

func main() {
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.10.3:6379",
		Password: "",
		DB:       0,
	})

	data := &myStruct2{
		UserId:   "123",
		UserName: "anker",
	}
	err := rdb.Set(ctx, "key1", data, 0).Err()
	if err != nil {
		panic(err)
	}

	result := &myStruct2{}
	err = rdb.Get(ctx, "key1").Scan(result)
	if err != nil {
		panic(err)
	}

	fmt.Println("get success:", data.UserId == result.UserId)
}
