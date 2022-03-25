package utils

// Menu 菜单
type Menu struct {
	Id int
	Pid int
	CateName string
	Desc string
	CreateTime int
	Children []Menu
}

// TreeList 菜单
type TreeList struct {
	Id int
	Pid int
	CateName string
	Desc string
	CreateTime int
	Children []TreeList
}

// GetMenu 获取菜单
func GetMenu(menuList []Menu,pid int)(treeList []TreeList) {
	for _, v := range menuList {
		if v.Pid == pid {
			child := GetMenu(menuList,v.Id)
			node := TreeList {
				Id: v.Id,
				Pid: v.Pid,
				CateName: v.CateName,
				Desc: v.Desc,
				CreateTime: v.CreateTime,
			}
			//node.Children = child
			treeList = append(treeList,node)
			treeList = append(treeList,child...)
		}
	}
	return treeList
}


