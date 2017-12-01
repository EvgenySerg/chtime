package chtime

import (
	"time"
	"errors"
)

type Ch chan interface{}

func NewCh(cap int) Ch {
	return make(chan interface{}, cap)
}

//Receive - receives data from a channel with timeout variable in seconds
func (ch Ch) Receive(timeout int) (interface{}, error){
	if timeout<=0{
		return nil,errors.New("timeout variable cannot be below or equals zero")
	}
	select {
	case b:=<- ch:
		return b, nil
	case <-time.After(time.Duration(timeout)*time.Second):
		return nil, errors.New("timeout data reception from channel")
	}
}

//Send - sends data to a channel using timeout variable
func (ch Ch) Send(value interface{}, timeout int) error{
	if timeout<=0{
		return errors.New("timeout variable cannot be below or equals zero")
	}
	select {
	case ch<-value:
		return errors.New("sending data to channel timeout")
	case <-time.After(time.Duration(timeout)*time.Second):
		return nil
	}
}

