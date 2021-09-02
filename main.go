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
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var parallels = flag.Bool("parallels", false, "parallels make projects")

func main() {
	flag.Parse()

	config, err := os.Open("release-config.json")
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

	if *parallels {
		parallelsMakeAll(projects)
	} else {
		serialMakeAll(projects)
	}
}

func gitCloneAll(project []project) error {
	for _, p := range project {
		if p.SkipClone {
			continue
		}
		if err := p.gitClone(); err != nil {
			return fmt.Errorf("[%s] git clone %s@%s failed: %v", p.Name, p.Repo, p.Branch, err)
		}
	}

	return nil
}

func parallelsMakeAll(projects []project) {
	w := &sync.WaitGroup{}
	w.Add(len(projects))

	for _, p := range projects {
		go func(project project) {
			if err := project.make(); err != nil {
				log.Printf("[%s] exec: %s failed; %v \n", project.Name, project.Exec, err)
			}
			w.Done()
		}(p)
	}

	w.Wait()
}

func serialMakeAll(projects []project) {
	for _, p := range projects {
		if err := p.make(); err != nil {
			log.Printf("[%s] exec: %s failed; %v \n", p.Name, p.Exec, err)
			continue
		}
	}
}
