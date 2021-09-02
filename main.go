/*
Copyright 2021 KubeCube Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const configName = "release-config.json"

func main() {
	config, err := os.Open(configName)
	if err != nil {
		panic(err)
	}

	defer config.Close()

	configBytes, err := ioutil.ReadAll(config)
	if err != nil {
		panic(err)
	}

	projects := make([]project, 0)

	err = json.Unmarshal(configBytes, &projects)
	if err != nil {
		panic(err)
	}

	err = gitCloneAll(projects)
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(1)
	}

	err = makeAll(projects)
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(1)
	}
}

func gitCloneAll(project []project) error {
	for _, p := range project {
		if p.Skip {
			continue
		}
		if err := p.gitClone(); err != nil {
			return fmt.Errorf("[%s] git clone %s@%s failed: %v", p.Name, p.Repo, p.Branch, err)
		}
	}

	return nil
}

// todo(weilaaa): parallels exec suppport
func makeAll(projects []project) error {
	for _, p := range projects {
		if p.Skip {
			continue
		}
		if err := p.make(); err != nil {
			return fmt.Errorf("[%s] exec: %s failed; %v, ", p.Name, p.Exec, err)
		}
	}

	return nil
}
