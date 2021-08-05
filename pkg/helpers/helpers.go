/**
  @author:panliang
  @data:2021/8/4
  @note
**/
package helpers

import (
	"bytes"
	"os/exec"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

//执行这个命令
func ShellExcel(shell string) {
	cmd := exec.Command("bin/bash", "-c", shell)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	CheckErr(err)
}
//执行这个api
func CurlAPi(api string){
	cmd := exec.Command("bin/bash", "-c", "curl "+api)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	CheckErr(err)
}


