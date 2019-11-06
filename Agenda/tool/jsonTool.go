package tool

import (
	"bytes"
	"encoding/json"
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

var Names Registers

func ReadJson(a interface{}, path string) {
	Bytes, err := ioutil.ReadFile(path)
	if err != nil {
		MakeLog("json read error", err.Error(), Files["sys"], false)
		os.Exit(1)
	}
	err = json.Unmarshal(Bytes, a)
	if err != nil {
		MakeLog("Unmarshal error", err.Error(), Files["sys"], false)
		os.Exit(1)
	}
}

func WriteJson(a interface{}, path string) {
	Bytes, err := json.Marshal(a)
	if err != nil {
		MakeLog("json error", err.Error(), Files["sys"], false)
		os.Exit(1)
	}
	var out bytes.Buffer
	err = json.Indent(&out, Bytes, "", "\t")
	if err != nil {
		MakeLog("json convert error", err.Error(), Files["sys"], false)
		os.Exit(1)
	}
	err = ioutil.WriteFile(path, out.Bytes(), 666)
	if err != nil {
		MakeLog("json write error", err.Error(), Files["sys"], false)
		os.Exit(1)
	}
}
