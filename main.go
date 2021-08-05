/**
  @author:panliang
  @data:2021/8/3
  @note
**/
package main

import (
	"github.com/gin-gonic/gin"
	"go_job/action"
	"go_job/pkg/crontab"
	"go_job/db"
)

//# For details see man 4 crontabs
//# Example of job definition:
//# .---------------- minute (0 - 59)
//# |  .------------- hour (0 - 23)
//# |  |  .---------- day of month (1 - 31)
//# |  |  |  .------- month (1 - 12) OR jan,feb,mar,apr ...
//# |  |  |  |  .---- day of week (0 - 6) (Sunday=0 or 7) OR sun,mon,tue,wed,thu,fri,sat
//# |  |  |  |  |
//# *  *  *  *  * user-name  command to be executed

func main() {
	//加载sqlite
	db.Conn()
	crontab.New()
	app := gin.Default()
	//启动任务
	crontab.Start()
	app.GET("/addJob")
	job := new(action.JobApi)
	app.GET("/api/addJob", job.AddJob)       //添加一个任务
	app.GET("/api/removeJob", job.RemoveJob) //剔除一个任务
	app.GET("/api/stopJob", job.StopJob)     //关闭调度器
	app.GET("/api/startJob", job.StartJob)   //启动调度器
	//启动端口
	_ = app.Run(":9201")
}
