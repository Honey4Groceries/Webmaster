package main

import (
    "fmt"
    "go-redis/redis"
)

func main() {
    client := redis.NewClient(&redis.Options{
        Addr:       "localhost:6379",
        Password:   "",
        DB:         0
    })


}
