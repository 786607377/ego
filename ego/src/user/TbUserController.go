package user

import (
    "ego/ego/src/commons"
    "encoding/json"
    "net/http"
)
//
func UserHandler(){
    commons.Router.HandleFunc("/login",loginController)
}
//登录
func loginController(w http.ResponseWriter,r *http.Request){
    username:=r.FormValue("username")
    password:=r.FormValue("password")
    er:=LoginService(username,password)
    //把结构体转换为json数据
    b,_:=json.Marshal(er)
    w.Header().Set("Content-Type", " application/json;charset=utf-8") //将内容指定为JSON格式，以UTF-8字符编码进行编码。
    w.Write(b)
}