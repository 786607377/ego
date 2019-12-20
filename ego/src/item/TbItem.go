package item

//商品
type TbItem struct{
    Id int
    Title string
    SellPoint string
    Price int
    Num int
    Barcode string
    Image string
    Cid int
    Status int8
    Created string
    Updated string
}


//给页面使用实现商品的类目
type TbItemChild struct{
    TbItem
    CategoryName string
}