package main

import (
	"net"
	"testing"
)

func Test_handleConn(t *testing.T) {
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleConn(tt.args.conn)
		})
	}
}
