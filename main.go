package main

import (
	//"fmt"
	color "github.com/fatih/color"
	greeting "github.com/Sazikoff/hexlet-go/greeting"
)

func main() {
    red := color.New(color.FgRed)

boldRed := red.Add(color.BgWhite)
boldRed.Println(greeting.Hello())

}
