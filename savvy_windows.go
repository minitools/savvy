package main

import (
	"fmt"
	"log"
	"os/exec"
)

func Archive(destFilename string, dirPath string) {

	/* Archive directory */
	out, err := exec.Command("c:/Program Files/7-Zip/7z.exe", "a", "-r", destFilename, dirPath).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-----------------")
	fmt.Println(string(out))
	fmt.Println("-----------------")

}
