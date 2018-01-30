package handler

import (
	//"fmt"
    "log"
	"net/http"
    "github.com/gin-gonic/gin" 
    mysqlPool "github.com/fecshopsoft/golang-db/mysql"    
    //"github.com/fecshop/go_fec_api/config"
)

type article struct {
    Id                  int         `form:"id" json:"id" `
    UrlKey              string      `form:"url_key" json:"url_key" `   // binding:"required"
    Title               string      `form:"title" json:"title" `
    MetaKeywords        string      `form:"meta_keywords" json:"meta_keywords" `
    MetaDescription     string      `form:"meta_description" json:"meta_description" `
    Content             string      `form:"content" json:"content" `
    Status              int         `form:"status" json:"status" `
    CreatedAt           int         `form:"created_at" json:"created_at" `
    UpdatedAt           int         `form:"updated_at" json:"updated_at" `
    CreatedUserId       int         `form:"created_user_id" json:"created_user_id" `
}

type ArticleData struct{}

var Article ArticleData

// 得到article表的id
func (Article ArticleData) PrimaryKey(c *gin.Context, mysqlDB *mysqlPool.SQLConnPool){
    c.AbortWithStatusJSON(http.StatusOK, gin.H{
        "status":HStatusOK,
        "data":"id",
    })
}

// 通过url key 得到article
func (Article ArticleData) OneByUrlKey(c *gin.Context, mysqlDB *mysqlPool.SQLConnPool){
    url_key := string(c.Query("url_key"));
    log.Println(url_key)
    if url_key == "" {
        ParamErr(c, `request param url_key is empty`)
        return 
    }
    log.Println(3333)
    rows, err := mysqlDB.SQLDB.Query("SELECT id, url_key, title, meta_keywords, meta_description, content, status, created_at, updated_at, created_user_id FROM `article` where `url_key` = ? limit 1", url_key);
    if err != nil {
        MysqlErr(c, err)
        return
    }
    defer rows.Close()
    //articles := make([]article, 0)
    var m article
    for rows.Next() {
        
        err := rows.Scan(&m.Id, &m.UrlKey, &m.Title, &m.MetaKeywords, &m.MetaDescription, &m.Content, &m.Status, &m.CreatedAt, &m.UpdatedAt, &m.CreatedUserId)
        if err != nil {
            MysqlErr(c, err)
        }
        //articles = append(articles, m)
    }
    c.AbortWithStatusJSON(http.StatusOK, gin.H{
        "status":HStatusOK,
        "data":m,
    })
}

// 通过id 得到article 
func (Article ArticleData) OneById(c *gin.Context, mysqlDB *mysqlPool.SQLConnPool){
    id := string(c.Query("id"));
    log.Println(id)
    if id == "" {
        ParamErr(c, `request param id is empty`)
        return 
    }
    log.Println(3333)
    rows, err := mysqlDB.SQLDB.Query("SELECT id, url_key, title, meta_keywords, meta_description, content, status, created_at, updated_at, created_user_id FROM `article` where `id` = ? limit 1", id);
    if err != nil {
        MysqlErr(c, err)
        return
    }
    defer rows.Close()
    //articles := make([]article, 0)
    var m article
    for rows.Next() {
        
        err := rows.Scan(&m.Id, &m.UrlKey, &m.Title, &m.MetaKeywords, &m.MetaDescription, &m.Content, &m.Status, &m.CreatedAt, &m.UpdatedAt, &m.CreatedUserId)
        if err != nil {
            MysqlErr(c, err)
        }
        //articles = append(articles, m)
    }
    c.AbortWithStatusJSON(http.StatusOK, gin.H{
        "status":HStatusOK,
        "data":m,
    })
}

type collParam struct {
    CountSql             string      `form:"countSql" json:"countSql" `   // binding:"required"
    AllSql               string      `form:"allSql" json:"allSql" `
}

// 通过id 得到article 
func (Article ArticleData) Coll(c *gin.Context, mysqlDB *mysqlPool.SQLConnPool){
    var p collParam
    c.BindJSON(&p)
    countSql := p.CountSql;
    allSql   := p.AllSql;
    
    if countSql == "" {
        ParamErr(c, `countSql is empty`)
        return 
    }
    if allSql == "" {
        ParamErr(c, `allSql is empty`)
        return 
    }
    
    rows, err := mysqlDB.SQLDB.Query(allSql);
    if err != nil {
        MysqlErr(c, err)
        return
    }
    defer rows.Close()
    articles := make([]article, 0)
    var m article
    for rows.Next() {
        
        err := rows.Scan(&m.Id, &m.UrlKey, &m.Title, &m.MetaKeywords, &m.MetaDescription, &m.Content, &m.Status, &m.CreatedAt, &m.UpdatedAt, &m.CreatedUserId)
        if err != nil {
            MysqlErr(c, err)
        }
        articles = append(articles, m)
    }
    c.AbortWithStatusJSON(http.StatusOK, gin.H{
        "status":HStatusOK,
        "data":gin.H{
            "coll":articles,
            "count":1000,
        },
    })
}


/*
// 得到article 的List
func (Article ArticleData) List(c *gin.Context, mysqlDB *mysqlPool.SQLConnPool){
    rows, err := mysqlDB.Query("SELECT * FROM `article` ")
    if err != nil {
        MysqlErr(c, err)
    }
    var dbdata []gin.H
    if rows != nil {
        for _, row := range rows {
            dbdata = append(dbdata, gin.H(row))
        }
    }
    jsonBody := make(gin.H) 
    jsonBody["status"] = HStatusOK
    jsonBody["data"]   = dbdata
    c.AbortWithStatusJSON(http.StatusOK, jsonBody)
}
*/






