package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
)

var cnt = 0

func main() {
	hostname := []byte(GetHostName())
	r := gin.New()
	r.GET("/", func (c *gin.Context) {
		fmt.Println(cnt)
		cnt++
		c.Data(200, "plaintext", hostname)
	})
	r.Run(":8080")
}

func GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "default"
	}
	return hostname
}
