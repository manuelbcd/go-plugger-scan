// Copyright 2021 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package staticplugin

import (
	"fmt"

	plugger "github.com/thediveo/go-plugger"
	_ "github.com/thediveo/go-plugger/examples/staticplugin/plugins/foo"
)

func Example() {
	plugs := plugger.New("plugins")
	for _, doit := range plugs.Func("DoIt") {
		fmt.Println(doit.(func() string)())
	}
	// Output: foo static plugin
}
