package main

import (
	"testing"
)

/* Tests for savvy backup utility */

/* 
	Empty directory can be backed up.
*/
func Test001_Empy(t *testing.T) {
	
	// make dir
	// run savvy on it
	// verify results
	t.Error("not implemented\n")
}

/*
	Directory with small number of files can be backed up. 
*/
func Test002_Simple(t *testing.T) {
	t.Error("not implemented\n")	
}

/* 
	Directory with names containing spaces & special characters can be backed.
*/
func Test003_SpecialCharName(t *testing.T) {
	t.Error("not implemented\n")
}

/*
	Directory with "savvyignore" inside is actually ignored.
*/
func Test004_Ignore(t *testing.T) {
	t.Error("not implemented\n")
}

/*
	Way to set destination directory is working.
	1. Set destination directory:  savvy --dest <dir>
	2. Invoke savvy as usual
*/
func Test005_DestinationDirectory(t *testing.T) {
	t.Error("not implemented\n")

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

	// remove any .savvy file

	// make dir and populate

	// invoke savvy to specify dest directory (--dest <dir> )

	// verify that .savvy file is created

	// invoke savvy

	// verify that archive is created into dest directory as expected
}
