package main

// #cgo CFLAGS: -I${SRCDIR}/../grt
// #cgo LDFLAGS: -lstdc++ -L${SRCDIR}/../grt -lwgrt -L${SRCDIR}/../grt -lgrt
/*
 #include "wgrt.h"
*/
import "C"
import (
	"fmt"
	"os"
	"os/signal"
	"unsafe"

	"github.com/Jeaung/go-grt/input"
	"github.com/Jeaung/go-grt/util"
	"github.com/pkg/errors"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic123", errors.WithStack(err.(error)))
		}
	}()

	C.init(C.int(util.ModelPredict))

	r := input.NewReader("/dev/cu.usbmodem141401")
	go r.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		fmt.Println("program exits")
		r.Stop()
		os.Exit(0)
	}()

	for {
		line := <-r.Channel()

		p, err := util.ParsePoint(line)
		if err == nil {
			result := float32(C.predict((*C.Point)(unsafe.Pointer(p))))
			if result != 0 {
				label := int(result)
				likelihood := (result - float32(label)) * 10
				fmt.Println("label", label, "likelihood", likelihood)
			}
		}
	}
}
