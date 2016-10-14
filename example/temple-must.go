package main

import (
	"fmt"
	"log"
	"os"

	temple "../"
)

func line() {
	fmt.Println("-------------------------------------------------------------------------------")
}

func main() {
	log.Println("Started")
	defer log.Println("Finished")

	line()
	tmpl, err := temple.NewTemple("templates", "base.html", false)
	log.Printf("tmpl=%#v, %#v\n", tmpl, err)
	line()

	trials := []string{"index.html", "about.html", "page.html", "post.html", "index.html"}

	for _, name := range trials {
		this := tmpl.MustGet(name)
		fmt.Printf("--- Template ---\n\n")
		fmt.Printf("this: %s, %#v, %#v\n\n", name, this)
		fmt.Printf("--- Output ---\n\n")
		this.Execute(os.Stdout, nil)
		line()
	}

	// and one that doesn't exist
	that := tmpl.MustGet("blah.html")
	log.Printf("that: %#v, %#v\n", that)
	that.Execute(os.Stdout, nil)
	line()
}
