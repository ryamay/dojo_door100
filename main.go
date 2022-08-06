package main

import (
	"door100/door100"
	"fmt"
	"strconv"
)

func main() {
	doors := door100.InitDoors()
	for i := 1; i <= 100; i++ {
		fmt.Print("===================Phase" + strconv.Itoa(i) + "===================\n")
		doors.VisitInPhase(i)
		fmt.Print(doors.ShowClosed())
		fmt.Print(doors.ShowOpen())
	}
}
