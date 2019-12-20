package main

import (
    "ego/ego/src/commons"
    "ego/ego/src/item"
    "ego/ego/src/item/cat"
    "ego/ego/src/user"
    "fmt"
    "github.com/gorilla/mux"
    "html/template"
    "net/http"
)

//显示登录页面
func welcome(w http.ResponseWriter,r *http.Request) {
    t,err:=template.ParseFiles("view/login.html")
    if err!=nil{
        fmt.Println("=template.ParseFiles failed err=",err)
        return
    }
    t.Execute(w,nil)
}
func showPage(w http.ResponseWriter,r *http.Request){
    vars:=mux.Vars(r)
    t,err:=template.ParseFiles("view/"+vars["page"]+".html")
    if err!=nil{
        fmt.Println("=template.ParseFiles failed err=",err)
        return
    }
    t.Execute(w,nil)
}

func main(){
    commons.Router.PathPrefix("/static").Handler(http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))
    commons.Router.HandleFunc("/",welcome)
    commons.Router.HandleFunc("/page/{page}", showPage)
    user.UserHandler()
    item.ItemHandler()
    cat.ItemCatHandler()
    http.ListenAndServe("localhost:80",commons.Router)
}