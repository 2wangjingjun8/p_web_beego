package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
)

// RedisDemoController is a struct
type RedisDemoController struct {
	beego.Controller
}

// Get redis
func (c *RedisDemoController) Get() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		beego.Info("redis 连接失败！", err)
	}
	defer conn.Close()

	// conn.Send("set", "name", "xiaowang")
	// conn.Send("mset", "age", 20, "sex", "man")
	// conn.Flush()
	// reply, err := conn.Receive()

	// reply, err := conn.Do("mset", "name", "xiaohong", "age", 21, "sex", "woman")

	// conn.Send("Multi")
	// conn.Send("lpush", "list", "xiaoxiao", "maomao", "dada", "xiaowang")
	// conn.Send("get", "name")
	// conn.Send("mset", "name", "hehe", "age", 20)
	// reply, err := conn.Do("Exec")

	// reply, err := redis.String(conn.Do("get", "name"))

	reply, err := redis.Values(conn.Do("mget", "name", "age"))

	if err != nil {
		c.Ctx.WriteString("获取内容失败")
	}
	beego.Info(reply)
	fmt.Printf("reply 的类型是%T", reply)
	fmt.Println()
	fmt.Println(reply[0], reply[1])
	var s string
	var i int
	redis.Scan(reply, &s, &i)
	beego.Info(s, i)

}
