#
#    boot - Copyright (c) 2020 by www.gatblau.org
#
#    Licensed under the Apache License, Version 2.0 (the "License");
#    you may not use this file except in compliance with the License.
#    You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
#    Unless required by applicable law or agreed to in writing, software distributed under
#    the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
#    either express or implied.
#    See the License for the specific language governing permissions and limitations under the License.
#
#    Contributors to this project, hereby assign copyright in this code to the project,
#    to be licensed under the same terms as the rest of the code.
#

BUILD_FOLDER=build
BINARY_NAME=boot

package:
	$(MAKE) build_linux
	$(MAKE) build_darwin
	$(MAKE) build_windows

build_linux:
	export GOPATH=$(HOME)/go; export CGO_ENABLED=0; export GOOS=linux; export GOARCH=amd64; go build -o $(BUILD_FOLDER)/$(BINARY_NAME) -v
	zip -mjT $(BUILD_FOLDER)/$(BINARY_NAME)_linux_amd64.zip $(BUILD_FOLDER)/$(BINARY_NAME)

build_darwin:
	export GOPATH=$(HOME)/go; export CGO_ENABLED=0; export GOOS=darwin; export GOARCH=amd64; go build -o $(BUILD_FOLDER)/$(BINARY_NAME) -v
	zip -mjT $(BUILD_FOLDER)/$(BINARY_NAME)_darwin_amd64.zip $(BUILD_FOLDER)/$(BINARY_NAME)

build_windows:
	export GOPATH=$(HOME)/go; export CGO_ENABLED=0; export GOOS=windows; export GOARCH=amd64; go build -o $(BUILD_FOLDER)/$(BINARY_NAME) -v
	zip -mjT $(BUILD_FOLDER)/$(BINARY_NAME)_windows_amd64.zip $(BUILD_FOLDER)/$(BINARY_NAME)