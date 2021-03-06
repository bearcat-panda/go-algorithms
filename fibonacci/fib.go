package fibonacci

import (
	"fmt"
	"strings"
	"io"
	"bufio"
)

func dup3(in <-chan int) (<-chan int, <-chan int, <-chan int) {
	a, b, c := make(chan int, 2), make(chan int, 2), make(chan int, 2)
	go func() {
		for {
			x := <-in
			a <- x
			b <- x
			c <- x
		}
	}()
	return a, b, c
}

func fibonacci1() <-chan int {
	x := make(chan int, 2)
	a, b, out := dup3(x)
	go func() {
		x <- 0
		x <- 1
		<-a
		for {
			x <- <-a + <-b
		}
	}()
	return out
}

// 1,1,2,3,5,8,13...
//     a,b
//       a,b
func fibonacci2() func() int {
	x1, x2 := 0, 1
	return func() int {
		// 重新赋值
		x1, x2 = x2, (x1 + x2)
		return x1
	}
}

func fibonacci3() fibGen {
	x1, x2 := 0, 1
	return func() int {
		// 重新赋值
		x1, x2 = x2, (x1 + x2)
		return x1
	}
}

type fibGen func() int

// 函数实现接口
func (g fibGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}