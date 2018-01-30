package handler
/**
 * json里面的状态值。
 */

// 正常200
var HStatusOK                =  200

//  **** 系统报错  **** 
// mysql连接报错
var HStatusMysqlConnectError =  1000001


//  **** 程序异常  **** 
// token 不存在
var HStatusTokenNotRight     =  1010002
// 请求参数不正确
var HStatusParamNotRight     =  1010003
//HttpStatus.XX = 101000001


//HttpStatus.XX = 102000001