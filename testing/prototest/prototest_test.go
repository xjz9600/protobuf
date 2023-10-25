// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package prototest_test

import (
	"fmt"
	"testing"

	"xjz9600/protobuf/internal/flags"
	"xjz9600/protobuf/proto"
	"xjz9600/protobuf/runtime/protoimpl"
	"xjz9600/protobuf/testing/prototest"

	irregularpb "xjz9600/protobuf/internal/testprotos/irregular"
	legacypb "xjz9600/protobuf/internal/testprotos/legacy"
	legacy1pb "xjz9600/protobuf/internal/testprotos/legacy/proto2_20160225_2fc053c5"
	testpb "xjz9600/protobuf/internal/testprotos/test"
	_ "xjz9600/protobuf/internal/testprotos/test/weak1"
	_ "xjz9600/protobuf/internal/testprotos/test/weak2"
	test3pb "xjz9600/protobuf/internal/testprotos/test3"
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
