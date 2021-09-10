package main

func CreateService(msg string) SomeInterface {
	return NewSomeService(msg)
}
