package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
)

type Config struct {
	DestPath string
	Includes []string
	Excludes []string
}

var (
	destPath = flag.String("dest", defaultDestPath(), "destination directory for all backed-up files")

	noConfigErr = errors.New("config file not present")
)

func flagsAndConfig() (*Config, error) {

	flag.Parse()

	cfg := &Config{}

	// 1) read configuration file, if present
	err := cfg.load(configPath())
	if err != nil {
		log.Println(err)
	}

	// 2) if destination path specified on command line,
	//    it overrides entry in configuration file
	if destPath != nil && *destPath != "" {
		cfg.DestPath = path.Clean(*destPath)

	}

	// 3) write back configuration
	//    (NOTE: overriden DestPath is saved)
	err = cfg.save(configPath())

	return cfg, err
}

func (cfg *Config) load(configPath string) error {

	// check existence of config file
	if _, err := os.Stat(configPath); err != nil {
		return noConfigErr
	}

	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	log.Printf("%s\n", string(file))

	err = json.Unmarshal(file, cfg)
	log.Printf("cfg.load results: %v\nErr: %v\n", cfg, err)

	return err
}

func (cfg *Config) save(configPath string) error {

	data, err := json.Marshal(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(configPath, data, os.ModePerm)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	return err
}

func configPath() string {

	// get user home dir
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("home=%q", usr.HomeDir)

	return path.Join(usr.HomeDir, "/.savvy.config")
}

func defaultDestPath() string {

	// get user home dir
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("home=%q", usr.HomeDir)

	return path.Join(usr.HomeDir, "BACKUP/")
}
