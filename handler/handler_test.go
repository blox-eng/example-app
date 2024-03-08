package handler

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/blox-eng/service/model"
)

func Test_ctx_handle(t *testing.T) {
	type fields struct {
		store Service
		h     func(Service, http.ResponseWriter, *http.Request)
	}
	tests := []struct {
		name   string
		fields fields
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &ctx{
				store: tt.fields.store,
				h:     tt.fields.h,
			}
			if got := g.handle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ctx.handle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler(t *testing.T) {
	type args struct {
		store Service
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Handler(tt.args.store); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createBlogs(t *testing.T) {
	type args struct {
		store Service
		w     http.ResponseWriter
		r     *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createBlogs(tt.args.store, tt.args.w, tt.args.r)
		})
	}
}

func Test_getRecordSetPost(t *testing.T) {
	type args struct {
		store Service
		w     http.ResponseWriter
		r     *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getRecordSetPost(tt.args.store, tt.args.w, tt.args.r)
		})
	}
}

func Test_updateBlogs(t *testing.T) {
	type args struct {
		store Service
		w     http.ResponseWriter
		r     *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateBlogs(tt.args.store, tt.args.w, tt.args.r)
		})
	}
}

func Test_deleteBlogs(t *testing.T) {
	type args struct {
		store Service
		w     http.ResponseWriter
		r     *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteBlogs(tt.args.store, tt.args.w, tt.args.r)
		})
	}
}

func TestRequest_Bind(t *testing.T) {
	type fields struct {
		Blogs *model.Blogs
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Request{
				Blogs: tt.fields.Blogs,
			}
			if err := a.Bind(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Request.Bind() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createWorkReport(t *testing.T) {
	type args struct {
		store Service
		w     http.ResponseWriter
		r     *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createWorkReport(tt.args.store, tt.args.w, tt.args.r)
		})
	}
}
