package item

import (
    "database/sql"
    "ego/ego/src/commons"
    "fmt"
)

/*
rows：每页显示的条数
page:当前第几页
*/
func selByPageDao(rows ,page int)[]TbItem{
    //sql分页 表示每页有rows条从第 rows*(page-1)条开始查询rows条数据
    r,err:=commons.Dql("select * from tb_item limit ?,?",rows*(page-1),rows)
    if err!=nil{
        fmt.Println(err)
        return nil
    }
    ts:=make([]TbItem,0)
    for r.Next(){
        var t TbItem
        var s sql.NullString
        //如果直接使用t.Barcode由于数据库中列为Null导致填充错误
        r.Scan(&t.Id,&t.Title,&t.SellPoint,&t.Price,&t.Num,&s,&t.Image,&t.Cid,&t.Status,&t.Created,&t.Updated)
        t.Barcode=s.String
        ts=append(ts,t)
    }
    defer commons.CloseConn()
    return ts
}
/*
如果返回值<0表示查询失败
*/
//查询总条数
func selCount()(count int){
    rows,err:=commons.Dql("select count(*) from tb_item")
    defer commons.CloseConn()
    if err!=nil{
       fmt.Println(err)
       return -1
    }
    if rows.Next(){
        rows.Scan(&count)
        return
    }
    return -1
}

/*
返回值小于零表示更新失败
*/
func updStatusByIdsDao(ids []string, status int)int{
    if len(ids)<0{
        return -1
    }
    sql:="update tb_item set status=? where"
    for i:=0;i<len(ids);i++{
        sql+=" id="+ids[i]
        if i<len(ids)-1 {
            sql+=" or"
        }
    }
    count, err:=commons.Dml(sql,status)
    if err!=nil{
        fmt.Println(err)
        return -1
    }
    return int(count)
}
