package tags

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"blog/util"
	"encoding/json"
	"strconv"
)

func GetTagByArticle(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	util.CheckError(err)

	stmt, err := util.Db.Prepare("select Tag.* from Tag,ATrelation " +
		"where ATrelation.aid=? and ATrelation.tid=Tag.id")
	util.CheckError(err)

	rows, err := stmt.Query(id)
	util.CheckError(err)

	var (
		tag util.Tag
		tags []util.Tag
	)

	for rows.Next() {
		err := rows.Scan(&tag.Id, &tag.Name)
		util.CheckError(err)

		tags = append(tags, tag)
	}

	defer rows.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(util.GetOk{Code:http.StatusOK, Text: "Get tags ok", Body: tags}); err != nil {
		panic(err)
	}
}

func GetTags(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	rows, err := util.Db.Query("select * from Tag")
	util.CheckError(err)

	var (
		tag util.Tag
		tags []util.Tag
	)

	for rows.Next() {
		err := rows.Scan(&tag.Id, &tag.Name)
		util.CheckError(err)
		tags = append(tags, tag)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(util.GetOk{Code:http.StatusOK, Text: "Get tags ok", Body: tags}); err != nil {
		panic(err)
	}
}

func GetTag(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	var (
		id int
		tag util.Tag
	)

	id, err := strconv.Atoi(ps.ByName("id"))
	util.CheckError(err)

	stmt, err := util.Db.Prepare("select * from Tag where id=?")
	util.CheckError(err)

	row := stmt.QueryRow(id)

	err = row.Scan(&tag.Id, &tag.Name)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)

		if err := json.NewEncoder(w).Encode(util.JsonErr{Code:http.StatusNotFound, Text: "Tag not foynd"}); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(util.GetOk{Code: http.StatusOK, Text: "Get tags ok", Body: tag}); err != nil {
			panic(err)
		}
	}
}

func PostTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tag := struct {
		Name string
	}{}

	json.NewDecoder(r.Body).Decode(&tag)

	defer r.Body.Close()

	stmt, err := util.Db.Prepare("insert Tag SET name=?")
	util.CheckError(err)

	_, err = stmt.Exec(tag.Name)
	util.CheckError(err)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(util.JsonOk{Code: http.StatusOK, Text: "Post tag ok"}); err != nil {
		panic(err)
	}
}

func DeleteTag(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	stmt, err := util.Db.Prepare("delete from Tag where id=?")
	util.CheckError(err)

	_, err = stmt.Exec(id)
	util.CheckError(err)

	stmt, err = util.Db.Prepare("delete from ATrelation where tid=?")
	util.CheckError(err)

	_, err = stmt.Exec(id)
	util.CheckError(err)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(util.JsonOk{Code: http.StatusOK, Text: "DELETE tags ok"}); err != nil {
		panic(err)
	}
}
