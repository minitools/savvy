package main

import (
	"log"
	"os"
	"os/exec"
	"testing"
)

/* Tests for savvy backup utility */

var safeEnvironment = false

func setup(t *testing.T) {
	// TODO: clean up existing dir

	// make dir
	err := os.Mkdir("test/", os.ModeDir|0777)
	//if err != nil && err != os.ErrExist {
	//	t.Fatal(err)
	//}

	err = os.Chdir("test")
	if err != nil {
		log.Fatal(err)
	}

	wd, err := os.Getwd()
	if err == nil {
		safeEnvironment = true
	}
	log.Println("Working dir:", wd)
}

func teardown(t *testing.T) {

}

/*
	Empty directory can be backed up.
*/
func Test001_Empty(t *testing.T) {

	setup()
	defer teardown()

	err := os.Mkdir("a00", os.ModeDir|0777)
	if err != nil {
		t.Fatal(err)
	}

	// run savvy on it
	out, err := exec.Command("./savvy").Output()

	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(out))

	// verify results
	//t.Error("not implemented\n")

	//os.RemoveAll("*")
}

/*
	Directory with small number of files can be backed up.
*/
func Test002_Simple(t *testing.T) {
	t.Error("not implemented\n")

	setup(t)
	defer teardown(t)
}

/*
	Directory with names containing spaces & special characters can be backed.
*/
func Test003_SpecialCharName(t *testing.T) {
	t.Error("not implemented\n")

	setup(t)
	defer teardown(t)
}

func Test_SymbolicLink(t *testing.T) {
	t.Error("not implemented\n")
}

/*
	Directory with "savvyignore" inside is actually ignored.
*/
func Test004_Ignore(t *testing.T) {
	t.Error("not implemented\n")

	setup(t)
	defer teardown(t)
}

/*
	Way to set destination directory is working.
	1. Set destination directory:  savvy --dest <dir>
	2. Invoke savvy as usual
*/
func Test005_DestinationDirectory(t *testing.T) {
	t.Error("not implemented\n")

	setup(t)
	defer teardown(t)

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

	setup(t)
	defer teardown(t)

	// remove any .savvy file

	// make dir and populate

	// invoke savvy to specify dest directory (--dest <dir> )

	// verify that .savvy file is created

	// invoke savvy

	// verify that archive is created into dest directory as expected
}
