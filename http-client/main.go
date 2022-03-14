package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	addr = GetString("ADDR", "localhost:8888")
)

func main() {
	_duration, err := strconv.ParseInt(GetString("DURATION", "1"), 10, 64)
	if err != nil {
		log.Fatalf("Parse Error %v", err)
	}
	duration := time.Duration(_duration)
	for i := 0;; i++ {
		resp, err := http.Get(addr)
		if err != nil {
			fmt.Printf("Http Request Err : [%v]\n", err)
		} else {
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Read Body Err : [%v]\n", err)
			}
			fmt.Println(i, ":", string(data))
		}
		time.Sleep(duration * time.Second)
	}
}

func GetString(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return val
}
