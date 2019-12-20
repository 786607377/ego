package cat

import (
    "ego/ego/src/commons"
    "fmt"
)

func selByIdDao(id int)*TbItemCat{
    rows,err:=commons.Dql("select * from tb_item_cat where id=?",id)
    defer commons.CloseConn()
    if err!=nil{
        fmt.Println(err)
        return nil
    }
    if rows.Next(){
        t:=new(TbItemCat)
        rows.Scan(&t.Id,&t.ParentId,&t.Name,&t.Status,&t.SortOrder,&t.IsParent,&t.Created,&t.Updated)
        return t
    }
    return nil
}

/*
根据parent_id查询所有子类目
*/
func selByPid(pid int)(c []TbItemCat){
    rows,err:=commons.Dql("select * from tb_item_cat where parent_id=?",pid)
    defer commons.CloseConn()
    if err!=nil{
        fmt.Println(err)
        return nil
    }
    c = make([]TbItemCat,0)
    for rows.Next(){
        var t TbItemCat
        rows.Scan(&t.Id,&t.ParentId,&t.Name,&t.Status,&t.SortOrder,&t.IsParent,&t.Created,&t.Updated)
        c=append(c,t)
    }
    return
}
