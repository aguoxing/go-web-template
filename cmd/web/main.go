package main

import (
	"fmt"
	"strings"
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

	str := "http://www.baidu.com"

	str1 := strings.Replace(str, "http://", "", 1)
	str2 := strings.Replace(str1, "https://", "", 1)
	str3 := strings.Replace(str2, "www.", "", 1)
	str4 := strings.Replace(str3, ".", "/", 1)
	fmt.Println(str4)
}
