package observer

import (
	"fmt"
	"testing"
)

// 观察者模式

// 观察者
type IObserver interface {
	Update(msg string)
}

type ISubject interface {
	Register(obs ISubject)
	Remove(obs ISubject)
	Notify(msg string)
}

type Subject struct {
	obss []IObserver
}

func (s *Subject) Register(obs IObserver) {
	s.obss = append(s.obss, obs)
}

func (s *Subject) Remove(obs IObserver) {
	for i, ob := range s.obss {
		if ob == obs {
			s.obss = append(s.obss[:i], s.obss[i+1:]...)
		}
	}
}

func (s *Subject) Notify(msg string) {
	for _, o := range s.obss {
		o.Update(msg)
	}
}

type Obs1 struct{}

func (o Obs1) Update(msg string) {
	fmt.Printf("obs1: %s", msg)
}

type Obs2 struct{}

func (o Obs2) Update(msg string) {
	fmt.Printf("obs2: %s", msg)
}

func TestObs(t *testing.T) {
	sub := &Subject{}
	sub.Register(&Obs1{})
	sub.Register(&Obs2{})
	sub.Notify("hi")
}
