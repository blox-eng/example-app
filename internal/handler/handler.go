package handler

import (
	"net/http"

	"github.com/blox-eng/app/internal/model"

	"github.com/blox-eng/app/internal/service"
	"github.com/blox-eng/app/pkg/httputil"

	"github.com/go-chi/chi/v5"
	// log "github.com/sirupsen/logrus"
)

type ctx struct {
	store service.Service
	h     func(service.Service, http.ResponseWriter, *http.Request)
}

type Request struct {
	*model.WorkReport
}

func (a *Request) Bind(r *http.Request) error {
	//TODO: to be expanded
	return nil
}

func (g *ctx) handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		g.h(g.store, w, r)
	}
}

// Handler returns an HTTP handler that routes requests to the appropriate endpoints.
func Handler(store service.Service) http.Handler {
	r := chi.NewRouter()

	// Define context for each endpoint handler
	listWorkReports := ctx{store: store, h: listWorkReports}
	getWorkReport := ctx{store: store, h: getWorkReport}
	createWorkReport := ctx{store: store, h: createWorkReport}
	updateWorkReport := ctx{store: store, h: updateWorkReport}
	deleteWorkReport := ctx{store: store, h: deleteWorkReport}

	r.Get(httputil.WrapHandlerFunc("/work-reports", "list work reports", listWorkReports.handle()))
	r.Get(httputil.WrapHandlerFunc("/work-reports/{id}", "get work report", getWorkReport.handle()))
	r.Post(httputil.WrapHandlerFunc("/work-reports", "create work report", createWorkReport.handle()))
	r.Put(httputil.WrapHandlerFunc("/work-reports/{id}", "update work report", updateWorkReport.handle()))
	r.Delete(httputil.WrapHandlerFunc("/work-reports/{id}", "delete work report", deleteWorkReport.handle()))

	return r
}
