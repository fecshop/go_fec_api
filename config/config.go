package config

import (
	//"github.com/fecshop/go_fec_api/handler"  
	
)

func Get(key string) string{
    cf := make(map[string]string)
    cf["logFileDir"]    = `/www/web_logs/go_fec_api/`
    cf["logGlobalFile"] = `global.log`
    cf["server_ip"]     = `120.24.37.249`
    cf["server_port"]   = `3000`
    cf["token"]         = `4Dr4%i2mX&BP^7lFTd%0!1IX2^zx26F8`
    // mysql配置
    cf["mysql_host"]    = `127.0.0.1:3306`
    cf["mysql_user"]    = `root`
    cf["mysql_database"]         = `fecshop`
    cf["mysql_password"]         = `Zhaoyong2017fdsfds3f3GDs3fgsd`
    cf["mysql_charset"]          = `utf8`
    cf["mysql_maxOpenConns"]     = 200
    cf["mysql_maxIdleConns"]     = 100
    
    return cf[key]
}