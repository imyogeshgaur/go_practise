package main

import "fmt"

type animal interface {
	eat()
	sleep()
}

type monkey struct {
	jump func()
	bite func()
}

type Human struct {
	nameOfHuman string
}

func jump(h Human) {
	fmt.Printf("%v is Jumping !!!\n",h.nameOfHuman);
}

func bite(h Human) {
	fmt.Printf("%v is Biting !!!\n",h.nameOfHuman);
}

func eat(h Human) {
	fmt.Printf("%v is Eating !!!\n",h.nameOfHuman);
}

func sleep(h Human) {
	fmt.Printf("%v is Sleeping !!!",h.nameOfHuman);
}

func main() {
	h := Human{"Yogesh"};

	jump(h);
	bite(h);
	eat(h);
	sleep(h);

}