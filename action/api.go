/**
  @author:panliang
  @data:2021/8/4
  @note
**/
package action

import (
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
	cmd := c.Query("cmd")
	id  := c.Query("id")
	Id, err := crontab.CornCmd.AddFunc(cmd, func() {
		helpers.ShellExcel("pwd")
	})
	helpers.CheckErr(err)
	db.DB.Table("tasks").Where("id=?", id).Update("job_id", Id)
	c.JSON(200, map[string]interface{}{
		"msg": "success",
	})
}
//删除一个任务
func (*JobApi) RemoveJob(c *gin.Context) {
	id  := c.Query("job_id")
	jobId,_ := strconv.Atoi(id)
	crontab.CornCmd.Remove(cron.EntryID(jobId))
}
//停止任务
func (*JobApi) StopJob(c *gin.Context) {
	crontab.CornCmd.Stop()
}
//启动一个任务
func (*JobApi) StartJob(c *gin.Context) {
	crontab.CornCmd.Start()
}




