package config
/* 新建文件config.go, 然后将下面的内容复制进去，添加或者更改相应的配置即可
package config

import (
	//"github.com/fecshop/go_fec_api/handler"  
	
)

func Get(key string) string{
    cf := make(map[string]string)
    cf["logFileDir"]    = `/www/web_logs/go_fec_api/`   // log存放的地址
    cf["logGlobalFile"] = `global.log`                  // log文件
    cf["server_ip"]     = `120.24.37.249`               // http服务对应的ip
    cf["server_port"]   = `3000`                        // http服务对应的prot
    cf["token"]         = `xxxxx` // 验证token，这个和fecshop_enterprise 里面的token值要对应，这个相当于apiKey，因为是内网访问，因此做一个key值验证就行了。
    // mysql配置
    cf["mysql_host"]    = `127.0.0.1:3306`  // mysql host
    cf["mysql_user"]    = `root`            // mysql user
    cf["mysql_database"]         = `fecshop`
    cf["mysql_password"]         = `xxxx`
    cf["mysql_charset"]          = `utf8`
    cf["mysql_maxOpenConns"]     = `200`    // mysql 用于设置最大打开的连接数，默认值为0表示不限制。 
    cf["mysql_maxIdleConns"]     = `100`    // mysql 用于设置闲置的连接数。
    
    return cf[key]
}
*/