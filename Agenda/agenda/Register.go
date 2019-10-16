package agenda

import (
	"Agenda/tool"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	bytes, err := ioutil.ReadFile("./resources/user.json")
	if err != nil {
		tool.MakeLog("json read error", err.Error(), tool.Files["sys"], false)
		os.Exit(1)
	}
	err = json.Unmarshal(bytes, &names)
	if err != nil {
		tool.MakeLog("Unmarshal error", err.Error(), tool.Files["sys"], false)
		os.Exit(1)
	}
}

func ShowMessage(name, password string) {
	fmt.Println(name + ", " + password)
}

func Check(name string) bool {
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
	return true
}

func AddUser(name, password string) {
	user := Users{Name: name, Password: password}
	names.Users = append(names.Users, user)
	bytes, err := json.Marshal(names)
	if err != nil {
		tool.MakeLog("json error", err.Error(), tool.Files["sys"], false)
		os.Exit(1)
	}
	err = ioutil.WriteFile(tool.Paths["user"], bytes, 666)
	if err != nil {
		tool.MakeLog("json write error", err.Error(), tool.Files["sys"], false)
		os.Exit(1)
	}
	message := "add user: "
	message += name + " succeed"
	tool.MakeLog("add user", message, tool.Files["reg"], true)
}
