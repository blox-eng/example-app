package handler

import (
	"net/http"

	"github.com/blox-eng/backend/internal/model"
	"github.com/blox-eng/backend/pkg/httputil"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	// log "github.com/sirupsen/logrus"
)

type ctx struct {
	store Service
	h     func(Service, http.ResponseWriter, *http.Request)
}

type Request struct {
	*model.WorkReport
}

func (a *Request) Bind(r *http.Request) error {
	//TODO: to be expanded
	if err := render.Bind(r, a.WorkReport); err != nil {
		return err
	}
	return nil
}

func (g *ctx) handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		g.h(g.store, w, r)
	}
}

// Handler returns an HTTP handler that routes requests to the appropriate endpoints.
func Handler(store Service) http.Handler {
	r := chi.NewRouter()

	// Define context for each endpoint handler
	getAllWorkReports := ctx{store: store, h: getAllWorkReports}
	getWorkReport := ctx{store: store, h: getWorkReport}
	createWorkReport := ctx{store: store, h: createWorkReport}
	updateWorkReport := ctx{store: store, h: updateWorkReport}
	deleteWorkReport := ctx{store: store, h: deleteWorkReport}

	r.Get(httputil.WrapHandlerFunc("/work-reports", "get all work reports", getAllWorkReports.handle()))
	r.Get(httputil.WrapHandlerFunc("/work-reports/{id}", "get work report", getWorkReport.handle()))
	r.Post(httputil.WrapHandlerFunc("/work-reports", "create work report", createWorkReport.handle()))
	r.Put(httputil.WrapHandlerFunc("/work-reports/{id}", "update work report", updateWorkReport.handle()))
	r.Delete(httputil.WrapHandlerFunc("/work-reports/{id}", "delete work report", deleteWorkReport.handle()))

	return r
}

//type Response struct {
//	Meta interface{}
//	Data interface{}
//}
