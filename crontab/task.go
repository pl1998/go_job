/**
  @author:panliang
  @data:2021/8/4
  @note
**/
package crontab

import (
	"bytes"
	"fmt"
	"github.com/robfig/cron/v3"
	 "go_job/db"
	"go_job/pkg/helpers"
	"os/exec"
)

var CornCmd * cron.Cron

func New() *cron.Cron {
	CornCmd = cron.New()
	return CornCmd
}

type TaskList struct {
	ID int8 `json:"id"`
	TaskName string `json:"task_name"`
	Status int `json:"status"`
	Textarea string `json:"textarea"`
	Cycle string `json:"cycle"`
	Type int `json:"type"`
	OpName string `json:"op_name"`
	Email string `json:"email"`
}

//启动并执行任务
func Start()  {
	var TaskList []TaskList
	result := db.DB.Table("tasks").Find(&TaskList)
	helpers.CheckErr(result.Error)
	fmt.Println(TaskList)
	for _,value := range TaskList  {
		CornCmd.AddFunc(value.Cycle, func() {
			if value.Type == 1 {
				fmt.Println("测试")
			}
			if value.Type == 2 {
				ShellExcel(value.Textarea)
			}
			if value.Type == 3 {
			}
		})
	}
	select {
	}
}

func ShellExcel(shell string)  {
	cmd := exec.Command("bin/bash","-c",shell)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	helpers.CheckErr(err)
}
