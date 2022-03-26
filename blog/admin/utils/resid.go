package utils

import "github.com/garyburd/redigo/redis"

var Pool *redis.Pool  //全局连接池

func InitRedisPool(){
	conf,_ := ParseConfig("../config/admin.json")
	redisAddr := conf.RedisConfig.Addr
	redisPort := conf.RedisConfig.Port

	Pool = &redis.Pool{     //实例化一个连接池
		MaxIdle:16,    //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:0,    //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout:300,    //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn ,error){     //要连接的redis数据库
			return redis.Dial("tcp",redisAddr+":"+redisPort)
		},
	}
}
