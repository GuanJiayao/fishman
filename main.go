package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	userMoving = false
	isRunning  = false
	fish       = 0
	start      = time.Now()
	duration   = 3 * time.Second
)

func main() {
	fmt.Println(`
┌─┐┬┌─┐┬ ┬┌┬┐┌─┐┌┐┌
├┤ │└─┐├─┤│││├─┤│││
└  ┴└─┘┴ ┴┴ ┴┴ ┴┘└┘`)
	fmt.Printf("start now: %s \n", start.String())
	timer := time.NewTimer(duration)
	hook.Register(hook.MouseMove, []string{}, func(event hook.Event) {
		x, y := robotgo.Location()
		time.Sleep(time.Second)
		x1, y1 := robotgo.Location()
		if math.Abs(float64(x-x1)) > 10 || math.Abs(float64(y-y1)) > 10 {
			userMoving = true
			timer.Reset(duration)
			fmt.Println("You scared the fish away by moving just now...")
		} else {
			userMoving = false
		}
	})
	go moveListener(timer)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	go func() {
		<-sig
		fmt.Printf("\nYou caught a total of %d fish in %s. Enjoy it!\n", fish, time.Now().Sub(start))
		fmt.Println(`
┌─┐┌─┐┌─┐  ┬ ┬┌─┐┬ ┬
└─┐├┤ ├┤   └┬┘│ ││ │
└─┘└─┘└─┘   ┴ └─┘└─┘`)
		os.Exit(0)
	}()
	s := hook.Start()
	<-hook.Process(s)
}
func moveMouse() {
	if !isRunning {
		robotgo.MoveRelative(1, 1)
		robotgo.MilliSleep(500)
		robotgo.MoveRelative(-1, -1)
		isRunning = true
		fish++
		fmt.Printf("Caught %d fish in %v，start from: %v\n", fish, time.Now().Sub(start), start.String())
	}
}

func moveListener(timer *time.Timer) {
	for {
		select {
		case <-timer.C:
			if !userMoving {
				moveMouse()
				timer.Reset(duration)
				isRunning = false
			}
		}
	}
}
