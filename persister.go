package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

const TwimerDirectory = ".twimer"
const StateFilename = "state"
const ElapsedFilename = "elapsed"

type FormatCallback = func(duration time.Duration) string

type Persister struct {
	statePath   string
	elapsedPath string
	formatter   FormatCallback
}

func NewPersister(formatter FormatCallback) *Persister {
	var directory string

	home, err := os.UserHomeDir()
	if err != nil {
		directory = TwimerDirectory
	} else {
		directory = path.Join(home, TwimerDirectory)
	}

	err = os.Mkdir(directory, os.ModeDir|os.ModePerm)
	if err != nil {
		if !os.IsExist(err) {
			log.Panicln("failed to create data directory\n", directory, "\n", err)
		}
	}

	statePath := path.Join(directory, StateFilename)
	elapsedPath := path.Join(directory, ElapsedFilename)

	if formatter == nil {
		formatter = func(duration time.Duration) string {
			return duration.String()
		}
	}

	return &Persister{
		statePath:   statePath,
		elapsedPath: elapsedPath,
		formatter:   formatter,
	}
}

func (p *Persister) Elapsed() time.Duration {
	data, err := ioutil.ReadFile(p.statePath)
	if err != nil {
		if os.IsNotExist(err) {
			p.SetElapsed(0)
			return p.Elapsed()
		}
		log.Panicln("failed to read file\n", p.statePath, "\n", err)
	}

	duration, err := time.ParseDuration(string(data))
	if err != nil {
		log.Panicln("failed to parse duration\n", string(data), "\n", err)
	}

	return duration
}

func (p *Persister) SetElapsed(value time.Duration) {
	var err error

	stateData := []byte(fmt.Sprintf("%dns", value.Nanoseconds()))
	err = ioutil.WriteFile(p.statePath, stateData, os.ModePerm)
	if err != nil {
		log.Panicln("failed to write file\n", p.statePath, "\n", err)
	}

	elapsedData := []byte(p.formatter(value))
	err = ioutil.WriteFile(p.elapsedPath, elapsedData, os.ModePerm)
	if err != nil {
		log.Panicln("failed to write file\n", p.elapsedPath, "\n", err)
	}
}

func (p *Persister) ElapsedPath() string {
	return p.elapsedPath
}
