package main

/* TO DO:
   1. finish traversal by 
   - recording last modification time for top-level directory (accumulate in map)
   - add top-level dir to list of directories to backup, if backup criteria met

   2. perform backup to fixed location
*/

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type dirInfo struct {
	path 		string
	path0		string
	info 		os.FileInfo
	mostRecent	time.Time
	doBackup	bool
}

var (
	oldest		time.Time = time.Now()
	mostRecent	time.Time

	topDirs 	map[string]*dirInfo
	modified	map[string]time.Time
	cache 		*backupCache
)

func main() {

	/* Get configuration options */
	parseFlags()

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
		dir.mostRecent = modified[ dir.path0 ]

		previous, done := cache.Lookup(dir.path)
		if !done || previous.mostRecent.Before(dir.mostRecent) {
			scheduleBackup(dir)
		}
	}

	fmt.Println("== backup ===================================")

	/* Perform backups */
	for key, dir := range topDirs {
		fmt.Println(key, " : ", dir)
		if dir.doBackup {
			cache.StartBackup(dir)
			performBackup(dir)
			cache.EndBackup(dir)
		}
	}

}

func parseFlags() {
	/* nothing here */
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

		dir, ok = topDirs[ path0 ]
		if !ok {
			dir = &dirInfo{}
			topDirs[ path0 ] = dir
		}
		dir.path = path
		dir.path0= path0
		dir.info = info

		fmt.Println("TOP-LEVEL Dir :", splitPath, "mod:", modTime, "info:", info)
	}


	/* for every file and directory, accumulate most recent modification time */
	cachedModTime, ok := modified[splitPath[0]]
	if !ok {
		
		modified[ path0 ] = cachedModTime
		fmt.Println("Added ", path0 )
	}
	if modTime.After(cachedModTime) {
		modified[ path0 ] = modTime
	}

	return err
}


func printStats() {
	fmt.Println("----------------------------")

	for key, dir := range topDirs {
		fmt.Println(key, " : ", dir)
		if ! dir.info.IsDir() {
			panic("Stored as dir, it's not actually a dir")
		}
	}
}

func scheduleBackup(dir *dirInfo) {
	fmt.Println("scheduled backup for", dir.path0)
	dir.doBackup = true
}

func performBackup(dir *dirInfo) {
	fmt.Println("performing backup for", dir)
	if dir.doBackup != true || ! dir.info.IsDir() {
		panic("Unexpected backup")
	}
}
