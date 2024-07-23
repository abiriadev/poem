package main

import (
	"fmt"
	"os"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

const defaultContent = `---
title: "無際"
date: %s
---
`

func main() {
	id, err := gonanoid.New()
	if err != nil {
		panic(err)
	}

	f, err := os.Create("poems/" + id + ".md")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, defaultContent, time.Now().Format(time.RFC3339))
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}
