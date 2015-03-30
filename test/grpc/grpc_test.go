// Copyright (c) 2015, Vastech SA (PTY) LTD. All rights reserved.
// http://github.com/gogo/protobuf/gogoproto
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package grpc

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

type testServer struct{}

func (this *testServer) UnaryCall(c context.Context, s *SimpleRequest) (*SimpleResponse, error) {
	return &SimpleResponse{s.Value}, nil
}
func (this *testServer) Downstream(*SimpleRequest, Test_DownstreamServer) error {
	return nil
}
func (this *testServer) Upstream(Test_UpstreamServer) error {
	return nil
}
func (this *testServer) Bidi(Test_BidiServer) error {
	return nil
}

func setup(t *testing.T) (*grpc.Server, TestClient) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	_, port, err := net.SplitHostPort(lis.Addr().String())
	if err != nil {
		log.Fatalf("Failed to parse listener address: %v", err)
	}
	s := grpc.NewServer(grpc.MaxConcurrentStreams(10))
	RegisterTestServer(s, &testServer{})
	go s.Serve(lis)
	addr := "localhost:" + port
	conn, err := grpc.Dial(addr)
	if err != nil {
		log.Fatalf("Dial(%q) = %v", addr, err)
	}
	tc := NewTestClient(conn)
	return s, tc
}

func TestUnary(t *testing.T) {
	server, client := setup(t)
	defer server.Stop()
	want := int64(5)
	reply, err := client.UnaryCall(context.Background(), &SimpleRequest{want})
	if err != nil || reply.Value != want {
		t.Fatalf("UnaryCall(_, _) = %d, %v, want %d, <nil>", reply.Value, err, want)
	}
}
