package model

import (
    "time"

    "github.com/jinzhu/gorm"
)

type ResponseMeta struct {
    AppStatusCode int    `json:"code"`
    Message       string `json:"statusType,omitempty"`
    ErrorDetail   string `json:"errorDetail,omitempty"`
    ErrorMessage  string `json:"errorMessage,omitempty"`
    DevMessage    string `json:"devErrorMessage,omitempty"`
}

type ErrResponse struct {
    HTTPStatusCode int          `json:"-"` // http response status code
    Status         ResponseMeta `json:"status"`
    AppCode        int64        `json:"code,omitempty"` // application-specific error code
}

type Blogs struct{
    ID  int `json:"id"`
    BlogName string `json:"blog_name"`
    BlogDetails string `json:"blog_details,omitempty"`
    BlogDescription string `json:"blog_description,omitempty"`
}

type BlogData struct{
    Blog Blogs `json:"blog"`
    Message string `json:"message"`
}

type Site struct {
    gorm.Model
    SiteName        string `json:"site_name"`
    SiteDetails     string `json:"site_details"`
    SiteLocLong     float64 `json:"site_location_longitude"` // Longitutde/Latitude
    SiteLocLat      float64 `json:"site_location_latitude"` // Longitutde/Latitude
}
type SubSite struct {
    gorm.Model
    SubSiteName     string `json:"site_name"`
    SubSiteDetails  string `json:"site_details"`
    SiteID          uint
    Site            Site
    SubSites        []SubSite
    WorkReports     []WorkReport // https://gorm.io/docs/has_many.html
}

type WorkReport struct {
    gorm.Model           // Adds some metadata fields to the table
    SubSiteID       uint 
    StartTime       time.Time `json:"start_time"`
    EndTime         time.Time `json:"end_time"`
    Work            string `json:"work"`
    WorkQuantity    float64 `json:"work_quantity"`
    QuantityUnit    string `json:"quantity_unit"`
}

type WorkReportData struct {
    WorkReport  WorkReport `json:"work_report"`
    Message     string `json:"message"`
}


