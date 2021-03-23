// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package conntrack

import (
	"net"
	"os"
	"syscall"
	"testing"
)

func Test_reportDialerConnFailed(t *testing.T) {
	type args struct {
		dialerName string
		err        error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "does_os_work",
			args: args{
				dialerName: "foo_bar",
				err: &net.OpError{
					Err: os.NewSyscallError("mock", syscall.ECONNREFUSED),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reportDialerConnFailed(tt.args.dialerName, tt.args.err)
		})
	}
}
