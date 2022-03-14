package main

import (
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)
var relPath = "resourcesFolder path"
var configPath = "server.cfg path"


var totalResources = 50000
func main()  {
	for i := 0; totalResources > i; i++ {
		createResource()
	}
}

func createResource() {
	dirName, err := uuid.NewUUID()
	
	err = os.Mkdir(relPath + dirName.String(), 0755)
	if err != nil {
		panic(err.Error())
	}
	
	source, err := os.Open("resource/fxmanifest.lua")
	if err != nil {
		log.Println(err.Error())
	}
	
	defer source.Close()
	
	destination, err := os.Create(relPath + dirName.String() + "/fxmanifest.lua")
	if err != nil {
		panic(err.Error())
	}
	
	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		panic(err.Error())
	}
	
	sourceClient, err := os.Open("resource/client.lua")
	if err != nil {
		log.Println(err.Error())
	}
	
	defer sourceClient.Close()
	
	destinationClient, err := os.Create(relPath + dirName.String() + "/client.lua")
	if err != nil {
		panic(err.Error())
	}
	
	defer destinationClient.Close()
	_, err = io.Copy(destinationClient, sourceClient)
	if err != nil {
		panic(err.Error())
	}
	
	input, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	
	newLine := "ensure " + dirName.String()
	
	lines := strings.Split(string(input), "\n")
	ls := append(lines, newLine)
	
	output := strings.Join(ls, "\n")
	err = ioutil.WriteFile(configPath, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}