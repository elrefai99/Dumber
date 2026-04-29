package dumb

import (
	"fmt"
	"os"
)

func Dumb(db string, out string, uri string, gzip bool) {
	if db == "" {
		fmt.Println("Error: db is required")
		os.Exit(1)
	}
	fmt.Println("db", db)
	fmt.Println("out", out)
	fmt.Println("uri", uri)
	fmt.Println("gzip", gzip)

}
