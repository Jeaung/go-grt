package main

// #cgo CFLAGS: -I${SRCDIR}/../grt
// #cgo LDFLAGS: -lstdc++ -L${SRCDIR}/../grt -lwgrt -L/usr/local/ -lgrt
/*
 #include "wgrt.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/Jeaung/go-grt/input"
	"github.com/Jeaung/go-grt/keyboard"
	"github.com/Jeaung/go-grt/util"
	"github.com/nsf/termbox-go"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic", err)
			termbox.Close()
		}
	}()

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	w := keyboard.NewWatcher()
	w.Init()

	r := input.NewReader("/dev/cu.usbmodem141401")
	go r.Start()

	C.init(C.int(util.ModeModel))

	go func() {
		<-w.Exit()
		fmt.Println("program exits")
		termbox.Close()
		r.Stop()
	}()

	fmt.Println(`program initialize
press f1 ~ f12 to set training data label
press mouse left key to record sample and release it to stop recording
press enter to train the model
press esc to exit`)

	sample := make([]util.Point, 0, 50)

	for {
		line := <-r.Channel()

		if w.Recording() {
			p, err := util.ParsePoint(line)
			if err == nil {
				sample = append(sample, *p)
			}
		}

		select {
		case <-w.Train():
			go func() {
				fmt.Println("start training")
				C.train()
			}()
		case <-w.Start():
			fmt.Println("start recording")
		case <-w.Stop():
			fmt.Println("stop recording", len(sample), "points in the sample")
			C.addSample(C.uint(w.Label()), (*C.Point)(unsafe.Pointer(&sample[0])), C.int(len(sample)))
			sample = make([]util.Point, 0, 50)
		default:
		}
	}
}
