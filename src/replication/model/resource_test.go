// Copyright Project Harbor Authors
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

package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResourceName(t *testing.T) {
	r := &ResourceMetadata{}
	assert.Equal(t, "", r.GetResourceName())

	r = &ResourceMetadata{
		Namespace: &Namespace{
			Name: "library",
		},
	}
	assert.Equal(t, "library", r.GetResourceName())

	r = &ResourceMetadata{
		Repository: &Repository{
			Name: "hello-world",
		},
	}
	assert.Equal(t, "hello-world", r.GetResourceName())

	r = &ResourceMetadata{
		Namespace: &Namespace{
			Name: "library",
		},
		Repository: &Repository{
			Name: "hello-world",
		},
	}
	assert.Equal(t, "library/hello-world", r.GetResourceName())
}