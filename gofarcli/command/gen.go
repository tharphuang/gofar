package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/TharpHuang/gofar/tools/text"
)

var protoLines = `syntax = "proto3";
package proto;

//Todo change the name of messages and service 

message $Request{
  string msg = 1;
}

message $Response{
  string response_msg = 1;
}

service  $Service{
  rpc Transfer($Request) returns ($Response){}
}
`

var databaseLines = `package migration



`

var warningGen = text.TrimLeft(`
Gofar gen: no job tyep input.
Use "gofar gen <type> [arguments]"
The type are:
	proto	protobuf files
	migrate	databases migrateion files`)

func GenerateFile(jobType, fileName string) {
	if jobType == "" {
		fmt.Println(warningGen)
		return
	}

	switch jobType {
	case "proto":
		protoFile(fileName)
	case "migration":
		migrationFile(fileName)
	default:
		fmt.Println(warningGen)
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

func protoFile(fileName string) {
	protoPath, err := filepath.Abs("proto")

	_, err = os.Stat(protoPath)
	if os.IsNotExist(err) {
		os.Mkdir(protoPath, 0777)
		os.Chmod(protoPath, 0777)
	}

	filePath := filepath.Join(protoPath, fileName+".proto")

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}

	lines := strings.Replace(protoLines, "$", formatName(fileName), -1)

	_, err = f.Write([]byte(lines))

	if err != nil {
		panic(err)
	}

	f.Close()

	fmt.Println("create proto file success!")
}

func migrationFile(fileName string) {
	dbPath, err := filepath.Abs("migration")

	_, err = os.Stat(dbPath)
	if os.IsNotExist(err) {
		os.Mkdir(dbPath, 0777)
		os.Chmod(dbPath, 0777)
	}

	timeStr := strconv.FormatInt(time.Now().Unix(), 10)
	filePath := filepath.Join(dbPath, fileName+"_"+timeStr+".go")

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}

	lines := strings.Replace(databaseLines, "$", formatName(fileName), -1)

	_, err = f.Write([]byte(lines))

	if err != nil {
		panic(err)
	}

	f.Close()

	fmt.Println("Created Migration： " + filePath)

}
