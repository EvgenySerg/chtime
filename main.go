package main

import (
	"time"
	"log"
	"ttt/chtime"
)

func main()  {
		cc:=chtime.NewCh(0)
		//cc = make(chan int)
		go Add("Оля + ", chtime.Ch(cc))
		time.Sleep(time.Millisecond*10)
		cc.Send("Вася",1)
		time.Sleep(1*time.Second)
}


// Add new value
func Add(a string, ch chtime.Ch){
	b,err:=ch.Receive(2)
	if err!=nil{
		log.Printf("Timeout b= %d",b)
		return
	}
	log.Printf(a+ b.(string))
}

