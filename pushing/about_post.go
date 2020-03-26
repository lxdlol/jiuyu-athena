package pushing

import (
	"athena/common"
	mgo_db "athena/db"
	"athena/models"
	"athena/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

/*
	从kafka读数据,
			判断操作类型 - 做出逻辑处理 -更新 - 点赞数/阅读量/臭鸡蛋数
									- 新增  - 评论
											- 一整个文章
*/

func ForReadKafka() {
	for {
		ReadKafka()
	}
}

func ReadKafka() {
	/* k_r := */ utils.GetKafkaReader("订阅的主题", []string{})
	/*
		k_r.Reader.ReadMessage  //消费数据

	*/
}

/*
	根据要求展示文章
					广场

*/

// 展示 广场动态

// 展示 关注动态

// 展示 tag 数据

// 展示 动态info (评论)

// 插入完整 文章
func InsertPost(c *gin.Context) {
	var (
		p         models.Post
		mgo, conn = mgo_db.Connect("post")
	)
	c.BindWith(&p, binding.Form)

	//从http.Request中读取值到User结构体中,根据请求方法类型和请求内容格式类型自动确定绑定类型
	c.Bind(p)
	if err := conn.Insert(&p); err != nil {
		mgo.Close()
		common.GResp.Failure(c, common.CodeInternalServerError, errors.New("发布文章 出现错误"))
		return
	}
	mgo.Close()
	common.GResp.Success(c, nil)
	return
}

// @Tags 更新(维护)文章 点赞
// @点赞
// @Accept  json
// @Produce json
// @Param fabulous_action query string true "减:- 增:+"
// @Param news_id query string true "文章的id"
// @Success 200
// @Router /update_post_fabulous [post]
func UpdatePostFabulous(c *gin.Context) {
	var (
		p         models.Post
		mgo, conn = mgo_db.Connect("post")
		err       error

		fabulous_action = c.PostForm("fabulous_action")
		news_id         = c.PostForm("news_id")
	)

	Select(conn, news_id, &p)
	switch fabulous_action {
	case "-":
		if num := p.Flowers - 1; num < 0 {
			p.Flowers = 0
		}
	case "+":
		p.Flowers += 1
	}

	bsonid := bson.M{"_id": p.Id}
	bsonobj := bson.M{"$set": &bson.M{"flowers": p.Flowers}}

	if err = conn.Update(&bsonid, &bsonobj); err != nil {
		mgo.Close()
		common.GResp.Failure(c, common.CodeInternalServerError, errors.New("文章(点赞)  出现错误"))
		return
	}

	mgo.Close()
	common.GResp.Success(c, nil)
	return
}

// @Tags 更新(维护)文章 踩
// @踩
// @Accept  json
// @Produce json
// @Param step_on_action query string true "减:- 增:+"
// @Param news_id query string true "文章的id"
// @Success 200
// @Router /update_post_step_on [post]
func UpdatePostStepOn(c *gin.Context) {
	var (
		p         models.Post
		mgo, conn = mgo_db.Connect("post")
		err       error

		step_on_action = c.PostForm("step_on_action")
		news_id        = c.PostForm("news_id")
	)

	Select(conn, news_id, &p)
	switch step_on_action {
	case "-":
		if num := p.Eggs - 1; num < 0 {
			p.Eggs = 0
		}
	case "+":
		p.Eggs += 1
	}

	bsonid := bson.M{"_id": p.Id}
	bsonobj := bson.M{"$set": &bson.M{"eggs": p.Eggs}}

	if err = conn.Update(&bsonid, &bsonobj); err != nil {
		mgo.Close()
		common.GResp.Failure(c, common.CodeInternalServerError, errors.New("文章(踩)  出现错误"))
		return
	}

	mgo.Close()
	common.GResp.Success(c, nil)
	return
}

// @Tags 更新(维护)文章 评论
// @评论
// @Accept  json
// @Produce json
// @Param action query string true "删除:delete 点赞:fabulous 踩:step_on"
// @Param fabulous_action query string true "点赞的更新是  增  还是  减"
// @Param comment_id query string true "评论的id"
// @Success 200
// @Router /update_post_comment [post]
func UpdatePostComment(c *gin.Context) {
	var (
		p         models.Post
		mgo, conn = mgo_db.Connect("post")
		err       error
		err_msg   string

		action          = c.PostForm("action")
		fabulous_action = c.PostForm("fabulous_action")
		comment_id_str  = c.PostForm("comment_id")
		comment_id, _   = strconv.Atoi(comment_id_str)
		news_id         = c.PostForm("news_id")
	)

	switch action {
	case "delete":
		Select(conn, news_id, &p)
		err = CommentAddOrReduce(conn, p, "-", models.Comment{CId: comment_id})
		err_msg = "文章评论(删除)  出现错误"
	case "fabulous":
		Select(conn, news_id, &p)
		err = FabulousAddOrReduce(conn, p, fabulous_action, models.Comment{CId: comment_id})
		err_msg = "文章评论(点赞)  出现错误"
	case "step_on":
		Select(conn, news_id, &p)
		err = StepOnAddOrReduce(conn, p, fabulous_action, models.Comment{CId: comment_id})
		err_msg = "文章评论(踩)  出现错误"
	}
	mgo.Close()

	if err != nil {
		common.GResp.Failure(c, common.CodeInternalServerError, errors.New(err_msg))
		return
	}

	common.GResp.Success(c, nil)
	return
}
