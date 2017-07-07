package logger

import (
	"fmt"
	"os"
	"sync"
)

var m sync.Mutex

func StdoutInline(s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	os.Stdout.Write([]byte(fmt.Sprint(s...)))
}

func Stdout(s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	os.Stdout.Write([]byte(fmt.Sprintln(s...)))
}

func StderrInline(s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	os.Stderr.Write([]byte(fmt.Sprint(s...)))
}

func Stderr(s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	os.Stderr.Write([]byte(fmt.Sprintln(s...)))
}
