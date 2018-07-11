package main

import (
	"ipfsarkdown/src/markdowner"
	"os"
)

func main() {

    previewer := markdowner.NewPreviewer(8893)
	previewer.UseBasic()
	previewer.Run(os.Args...)
    ////os.Args...


}

