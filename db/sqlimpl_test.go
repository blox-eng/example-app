package db

import (
	"reflect"
	"testing"

	"github.com/blox-eng/work/model"
)

func Test_createBlog(t *testing.T) {
	type args struct {
		bl *model.Blogs
	}
	tests := []struct {
		name    string
		args    args
		want    model.BlogData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createBlog(tt.args.bl)
			if (err != nil) != tt.wantErr {
				t.Errorf("createBlog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAllBlog(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    model.Blogs
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getAllBlog(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAllBlog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateBlog(t *testing.T) {
	type args struct {
		id string
		bl *model.Blogs
	}
	tests := []struct {
		name    string
		args    args
		want    model.BlogData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := updateBlog(tt.args.id, tt.args.bl)
			if (err != nil) != tt.wantErr {
				t.Errorf("updateBlog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteBlog(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deleteBlog(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("deleteBlog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("deleteBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}
