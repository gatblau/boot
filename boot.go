/*
*    boot - Copyright (c) 2020 by www.gatblau.org
*
*    Licensed under the Apache License, Version 2.0 (the "License");
*    you may not use this file except in compliance with the License.
*    You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
*    Unless required by applicable law or agreed to in writing, software distributed under
*    the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
*    either express or implied.
*    See the License for the specific language governing permissions and limitations under the License.
*
*    Contributors to this project, hereby assign copyright in this code to the project,
*    to be licensed under the same terms as the rest of the code.
 */
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

	// loop through the specified configuration files
	for _, file := range files {
		merged := false
		// read the file content
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal().Msgf("cannot read file %s: %s", file, err)
		}
		content := string(bytes)
		// find all environment variable placeholders in the content
		vars := regex.FindAll(bytes, 1000)
		// loop though the found vars to merge
		for _, v := range vars {
			defValue := ""
			// removes placeholder marks: ${...}
			vname := strings.TrimSuffix(strings.TrimPrefix(string(v), "${"), "}")
			// is a default value defined?
			cut := strings.Index(vname, ":")
			// split default value and var name
			if cut > 0 {
				// get the default value
				defValue = vname[cut+1:]
				// get the name of the var without the default value
				vname = vname[0:cut]
			}
			// check the name of the env variable is not "PWD" as it can return the current directory in some OSs
			if vname == "PWD" {
				log.Fatal().Msgf("environment variable cannot be PWD, choose a different name")
			}
			// fetch the env variable value
			ev := os.Getenv(vname)
			// if the variable is not defined in the environment
			if len(ev) == 0 {
				// if no default value has been defined
				if len(defValue) == 0 {
					log.Fatal().Msgf("environment variable '%s' and/or default value not defined, skipping merging", vname)
				} else {
					// merge with the default value
					content = strings.Replace(content, string(v), defValue, 1000)
					merged = true
					log.Info().Msgf("merged placeholder %s with default value '%s'", string(v), defValue)
				}
			} else {
				// merge with the env variable value
				content = strings.Replace(content, string(v), ev, 1000)
				merged = true
				log.Info().Msgf("merged placeholder %s with value '%s'", string(v), ev)
			}
		}
		// if variables have been merged at all
		if merged {
			// override file with merged values
			err = WriteToFile(file, content)
			if err != nil {
				log.Fatal().Msgf("cannot update config file: %s", err)
			}
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
	log.Info().Msgf("'%v' bytes written to file '%s'", len(data), filename)
	return file.Sync()
}
