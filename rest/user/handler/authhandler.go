package handler

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var p UserParam
	if e := c.ShouldBind(&p); e != nil {

	}

}

func Register(c *gin.Context) {

}
