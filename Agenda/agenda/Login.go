package agenda

import (
	"Agenda/tool"
	"fmt"
	"os"
)

var LName string
//var LPassword string

func CheckLogin(name, password string) bool {
	if name == "" {
		tool.MakeLog("format error", "login with empty name", tool.Files["login"], true)
		return false
	}
	var num = -1
	for n, user := range tool.Names.Users {
		if name == user.Name {
			num = n
			break
		}
	}
	if num == -1 {
		tool.MakeLog("login error", "login with wrong name", tool.Files["login"], true)
		return false
	}
	if password == "_default_" {
		var pass string
		fmt.Println("Enter your password: ")
		_, err := fmt.Scanf("%s", &pass)
		if err == nil {
			password = pass
		}else {
			tool.MakeLog("input error", err.Error(), tool.Files["sys"], false)
			os.Exit(1)
		}
	}
	if tool.Names.Users[num].Password == password {
		message := tool.GetMessage("login with name: ", name, " succeed")
		tool.MakeLog("login normal", message, tool.Files["login"], false)
		LName = name
		//LPassword = password
		return true
	}else {
		tool.MakeLog("login fail", "login with wrong password", tool.Files["login"], true)
		return false
	}
}

func GotoAgenda() {
	fmt.Println("Welcome back, " + LName)
}
