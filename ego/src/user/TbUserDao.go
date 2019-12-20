package user

import(
    "ego/ego/src/commons"
    "fmt"
)
//根据用户名和密码做查询,如果返回值为nil表示查询失败否则成功
func SelByUnPwdDao(un,pwd string)*TbUser{
    sql:="select *from tb_user where username=? and password=? or email=? and password=? or phone=? and password=?"
    rows,err:=commons.Dql(sql,un,pwd,un,pwd,un,pwd)
    defer commons.CloseConn()
    if err!=nil{
        fmt.Println(err)
        return nil
    }
    if rows.Next(){
        user:=new(TbUser)
        rows.Scan(&user.Id,&user.Username,&user.Password,&user.Phone,&user.Email,&user.Greate,&user.Update)
        return user
    }
    return nil
}