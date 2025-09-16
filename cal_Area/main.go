package main

import (
	"github.com/zoklk/Go-prac/cal_Area/Area"
)

func main() {
	var list []area.Shape
	list = append(list, &area.Rectangle{Width: 5, Height: 3})
	list = append(list, &area.Rectangle{Width: 7, Height: 8})
	list = append(list, &area.Circle{Radius: 3})


	for _, shape := range list {
	area.PrintArea(shape)
    }
}