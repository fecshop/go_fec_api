package initialization

import (
	"log"
	"github.com/fecshop/go_fec_api/config"
	"os"
    //"github.com/fecshop/go_fec_api/filepath"
)

func InitGlobalLog() {
    
    logDir := config.Get("logFileDir")  //   `/www/web_logs/go_fec_api/`
    logGlobalFileName := config.Get("logGlobalFile")
    
    os.MkdirAll(logDir, 0777)
    globalLogFile := logDir + logGlobalFileName
    
    logFile, err := os.OpenFile(globalLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
    
	log.SetOutput(logFile)
	log.Println()
    
}