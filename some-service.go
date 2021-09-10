package main

import "fmt"

type SomeService struct {
	msg string
}

func NewSomeService(msg string) *SomeService {
	return &SomeService{msg: msg}
}

func (s *SomeService) PrintSomething() {
	fmt.Println(s.msg)
}
