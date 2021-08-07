/**
  @author:panliang
  @data:2021/8/4
  @note
**/
package action

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"go_job/db"
	"go_job/pkg/crontab"
	"go_job/pkg/helpers"
	"strconv"
)

type JobApi struct{}

//添加一个任务
func (*JobApi) AddJob(c *gin.Context) {
	cycle := c.PostForm("cycle")
	textarea := c.PostForm("textarea")
	id  := c.PostForm("id")
	fmt.Println(cycle,textarea,id)
	Id, err := crontab.CornCmd.AddFunc(cycle, func() {
		helpers.ShellExcel(textarea)
	})
	fmt.Println("任务投递",Id)
	helpers.CheckErr(err)
	db.DB.Table("tasks").Where("id=?", id).Update("job_id", Id)
	c.JSON(200, map[string]interface{}{
		"msg": "success",
	})
}
//删除一个任务
func (*JobApi) RemoveJob(c *gin.Context) {
	id  := c.Query("job_id")
	fmt.Println(id)
	jobId,_ := strconv.Atoi(id)
	crontab.CornCmd.Remove(cron.EntryID(jobId))
	db.DB.Table("tasks").Where("job_id=?",jobId).Updates(map[string]interface{
	}{"status":0, "job_id":""})
	c.JSON(200, map[string]interface{}{
		"msg": "success",
	})
}
//停止任务
func (*JobApi) StopJob(c *gin.Context) {
	crontab.CornCmd.Stop()
}
//启动一个任务
func (*JobApi) StartJob(c *gin.Context) {
	crontab.CornCmd.Start()
}




