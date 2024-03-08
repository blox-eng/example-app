package handler

import (
	"reflect"
	"testing"

	"github.com/blox-eng/service/db"
	"github.com/blox-eng/service/model"
)

func TestNewService(t *testing.T) {
	type args struct {
		sqlDB db.SqlClient
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.sqlDB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateRecordCoreTeam(t *testing.T) {
	type fields struct {
		sqlDB db.SqlClient
	}
	type args struct {
		bl *model.Blogs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.BlogData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				sqlDB: tt.fields.sqlDB,
			}
			got, err := s.CreateRecordCoreTeam(tt.args.bl)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateRecordCoreTeam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.CreateRecordCoreTeam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetRecordSetPost(t *testing.T) {
	type fields struct {
		sqlDB db.SqlClient
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Blogs
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				sqlDB: tt.fields.sqlDB,
			}
			got, err := s.GetRecordSetPost(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetRecordSetPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetRecordSetPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_UpdateBlog(t *testing.T) {
	type fields struct {
		sqlDB db.SqlClient
	}
	type args struct {
		id string
		bl *model.Blogs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.BlogData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				sqlDB: tt.fields.sqlDB,
			}
			got, err := s.UpdateBlog(tt.args.id, tt.args.bl)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateBlog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.UpdateBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_DeleteBlogs(t *testing.T) {
	type fields struct {
		sqlDB db.SqlClient
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				sqlDB: tt.fields.sqlDB,
			}
			got, err := s.DeleteBlogs(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteBlogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("service.DeleteBlogs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateWorkReport(t *testing.T) {
	type args struct {
		wr *model.WorkReport
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		want    model.WorkReportData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateWorkReport(tt.args.wr)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateWorkReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.CreateWorkReport() = %v, want %v", got, tt.want)
			}
		})
	}
}
