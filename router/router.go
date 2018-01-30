package router

import (
	"github.com/gin-gonic/gin"
	//"io"
	"github.com/fecshop/go_fec_api/config"
	"github.com/fecshop/go_fec_api/handler"
    _ "github.com/go-sql-driver/mysql" 
    mysqlPool "github.com/fecshopsoft/golang-db/mysql"
	//"os"
	//"github.com/fecshop/go_fec_api/filepath"
)

func mysqlDBPool() *mysqlPool.SQLConnPool{
    host         := config.Get("mysql_host")
    database     := config.Get("mysql_database")
    user         := config.Get("mysql_user")
    password     := config.Get("mysql_password")
    charset      := config.Get("mysql_charset")
    // 用于设置最大打开的连接数
    maxOpenConns := config.Get("mysql_maxOpenConns")
    // 用于设置闲置的连接数
    maxIdleConns := config.Get("mysql_maxIdleConns")
    mysqlDB := mysqlPool.InitMySQLPool(host, database, user, password, charset, maxOpenConns, maxIdleConns)
    return mysqlDB
}


func Begin() {
    mysqlDB := mysqlDBPool();
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()
	//initLog()
	router := gin.Default()
	router.NoRoute(handler.NotFound)
    v1 := router.Group("/v1", handler.ApiTokenValidate)
    {
        //v1.GET("/cms/article/list",  func(c *gin.Context) {
        //   handler.Article.List(c, mysqlDB)
        //})
        v1.GET("/cms/articles/primaryKey",  func(c *gin.Context) {
            handler.Article.PrimaryKey(c, mysqlDB)
        })
        v1.GET("/cms/articles/oneByUrlKey",  func(c *gin.Context) {
            handler.Article.OneByUrlKey(c, mysqlDB)
        })
        v1.GET("/cms/articles/oneById",  func(c *gin.Context) {
            handler.Article.OneById(c, mysqlDB)
        })
        v1.POST("/cms/articles/coll",  func(c *gin.Context) {
            handler.Article.Coll(c, mysqlDB)
        })
    }
        
        
    /*
	mi := router.Group("/mi", handler.ApiGlobal, handler.AdminCheckLogin)

	mi.POST("/login", handler.AdminLogin)
	mi.POST("/logout", handler.AdminLogout)

	miAdmin := mi.Group("/admin")
	miAdmin.GET("/", handler.PermissionAdminR, handler.AdminList)
	miAdmin.GET("/:id", handler.PermissionAdminR, handler.AdminGet)
	miAdmin.POST("/", handler.PermissionAdminC, handler.AdminPost)
	miAdmin.PATCH("/:id", handler.PermissionAdminU, handler.AdminPatch)
	miAdmin.DELETE("/", handler.PermissionAdminD, handler.AdminDelete)

	mi.GET("/adminAction/self", handler.AdminGetSelf)
	mi.PATCH("/adminAction/self", handler.AdminPatchSelf)

	miRole := mi.Group("/role")
	miRole.GET("/", handler.PermissionRoleR, handler.RoleList)
	miRole.GET("/:id", handler.PermissionRoleR, handler.RoleGet)
	miRole.POST("/", handler.PermissionRoleC, handler.RolePost)
	miRole.PATCH("/:id", handler.PermissionRoleU, handler.RolePatch)
	miRole.DELETE("/", handler.PermissionRoleD, handler.RoleDelete)

	open := router.Group("/open", handler.ApiGlobal)
	open.GET("/getServerRootUrl", handler.GetServerRootUrl)
	
	router.Static("/static", "html/assets")
	router.Static("/html", "html")

	router.StaticFile("/", "html/ui/index.html")
	
	router.StaticFile("/MP_verify_IrIxUvnx9Bob0ktY.txt", "/home/zhaohao/app/MP_verify_IrIxUvnx9Bob0ktY.txt")
	*/
	router.Run(config.Get("server_ip") + ":" + config.Get("server_port"))
}
/*
func initLog() {
	if "false" == config.Get("output_log") {
		return
	}
	gin.SetMode(gin.ReleaseMode)
	// Logging to a file.
	//gin_file, _ := os.Create("gin.log")
	routerInfoLogUrl := config.Get("router_info_log")
	if routerInfoLogUrl == "" {
		routerInfoLogUrl = "logs/router_info.log"

	}
	path := filepath.Dir(routerInfoLogUrl)
	os.MkdirAll(path, 0777)

	routerErrorLogUrl := config.Get("router_error_log")
	if routerErrorLogUrl == "" {
		routerErrorLogUrl = "logs/router_error.log"
	}
	path = filepath.Dir(routerErrorLogUrl)
	os.MkdirAll(path, 0777)
	infoFile, err := os.OpenFile(routerInfoLogUrl, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	errorFile, err2 := os.OpenFile(routerErrorLogUrl, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err2 != nil {
		panic(err2)
	}
	//ginLogUrl := "gin.log"
	//var gin_file *os.File
	//if _,err := os.Stat(ginLogUrl);err!=nil{
	//	gin_file, _ = os.Create(ginLogUrl)
	//}else{
	//	gin_file, _ = os.OpenFile("gin.log", os.O_RDWR|os.O_APPEND, 0666)
	//}
	//gin_error_file, _ := os.Create(routerErrorLogUrl)
	gin.DefaultWriter = io.MultiWriter(infoFile)
	gin.DefaultErrorWriter = errorFile
	//gin.RecoveryWithWriter(gin_error_file)
}
*/