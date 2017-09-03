package timeline

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"blog/util"
	"time"
	"encoding/json"
)

const START_YEAR = 2017

func GetTimeline(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()

	rows, err := util.Db.Query("select event, time from Timeline order by time DESC")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "database select error")
	defer rows.Close()

	var timelines []util.Timeline

	for i:=START_YEAR; i<=time.Now().Year(); i++ {
		timeline := util.Timeline{Year:i}
		timelines = append(timelines, timeline)
	}

	for rows.Next() {
		var event util.Event

		err = rows.Scan(&event.Even, &event.Times)
		util.CheckAndResponse(w, err, http.StatusInternalServerError, "database scan error")
		event.Months = int(event.Times.Month())

		timelines[time.Now().Year() - START_YEAR].Events = append(timelines[time.Now().Year() - START_YEAR].Events, event)
	}

	// set the header, and write the statusOK with body
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.Response{Code: http.StatusOK, Text: "Get paged articles successfully", Body: timelines}); err != nil {
		panic(err)
	}
}

func PostTimeline(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()

	var event struct{
		Event string
	}
	err := json.NewDecoder(r.Body).Decode(&event)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "request body json error")

	stmt, err := util.Db.Prepare("insert Timeline SET event=?,time=Now()")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")
	_, err = stmt.Exec(event.Event)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database exec error")

	// set the header, and write the statusOK with body
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.Response{Code: http.StatusOK, Text: "POST SUCCESSFULLY"}); err != nil {
		panic(err)
	}
}