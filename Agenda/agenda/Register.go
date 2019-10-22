package agenda

import (
	"Agenda/tool"
	"fmt"
	"os"
)

func Check(name, password string) bool {
	if name == "" {
		tool.MakeLog("format error", "empty name", tool.Files["reg"], true)
		return false
	}
	for _, user := range tool.Names.Users {
		if name == user.Name {
			message := tool.GetMessage("duplicate name: ", name)
			tool.MakeLog("repetition error", message, tool.Files["reg"], true)
			return false
		}
	}
	if password == tool.DefaultPassword {
		var check string
		fmt.Println("Sure to use the default password: " + tool.DefaultPassword + " ?(y/n)")
		_, err := fmt.Scanf("%s", &check)
		if err == nil {
			if check == "y" || check == "Y" {
				return true
			}else {
				tool.MakeLog("failed try", "fail to register", tool.Files["reg"], true)
				return false
			}
		}else {
			tool.MakeLog("input error", err.Error(), tool.Files["sys"], false)
			os.Exit(1)
		}
	}
	return true
}

func AddUser(name, password string) {
	user := tool.Users{Name: name, Password: password}
	tool.Names.Users = append(tool.Names.Users, user)
	tool.WriteJson(tool.Names, tool.Paths["user"])
	tool.CreateFile(name, tool.GetJsonPath(name))
	short := tool.GetMessage(name, "_l")
	tool.CreateFile(short, tool.GetLogPath(short))
	message := tool.GetMessage("add user: ", name, " succeed")
	tool.MakeLog("add user", message, tool.Files["reg"], true)
}
