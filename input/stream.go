package input

import (
	"strings"

	"github.com/tarm/serial"
	"go.uber.org/atomic"
)

type Reader struct {
	s    *serial.Port
	ch   chan string
	stop *atomic.Bool
}

func NewReader(name string) *Reader {
	c := &serial.Config{Name: name, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		panic(err)
	}

	return &Reader{
		s:    s,
		ch:   make(chan string, 50),
		stop: atomic.NewBool(false),
	}
}

func (r *Reader) Start() {
	var rest string
	for {
		if r.stop.Load() {
			r.s.Close()
			break
		}

		buf := make([]byte, 256)
		n, err := r.s.Read(buf)
		if err != nil {
			panic(err)
		}

		buf = append([]byte(rest), buf[:n]...)
		read := string(buf)
		rest = ""

		parts := strings.Split(read, "\r\n")
		if len(parts) == 1 {
			rest = read
		} else {
			for i := 0; i < len(parts)-1; i++ {
				line := parts[i]
				r.ch <- line
			}

			rest = parts[len(parts)-1]
		}
	}
}

func (r *Reader) Channel() <-chan string {
	return r.ch
}

func (r *Reader) Stop() {
	r.stop.Store(true)
}
