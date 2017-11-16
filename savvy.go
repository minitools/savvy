package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type dirInfo struct {
	path       string
	path0      string
	info       os.FileInfo
	mostRecent time.Time
	doBackup   bool
}

var (
	topDirs  map[string]*dirInfo
	modified map[string]time.Time
	cache    *backupCache

	config *Config
)

func main() {

	var err error

	/* Get configuration options */
	config, err = flagsAndConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	cache.Load()

	fmt.Println("== schedule ================================")

	/* Traverse tree and collect modification times */
	topDirs = make(map[string]*dirInfo, 10)
	modified = make(map[string]time.Time)

	traverseLocal(walkAndGetTime)
	printStats()

	/* Get recent modified date from all files into directory
	   Schedule backups, marking top directories as necessary */
	for _, dir := range topDirs {
		dir.mostRecent = modified[dir.path0]

		previous, done := cache.Lookup(dir.path)
		if !done || previous.mostRecent.Before(dir.mostRecent) {
			scheduleBackup(dir)
		}
	}

	fmt.Println("== backup ===================================")

	/* Perform backups */
	for key, dir := range topDirs {
		if *flagVerbose {
			fmt.Println(key, " : ", dir)
		}

		if dir.doBackup {
			cache.StartBackup(dir)
			performBackup(dir)
			cache.EndBackup(dir)
		}
	}

}

func traverseLocal(doAction filepath.WalkFunc) {
	filepath.Walk(".", doAction)
}

func walkAndGetTime(path string, info os.FileInfo, err error) error {
	var dir *dirInfo
	var ok bool

	modTime := info.ModTime()
	seps := string(os.PathSeparator)

	splitPath := strings.SplitN(path, seps, 2)
	path0 := splitPath[0]

	/* if top-level directory... */
	if info.IsDir() && len(splitPath) == 1 {

		/* If directory contains a .savvyignore file, skip */
		fi, err := os.Stat(filepath.Join(path0, ".savvyignore"))
		if fi != nil && err == nil {
			log.Println("Directory contains \".savvyignore\", skipping :", path0)
			return nil
		}

		dir, ok = topDirs[path0]
		if !ok {
			dir = &dirInfo{}
			topDirs[path0] = dir
		}
		dir.path = path
		dir.path0 = path0
		dir.info = info

		if *flagVerbose {
			fmt.Println("TOP-LEVEL Dir :", splitPath, "mod:", modTime, "info:", info)

		}
	}

	/* for every file and directory, accumulate most recent modification time */
	cachedModTime, ok := modified[splitPath[0]]
	if !ok {

		modified[path0] = cachedModTime
		fmt.Println("Scanned ", path0)
	}
	if modTime.After(cachedModTime) {
		modified[path0] = modTime
	}

	return err
}

func printStats() {
	fmt.Println("----------------------------")

	for key, dir := range topDirs {
		if *flagVerbose {
			fmt.Println(key, " : ", dir)
		}

		if !dir.info.IsDir() {
			panic("Stored as dir, it's not actually a dir")
		}
	}
}

func scheduleBackup(dir *dirInfo) {
	fmt.Println("scheduled backup for", dir.path0)
	dir.doBackup = true
}

func performBackup(dir *dirInfo) {

	baseBackupDir := config.DestPath + "/"
	log.Println("backup: ", dir.path, " -> ", baseBackupDir)

	if flagNoOp != nil && *flagNoOp {
		log.Println("invoked with -n, skipping")
		return
	}

	/* TODO: rework filtering of directories */
	if dir.path == "." || dir.path == ".hg" {
		log.Println("skipping ", dir.path)
		return
	}

	if dir.doBackup != true || !dir.info.IsDir() {
		log.Panic("Unexpected backup")
	}

	/* Check existence of base directory every time.
	   TODO: optimize
	*/
	fi, err := os.Stat(baseBackupDir)
	if fi == nil && err != nil {
		log.Fatal("Destination directory does not exist :", baseBackupDir)
		return
	}

	destDate := dir.mostRecent.Format("20060102-150405")
	destFilename := baseBackupDir + dir.path0 + "-" + destDate + ".zip"

	/* Check if archive already exists */
	fi, err = os.Stat(destFilename)
	if fi != nil && err == nil {
		log.Println("Archive exists, skipping :", destFilename)
		return
	}

	fmt.Println("performing backup for", dir.path, " --> ", destFilename)

	Archive(destFilename, dir.path)
}
