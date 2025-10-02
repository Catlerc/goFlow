package main

import (
	"awesomeProject/Flow"
	"strconv"
)

func toStr(i int, out chan string) {
	out <- strconv.Itoa(i)
}

func toStr2(i int, out chan string) {
	out <- strconv.Itoa(i) + "!!"
}

func printer(str string) {
	println(str)
}

func printer2(str string) {
	println("_" + str)
}

func main() {
	input := make(chan int)

	start := Flow.Start(input, nil)
	flow := Flow.To(start, toStr)
	flow2 := Flow.To(start, toStr2)
	join := Flow.Join[string](flow, flow2)

	end := Flow.End(join, printer)

    //// TEST ////
	
	go func() {
		for i := 0; i < 999999; i++ {
			input <- i
		}
		close(input)
	}()

	end.Close()
}
