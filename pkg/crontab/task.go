/**
  @author:panliang
  @data:2021/8/5
  @note
**/
package crontab

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"go_job/db"
	"go_job/pkg/helpers"
)

var CornCmd *cron.Cron

func New() *cron.Cron {
	CornCmd = cron.New()
	return CornCmd
}

type TaskList struct {
	ID       int8   `json:"id"`
	TaskName string `json:"task_name"`
	Status   int    `json:"status"`
	Textarea string `json:"textarea"`
	Cycle    string `json:"cycle"`
	Type     int    `json:"type"`
	OpName   string `json:"op_name"`
	Email    string `json:"email"`
}

//启动并执行任务
func Start() {
	CornCmd.Start()
	var TaskList []TaskList
	//将所有任务放入队列
	result := db.DB.Table("tasks").Where("status=?",1).Find(&TaskList)
	helpers.CheckErr(result.Error)
	fmt.Println(TaskList)
	for _,value := range TaskList  {
		if value.Type == 1 {
			Id, err := CornCmd.AddFunc(value.Cycle, func() {
				helpers.Cmd("curl "+value.Textarea,true)
			})
			helpers.CheckErr(err)
			db.DB.Table("tasks").Where("id=?", value.ID).Update("job_id", Id)
		}

		if value.Type == 2 {
			//执行一个shell
			fmt.Println(value.Cycle)
			Id, err := CornCmd.AddFunc(value.Cycle, func() {
				helpers.Cmd(value.Textarea,true)
			})
			helpers.CheckErr(err)
			db.DB.Table("tasks").Where("id=?", value.ID).Update("job_id", Id)
		}
		if value.Type == 3 {

		}


	}

}

