package main

import (
	"fmt"
	"math/rand"
	"time"

	"honnef.co/go/js/dom"
)

func main() {
	d := dom.GetWindow().Document()

	clickme := d.GetElementByID("clickme").(*dom.HTMLButtonElement)
	display := d.GetElementByID("display").(*dom.HTMLDivElement)

	clickme.AddEventListener("click", false, func(event dom.Event) {
		for n := 0; n < 5; n++ {
			go func() {
				n := n
				time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
				text := display.InnerHTML()
				text += fmt.Sprintf("%d<br>", n)
				display.SetInnerHTML(text)
			}()
		}
	})
}
