package timeline

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"blog/util"
	"time"
	"encoding/json"
	"strconv"
)

const START_YEAR = 2017

func GetTimeline(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()

	rows, err := util.Db.Query("select id, event, time from Timeline order by time DESC")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "database select error")
	defer rows.Close()

	var timelines []util.Timeline

	for i:=START_YEAR; i<=time.Now().Year(); i++ {
		timeline := util.Timeline{Year:i}
		timelines = append(timelines, timeline)
	}

	for rows.Next() {
		var event util.Event

		err = rows.Scan(&event.Id, &event.Even, &event.Times)
		util.CheckAndResponse(w, err, http.StatusInternalServerError, "database scan error")
		event.Months = int(event.Times.Month())

		timelines[time.Now().Year() - START_YEAR].Events = append(timelines[time.Now().Year() - START_YEAR].Events, event)
	}

	// set the header, and write the statusOK with body
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.Response{Code: http.StatusOK, Text: "POST SUCCESSFULLY"}); err != nil {
		panic(err)
	}
}

func UpdateTimeline(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	// get the id argument for find
	id, err := strconv.Atoi(ps.ByName("id"))
	util.CheckAndResponse(w, err, http.StatusBadRequest, "request's id argument error")

	var event struct{
		Event string
	}
	err = json.NewDecoder(r.Body).Decode(&event)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "request body json error")

	stmt, err := util.Db.Prepare("update Timeline SET event=? where id=?")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")
	_, err = stmt.Exec(event.Event, id)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database exec error")

	// set the header, and write the statusOK with body
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.Response{Code: http.StatusOK, Text: "UPDATE SUCCESSFULLY"}); err != nil {
		panic(err)
	}
}

func DeleteTimeline(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	// get the id argument for find
	id, err := strconv.Atoi(ps.ByName("id"))
	util.CheckAndResponse(w, err, http.StatusBadRequest, "request's id argument error")

	stmt, err := util.Db.Prepare("delete from Timeline where id=?")
	_, err = stmt.Exec(id)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database exec error")

	// set the header, and write the statusOK with body
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.Response{Code: http.StatusOK, Text: "DELETE SUCCESSFULLY"}); err != nil {
		panic(err)
	}
}