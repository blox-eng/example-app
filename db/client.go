package db

import (
	"github.com/blox-eng/work/model"
)

func NewClient(config *Config) SqlClient {
	return &sqlClient{
		config: config,
	}
}

type SqlClient interface {
	CreateBlogRecord(bl *model.Blogs) (model.BlogData, error)
	GetBlogs(string) (model.Blogs, error)
	UpdateBlogs(id string, bl *model.Blogs) (model.BlogData, error)
	DeleteBlog(id string) (string, error)

	CreateWorkReport(wr *model.WorkReport) (model.WorkReportData, error)
	GetWorkReport(string) (model.WorkReport, error)
	UpdateWorkReport(id string, wr *model.WorkReport) (model.WorkReportData, error)
	DeleteWorkReport(id string) (string, error)
}

type Config struct {
	DBConnection string
}

type sqlClient struct {
	config *Config
}

func (c *sqlClient) CreateWorkReport(wr *model.WorkReport) (model.WorkReportData, error) {
	return createWorkReport(wr)
}

func (c *sqlClient) GetWorkReport(id string) (model.WorkReport, error) {
	return getAllWorkReports(id)
}

func (c *sqlClient) UpdateWorkReport(id string, wr *model.WorkReport) (model.WorkReportData, error) {
	return updateWorkReport(id, wr)
}

func (c *sqlClient) DeleteWorkReport(id string) (string, error) {
	return deleteWorkReport(id)
}

func (c *sqlClient) CreateBlogRecord(bl *model.Blogs) (model.BlogData, error) {
	return createBlog(bl)
}

func (c *sqlClient) GetBlogs(id string) (model.Blogs, error) {
	return getAllBlog(id)
}

func (c *sqlClient) UpdateBlogs(id string, bl *model.Blogs) (model.BlogData, error) {
	return updateBlog(id, bl)
}

func (c *sqlClient) DeleteBlog(id string) (string, error) {
	return deleteBlog(id)
}
