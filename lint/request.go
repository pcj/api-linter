// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lint

import (
	"github.com/golang/protobuf/v2/reflect/protodesc"
	"github.com/golang/protobuf/v2/reflect/protoreflect"
	descriptorpb "github.com/golang/protobuf/v2/types/descriptor"
)

type LintRequest struct {
	Files   []File
	Configs Configs
}

type File struct {
	Name     string
	Contents string
}

// protoRequest defines input data for a rule to perform linting.
type protoRequest struct {
	fileDesc   protoreflect.FileDescriptor
	descSource DescriptorSourceMap
}

// NewProtoRequest creates a linting protoRequest for a .proto file.
func NewProtoRequest(fd *descriptorpb.FileDescriptorProto) (protoRequest, error) {
	f, err := protodesc.NewFile(fd, nil)
	if err != nil {
		return protoRequest{}, err
	}
	s, err := newDescriptorSourceMap(fd)
	return protoRequest{
		fileDesc:   f,
		descSource: s,
	}, err
}
