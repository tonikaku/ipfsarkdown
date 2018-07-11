package ipfstool

import (
"os/exec"
"fmt"
"os"
)

func IPFScmd(operation string)string {
	var result []byte
	var err error
	var cmd *exec.Cmd
	// 执行单个shell命令时, 直接运行即可
	cmd = exec.Command("/bin/bash","-c",operation)
	if result, err = cmd.Output(); err != nil {
		fmt.Println(err)
		//fmt.Println("XX")
		os.Exit(1)
	}
	// 默认输出有一个换行
	//fmt.Println(string(result))
	// 指定参数后过滤换行符
	//fmt.Println(strings.Trim(string(result), "\n"))

	return  string(result);
}

