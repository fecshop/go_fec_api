package main

import (
	"log"
	//"mimi/djq/config"
	//"mimi/djq/constant"
	//"mimi/djq/initialization"
	//"mimi/djq/router"
	//"mimi/djq/task"
	"time"
    "github.com/fecshop/go_fec_api/initialization"
    "github.com/fecshop/go_fec_api/router"
)

func main() {
	log.Println("------start：" + time.Now().String())
	initialization.InitGlobalLog()
	log.Println("------start：" + time.Now().String())
	log.SetFlags(log.LstdFlags | log.Llongfile)
	//initData()
	//initTestData()
	//beginTask()
	router.Begin()
}

/*
func beginTask() {
	if "true" == config.Get("task_run") {
		go task.CheckPayingOrder()
		go task.CheckRefundingOrder()
		go task.AgreeNotUsedRefunding()
		go task.CountCashCoupon()
		go task.CountForPromotionalPartner()
		go task.CheckExpiredCashCoupon()
		go task.CheckExpiredPresent()
	}
}

func initData() {
	constant.AdminId = initialization.InitData()
}

func initTestData() {
	if config.Get("buildTestData") == "true" {
		initialization.InitTestData()
	}
}
*/