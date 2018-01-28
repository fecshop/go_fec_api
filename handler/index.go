package handler

import (
	//"mimi/djq/util"
	"net/http"
    "github.com/gin-gonic/gin"  
)


func NotFound(c *gin.Context) {
    body := make(gin.H) 
    body["content"] = "Not Found Page"
	c.AbortWithStatusJSON(http.StatusNotFound, body)
}