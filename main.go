package main

import (
	"github.com/lixiang4u/aliyundrive-sdk/cmd"
	"github.com/lixiang4u/aliyundrive-sdk/utils"
	"log"
	"os"
	"strings"
)

func SubStr(html, key string) string {
	var findkey = "\"" + key + "\":\""
	var findindex = strings.Index(html, findkey)
	log.Println(findindex)
	if findindex < 0 {
		return ""
	}
	var find = html[findindex+len(findkey):]
	log.Println(find)
	return find[0:strings.Index(find, "\"")]
}

func main() {

	loginConf, err := utils.GetLoginFormConfig()
	if err != nil {
		log.Fatal(err)
	}
	utils.Login("18552072610", "111111", loginConf)
	if err != nil {
		log.Fatal(err)
	}

	//b, _ := json.MarshalIndent(loginConf, "", "	")
	//log.Printf("%s", string(b))

	return

	if err := cmd.Execute(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
