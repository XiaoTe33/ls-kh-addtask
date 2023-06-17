package router

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"ls-kh-addtask/dao"
	"ls-kh-addtask/log"
	"ls-kh-addtask/model"
	"strings"
)

var (
	myLog = log.Log
	db    = dao.DB
)
var redisCli = redis.NewClient(&redis.Options{
	Addr: viper.GetString("Redis.Addr"),
})

func InitRouters() {
	r := gin.Default()
	r.Use(Cors())
	r.POST("/addTask", JWT(), func(c *gin.Context) {
		username := c.GetString("username")
		userId := c.GetString("userId")
		task := c.PostForm("task")
		topic := c.PostForm("topic")
		objects := c.PostForm("objects")
		mq := model.NewListMQ(redisCli)
		mq.SendTask(model.NewTask(username, topic, userId, task, strings.Split(objects, ",")))
	})
	_ = r.Run(":" + viper.GetString("Router.Port"))
}
