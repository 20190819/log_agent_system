package main

import (
	"fmt"

	"github.com/hpcloud/tail"
)

const (
	LOG_FILE = "./xx.log"
)

func main() {
	t, err := tail.TailFile(LOG_FILE, tail.Config{
		Location: &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:   true,
		Follow:   true,
		Poll:     true,
	})
	if err != nil {
		fmt.Println("TailFile error: ", err.Error())
	}
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
