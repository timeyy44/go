package tool

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadJson(a interface{}, path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		MakeLog("json read error", err.Error(), Files["sys"], false)
		os.Exit(1)
	}
	err = json.Unmarshal(bytes, a)
	if err != nil {
		MakeLog("Unmarshal error", err.Error(), Files["sys"], false)
		os.Exit(1)
	}
}

func WriteJson(a interface{}, path string) {
	bytes, err := json.Marshal(a)
	if err != nil {
		MakeLog("json error", err.Error(), Files["sys"], false)
		os.Exit(1)
	}
	err = ioutil.WriteFile(path, bytes, 666)
	if err != nil {
		MakeLog("json write error", err.Error(), Files["sys"], false)
		os.Exit(1)
	}
}
