package markdowner

import (
	"time"
	"os"
	"io/ioutil"
)

const (
	MonitorInterval = 300
	DataChannelSize    = 10
)

type DataChannel struct {
	Origin chan *[]byte
	Request chan bool
}

type Monitor struct {
	path   string
	ticker *time.Ticker
	stop   chan bool
	C      *DataChannel
}

func GetNewMonitor(path string) *Monitor {
	dataChan := DataChannel{make(chan *[]byte, DataChannelSize), make(chan bool)}
	return &Monitor{path, nil, nil, &dataChan}
}

func (w *Monitor) Start() {
	go func() {
		w.ticker = time.NewTicker(time.Millisecond * MonitorInterval)
		defer w.ticker.Stop()
		w.stop = make(chan bool)
		var currentTimestamp int64
		for {
			select {
			case <-w.stop:
				return
			case <-w.ticker.C:
				reload := false
				select {
				case <-w.C.Request:
					reload = true
				default:
				}

				info, err := os.Stat(w.path)
				if err != nil {
					continue
				}

				timestamp := info.ModTime().Unix()
				if currentTimestamp < timestamp || reload {
					currentTimestamp = timestamp

					origin, err := ioutil.ReadFile(w.path)
					if err != nil {
						continue
					}

					w.C.Origin <- &origin
				}
			}
		}
	}()
}

func (w *Monitor) Stop() {
	w.stop <- true
}
