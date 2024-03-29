package cat

import (
    "ego/ego/src/commons"
    "encoding/json"
    "net/http"
    "strconv"
)

func ItemCatHandler(){
    commons.Router.HandleFunc("/item/cat/show",showItemCatController)
}
func showItemCatController(w http.ResponseWriter,r *http.Request){
    id:=r.FormValue("id")
    if id==""{
        id="0"
    }
    idInt,_:=strconv.Atoi(id)
    t:=showCatByPid(idInt)
    b,_:=json.Marshal(t)
    w.Header().Set("Content-type","application/json;charset=utf-8")
    w.Write(b)
}