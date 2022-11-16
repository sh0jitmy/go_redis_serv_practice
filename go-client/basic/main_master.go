package main

import (
    "context"
    "github.com/go-redis/redis/v8"
    "fmt"
    "time"
    "encoding/json"
)

var ctx = context.Background()

type Message struct {
	ID string `json:"ID"` 
	Code string `json:"Code"` 
	Payload string `json:"Payload"` 
}


func NewClient() (*redis.Client){
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
   return rdb
}

func  Set(rdb *redis.Client,key string,value string) {
    err := rdb.RPush(ctx, key, value).Err()
    if err != nil {
        panic(err)
    }
    return 
}

func  Get(rdb *redis.Client,key string) ([]string) {
    value,err := rdb.BLPop(ctx, 5*time.Second,key).Result()
    if err != nil {
        panic(err)
    }
    return value
}

func  Len(rdb *redis.Client,key string) (int64) {
    
    str,err := rdb.LRange(ctx,key,0,-1).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("list",str) 


    value,err := rdb.LLen(ctx,key).Result()
    if err != nil {
        panic(err)
    }
    return value
    
}

func Delete(rdb *redis.Client,key string) (int64) {
    value,err := rdb.Del(ctx,key).Result()
    if err != nil {
        panic(err)
    }
    return value
}

func Publish(rdb *redis.Client,key string, value string) (error) {
    //var ctx = context.Background()
    rderr := redisClient.Publish(ctx,key,value).Err()
    if rderr != nil {
	return rderr
    }
    return nil
}

func Subscribe(rdb *redis.Client,key string, value string) (error) {
    //var ctx = context.Background()
    sub := redisClient.Subscribe(ctx,key)
    _,err := sub.Receive(ctx)
    if err != nil {
        return err
    }
   ch := sub.Channel()

   for msg : =  range ch {
      fmt.Println(msg.Channel,msg.Payload)
   }
   return nil
}



func main() {
	client :=  NewClient()
	Delete(client,"Message")
	length  := Len(client,"Message")
	fmt.Println("length",length)
	
	msgdata := Message{ID:"1",Code:"0xFFFF",Payload:"Payload1"}
	jsonbin,_ := json.Marshal(msgdata)
	Set(client,"Message",string(jsonbin))
	fmt.Println("put1",string(jsonbin))
	
	length  = Len(client,"Message")
	fmt.Println("length",length)

	msgdata = Message{ID:"2",Code:"0xFCFF",Payload:"Payload2"}
	jsonbin,_ = json.Marshal(msgdata)
	Set(client,"Message",string(jsonbin))
	fmt.Println("put2",string(jsonbin))
	length  = Len(client,"Message")
	fmt.Println("length",length)


	getdata := Get(client,"Message")
	fmt.Println("getdata1",getdata)
	getdata = Get(client,"Message")
	fmt.Println("getdata2",getdata)

}
