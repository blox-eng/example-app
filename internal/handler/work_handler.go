package handler

import (
	"net/http"
	"strconv"

	"github.com/blox-eng/backend/internal/model"
	"github.com/blox-eng/backend/pkg/httputil"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
)

// TODO: Add validation to all routes
// TODO: Add pagination to list routes
// TODO: Add authentication to all routes
// TODO: Add authorization to all routes

// ListWorkReports godoc
//
//	@Summary		List work reports
//	@Description	Retrieve a list of work reports
//	@Produce		json
//	@Param			worker			query	string	false	"Filter by worker name"
//	@Param			work			query	string	false	"Filter by type of work"
//	@Param			work_quantity	query	number	false	"Filter by work quantity"
//	@Param			quantity_unit	query	string	false	"Filter by quantity unit"
//	@Param			site_id			query	integer	false	"Filter by site ID"
//	@Tags			Work Report
//	@Router			/api/work-reports [get]
//	@Success		200	{object}	httputil.Response{data=[]model.WorkReport}	"A list of work reports"
//	@Failure		400	{object}	httputil.HTTPErr							"Bad Request"
//	@Failure		404	{object}	httputil.HTTPErr							"Not Found"
//	@Failure		500	{object}	httputil.HTTPErr							"Internal Server Error"
//	@Failure		401	{object}	httputil.HTTPErr							"Unauthorized"
//	@Failure		403	{object}	httputil.HTTPErr							"Forbidden"
func listWorkReports(store Service, w http.ResponseWriter, r *http.Request) {
	// Parse query parameters from the request
	queryParams := r.URL.Query()
	// Create a query object based on the parsed parameters
	query := model.WorkReport{
		Worker:       queryParams.Get("worker"),
		Work:         queryParams.Get("work"),
		WorkQuantity: parseFloatQueryParam(queryParams.Get("work_quantity")),
		QuantityUnit: queryParams.Get("quantity_unit"),
		SiteID:       parseUintQueryParam(queryParams.Get("site_id")),
	}

	services, err := store.ListWorkReports(&query)
	if err != nil {
		log.Error("Unable to fetch work report ", httputil.Error(err).Code, services, err)
		httputil.ErrInvalidRequest(err, "Unable to Fetch Services")
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, httputil.NewSuccessResponse(http.StatusOK, services))

}

// CreateWorkReport godoc
//
//	@Summary		Create a new work report
//	@Description	Create a new work report
//	@Accept			json
//	@Tags			Work Report
//	@Produce		json
//	@Param			workReport	body	model.CreateWorkReport	true	"Add Work Report"
//	@Router			/api/work-reports [post]
//	@Success		201	{object}	httputil.Response{data=model.WorkReport}	"Created"
//	@Failure		400	{object}	httputil.HTTPErr							"Bad Request"
//	@Failure		404	{object}	httputil.HTTPErr							"Not Found"
//	@Failure		500	{object}	httputil.HTTPErr							"Internal Server Error"
//	@Failure		401	{object}	httputil.HTTPErr							"Unauthorized"
//	@Failure		403	{object}	httputil.HTTPErr							"Forbidden"
func createWorkReport(store Service, w http.ResponseWriter, r *http.Request) {
	data := &Request{}
	//	if err := render.Bind(r, data); err != nil {
	//		render.Render(w, r, httputil.ErrInvalidRequest(err, "Invalid Request"))
	//		return
	//	}

	recordSchema := data.WorkReport
	services, err := store.CreateWorkReport(recordSchema)
	if err != nil {
		log.Error("Unable To Fetch stats ", httputil.Error(err).Code, services, err)
		httputil.ErrInvalidRequest(err, "Unable To Fetch Services ")
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, httputil.NewSuccessResponse(http.StatusOK, services))
}

// GetWorkReport godoc
//
//	@Summary		Get a specific work report
//	@Description	Retrieve details of a specific work report by its ID
//	@Param			id	path	string	true	"Work report ID"
//	@Tags			Work Report
//	@Accept			json
//	@Produce		json
//	@Router			/api/work-reports/{id} [get]
//	@Success		200	{object}	httputil.Response{data=model.WorkReport}	"OK"
//	@Failure		400	{object}	httputil.HTTPErr							"Bad Request"
//	@Failure		404	{object}	httputil.HTTPErr							"Not Found"
//	@Failure		500	{object}	httputil.HTTPErr							"Internal Server Error"
//	@Failure		401	{object}	httputil.HTTPErr							"Unauthorized"
//	@Failure		403	{object}	httputil.HTTPErr							"Forbidden"
func getWorkReport(store Service, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	services, err := store.GetWorkReport(id)
	if err != nil {
		log.Error("Unable to fetch work report ", httputil.Error(err).Code, services, err)
		httputil.ErrInvalidRequest(err, "Unable to Fetch Services")
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, httputil.NewSuccessResponse(http.StatusOK, services))

}

// UpdateWorkReport godoc
//
//	@Summary		Update an existing work report
//	@Description	Update an existing work report by its ID
//	@Param			id	path	string	true	"Work report ID"
//	@Tags			Work Report
//	@Accept			json
//	@Produce		json
//	@Router			/api/work-reports/{id} [put]
//	@Success		200	{object}	httputil.Response{data=model.WorkReport}	"OK"
//	@Failure		400	{object}	httputil.HTTPErr							"Bad Request"
//	@Failure		404	{object}	httputil.HTTPErr							"Not Found"
//	@Failure		500	{object}	httputil.HTTPErr							"Internal Server Error"
//	@Failure		401	{object}	httputil.HTTPErr							"Unauthorized"
//	@Failure		403	{object}	httputil.HTTPErr							"Forbidden"
func updateWorkReport(store Service, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	payload := &Request{}
	if err := render.Bind(r, payload); err != nil {
		render.Render(w, r, httputil.ErrInvalidRequest(err, "Invalid Request"))
		return
	}
	recordSchema := payload.WorkReport
	services, err := store.UpdateWorkReport(id, recordSchema)
	if err != nil {
		log.Error("Unable To Fetch stats ", httputil.Error(err).Code, services, err)
		httputil.ErrInvalidRequest(err, "Unable To Fetch Services ")
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, httputil.NewSuccessResponse(http.StatusOK, services))
}

// DeleteWorkReport godoc
//
//	@Summary		Delete a work report
//	@Description	Delete a work report by its ID
//	@Param			id	path	string	true	"Work report ID"
//	@Tags			Work Report
//	@Accept			json
//	@Produce		json
//	@Router			/api/work-reports/{id} [delete]
//	@Success		200	{object}	httputil.Response{}	"OK"
//	@Failure		400	{object}	httputil.HTTPErr	"Bad Request"
//	@Failure		404	{object}	httputil.HTTPErr	"Not Found"
//	@Failure		500	{object}	httputil.HTTPErr	"Internal Server Error"
//	@Failure		401	{object}	httputil.HTTPErr	"Unauthorized"
//	@Failure		403	{object}	httputil.HTTPErr	"Forbidden"
func deleteWorkReport(store Service, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	services, err := store.DeleteWorkReport(id)
	if err != nil {
		log.Error("Unable To Fetch stats ", httputil.Error(err).Code, services, err)
		httputil.ErrInvalidRequest(err, "Unable To Fetch Services ")
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, httputil.NewSuccessResponse(http.StatusOK, services))
}

// Helper function to parse float query parameters
func parseFloatQueryParam(param string) float64 {
	value, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return 0.0
	}
	return value
}

// Helper function to parse uint query parameters
func parseUintQueryParam(param string) uint64 {
	value, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return 0
	}
	return value
}
