package tags

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"blog/util"
	"encoding/json"
	"strconv"
	"database/sql"
)


/*
	TODO:
	May be we can create an Article id list to
	return tag, and it may be more effective
*/
func GetTagByArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	// get article id from request
	id, err := strconv.Atoi(ps.ByName("id"))
	util.CheckAndResponse(w, err, http.StatusBadRequest, "request's id argument error")

	// select all corresponding Tag structure using
	// ATrelation Table and Tag Table with id from request
	stmt, err := util.Db.Prepare("select Tag.* from Tag,ATrelation " +
		"where ATrelation.aid=? and ATrelation.tid=Tag.id")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")
	rows, err := stmt.Query(id)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database query error")
	defer rows.Close()

	var tags []util.Tag

	// construct the tags
	for rows.Next() {
		var tag util.Tag
		err := rows.Scan(&tag.Id, &tag.Name)
		util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database rows.scan error")
		tags = append(tags, tag)
	}

	// set the header, and write the statusOK with body
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.Response{Code:http.StatusOK, Text: "GET tags through article id successfully", Body: tags}); err != nil {
		panic(err)
	}
}

func GetTags(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()

	// query all tags from db
	rows, err := util.Db.Query("select * from Tag")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database query error")
	defer rows.Close()

	var (
		tag util.Tag
		tags []util.Tag
	)

	// fill the tags structure
	for rows.Next() {
		err := rows.Scan(&tag.Id, &tag.Name)
		util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database rows.scan error")
		tags = append(tags, tag)
	}

	// set the header, and write the statusOK with body
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.Response{Code:http.StatusOK, Text: "Get all tags successfully", Body: tags}); err != nil {
		panic(err)
	}
}


/*
	TODO:
	May be we can create a Tag id list to
	return tag, and it may be more effective
*/
func GetTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	var (
		id int
		tag util.Tag
	)

	// get id from url and convert it to int
	id, err := strconv.Atoi(ps.ByName("id"))
	util.CheckAndResponse(w, err, http.StatusBadRequest, "request's id argument error")

	// select corresponding Tg
	stmt, err := util.Db.Prepare("select * from Tag where id=?")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")
	// just find one row, because id is unique
	row := stmt.QueryRow(id)
	err = row.Scan(&tag.Id, &tag.Name)

	if err == sql.ErrNoRows {
		// not found, set the header, and write the statusNotFound with body
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(util.Response{Code:http.StatusNotFound, Text: "Tag not found"}); err != nil {
			panic(err)
		}
		return
	} else if err != nil {
		// other error
		util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database rows.scan error")
		return
	}

	// err == nil
	// set the header, and write the statusOK with body
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.Response{Code: http.StatusOK, Text: "Get tags ok", Body: tag}); err != nil {
		panic(err)
	}
}

func PostTag(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()

	tag := struct {
		Name string
	}{}

	// read the request body to tag structure
	json.NewDecoder(r.Body).Decode(&tag)
	defer r.Body.Close()

	// insert it
	stmt, err := util.Db.Prepare("insert Tag SET name=?")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")

	_, err = stmt.Exec(tag.Name)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database exec error")

	// set the header, and write the statusOK with body
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.Response{Code: http.StatusOK, Text: "Post tag ok"}); err != nil {
		panic(err)
	}
}


/*
	TODO:
	May be we can create an tag id list to
	delete all, and it may be more effective
 */
func DeleteTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	// read id from request
	id, err := strconv.Atoi(ps.ByName("id"))
	util.CheckAndResponse(w, err, http.StatusBadRequest, "request's id argument error")

	// delete the tag
	stmt, err := util.Db.Prepare("delete from Tag where id=?")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")

	_, err = stmt.Exec(id)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database exec error")

	stmt, err = util.Db.Prepare("delete from ATrelation where tid=?")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")

	_, err = stmt.Exec(id)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database exec error")

	// set the header, and write the statusOK with body
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.Response{Code: http.StatusOK, Text: "DELETE tags ok"}); err != nil {
		panic(err)
	}
}
