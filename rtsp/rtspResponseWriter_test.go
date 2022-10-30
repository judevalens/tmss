package rtsp

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net"
	"net/http"
	"testing"
	mock_rtsp "tmss/rtsp/mocks"
)

func TestResponseWriter_Write(t *testing.T) {
	type fields struct {
		Response    *http.Response
		conn        net.Conn
		isHeaderSet bool
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
		setup   func(fields2 *fields, args2 *args, mockCtrl *gomock.Controller)
	}{
		{
			setup: func(fields2 *fields, args2 *args, mockCtrl *gomock.Controller) {
				connMock := mock_rtsp.NewMockConn(mockCtrl)
				fields2.conn = connMock
				fields2.Response = &http.Response{
					Proto: RtspVersion,
				}
				args2.bytes = []byte("hello world")
				connMock.EXPECT().Write(gomock.Any()).Return(len(args2.bytes), nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ResponseWriter{
				Response:    tt.fields.Response,
				conn:        tt.fields.conn,
				isHeaderSet: tt.fields.isHeaderSet,
			}
			mockCtrl := gomock.NewController(t)
			tt.setup(&tt.fields, &tt.args, mockCtrl)
			got, err := r.Write(tt.args.bytes)

			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v\n", err, tt.wantErr)
			}
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestResponseWriter_WriteHeader(t *testing.T) {
	type fields struct {
		Response    *http.Response
		conn        net.Conn
		isHeaderSet bool
	}
	type args struct {
		statusCode int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ResponseWriter{
				Response:    tt.fields.Response,
				conn:        tt.fields.conn,
				isHeaderSet: tt.fields.isHeaderSet,
			}
			r.WriteHeader(tt.args.statusCode)
		})
	}
}
