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

func Test_createWorkReport(t *testing.T) {
	type args struct {
		wr *model.WorkReport
	}
	tests := []struct {
		name    string
		args    args
		want    model.WorkReportData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createWorkReport(tt.args.wr)
			if (err != nil) != tt.wantErr {
				t.Errorf("createWorkReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createWorkReport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAllWorkReports(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    model.WorkReport
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getAllWorkReports(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAllWorkReports() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllWorkReports() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateWorkReport(t *testing.T) {
	type args struct {
		id string
		wr *model.WorkReport
	}
	tests := []struct {
		name    string
		args    args
		want    model.WorkReportData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := updateWorkReport(tt.args.id, tt.args.wr)
			if (err != nil) != tt.wantErr {
				t.Errorf("updateWorkReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateWorkReport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteWorkReport(t *testing.T) {
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
			got, err := deleteWorkReport(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("deleteWorkReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("deleteWorkReport() = %v, want %v", got, tt.want)
			}
		})
	}
}
