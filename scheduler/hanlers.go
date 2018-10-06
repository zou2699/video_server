package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/zou2699/video_server/scheduler/dbops"
)

func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")

	if len(vid) == 0 {
		sendResponse(w, 400, "video is should not be empty")
		return
	}

	err := dbops.AddvideoDeletionRecord(vid)
	if err != nil {
		sendResponse(w, 500, "Internal server error")
		return
	}

	sendResponse(w, 200, "")
	return
}
