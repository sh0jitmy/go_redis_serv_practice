package main

import (
    "context"
    "github.com/go-redis/redis/v8"
    "fmt"
//    "time"
//    "encoding/json"
)

var ctx = context.Background()

type Message struct {
	ID string `json:"ID"` 
	Code string `json:"Code"` 
	Payload string `json:"Payload"` 
}


func NewClient() (*redis.Client){
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6380",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
   return rdb
}


func Publish(rdb *redis.Client,key string, value string) (error) {
    //var ctx = context.Background()
    rderr := rdb.Publish(ctx,key,value).Err()
    if rderr != nil {
	return rderr
    }
    return nil
}

func Subscribe(rdb *redis.Client,key string) (string,error) {
    //var ctx = context.Background()
    var data string =""
    sub := rdb.Subscribe(ctx,key)
    _,err := sub.Receive(ctx)
    if err != nil {
        return "",err
    }
   ch := sub.Channel()

   for msg :=  range ch {
      //fmt.Println(msg.Channel,msg.Payload)
      data = msg.Payload
      break
   }
   //close(ch)
   return data,nil
}



func main() {
	client :=  NewClient()
	//msgdata := Message{ID:"1",Code:"0xFFFF",Payload:"Payload1"}
	data ,_ := Subscribe(client,"Message")
        fmt.Println(data)
}
