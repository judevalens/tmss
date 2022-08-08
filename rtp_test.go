package main

import (
	"testing"
)

func Test_parseRtpHeader(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		args   args
		want   RtpHeader
		assert func(args2 args, got RtpHeader, want RtpHeader)
	}{
		{
			name: "Should parse the correct Version field",
			args: args{
				[]byte{3},
			},
			want: RtpHeader{
				Version: 9,
			},
			assert: func(args2 args, got RtpHeader, want RtpHeader) {
				if got.Version != want.Version {
					t.Errorf("got %v, want %v", got.Version, want.Version)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseRtpHeader(tt.args.data)
			tt.assert(tt.args, got, tt.want)
		})
	}
}
