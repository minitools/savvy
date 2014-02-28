package main

import (
	"fmt"
	"log"
	"os/exec"
)

func Archive(destFilename string, dirPath string) {

	/* Archive directory */
	out, err := exec.Command("zip", "-r", destFilename, dirPath).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-----------------")
	fmt.Println(string(out))
	fmt.Println("-----------------")
}

func DestPath() string {
	return "/Users/Shared/BACKUP"
}