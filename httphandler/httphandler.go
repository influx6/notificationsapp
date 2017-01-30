// Package httphandler contains an API handler for notifications.Service.
package httphandler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/shurcooL/httperror"
	"github.com/shurcooL/notifications"
)

// Notifications is an API handler for notifications.Service.
// It returns errors compatible with httperror package.
type Notifications struct {
	Notifications notifications.Service
}

func (h Notifications) Count(w http.ResponseWriter, req *http.Request) error {
	if req.Method != "GET" {
		return httperror.Method{Allowed: []string{"GET"}}
	}
	n, err := h.Notifications.Count(req.Context(), nil)
	if err != nil {
		return err
	}
	return httperror.JSONResponse{V: n}
}

func (h Notifications) MarkRead(w http.ResponseWriter, req *http.Request) error {
	if req.Method != "POST" {
		return httperror.Method{Allowed: []string{"POST"}}
	}
	q := req.URL.Query() // TODO: Automate this conversion process.
	appID := q.Get("AppID")
	repo := notifications.RepoSpec{URI: q.Get("RepoURI")}
	threadID, err := strconv.ParseUint(q.Get("ThreadID"), 10, 64)
	if err != nil {
		return httperror.HTTP{Code: http.StatusBadRequest, Err: fmt.Errorf("parsing ThreadID query parameter: %v", err)}
	}
	err = h.Notifications.MarkRead(req.Context(), appID, repo, threadID)
	return err
}

func (h Notifications) MarkAllRead(w http.ResponseWriter, req *http.Request) error {
	if req.Method != "POST" {
		return httperror.Method{Allowed: []string{"POST"}}
	}
	q := req.URL.Query() // TODO: Automate this conversion process.
	repo := notifications.RepoSpec{URI: q.Get("RepoURI")}
	err := h.Notifications.MarkAllRead(req.Context(), repo)
	return err
}