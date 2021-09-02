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
	"log"
	"os"
	"os/exec"
)

// doExec execute exec cmd return nil if exit 0
// otherwise return error
func doExec(cmd string, dir string) error {
	_cmd := exec.Command("sh", "-c", cmd)
	_cmd.Dir = dir
	_cmd.Stdout = os.Stdout
	_cmd.Stderr = os.Stderr

	log.Println(_cmd.String())

	// block until cmd execute completed
	return _cmd.Run()
}
