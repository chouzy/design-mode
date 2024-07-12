package observer

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
)

// EventBus

type Bus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

// 异步事件总线
type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.Mutex
}

// new
func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: map[string][]reflect.Value{},
		lock:     sync.Mutex{},
	}
}

// 订阅
func (e *AsyncEventBus) Subscribe(topic string, f interface{}) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	v := reflect.ValueOf(f)
	if v.Type().Kind() != reflect.Func {
		return fmt.Errorf("hander is not a funcation")
	}

	handler, ok := e.handlers[topic]
	if !ok {
		handler = []reflect.Value{}
	}
	handler = append(handler, v)
	e.handlers[topic] = handler

	return nil
}

// 发布
// 这里异步执行，且不会等待返回结果
func (e *AsyncEventBus) Publish(topic string, args ...interface{}) {
	handlers, ok := e.handlers[topic]
	if !ok {
		fmt.Println("not found handlers in topic: ", topic)
		return
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	for i := range handlers {
		go handlers[i].Call(params)
	}
}

func sub1(msg1, msg2 string) {
	time.Sleep(time.Microsecond)
	fmt.Printf("sub1, %s %s\n", msg1, msg2)
}

func sub2(msg1, msg2 string) {
	fmt.Printf("sub2, %s %s\n", msg1, msg2)
}

func TestEvent(t *testing.T) {
	bus := NewAsyncEventBus()
	bus.Subscribe("topic:1", sub1)
	bus.Subscribe("topic:2", sub2)

	bus.Publish("topic:1", "test1", "test2")
	bus.Publish("topic:2", "testA", "testB")
	time.Sleep(time.Second)
}
