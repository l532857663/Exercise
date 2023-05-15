package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"gorm.io/gorm"
)

func main() {

	var age *uint64
	v := uint64(0)
	age = &v
	fmt.Printf("wch---- :%+v\n", age)

	a := gorm.Model{
		ID: 1,
	}

	student := map[string]interface{}{
		"id":   "st01",
		"name": 0,
		"age":  age,
		"time": time.Now(),
		"dao":  a,
	}

	set("key1", student, 0)
	get("key1")

}

func set(key string, value map[string]interface{}, ttl int) bool {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	//ctx := context.Background()
	// err := client.HMSet(ctx, key, value).Err()
	err := client.HMSet(key, value).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func get(key string) bool {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	// ctx := context.Background()
	// val, err := client.HMGet(ctx, key, "id").Result()
	val, err := client.HMGet(key, "id").Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(val)
	return true
}
