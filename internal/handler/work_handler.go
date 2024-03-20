package handler

import (
	"net/http"

	_ "github.com/blox-eng/backend/internal/model"
	"github.com/blox-eng/backend/pkg/httputil"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
)

// TODO: Add validation to all routes
// TODO: Add pagination to list routes
// TODO: Add authentication to all routes
// TODO: Add authorization to all routes

// GetAllWorkReports godoc
//
//	@Summary		Get all work reports
//	@Description	Retrieve a list of all work reports
//	@Produce		json
//	@Tags			Work Report
//	@Router			/api/work-reports [get]
//	@Success		200	{object}	httputil.Response{data=[]model.WorkReport}	"A list of work reports"
//	@Failure		400	{object}	httputil.HTTPErr							"Bad Request"
//	@Failure		404	{object}	httputil.HTTPErr							"Not Found"
//	@Failure		500	{object}	httputil.HTTPErr							"Internal Server Error"
//	@Failure		401	{object}	httputil.HTTPErr							"Unauthorized"
//	@Failure		403	{object}	httputil.HTTPErr							"Forbidden"
func getAllWorkReports(store Service, w http.ResponseWriter, r *http.Request) {
	services, err := store.GetAllWorkReports()
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
//	@Param			account	body	model.CreateWorkReport	true	"Add Work Report"
//	@Router			/api/work-reports [post]
//	@Success		201	{object}	httputil.Response{data=model.WorkReport}	"Created"
//	@Failure		400	{object}	httputil.HTTPErr							"Bad Request"
//	@Failure		404	{object}	httputil.HTTPErr							"Not Found"
//	@Failure		500	{object}	httputil.HTTPErr							"Internal Server Error"
//	@Failure		401	{object}	httputil.HTTPErr							"Unauthorized"
//	@Failure		403	{object}	httputil.HTTPErr							"Forbidden"
func createWorkReport(store Service, w http.ResponseWriter, r *http.Request) {
	data := &Request{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httputil.ErrInvalidRequest(err, "Invalid Request"))
	}

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
