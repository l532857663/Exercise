package main

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

// MQ is message queue listener handle
type MQ struct {
	context   *zmq.Context
	socket    *zmq.Socket
	isRunning bool
	finished  chan error
	binding   string
}

type Zmq struct {
	Address string
}

const (
	// Topic
	// TOPIC_TRANSFER = "transfer"
	TOPIC_TRANSFER = "ethereum"

	// Message
	MESSAGE_TOPIC_INDEX   = 0
	MESSAGE_CONTENT_INDEX = 1
)

func (mq *MQ) Shutdown() {
	if mq.isRunning {
		go func() {
			// if errors in the closing sequence, let it close ungracefully
			if err := mq.socket.SetUnsubscribe(TOPIC_TRANSFER); err != nil {
				mq.finished <- err
				return
			}
			if err := mq.socket.Unbind(mq.binding); err != nil {
				mq.finished <- err
				return
			}
			if err := mq.socket.Close(); err != nil {
				mq.finished <- err
				return
			}
			if err := mq.context.Term(); err != nil {
				mq.finished <- err
				return
			}
		}()

		var err error
		select {
		// case <-ctx.Done():
		// 	err = ctx.Err()
		case err = <-mq.finished:
		}
		if err != nil {
			fmt.Errorf("Shutdown MQ server error %v, %v\n", err, zmq.AsErrno(err))
			return
		}
	}

	fmt.Printf("MQ server shutdown finished\n")
	return
}
func New(conf *Zmq) *MQ {
	context, err := zmq.NewContext()
	if err != nil {
		fmt.Errorf("Create new MQ context error: %+v", err)
		panic(err)
	}
	socket, err := context.NewSocket(zmq.SUB)
	if err != nil {
		fmt.Errorf("Create new MQ socket error: %+v", err)
		panic(err)
	}

	err = socket.SetSubscribe(TOPIC_TRANSFER)
	if err != nil {
		fmt.Errorf("Set MQ topic error: %+v", err)
		panic(err)
	}
	// 订阅其他topic
	// err = socket.SetSubscribe("topic n")
	// if err != nil {
	// 	fmt.Errorf("set MQ topic error: %+v", err)
	// 	panic(err)
	// }

	// socket.SetLinger(0)
	// for now do not use raw subscriptions - we would have to handle skipped/lost notifications from zeromq
	// on each notification we do sync or syncmempool respectively
	err = socket.Bind(conf.Address)
	if err != nil {
		fmt.Errorf("Create new MQ server error: %+v", err)
		panic(err)
	}
	fmt.Println("Starting MQ server on: ", conf.Address)

	mq := &MQ{context, socket, true, make(chan error), conf.Address}
	go mq.run()

	return mq
}

func (mq *MQ) run() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("MQ loop recovered from %v\n", r)
		}
		mq.isRunning = false
		fmt.Println("MQ loop terminated")
		mq.finished <- nil
	}()
	mq.isRunning = true
	repeatedError := false
	for {
		msg, err := mq.socket.RecvMessage(0)
		if err != nil {
			if zmq.AsErrno(err) == zmq.Errno(zmq.ETERM) || err.Error() == "Socket is closed" {
				fmt.Println("Recv from zmq error: ", err)
				break
			}
			// suppress logging of error for the first time
			// programs built with Go 1.14 will receive more signals
			// the error should be resolved by retrying the call
			// see https://golang.org/doc/go1.14#runtime
			if repeatedError {
				fmt.Errorf("MQ RecvMessageBytes error %v, %v", err, zmq.AsErrno(err))
			}
			repeatedError = true
			time.Sleep(100 * time.Millisecond)
		}
		if msg != nil && len(msg) >= 3 {
			switch msg[MESSAGE_TOPIC_INDEX] {
			case TOPIC_TRANSFER:
				fmt.Printf("receive msg[%s] from topic[%s]\n", msg[MESSAGE_CONTENT_INDEX], TOPIC_TRANSFER)
				break
			default:
				fmt.Printf("MQ: Unknown topic %s\n", msg[MESSAGE_TOPIC_INDEX])
			}
			// TODO: 封装消息类型对应的dao层功能函数
			// callback(nt)
		}
	}

	fmt.Println("MQ server exit!")
}

func main() {
	conf := &Zmq{
		Address: "tcp://127.0.0.1:5557",
	}
	mq := New(conf)
	defer mq.Shutdown()

	context, err := zmq.NewContext()
	if err != nil {
		fmt.Errorf("Create new mq context error: %+v\n", err)
		panic(err)
	}
	defer context.Term()

	socket, err := context.NewSocket(zmq.PUB)
	if err != nil {
		fmt.Errorf("Create new mq socket error: %+v\n", err)
		panic(err)
	}

	err = socket.Connect(conf.Address)
	if err != nil {
		fmt.Errorf("Create new mq connection error: %+v\n", err)
		panic(err)
	}
	defer socket.Close()
	fmt.Printf("MQ listening to %v\n", conf.Address)

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("MQ sending to %v\n", conf.Address)
	_, err = socket.SendMessage(TOPIC_TRANSFER, "test message ...", 0)
	if err != nil {
		fmt.Errorf("Send message to zmq error: %+v\n", err)
		return
	}
	fmt.Printf("wch------ test1\n")

	// 未订阅的主题无法被正常接收
	_, err = socket.SendMessage("", "test message null", 0)
	if err != nil {
		fmt.Errorf("Send message to zmq error: %+v\n", err)
		return
	}
	fmt.Printf("wch------ test2\n")

	_, err = socket.SendMessage("other...", "test message other", 0)
	if err != nil {
		fmt.Errorf("Send message to zmq error: %+v\n", err)
		return
	}
	// 未订阅的主题测试结束
	fmt.Printf("wch------ test3\n")
	time.Sleep(30 * time.Second)

	return
}
