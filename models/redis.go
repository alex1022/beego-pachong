package models

import (
	"github.com/gomodule/redigo/redis"
	"github.com/monnand/goredis"
)

var (
	client goredis.Client
	con    redis.Conn
)

const (
	URL_QUEUE     = "url_queue"     //list名称，搞个const常量，表示严肃一点
	URL_VISIT_SET = "url_visit_set" //set名称，已经访问过的set
)

func ConnectRedis(addr string) {
	client.Addr = addr

}

//添加到队列Lpush，从左边插入一个,复杂度O(1)
func PutinSuperQueue(url string) {
	err := client.Lpush(URL_QUEUE, []byte(url)) //数组类型的url 传入队列
	if err != nil {
		panic(err)
	}
}

//删除一个list元素，最右边删除，rpop，并且返回这个删除的key，复杂度O(1)
func PopfromQueue() string {
	res, err := client.Rpop(URL_QUEUE)
	if err != nil {
		panic(err)
	}
	return string(res)
}

//设置一个集合，存放已经处理过的 SuperUrl
func AddtoSet(SuperUrl string) {
	client.Sadd(URL_VISIT_SET, []byte(SuperUrl))
}

//判断队列长度
func GetQueueLength() int {
	length, err := client.Llen(URL_QUEUE)
	if err != nil {
		panic(err)
	}
	return length
}

//判断url是否已存在
func IsVisit(url string) bool {
	bIsVisit, err := client.Sismember(URL_VISIT_SET, []byte(url))
	if err != nil {
		return false
	}
	return bIsVisit
}
