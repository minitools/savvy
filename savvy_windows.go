package main

import (
	"fmt"
	"os/exec"
	"log"
)

/*
type tool struct {

}


func find(execName string, path string) {

}

func DetectTools() []tool {
	
	// find one of 7zip, zip in Program Files directory
	exe := find("7zip.exe", "c:\Program Files")
	if exe != nil {
		return []
	}

}
*/

func Archive(destFilename string, dirPath string) {
	out, err := exec.Command("c:/Program Files/7-Zip/7z.exe", "a", "-r", destFilename, dirPath).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-----------------")
	fmt.Println(string(out))
	fmt.Println("-----------------")

}

func DestPath() string {
	return "/BACKUP"
}