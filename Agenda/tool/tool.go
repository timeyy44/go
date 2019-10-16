package tool

import (
	"fmt"
	"io"
	"log"
	"os"
)

var Paths = map[string]string{"sys":"./log/sys.log", "reg":"./log/register.log", "user":"./resources/user.json"}
var Files = make(map[string]*os.File)

func init() {
	for key, value := range Paths {
		file, err := os.OpenFile(value, os.O_RDWR | os.O_APPEND, 666)
		if err != nil {
			if os.IsNotExist(err) {
				file, err = os.Create(value)
			}else {
				if key == "sys" {
					panic(err)
				}else {
					MakeLog("file error", err.Error(), Files["sys"], false)
					os.Exit(1)
				}
			}
		}
		Files[key] = file
		message := "file opened: "
		message += Paths[key]
		MakeLog("file normal", message, Files["sys"], false)
	}
}

func MakeLog(prefix, message string, writer io.Writer, output bool) {
	prefix = fmt.Sprintf("[%s] ", prefix)
	log.SetPrefix(prefix)
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.SetOutput(writer)
	log.Println(message)
	if output {
		fmt.Println(prefix + message)
	}
}

func CloseFiles() {
	for _, value := range Files {
		err := value.Close()
		if err != nil {
			panic(err)
		}
	}
}
