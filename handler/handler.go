package handler

import (
	//"mimi/djq/util"
    "fmt"
	"net/http"
    "github.com/gin-gonic/gin"  
    "github.com/fecshop/go_fec_api/config"
)

// 404 , 找不到错误
func NotFound(c *gin.Context) {
    jsonBody := make(gin.H) 
    jsonBody["status"] = http.StatusNotFound
    jsonBody["content"] = "Not Found Page"
	c.AbortWithStatusJSON(http.StatusNotFound, jsonBody)
}

// api token 验证
func ApiTokenValidate(c *gin.Context) {
    headerToken := c.Request.Header.Get("token");
    localToken := config.Get("token");
    if headerToken != localToken {
        jsonBody := make(gin.H) 
        jsonBody["status"] = HStatusOK
        jsonBody["content"] = "token is not right"
        c.AbortWithStatusJSON(http.StatusOK, jsonBody)
		return
    }
}
// 数据库异常
func MysqlErr(c *gin.Context, err error){ 
    fmt.Printf("%s\r\n","mysql query error")
    jsonBody := make(gin.H) 
    data := make(gin.H) 
    data["desc"] = `connect mysql error`
    data["error"] = err.Error()
    jsonBody["status"] = HStatusMysqlConnectError
    jsonBody["data"]   = data
    c.AbortWithStatusJSON(http.StatusOK, jsonBody)
}

// 参数不正确异常
func ParamErr(c *gin.Context, errStr string){ 
    fmt.Printf("%s\r\n","request param error")
    jsonBody := make(gin.H) 
    data := make(gin.H) 
    data["desc"] = errStr
    jsonBody["status"] = HStatusParamNotRight
    jsonBody["data"]   = data
    c.AbortWithStatusJSON(http.StatusOK, jsonBody)
}












