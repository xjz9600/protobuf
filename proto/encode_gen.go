// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by generate-types. DO NOT EDIT.

package proto

import (
	"math"
	"unicode/utf8"

	"xjz9600/protobuf/encoding/protowire"
	"xjz9600/protobuf/internal/errors"
	"xjz9600/protobuf/internal/strs"
	"xjz9600/protobuf/reflect/protoreflect"
)

var wireTypes = map[protoreflect.Kind]protowire.Type{
	protoreflect.BoolKind:     protowire.VarintType,
	protoreflect.EnumKind:     protowire.VarintType,
	protoreflect.Int32Kind:    protowire.VarintType,
	protoreflect.Sint32Kind:   protowire.VarintType,
	protoreflect.Uint32Kind:   protowire.VarintType,
	protoreflect.Int64Kind:    protowire.VarintType,
	protoreflect.Sint64Kind:   protowire.VarintType,
	protoreflect.Uint64Kind:   protowire.VarintType,
	protoreflect.Sfixed32Kind: protowire.Fixed32Type,
	protoreflect.Fixed32Kind:  protowire.Fixed32Type,
	protoreflect.FloatKind:    protowire.Fixed32Type,
	protoreflect.Sfixed64Kind: protowire.Fixed64Type,
	protoreflect.Fixed64Kind:  protowire.Fixed64Type,
	protoreflect.DoubleKind:   protowire.Fixed64Type,
	protoreflect.StringKind:   protowire.BytesType,
	protoreflect.BytesKind:    protowire.BytesType,
	protoreflect.MessageKind:  protowire.BytesType,
	protoreflect.GroupKind:    protowire.StartGroupType,
}

func (o MarshalOptions) marshalSingular(b []byte, fd protoreflect.FieldDescriptor, v protoreflect.Value) ([]byte, error) {
	switch fd.Kind() {
	case protoreflect.BoolKind:
		b = protowire.AppendVarint(b, protowire.EncodeBool(v.Bool()))
	case protoreflect.EnumKind:
		b = protowire.AppendVarint(b, uint64(v.Enum()))
	case protoreflect.Int32Kind:
		b = protowire.AppendVarint(b, uint64(int32(v.Int())))
	case protoreflect.Sint32Kind:
		b = protowire.AppendVarint(b, protowire.EncodeZigZag(int64(int32(v.Int()))))
	case protoreflect.Uint32Kind:
		b = protowire.AppendVarint(b, uint64(uint32(v.Uint())))
	case protoreflect.Int64Kind:
		b = protowire.AppendVarint(b, uint64(v.Int()))
	case protoreflect.Sint64Kind:
		b = protowire.AppendVarint(b, protowire.EncodeZigZag(v.Int()))
	case protoreflect.Uint64Kind:
		b = protowire.AppendVarint(b, v.Uint())
	case protoreflect.Sfixed32Kind:
		b = protowire.AppendFixed32(b, uint32(v.Int()))
	case protoreflect.Fixed32Kind:
		b = protowire.AppendFixed32(b, uint32(v.Uint()))
	case protoreflect.FloatKind:
		b = protowire.AppendFixed32(b, math.Float32bits(float32(v.Float())))
	case protoreflect.Sfixed64Kind:
		b = protowire.AppendFixed64(b, uint64(v.Int()))
	case protoreflect.Fixed64Kind:
		b = protowire.AppendFixed64(b, v.Uint())
	case protoreflect.DoubleKind:
		b = protowire.AppendFixed64(b, math.Float64bits(v.Float()))
	case protoreflect.StringKind:
		if strs.EnforceUTF8(fd) && !utf8.ValidString(v.String()) {
			return b, errors.InvalidUTF8(string(fd.FullName()))
		}
		b = protowire.AppendString(b, v.String())
	case protoreflect.BytesKind:
		b = protowire.AppendBytes(b, v.Bytes())
	case protoreflect.MessageKind:
		var pos int
		var err error
		b, pos = appendSpeculativeLength(b)
		b, err = o.marshalMessage(b, v.Message())
		if err != nil {
			return b, err
		}
		b = finishSpeculativeLength(b, pos)
	case protoreflect.GroupKind:
		var err error
		b, err = o.marshalMessage(b, v.Message())
		if err != nil {
			return b, err
		}
		b = protowire.AppendVarint(b, protowire.EncodeTag(fd.Number(), protowire.EndGroupType))
	default:
		return b, errors.New("invalid kind %v", fd.Kind())
	}
	return b, nil
}
