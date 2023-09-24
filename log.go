package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func setupLogging() {
	//setup our handy loggerino
	file, err := openLogFile("./logfile.txt")
	if err != nil {
		panic(err)
	}
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	log.SetFlags(log.Ltime | log.LUTC)
	log.Println("Logging started")

	// Make a messagebox with the logfile path
	fullpath := filepath.Join(wd, file.Name())
	fmt.Printf("Logging to file %s\n", fullpath)
}

func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func dumpStackTraces() {
	trace := make([]byte, 8192)
	length := runtime.Stack(trace, true)
	log.Println(string(trace[:length]))
}
