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
    
    return cf[key]
}