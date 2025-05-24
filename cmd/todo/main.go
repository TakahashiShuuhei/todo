package main

import (
	"log"
	"os"

	"github.com/TakahashiShuuhei/todo/internal/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
