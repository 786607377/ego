package item

import (
    "ego/ego/src/commons"
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
)

func ItemHandler(){
    commons.Router.HandleFunc("/showItem",showItemController)
    commons.Router.HandleFunc("/item/delete",delByIdsController)
    commons.Router.HandleFunc("/item/instock",inStockController)
    commons.Router.HandleFunc("/item/offstock",offStockController)
    commons.Router.HandleFunc("/item/imageupload",imagesUploadController)
}
//显示商品信息
func showItemController(w http.ResponseWriter,r *http.Request){
    page,err:=strconv.Atoi(r.FormValue("page"))
    if err!=nil{
        fmt.Println("转换数据失败")
        return
    }
    rows,err:=strconv.Atoi(r.FormValue("rows"))
    if err!=nil{
        fmt.Println("转换数据失败")
        return
    }
    datagrid:=showItemService(rows,page)
    b,_:=json.Marshal(datagrid)
    w.Header().Set("Content-Type", "application/json;charset=utf-8") //将内容指定为JSON格式，以UTF-8字符编码进行编码。
    w.Write(b)
}

//商品删除
func delByIdsController(w http.ResponseWriter,r*http.Request){
    ids:=r.FormValue("ids")
    er:=delByIdsService(ids)
    b,_:=json.Marshal(er)
    w.Header().Set("Content-type","application/json;charset=utf-8")
    w.Write(b)
}

//商品删除
func inStockController(w http.ResponseWriter,r*http.Request){
    ids:=r.FormValue("ids")
    er:=inStockService(ids)
    b,_:=json.Marshal(er)
    w.Header().Set("Content-type","application/json;charset=utf-8")
    w.Write(b)
}
func offStockController(w http.ResponseWriter,r*http.Request){
    ids:=r.FormValue("ids")
    er:=offStockService(ids)
    b,_:=json.Marshal(er)
    w.Header().Set("Content-type","application/json;charset=utf-8")
    w.Write(b)
}

func imagesUploadController(w http.ResponseWriter,r*http.Request){
    //file,fileHeader,err:=r.FormValue("imageFile")
}