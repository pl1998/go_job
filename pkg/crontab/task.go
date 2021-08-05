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
	result := db.DB.Table("tasks").Find(&TaskList)
	helpers.CheckErr(result.Error)
	fmt.Println(TaskList)
	for _,value := range TaskList  {
		Id, err := CornCmd.AddFunc(value.Cycle, func() {
			if value.Type == 1 {
				//执行一个url
				helpers.CurlAPi(value.Cycle)
			}
			if value.Type == 2 {
				//执行一个shell
			  helpers.ShellExcel(value.Cycle)
			}
			if value.Type == 3 {
				//发送邮件
			}
		})
		helpers.CheckErr(err)
		//将Id存入数据
		db.DB.Table("tasks").Where("id=?", value.ID).Update("job_id", Id)
	}

}

