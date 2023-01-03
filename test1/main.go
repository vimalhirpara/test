package main

import "fmt"

type human interface {
	walk()
	talk()
}
type person struct {
	firstName string
	lastName  string
}

type superhuman interface {
	human
	fly()
}

func (p person) walk() {
	fmt.Printf("%s %s walks\n", p.firstName, p.lastName)
}
func (p person) talk() {
	fmt.Printf("%s %s talks!\n", p.firstName, p.lastName)
}

type superman struct {
	wearUWoutside bool
}

func (p superman) walk() {
	fmt.Println("superman walks")
}
func (p superman) talk() {
	fmt.Println("superman talks!")
}

func (s superman) fly() {
	fmt.Println("Superman is wearing UW outside: ", s.wearUWoutside)
	fmt.Println("Superman in flying ~~~~~~~~~")
}

type dog struct {
	name string
}

func (d dog) walk() {
	fmt.Printf("%s walks!\n", d.name)
}

func (d dog) talk() {
	fmt.Printf("%s talks!\n", d.name)
}

func walkAndTalk(h human) {
	h.talk()
	h.walk()
}
func doeverything(s superhuman) {
	s.fly()
	s.talk()
	s.walk()
}
func doit(i interface{}) {
	var str string
	str, ok := i.(string)
	if !ok {
		return
	}
	fmt.Println("Val received: ", str)
}

func main() {
	p := person{
		firstName: "souvik",
		lastName:  "haldar",
	}
	d := dog{
		name: "tom",
	}
	s := superman{
		wearUWoutside: true,
	}
	walkAndTalk(d)
	walkAndTalk(p)
	doit(5)
	doeverything(s)

}
