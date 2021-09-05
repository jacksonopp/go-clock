package main

import (
	"os"
	"time"

	"github.com/jacksonopp/go-clock/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
