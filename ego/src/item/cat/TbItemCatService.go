package cat

import "ego/ego/src/commons"

//根据id返回数据
func ShowCatByIdService(id int)*TbItemCat{
    return selByIdDao(id)
}

func showCatByPid(pid int)(tree []commons.EasyUITree){
    cats:=selByPid(pid)
    tree=make([]commons.EasyUITree,0)
    for _,n:=range cats{
        state:="open"
        if n.IsParent!=0{
            state="closed"
        }
        tree=append(tree,commons.EasyUITree{n.Id,n.Name,state})
    }
    return
}