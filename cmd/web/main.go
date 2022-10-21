package main

import (
	"fmt"
	"go-web-template/util"
)

func main() {
	/*	menuMap := make(map[int64]system.SysMenu)
		var m []system.SysMenu
		menu := system.SysMenu{
			MenuID:   1,
			MenuName: "test",
		}

		for i := 0; i < 10; i++ {
			m = append(m, menu)
		}

		fmt.Println(m)

		for _, menu := range m {
			menuMap[menu.MenuID] = menu
		}*/
	str := util.FirstUpper("asdasd")
	fmt.Println(str)
}
