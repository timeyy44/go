package agenda

import (
	"Agenda/tool"
	"fmt"
	"os"
)

type Users struct{
	Name string `json:"name"`
	Password string `json:"password"`
}

type Registers struct {
	 Users []Users `json:"users"`
}

var names Registers

func init() {
	tool.ReadJson(&names, tool.Paths["user"])
}

func Check(name, password string) bool {
	if name == "" {
		tool.MakeLog("format error", "empty name", tool.Files["reg"], true)
		return false
	}
	for _, user := range names.Users {
		if name == user.Name {
			message := "duplicate name: "
			message += name
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
	user := Users{Name: name, Password: password}
	names.Users = append(names.Users, user)
	tool.WriteJson(names, tool.Paths["user"])
	tool.CreateFile(name, tool.GetJsonPath(name))
	short := name
	short += "_l"
	tool.CreateFile(short, tool.GetLogPath(short))
	message := "add user: "
	message += name + " succeed"
	tool.MakeLog("add user", message, tool.Files["reg"], true)
}
