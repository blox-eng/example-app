package server

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/blox-eng/work/handler"
	"github.com/go-chi/chi/v5"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setupRoutesForUpdate(t *testing.T) {
	type args struct {
		service handler.Service
		r       *chi.Mux
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupRoutesForUpdate(tt.args.service, tt.args.r)
		})
	}
}

func TestServer_ListenAndServe(t *testing.T) {
	type fields struct {
		httpServer *http.Server
		router     *chi.Mux
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				httpServer: tt.fields.httpServer,
				router:     tt.fields.router,
			}
			if err := s.ListenAndServe(); (err != nil) != tt.wantErr {
				t.Errorf("Server.ListenAndServe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newServer(t *testing.T) {
	type args struct {
		r *chi.Mux
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newServer(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
