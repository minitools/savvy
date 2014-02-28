package main

/* cache of backups performed in the current directory */
type backupCache struct {
	previous map[string]dirInfo
}

func (bk *backupCache) Load() {

}

func (bk *backupCache) Lookup(path string) (*dirInfo, bool) {
	return nil, false
}

func (bk *backupCache) Add(dir *dirInfo) {

}

func (bk *backupCache) StartBackup(dir *dirInfo) {

}

func (bk *backupCache) EndBackup(dir *dirInfo) {

}
