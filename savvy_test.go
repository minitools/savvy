package main

import (
	"os"
	"os/exec"
	"testing"
)

/* Tests for savvy backup utility */

var safeEnvironment = false

func safeToRun(t *testing.T) {
	if !safeEnvironment {
		t.Fatal("unsafe env")
	}
}

func Test000_SetupDirs(t *testing.T) {

	// make dir
	err := os.Mkdir("test/", os.ModeDir|0777)
	//if err != nil && err != os.ErrExist {
	//	t.Fatal(err)
	//}

	err = os.Chdir("test")
	if err != nil {
		t.Fatal(err)
	}

	wd, err2 := os.Getwd()
	if err2 == nil {
		safeEnvironment = true
	}
	t.Log("Working dir:", wd)
}

/*
	Empty directory can be backed up.
*/
func Test001_Empty(t *testing.T) {

	safeToRun(t)

	err := os.Mkdir("a00", os.ModeDir|0777)
	err = err
	//if err != nil {
	//	t.Fatal(err)
	//}

	// run savvy on it
	out, err2 := exec.Command("../savvy").Output()

	if err2 != nil {
		t.Fatal(err2)
	}

	t.Log(string(out))
	err2 = err2

	// verify results
	//t.Error("not implemented\n")

	//os.RemoveAll("*")
}

/*
	Directory with small number of files can be backed up.
*/
func Test002_Simple(t *testing.T) {
	t.Error("not implemented\n")

	safeToRun(t)
}

/*
	Directory with names containing spaces & special characters can be backed.
*/
func Test003_SpecialCharName(t *testing.T) {
	t.Error("not implemented\n")

	safeToRun(t)
}

/*
	Directory with "savvyignore" inside is actually ignored.
*/
func Test004_Ignore(t *testing.T) {
	t.Error("not implemented\n")

	safeToRun(t)
}

/*
	Way to set destination directory is working.
	1. Set destination directory:  savvy --dest <dir>
	2. Invoke savvy as usual
*/
func Test005_DestinationDirectory(t *testing.T) {
	t.Error("not implemented\n")

	safeToRun(t)

	// remove any .savvy file

	// make dir and populate

	// invoke savvy to specify dest directory (--dest <dir> )

	// verify that .savvy file is created

	// invoke savvy

	// verify that archive is created into dest directory
}

/*
	Cache of backed-up files is working correctly.
	(Useful for write-only destination directories)
*/
func Test006_CacheIsWorking(t *testing.T) {
	t.Error("not implemented\n")

	safeToRun(t)

	// remove any .savvy file

	// make dir and populate

	// invoke savvy to specify dest directory (--dest <dir> )

	// verify that .savvy file is created

	// invoke savvy

	// verify that archive is created into dest directory as expected
}
