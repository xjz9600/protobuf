// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"xjz9600/protobuf/internal/filedesc"
	"xjz9600/protobuf/reflect/protoreflect"
	"xjz9600/protobuf/reflect/protoregistry"
)

func TestRegistry(t *testing.T) {
	var hasFiles bool
	protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		if fd.(*filedesc.File).L2 != nil {
			t.Errorf("file %q eagerly went through lazy initialization", fd.Path())
		}
		hasFiles = true
		return true
	})
	if !hasFiles {
		t.Errorf("protoregistry.GlobalFiles is empty")
	}
}
