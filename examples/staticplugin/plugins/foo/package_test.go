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

package foo

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	plugger "github.com/thediveo/go-plugger"
)

func TestFooPlugin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "plugger/examples/staticplugin/plugins/foo suite")
}

var _ = Describe("static plugin", func() {

	It("registers its plugin function", func() {
		Expect(plugger.New("plugins").Func("DoIt")).To(HaveLen(1))
	})

	It("successfully calls a registered plugin function", func() {
		Expect(plugger.New("plugins").Func("DoIt")[0].(func() string)()).To(Equal("foo static plugin"))
	})

})
