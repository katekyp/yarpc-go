// Code generated by thriftrw v1.17.0. DO NOT EDIT.
// @generated

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package kv

import thriftreflect "go.uber.org/thriftrw/thriftreflect"

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "kv",
	Package:  "go.uber.org/yarpc/internal/examples/thrift-keyvalue/keyvalue/kv",
	FilePath: "kv.thrift",
	SHA1:     "9e8c1c30d0b6bd7d83426a92962269aaef706295",
	Raw:      rawIDL,
}

const rawIDL = "exception ResourceDoesNotExist {\n    1: required string key\n    2: optional string message\n}\n\nservice KeyValue {\n    string getValue(1: string key)\n        throws (1: ResourceDoesNotExist doesNotExist)\n    void setValue(1: string key, 2: string value)\n}\n"
