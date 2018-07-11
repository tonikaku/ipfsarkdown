package ipfstool

import (
	"io/ioutil"
	"fmt"
	"strings"
	"time"
)

func WaitUpload(file string)string{

	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
		return "false"
	}
	str := string(b)
	fmt.Println(str)
	if strings.Contains(str,"$$uploadthistoipfs$$") {
		str = strings.Replace(str,"$$uploadthistoipfs$$","",-1)
		err := ioutil.WriteFile(file, []byte(str), 0644)
		result := IPFScmd("ipfs add "+file) + time.Now().Format("2006-01-02 15:04:05")
		//fmt.Println(result)
		return  result
		check(err)
	}

	return "false"
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}