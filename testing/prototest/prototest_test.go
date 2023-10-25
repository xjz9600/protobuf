// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package prototest_test

import (
	"fmt"
	"testing"

	"github.com/xjz9600/protobuf/internal/flags"
	"github.com/xjz9600/protobuf/proto"
	"github.com/xjz9600/protobuf/runtime/protoimpl"
	"github.com/xjz9600/protobuf/testing/prototest"

	irregularpb "github.com/xjz9600/protobuf/internal/testprotos/irregular"
	legacypb "github.com/xjz9600/protobuf/internal/testprotos/legacy"
	legacy1pb "github.com/xjz9600/protobuf/internal/testprotos/legacy/proto2_20160225_2fc053c5"
	testpb "github.com/xjz9600/protobuf/internal/testprotos/test"
	_ "github.com/xjz9600/protobuf/internal/testprotos/test/weak1"
	_ "github.com/xjz9600/protobuf/internal/testprotos/test/weak2"
	test3pb "github.com/xjz9600/protobuf/internal/testprotos/test3"
)

func Test(t *testing.T) {
	ms := []proto.Message{
		(*testpb.TestAllTypes)(nil),
		(*test3pb.TestAllTypes)(nil),
		(*testpb.TestRequired)(nil),
		(*irregularpb.Message)(nil),
		(*testpb.TestAllExtensions)(nil),
		(*legacypb.Legacy)(nil),
		protoimpl.X.MessageOf((*legacy1pb.Message)(nil)).Interface(),
	}
	if flags.ProtoLegacy {
		ms = append(ms, (*testpb.TestWeak)(nil))
	}

	for _, m := range ms {
		t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
			prototest.Message{}.Test(t, m.ProtoReflect().Type())
		})
	}
}
