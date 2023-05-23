package model

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

// https://www.w3cschool.cn/redis/redis-yj3f2p0c.html
// redis 分布式锁实现
func init() {
	initRedis()
}

var redisPool *redis.Pool

func initRedis() {
	redisPool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   10,
		IdleTimeout: 5 * time.Second,
		Dial: func() (conn redis.Conn, e error) {
			conn, err := redis.Dial("tcp", "127.0.1.1:6379")
			if err != nil {
				return nil, err
			}
			conn.Do("select", 0)
			return conn, err
		},
	}
}

const LockName = "LockName"
const StringSetIfNotExist = "NX"
const StringSetWithExpireTime = "EX"

func getLock(id string) bool {
	// 获得锁
	// 一个命令原子操作  为了防止执行了nx 之后程序突然奔溃 则就会无法设置过期时间发生死锁
	conn := redisPool.Get()
	defer conn.Close()
	_, err := redis.String(conn.Do("set", LockName, id, StringSetIfNotExist, StringSetWithExpireTime, 10))
	if err == nil {
		return true
	}
	return false

}

const luaScript = `
if redis.call('get', KEYS[1])==ARGV[1] then
	return redis.call('del', KEYS[1])
else
	return 0
end
`

func releaseDistributedLock(id string) bool {
	// 释放分布式锁
	// 也必须是原子操作 借助于lua 脚本实现
	// 谁 上锁谁解锁
	conn := redisPool.Get()
	defer conn.Close()
	lua := redis.NewScript(1, luaScript)            // 定义参数的个数
	_, err := redis.Int(lua.Do(conn, LockName, id)) // 上面定义几个参数conn 后面几个都是参数， 参数的后面就是值 按照顺序
	if err == nil {
		return true
	}
	return false
}
func main() {
	theId := "yang"
	flag := getLock(theId)
	fmt.Println(flag)
	releaseFlag := releaseDistributedLock(theId)
	fmt.Println("releaseFlag: ", releaseFlag)
}
