package tool

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var configurePath = "./resources/configure.json"

type File struct {
	Short string `json:"short"`
	Path string `json:"path"`
}

type Path struct {
	Files []File `json:"paths"`
}

var paths Path
var Paths = make(map[string]string)
var Files = make(map[string]*os.File)
var DefaultPassword = "000000"

func init() {
	readConfigure()
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
		message := GetMessage("file opened: ", Paths[key])
		MakeLog("file normal", message, Files["sys"], false)
	}
	ReadJson(&Names, Paths["user"])
}

func MakeLog(prefix, message string, writer io.Writer, output bool) {
	if writer == nil {
		panic(errors.New("file to write can't find"))
	}
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

func CreateFile(short, path string) {
	_, exists := Paths[short]
	if !exists {
		file, err := os.Create(path)
		if err != nil {
			MakeLog("create file error", err.Error(), Files["sys"], false)
			os.Exit(1)
		}
		Files[short] = file
		Paths[short] = path
		paths.Files = append(paths.Files, File{short, path})
		message := GetMessage("File created: ", path)
		MakeLog("file create", message, Files["sys"], false)
	}
}

func GetLogPath(name string) string {
	return "./log/" + name + ".log"
}
func GetJsonPath(name string) string {
	return "./resources/" + name + ".json"
}

func readConfigure() {
	ReadJson(&paths, configurePath)
	for _, value := range paths.Files {
		Paths[value.Short] = value.Path
	}
}

func WriteConfigure() {
	WriteJson(paths, configurePath)
}

func GetMessage(messages ...string) string {
	message := ""
	for _, str := range messages {
		message += str
	}
	return message
}

