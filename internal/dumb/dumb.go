package dumb

import (
	"fmt"
	"os"
	"os/exec"
)

func Dumb(db string, out string, uri string, gzip bool) {
	if db == "" {
		fmt.Println("Error: db is required")
		os.Exit(1)
	}

	args := []string{}

	if uri != "" {
		args = append(args, "--uri", uri)
	} else {
		args = append(args, "--db", db)
	}

	if gzip {
		args = append(args, "--gzip")
		args = append(args, "--archive="+out+".gz")
	} else {
		args = append(args, "--out", out)
	}

	cmd := exec.Command("mongodump", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Running mongodump...")

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Backup saved to:", out)

}
