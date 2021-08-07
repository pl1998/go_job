/**
  @author:panliang
  @data:2021/8/4
  @note
**/
package helpers

import (
	"fmt"
	"os/exec"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

//执行这个命令
func ShellExcel(shell string) {
	out := string(Cmd(shell,true))
	fmt.Println(out)

}
//执行一个命令
func Cmd(cmd string, shell bool) []byte {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return out
	} else {
		out, err := exec.Command(cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return out
	}
}


