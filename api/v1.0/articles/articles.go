package articles

import (
	"blog/util"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const PERPAGE int = 20


func GetArticles(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	// get the page argument for pagination
	page := int(ps.ByName("page"))

	stmt, err := util.Db.Prepare("SELECT * FROM Article LIMIT ?, ? BY time DESC")
	util.CheckError(err)

	rows, err := stmt.Query((page-1)*PERPAGE, page*PERPAGE)
	util.CheckError(err)

	var (
		articles []util.Article
		article  util.Article
	)
	for rows.Next() {
		err = rows.Scan(&article.Id, &article.Title, &article.Content, &article.Time)
		util.CheckError(err)

		articles = append(articles, article)
	}

	defer rows.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(util.GetOk{Code:http.StatusOK, Text: "Get articles ok", Body: articles}); err != nil {
		panic(err)
	}
}

func GetArticleByTag(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := int(ps.ByName("id"))
	page := int(ps.ByName("page"))

	stmt, err := util.Db.Prepare("select Article.* from Article,ATrelation " +
		"where ATrelation.tid=? and ATrelation.aid=Article.id limit ?, ? BY time DESC")
	util.CheckError(err)

	rows, err := stmt.Query(id, (page-1)*PERPAGE, (page-1)*PERPAGE)
	util.CheckError(err)

	var (
		articles []util.Article
		article  util.Article
	)

	for rows.Next() {
		err = rows.Scan(&article.Id, &article.Title, &article.Content, &article.Time)
		util.CheckError(err)

		articles = append(articles, article)
	}

	defer rows.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.GetOk{Code:http.StatusOK, Text: "Get articles ok", Body: articles}); err != nil {
		panic(err)
	}
}

func GetArticle(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	// get the page argument for pagination
	id := int(ps.ByName("id"))

	var article util.Article

	stmt, err := util.Db.Prepare("select Tag.name from Tag, ATrelation " +
		"where ATrelation.aid=? and ATrelation.tid=Tag.id")
	util.CheckError(err)

	tags, err := stmt.Query(id)
	util.CheckError(err)

	for tags.Next() {
		var tag string
		tags.Scan(&tag)
		article.Tag = append(article.Tag, tag)
	}

	tags.Close()

	stmt, err = util.Db.Prepare("select * from Article where Article.id=?")
	util.CheckError(err)

	rows, err := stmt.Query(id)
	util.CheckError(err)

	ok := false

	for rows.Next {
		ok = true
		rows.Scan(&article.Id, &article.Title, &article.Content, &article.Time)
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if ok {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(util.GetOk{Code: http.StatusOK, Text: "Get article ok", Body: article}); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(util.JsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	}
}

func PostArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rb := util.GetRequest{}
	json.NewDecoder(r.Body).Decode(&rb)

	stmt, err := util.Db.Prepare("insert Article SET title=?,content=?,time=NOW()")
	util.CheckError(err)

	res, err := stmt.Exec(rb.Title, rb.Content)
	util.CheckError(err)

	id, err := res.LastInsertId()
	util.CheckError(err)

	// TODO:
	//      There may be some improvement through
	//		adding cache for tags
	for i := range rb.Tag {
		stmt, err = util.Db.Prepare("select name from Tag where id=?")
		util.CheckError(err)

		row := stmt.QueryRow(i)

		var tmp string
		err := row.Scan(&tmp)
		util.CheckError(err)

		stmt, err = util.Db.Prepare("INSERT ATrelation SET aid=?,tid=?")
		util.CheckError(err)

		_, err = stmt.Exec(id, i)
		util.CheckError(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.JsonOk{Code: http.StatusOK, Text: "POST SUCCESSFULLY"}); err != nil {
		panic(err)
	}
}

func UpdateArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	postRequest := util.PostRequest{}
	json.NewDecoder(r.Body).Decode(&postRequest)

	stmt, err := util.Db.Prepare("update Article set title=?,content=?,time=NOW() where id=?")
	util.CheckError(err)

	_, err = stmt.Exec(postRequest.Title, postRequest.Content, postRequest.Id)
	util.CheckError(err)

	stmt, err = util.Db.Prepare("delete from ATrelation where aid=?")
	util.CheckError(err)

	_, err = stmt.Exec(postRequest.Id)
	util.CheckError(err)

	// TODO:
	//      There may be some improvement through
	//		adding cache for tags
	for i := range postRequest.Tag {
		stmt, err = util.Db.Prepare("select name from Tag where id=?")
		util.CheckError(err)

		row := stmt.QueryRow(i)

		var tmp string
		err := row.Scan(&tmp)
		util.CheckError(err)

		stmt, err = util.Db.Prepare("INSERT ATrelation SET aid=?,tid=?")
		util.CheckError(err)

		_, err = stmt.Exec(postRequest.Id, i)
		util.CheckError(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.JsonOk{Code: http.StatusOK, Text: "UPDATE SUCCESSFULLY"}); err != nil {
		panic(err)
	}

}

func DeleteArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var id int
	json.NewDecoder(r.Body).Decode(&id)

	stmt, err := util.Db.Prepare("delete from Article where id=?")
	util.CheckError(err)

	_, err = stmt.Exec(id)
	util.CheckError(err)

	stmt, err = util.Db.Prepare("delete from ATrelation where aid=?")
	util.CheckError(err)

	_, err = stmt.Exec(id)
	util.CheckError(err)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.JsonOk{Code: http.StatusOK, Text: "DELETE SUCCESSFULLY"}); err != nil {
		panic(err)
	}
}