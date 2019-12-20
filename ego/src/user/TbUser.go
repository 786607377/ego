package user

type TbUser struct{
    //属性首字母大写:1/要转换为Json 2.可能出现挎包访问
    Id int64
    Username string
    Password string
    Phone string
    Email string
    Greate string
    Update string
}