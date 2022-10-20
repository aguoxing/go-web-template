package main

import (
	"encoding/json"
	"fmt"
	"go-web-template/app/model/system"
	"go-web-template/util"
)

func main() {
	p, e := util.PasswordHash("123456")
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(p)

	test := "{\"action\":\"reponse_common_policy_config\",\"retCode\":0,\"common_policy_config\":[[3,15],\"1234\"]}"

	var config *system.SysUser

	err := json.Unmarshal([]byte(test), &config)

	if err != nil {
		//logs.Info(err)
	}
}
