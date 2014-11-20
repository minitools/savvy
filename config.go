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
	Excludes []string
	Recurses []string
}

var (
	flagDestPath = flag.String("dest", "", "destination directory for all backed-up files")
	flagNoOp     = flag.Bool("n", false, "don't perform backup operations (dry run)")
	flagVerbose  = flag.Bool("v", false, "generate verbose output")

	errNoConfig        = errors.New("config file not present")
	errConfigCorrupted = errors.New("config file corrupted, not usable")
	errInvalidConfig   = errors.New("config: expected JSON, got invalid format")
)

func flagsAndConfig() (*Config, error) {

	flag.Parse()

	cfg := &Config{}

	// 1) read configuration file, if present
	cfg, err := load(configPath())
	if err != nil {
		log.Println(err)
	}

	// 2) if destination path specified on command line,
	//    it overrides entry in configuration file
	log.Println("flagDestPath:", *flagDestPath)
	if flagDestPath != nil && *flagDestPath != "" {
		cfg.DestPath = path.Clean(*flagDestPath)

	}

	// 3) check backup destination
	// if not present, exit
	if !exists(cfg.DestPath) {
		fmt.Println("BACKUP DESTINATION DOES NOT EXIST: '", cfg.DestPath, "'")
		os.Exit(-1)
	}

	// 3) write back configuration
	//    (NOTE: overriden DestPath is saved)
	err = cfg.save(configPath())

	return cfg, err
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func load(configPath string) (*Config, error) {

	// check existence of config file
	if _, err := os.Stat(configPath); err != nil {
		return nil, errNoConfig
	}

	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, errConfigCorrupted
	}

	log.Printf("file: %s\n", string(file))

	cfg := &Config{}
	err = json.Unmarshal(file, cfg)
	if err != nil {
		return nil, errInvalidConfig
	}

	log.Printf("cfg.load results: %v\nErr: %v\n", cfg, err)

	return cfg, err
}

func (cfg *Config) isValid() bool {
	_, err := os.Stat(cfg.DestPath)
	return err == nil
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
