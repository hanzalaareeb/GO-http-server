package main

import "github.com/codegangsta/martini"

func main() {
	// creating first server
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
