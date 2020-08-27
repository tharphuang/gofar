package command

import (
	"fmt"
	"os"
	"strings"
)

var lines = `syntax = "proto3";
package proto;

//Todo change the name of messages and service 

message $Request{
  string msg = 1;
}

message $Response{
  string return_msg = 1;
}

service  $Service{
  rpc Transfer($Request) returns ($Response){}
}
`

func GenerateFile(fileType, fileName string) {
	switch fileType {
	case "proto":
		createProtoFile(fileName)
	default:
		fmt.Println("No such file type, please try again.")
	}

}

func formatName(fileName string) string {
	strings.ToLower(fileName)
	nameChars := []byte(fileName)
	if 'a' <= nameChars[0] || nameChars[0] <= 'z' {
		nameChars[0] = fileName[0] - 'a' + 'A'
	}

	return string(nameChars)
}

func createProtoFile(filename string) {
	_, err := os.Stat("../proto")
	if os.IsNotExist(err) {
		os.Mkdir("../proto", 0777)
		os.Chmod("../proto", 0777)
	}

	filePath := "../proto/" + filename + ".proto"

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}

	lines = strings.Replace(lines, "$", formatName(filename), -1)

	_, err = f.Write([]byte(lines))

	if err != nil {
		panic(err)
	}

	f.Close()

	fmt.Println("create proto file success!")
}
