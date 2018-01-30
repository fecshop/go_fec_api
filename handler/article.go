package handler

import (
	//"fmt"
    "log"
	"net/http"
    "time"
    "strconv"
    "strings"
    "github.com/gin-gonic/gin" 
    mysqlPool "github.com/fecshopsoft/golang-db/mysql"    
    //"github.com/fecshop/go_fec_api/config"
)

type article struct {
    Id                  string         `form:"id" json:"id" `
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
    rows := mysqlDB.SQLDB.QueryRow("SELECT id, url_key, title, meta_keywords, meta_description, content, status, created_at, updated_at, created_user_id FROM `article` where `url_key` = ?", url_key);
    var m article
    err := rows.Scan(&m.Id, &m.UrlKey, &m.Title, &m.MetaKeywords, &m.MetaDescription, &m.Content, &m.Status, &m.CreatedAt, &m.UpdatedAt, &m.CreatedUserId)
    if err != nil {
        MysqlErr(c, err)
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
    rows := mysqlDB.SQLDB.QueryRow("SELECT id, url_key, title, meta_keywords, meta_description, content, status, created_at, updated_at, created_user_id FROM `article` where `id` = ? ", id);
    var m article
    err := rows.Scan(&m.Id, &m.UrlKey, &m.Title, &m.MetaKeywords, &m.MetaDescription, &m.Content, &m.Status, &m.CreatedAt, &m.UpdatedAt, &m.CreatedUserId)
    if err != nil {
        MysqlErr(c, err)
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

// 通过 allSql 和 countSql 得到coll
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
    // 通过allSql得到coll
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
    // 通过countSql得到count
    var countVal uint
    rowsCount, err := mysqlDB.SQLDB.Query(countSql)
    if err != nil {
        MysqlErr(c, err)
        return
    }
    defer rowsCount.Close()
    if rowsCount.Next() {
        rowsCount.Scan(&countVal)
    }

    c.AbortWithStatusJSON(http.StatusOK, gin.H{
        "status":HStatusOK,
        "data":gin.H{
            "coll":articles,
            "count":countVal,
        },
    })
}

// 保存article ，包含update 和insert
func (Article ArticleData) Save(c *gin.Context, mysqlDB *mysqlPool.SQLConnPool){
    var p article
    c.BindJSON(&p)
    nowUnix := int(time.Now().Unix())
    p.UpdatedAt = nowUnix;
    if p.Id != "" {
        affect, err := mysqlDB.Update(`
            update article set  
            url_key = ? , 
            title = ? , 
            meta_keywords = ? ,
            meta_description = ? ,
            content = ? ,
            status = ? ,
            updated_at = ? ,
            created_user_id = ?
            where id = ?`,
            p.UrlKey, p.Title, p.MetaKeywords, p.MetaDescription,
            p.Content, p.Status, p.UpdatedAt, p.CreatedUserId,
            p.Id)
        if err != nil {
            MysqlErr(c, err)
            return
        }
        c.AbortWithStatusJSON(http.StatusOK, gin.H{
            "status":HStatusOK,
            "data":gin.H{
                "affect":affect,
                "one":p,
            },
        }) 
        return
    } else {
        p.CreatedAt = nowUnix;
        lastid, err := mysqlDB.Insert(`
            insert into article ( 
                url_key, 
                title, 
                meta_keywords,
                meta_description,
                content,
                status,
                created_at,
                updated_at,
                created_user_id
            ) VALUES
            (
                ?, ?, ?, ?, ?, ?, ?, ?, ?
            )
            `,
            p.UrlKey, p.Title, p.MetaKeywords, p.MetaDescription,
            p.Content, p.Status, p.CreatedAt, p.UpdatedAt, p.CreatedUserId )
        if err != nil {
            MysqlErr(c, err)
            return
        }
        p.Id = strconv.Itoa(int(lastid))
        c.AbortWithStatusJSON(http.StatusOK, gin.H{
            "status":HStatusOK,
            "data":gin.H{
                "lastid":lastid,
                "one":p,
            },
        }) 
        return
    }
    
    
}



// 保存article ，包含update 和insert
func (Article ArticleData) UpdateUrlKeyInfo(c *gin.Context, mysqlDB *mysqlPool.SQLConnPool){
    var p article
    c.BindJSON(&p)
    nowUnix := int(time.Now().Unix())
    p.UpdatedAt = nowUnix;
    if p.Id != "" {
        affect, err := mysqlDB.Update(`
            update article set  
            url_key = ? , 
            updated_at = ? ,
            where id = ?`,
            p.UrlKey, p.UpdatedAt,
            p.Id)
        if err != nil {
            MysqlErr(c, err)
            return
        }
        c.AbortWithStatusJSON(http.StatusOK, gin.H{
            "status":HStatusOK,
            "data":gin.H{
                "affect":affect,
                "one":p,
            },
        }) 
        return
    }
    
    
}





// 保存article ，包含update 和insert
func (Article ArticleData) DelById(c *gin.Context, mysqlDB *mysqlPool.SQLConnPool){
    var p article
    c.BindJSON(&p)
    if p.Id != "" {
        affect, err := mysqlDB.Delete(`
            delete from article 
            where id = ?`,
            p.Id)
        if err != nil {
            MysqlErr(c, err)
            return
        }
        c.AbortWithStatusJSON(http.StatusOK, gin.H{
            "status":HStatusOK,
            "data":gin.H{
                "affect":affect,
            },
        }) 
        return
    }
    
}



// 保存article ，包含update 和insert
func (Article ArticleData) DelByIds(c *gin.Context, mysqlDB *mysqlPool.SQLConnPool){
    var p article
    c.BindJSON(&p)
    
    if p.Id != "" {
        
        args := strings.Split(p.Id, ",")
        affect, err := mysqlDB.Delete(`
            delete from article 
            where id in (?` + strings.Repeat(",?", len(args)-1) + `)`,
            args)
        if err != nil {
            MysqlErr(c, err)
            return
        }
        c.AbortWithStatusJSON(http.StatusOK, gin.H{
            "status":HStatusOK,
            "data":gin.H{
                "affect":affect,
            },
        }) 
        return
    }
    
}













