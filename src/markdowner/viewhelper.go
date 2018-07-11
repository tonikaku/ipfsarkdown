package markdowner

import (
"fmt"
"github.com/skratchdot/open-golang/open"
	"ipfsarkdown/src/ipfstool"
	"time"
	"strings"
	"io/ioutil"
)

const (
	MarkdownChanSize = 3
	Version          = "0.1"
)

func NewPreviewer(port int) *Previewer {
	return &Previewer{port, nil, make(chan bool)}
}

type Previewer struct {
	port       int
	httpServer *HTTPServer
	stop       chan bool
}

func (p *Previewer) Run(files ...string) {
	p.httpServer = NewHTTPServer(p.port)
	p.httpServer.Listen()

	for _, file := range files {
		addr := fmt.Sprintf("http://localhost:%d/%s", p.port, file)
			go func() {
				for true{
					time.Sleep(time.Second * 10)
					result := ipfstool.WaitUpload(file)
					if strings.Contains(result,"added"){
						err := ioutil.WriteFile("/users/desktop/ipfsarkdown.txt", []byte(result), 0644)
						check(err)
					}
				}
			}()
		open.Run(addr)
	}

	<-p.stop
}

func (p *Previewer) UseBasic() {
	MdTransformer.UseBasic()
}



func check(e error) {
	if e != nil {
		panic(e)
	}
}