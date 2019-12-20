package item

import(
    "ego/ego/src/commons"
    "ego/ego/src/item/cat"
    "strings"
)

func showItemService(rows,page int)(e *commons.Datagrid){
  ts:=selByPageDao(rows,page)
  if ts!=nil{
      itemChildren:=make([]TbItemChild,0)
      for _,v:=range ts{
         var itemchild TbItemChild
         //itemchild.Id=v.Id
         //itemchild.Title=v.Title
         //itemchild.Barcode=v.Barcode
         //itemchild.Num=v.Num
         //itemchild.Image=v.Image
         //itemchild.SellPoint=v.SellPoint
         //itemchild.Cid=v.Cid
         //itemchild.Price=v.Price
         //itemchild.Status=v.Status
         //itemchild.Created=v.Created
         //itemchild.Updated=v.Updated
         itemchild.TbItem=v
         itemchild.CategoryName=cat.ShowCatByIdService(v.Cid).Name
         itemChildren = append(itemChildren,itemchild)
      }
      e=new(commons.Datagrid)
      e.Rows= itemChildren
      e.Total=selCount()
      return
  }
  return nil
}

//删除商品
func delByIdsService(ids string)(e commons.EgoResult){
    count:=updStatusByIdsDao(strings.Split(ids,","),3)
    if count>0{
        e.Status=200
    }
    return
}
//上架商品
func inStockService(ids string)(e commons.EgoResult){
    count:=updStatusByIdsDao(strings.Split(ids,","),1)
    if count>0{
        e.Status=200
    }
    return
}
//下架商品
func offStockService(ids string)(e commons.EgoResult){
    count:=updStatusByIdsDao(strings.Split(ids,","),2)
    if count>0{
        e.Status=200
    }
    return
}