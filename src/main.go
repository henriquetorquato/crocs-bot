package main

import (
	"flag"

	handlers "./handlers"
)

const defaultPeriod = 12

func main() {
	period := flag.Int("period", defaultPeriod, "period in hours for the weather forecast")
	handlers.CreatePost(*period)
}
