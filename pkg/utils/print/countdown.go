package print

import (
	"fmt"
	"os"
	"time"
)

func CountDown(duration time.Duration, name string) {
	for remaining := duration; remaining > 0; remaining -= time.Second {
		fmt.Println(name+", wait:", int(remaining.Seconds()))
		os.Stdout.Sync()
		time.Sleep(time.Second)
	}
}
