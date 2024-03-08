package httphandler

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/go-chi/render"
)

func TestHTTPErr_Error(t *testing.T) {
	tests := []struct {
		name string
		e    HTTPErr
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("HTTPErr.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		err  error
		code int
	}
	tests := []struct {
		name string
		args args
		want *HTTPErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.err, tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		err interface{}
	}
	tests := []struct {
		name string
		args args
		want *HTTPErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Error(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrResponse_Render(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		e       *ErrResponse
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.Render(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("ErrResponse.Render() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResponse_Render(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		rd      *Response
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rd.Render(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Response.Render() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWrapHandlerFunc(t *testing.T) {
	type args struct {
		route   string
		name    string
		handler http.HandlerFunc
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := WrapHandlerFunc(tt.args.route, tt.args.name, tt.args.handler)
			if got != tt.want {
				t.Errorf("WrapHandlerFunc() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("WrapHandlerFunc() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewSuccessResponse(t *testing.T) {
	type args struct {
		status int
		data   interface{}
	}
	tests := []struct {
		name string
		args args
		want *Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSuccessResponse(tt.args.status, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSuccessResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrInvalidRequest(t *testing.T) {
	type args struct {
		err     error
		message string
	}
	tests := []struct {
		name string
		args args
		want render.Renderer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrInvalidRequest(tt.args.err, tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrInvalidRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
