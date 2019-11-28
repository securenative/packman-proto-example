package main

import (
	"fmt"
	"github.com/securenative/GoProtobufReader/proto_reader"
	"github.com/securenative/packman/pkg"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type ReplyModel struct {
	Name        string
	Methods     map[string]*proto_reader.Method
	Messages    map[string]*proto_reader.Message
	Port        int32
	PackageName string
}

func main() {
	// Parse the flags:
	flags := packman.ReadFlags()
	protoPath := flags["proto"]
	packageName := flags[packman.PackageNameFlag]
	portString := flags["port"]
	port, err := strconv.ParseInt(portString, 10, 32)
	panicOnErr(err)

	// Read the protobuf file:
	protoContent := readFile(protoPath)

	// Parse the protobuf file:
	reader := proto_reader.NewReader()
	protoDef, err := reader.Read(protoContent)
	panicOnErr(err)

	// Copy the protobuf file to the pkg folder:
	err = ioutil.WriteFile(filepath.Join("..", "pkg", filepath.Base(protoPath)), []byte(protoContent), os.ModePerm)
	panicOnErr(err)

	// Run the protobuf compiler:
	pwd, err := os.Getwd()
	panicOnErr(err)
	pkgFolder := filepath.Join(filepath.Dir(pwd), "pkg")
	fileName := filepath.Base(protoPath)
	run(fmt.Sprintf("docker run -v %s:/defs namely/protoc-all -f %s -l go -o .", pkgFolder, fileName))

	// For the sake of the example, we only need the first service defined:
	srv := first(protoDef.Services)

	// We converting the protobuf lower-case field names to be capitalized:
	for _, v := range srv.Methods {
		for fk, _ := range v.Input.Fields {
			v.Input.Fields[fk].Name = strings.Title(v.Input.Fields[fk].Name)
		}

		for fk, _ := range v.Output.Fields {
			v.Output.Fields[fk].Name = strings.Title(v.Output.Fields[fk].Name)
		}
	}

	// Build the reply model:
	reply := ReplyModel{
		Name:        srv.Name,
		Methods:     srv.Methods,
		Messages:    protoDef.Messages,
		Port:        int32(port),
		PackageName: packageName,
	}

	// Write the reply
	packman.WriteReply(reply)
}

func readFile(path string) string {
	bytes, err := ioutil.ReadFile(path)
	panicOnErr(err)
	return string(bytes)
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func run(command string) {
	cmds := strings.Split(command, " ")
	cmd := exec.Command(cmds[0], cmds[1:]...)
	result, err := cmd.CombinedOutput()
	fmt.Println(string(result))
	panicOnErr(err)
}

func first(m map[string]*proto_reader.Service) *proto_reader.Service {
	for _, v := range m {
		return v
	}
	panic("the protobuf file need to define exactly one service")
}
