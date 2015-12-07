package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func rateLimit(c *gin.Context) {

	ip := c.ClientIP()
	value := int(ips.Add(ip, 1))
	if value%50 == 0 {
		fmt.Printf("ip: %s, count: %d\n", ip, value)
	}
	if value >= 200 {
		if value%200 == 0 {
			fmt.Println("ip blocked")
		}
		c.Abort()
		c.String(503, "you were automatically banned :)")
	}
}

func index(c *gin.Context) {
	c.Redirect(301, "/room/hn")
}

func test(c *gin.Context) {
	c.HTML(200, "room_login.templ.html", gin.H{})
}

func roomGET(c *gin.Context) {
	roomid := c.Param("roomid")
	nick := c.Query("nick")
	if len(nick) < 2 {
		nick = ""
	}
	if len(nick) > 13 {
		nick = nick[0:12] + "..."
	}
	c.HTML(200, "index.html", gin.H{
		"roomid":    roomid,
		"nick":      nick,
		"timestamp": time.Now().Unix(),
	})

}
