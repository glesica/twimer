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
const ElapsedFilename = "elapsed"

// TODO: Might make sense to make this thread safe as well
type Persister struct {
	elapsedPath string
}

func NewPersister() *Persister {
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
			log.Panicln("failed to create directory\n", directory, "\n", err)
		}
	}

	elapsedPath := path.Join(directory, ElapsedFilename)

	return &Persister{
		elapsedPath: elapsedPath,
	}
}

func (p *Persister) Elapsed() time.Duration {
	data, err := ioutil.ReadFile(p.elapsedPath)
	if err != nil {
		if os.IsNotExist(err) {
			p.SetElapsed(0)
			return p.Elapsed()
		}
		log.Panicln("failed to read file\n", p.elapsedPath, "\n", err)
	}

	duration, err := time.ParseDuration(string(data))
	if err != nil {
		log.Panicln("failed to parse duration\n", string(data), "\n", err)
	}

	return duration
}

func (p *Persister) SetElapsed(value time.Duration) {
	strData := fmt.Sprintf("%dns", value.Nanoseconds())
	data := []byte(strData)
	err := ioutil.WriteFile(p.elapsedPath, data, os.ModePerm)
	if err != nil {
		log.Panic("failed to write file\n", p.elapsedPath, "\n", err)
	}
}

func (p *Persister) ElapsedPath() string {
	return p.elapsedPath
}
