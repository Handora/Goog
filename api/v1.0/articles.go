package article

import (
	"blog/util"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

const PERPAGE int = 20

type Article struct {
	id      int
	title   string
	content string
	tag     []string
	time    time.Time
}

type jsonErr struct {
	Code int
	Text string
}

type RequestBody struct {
	title   string
	content string
	tag     []string
}

func GetArticles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// get the page argument for pagination
	page := int(ps.ByName("page"))

	stmt, err := util.Db.Prepare("SELECT * FROM Article WHERE Article. LIMIT ?, ? BY time DESC")
	util.CheckError(err)

	rows, err := stmt.Query((page-1)*PERPAGE, page*PERPAGE)
	util.CheckError(err)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var (
		articles []Article
		article  Article
	)
	for rows.Next() {
		err = rows.Scan(&article.id, &article.title, &article.content, &article.time)
		util.CheckError(err)

		stmt, err = util.Db.Prepare("select Tag.name from Tag, ATrelation " +
			"where ATrelation.aid=? and ATrelation.tid=Tag.id")
		util.CheckError(err)

		tags, err := stmt.Query(article.id)
		util.CheckError(err)

		for tags.Next() {
			var tag string
			tags.Scan(&tag)
			article.tag = append(article.tag, tag)
		}
		tags.Close()
		articles = append(articles, article)
	}

	defer rows.Close()

	if err := json.NewEncoder(w).Encode(articles); err != nil {
		panic(err)
	}
}

func GetArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// get the page argument for pagination
	id := int(ps.ByName("id"))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var article Article

	stmt, err := util.Db.Prepare("select Tag.name from Tag, ATrelation " +
		"where ATrelation.aid=? and ATrelation.tid=Tag.id")
	util.CheckError(err)

	tags, err := stmt.Query(id)
	util.CheckError(err)

	for tags.Next() {
		var tag string
		tags.Scan(&tag)
		article.tag = append(article.tag, tag)
	}

	tags.Close()

	stmt, err = util.Db.Prepare("select * from Article where Article.id=?")
	util.CheckError(err)

	rows, err := stmt.Query(id)
	util.CheckError(err)

	ok := false

	for rows.Next {
		ok = true
		rows.Scan(&article.id, &article.title, &article.content, &article.time)
	}
	defer rows.Close()

	if ok {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(article); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	}
}

func PostArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rb := RequestBody{}
	json.NewDecoder(r.Body).Decode(&rb)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var article Article

	stmt, err := util.Db.Prepare("select Tag.name from Tag, ATrelation " +
		"where ATrelation.aid=? and ATrelation.tid=Tag.id")
	util.CheckError(err)

	tags, err := stmt.Query(id)
	util.CheckError(err)

	for tags.Next() {
		var tag string
		tags.Scan(&tag)
		article.tag = append(article.tag, tag)
	}

	tags.Close()

	stmt, err = util.Db.Prepare("select * from Article where Article.id=?")
	util.CheckError(err)

	rows, err := stmt.Query(id)
	util.CheckError(err)

	ok := false

	for rows.Next {
		ok = true
		rows.Scan(&article.id, &article.title, &article.content, &article.time)
	}
	defer rows.Close()

	if ok {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(article); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	}
}
