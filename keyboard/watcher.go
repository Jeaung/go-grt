package keyboard

import (
	"github.com/nsf/termbox-go"
	"go.uber.org/atomic"
)

type Watcher struct {
	label     *atomic.Int32
	exit      chan bool
	start     chan bool
	stop      chan bool
	train     chan bool
	recording *atomic.Bool
}

func NewWatcher() *Watcher {
	return &Watcher{
		label:     atomic.NewInt32(0),
		stop:      make(chan bool, 1),
		start:     make(chan bool, 1),
		train:     make(chan bool, 1),
		exit:      make(chan bool, 1),
		recording: atomic.NewBool(false),
	}
}

func (w *Watcher) Init() {
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	go func() {
		for {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyEsc:
					println("pppppppesc")
					w.exit <- true
				case termbox.KeyF1:
					w.label.Store(1)
				case termbox.KeyF2:
					w.label.Store(2)
				case termbox.KeyF3:
					w.label.Store(3)
				case termbox.KeyF4:
					w.label.Store(4)
				case termbox.KeyF5:
					w.label.Store(5)
				case termbox.KeyF6:
					w.label.Store(6)
				case termbox.KeyF7:
					w.label.Store(7)
				case termbox.KeyF8:
					w.label.Store(8)
				case termbox.KeyF9:
					w.label.Store(9)
				case termbox.KeyF10:
					w.label.Store(10)
				case termbox.KeyF11:
					w.label.Store(11)
				case termbox.KeyF12:
					w.label.Store(12)
				case termbox.KeyEnter:
					w.train <- true
				}
			case termbox.EventMouse:
				switch ev.Key {
				case termbox.MouseLeft:
					w.start <- true
					w.recording.Store(true)
				case termbox.MouseRelease:
					w.recording.Store(false)
					w.stop <- true
				}
			case termbox.EventError:
				panic(ev.Err)
			}
		}
	}()
}

func (w *Watcher) Label() uint {
	return uint(w.label.Load())
}

func (w *Watcher) Exit() <-chan bool {
	return w.exit
}

func (w *Watcher) Train() <-chan bool {
	return w.train
}

func (w *Watcher) Stop() <-chan bool {
	return w.stop
}

func (w *Watcher) Start() <-chan bool {
	return w.start
}

func (w *Watcher) Close() {
	termbox.Close()
}

func (w *Watcher) Recording() bool {
	return w.recording.Load()
}
