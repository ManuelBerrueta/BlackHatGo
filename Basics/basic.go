package maing

import "fmt"

func strlen(s string, c chan int) {
	c <- len(s)
}

func main() {

	c := make(chan int)
	go strlen("Salutations", c)
	go strlen("World", c)
	x, y := <-c, <-c //Channels look to be a stack type data structure LIFO!
	fmt.Println(x, y, x+y)
}
