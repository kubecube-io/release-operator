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

import "fmt"

const localRepos = "./repos/"

type project struct {
	Name   string `json:"name"`
	Repo   string `json:"repo"`
	Branch string `json:"branch"`

	SkipClone bool `json:"skip_clone"`
	SkipMake  bool `json:"skip_make"`

	// exec do cmd of image build
	Exec string `json:"exec"`
}

// gitClone do "git clone" according to repo and tag
func (p project) gitClone() error {
	cmd := fmt.Sprintf("git clone -b %s %s %s", p.Branch, p.Repo, localRepos+p.Name)
	return doExec(cmd, "")
}

// make do exec with project to build image
func (p project) make() error {
	dir := fmt.Sprintf("%s%s", localRepos, p.Name)
	return doExec(p.Exec, dir)
}
