package main

import (
	mini_tiktok "github.com/iceberg-work/mini-tiktok/internal/mini-tiktok"
	"os"

	_ "go.uber.org/automaxprocs"
)

// Go 程序的默认入口函数(主函数).
func main() {
	command := mini_tiktok.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
