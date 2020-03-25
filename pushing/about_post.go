package pushing

import (
	"athena/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/*
	从kafka读数据,
			判断操作类型 - 做出逻辑处理 -更新 - 点赞数/阅读量/臭鸡蛋数
									- 新增  - 评论
											- 一整个文章
*/

// 插入完整 文章
func InsertPost(p models.Post) {

}

// 更新(维护)文章 点赞/踩/评论
func UpdatePost(c *gin.Context) {
	var p models.Post
	c.BindWith(&p, binding.Form)

	//从http.Request中读取值到User结构体中,根据请求方法类型和请求内容格式类型自动确定绑定类型
	c.Bind(p)
}

/*
	根据要求展示文章
					广场

*/

// 展示 广场动态

// 展示 关注动态

// 展示 tag 数据

// 展示 动态info (评论)
