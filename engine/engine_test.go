// Copyright 2017 The GoReporter Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package engine

import (
	"testing"
	"time"

	"github.com/360EntSecGroup-Skylar/goreporter/tools/processbar"
)

func Test_Engine(t *testing.T) {
	lintersProcessChans := make(chan int64, 20)
	lintersFinishedSignal := make(chan string, 10)
	go processbar.LinterProcessBar(lintersProcessChans, lintersFinishedSignal)
	start := time.Now()
	reporter := NewReporter(InitConfig{
		ProjectPath:           "../../../wgliang/logcool",
		ExceptPackages:        "",
		LintersProcessChans:   lintersProcessChans,
		LintersFinishedSignal: lintersFinishedSignal,
		StartTime:             start,
	})
	reporter.Engine()
	close(lintersFinishedSignal)
	close(lintersProcessChans)
}
