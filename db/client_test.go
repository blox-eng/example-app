package db

import (
	"reflect"
	"testing"

	"github.com/blox-eng/service/model"
)

func TestNewClient(t *testing.T) {
	type args struct {
		config *Config
	}
	tests := []struct {
		name string
		args args
		want SqlClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqlClient_CreateWorkReport(t *testing.T) {
	type args struct {
		wr *model.WorkReport
	}
	tests := []struct {
		name    string
		c       *sqlClient
		args    args
		want    model.WorkReportData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CreateWorkReport(tt.args.wr)
			if (err != nil) != tt.wantErr {
				t.Errorf("sqlClient.CreateWorkReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sqlClient.CreateWorkReport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqlClient_GetWorkReport(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		c       *sqlClient
		args    args
		want    model.WorkReport
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetWorkReport(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("sqlClient.GetWorkReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sqlClient.GetWorkReport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqlClient_UpdateWorkReport(t *testing.T) {
	type args struct {
		id string
		wr *model.WorkReport
	}
	tests := []struct {
		name    string
		c       *sqlClient
		args    args
		want    model.WorkReportData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.UpdateWorkReport(tt.args.id, tt.args.wr)
			if (err != nil) != tt.wantErr {
				t.Errorf("sqlClient.UpdateWorkReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sqlClient.UpdateWorkReport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqlClient_DeleteWorkReport(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		c       *sqlClient
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.DeleteWorkReport(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("sqlClient.DeleteWorkReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sqlClient.DeleteWorkReport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqlClient_CreateBlogRecord(t *testing.T) {
	type args struct {
		bl *model.Blogs
	}
	tests := []struct {
		name    string
		c       *sqlClient
		args    args
		want    model.BlogData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CreateBlogRecord(tt.args.bl)
			if (err != nil) != tt.wantErr {
				t.Errorf("sqlClient.CreateBlogRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sqlClient.CreateBlogRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqlClient_GetBlogs(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		c       *sqlClient
		args    args
		want    model.Blogs
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetBlogs(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("sqlClient.GetBlogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sqlClient.GetBlogs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqlClient_UpdateBlogs(t *testing.T) {
	type args struct {
		id string
		bl *model.Blogs
	}
	tests := []struct {
		name    string
		c       *sqlClient
		args    args
		want    model.BlogData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.UpdateBlogs(tt.args.id, tt.args.bl)
			if (err != nil) != tt.wantErr {
				t.Errorf("sqlClient.UpdateBlogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sqlClient.UpdateBlogs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqlClient_DeleteBlog(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		c       *sqlClient
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.DeleteBlog(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("sqlClient.DeleteBlog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sqlClient.DeleteBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}
