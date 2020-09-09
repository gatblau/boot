package main

import (
	"github.com/rs/zerolog/log"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	files := os.Args[1:]
	regex, err := regexp.Compile("\\${(?P<NAME>.*)}")
	if err != nil {
		log.Fatal().Msgf("cannot compile regex: %s", err)
	}
	if len(files) == 0 {
		log.Fatal().Msgf("you must provide files to merge!")
	}

	for _, file := range files {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal().Msgf("cannot read file %s: %s", file, err)
		}
		content := string(bytes)

		vars := regex.FindAll(bytes, 1000)
		for _, v := range vars {
			vname := strings.TrimSuffix(strings.TrimPrefix(string(v), "${"), "}")
			ev := os.Getenv(vname)
			if len(ev) == 0 {
				log.Error().Msgf("environment variable %s not defined", vname)
			} else {
				content = strings.Replace(content, string(v), ev, 1000)
			}
		}
		err = WriteToFile(file, content)
		if err != nil {
			log.Error().Msgf("cannot update config file: %s", err)
		}
	}
}

// checking for errors and syncing at the end.
func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
