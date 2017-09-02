package articles

import (
	"blog/util"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"database/sql"
	"fmt"
	"os"
)

const PER_PAGE int = 5

func GetArticles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// get the page argument for pagination
	pages, ok := r.URL.Query()["page"]
	var (
		page int
		err  error
		count int
	)

	if ok {
		page, err = strconv.Atoi(pages[0])
		if err != nil {
			page = 1
		}
	} else {
		page = 1
	}

	stmt, err := util.Db.Prepare("SELECT COUNT(*) FROM Article")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")

	row := stmt.QueryRow()
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database query error")

	err = row.Scan(&count)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database scan error")

	var articles util.PagedArticles
	articles.Total = count / PER_PAGE
	articles.CurrentPage = page

	// select the corresponding articles
	stmt, err = util.Db.Prepare("SELECT * FROM Article ORDER BY time DESC LIMIT ?, ?")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")

	rows, err := stmt.Query((page-1)*PER_PAGE, page*PER_PAGE)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database query error")
	defer rows.Close()

	var (
		article  util.Article
		i int
	)

	// scan the rows for all articles
	for rows.Next() {
		err = rows.Scan(&article.Id, &article.Title, &article.Intro, &article.Content, &article.Time)
		util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database rows.scan error")
		if i == 0 {
			articles.First = article
		} else if i == 1 {
			articles.Second = article
		} else if i == 2 {
			articles.Third = article
		} else if i == 3 {
			articles.Fouth = article
		} else if i == 4 {
			articles.Fifth = article
		} else {
			panic("something error about rows scan")
		}
		i++
	}

	// set the header, and write the statusOK with body
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.Response{Code:http.StatusOK, Text: "Get articles ok", Body: articles}); err != nil {
		panic(err)
	}
}

func GetArticleByTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		articles util.PagedArticles
		article  util.Article
		count int
		i int
	)

	// get the id from request
	id, err := strconv.Atoi(ps.ByName("id"))
	util.CheckAndResponse(w, err, http.StatusBadRequest, "request's id argument error")

	// get the page if page is empty we just use 1 as page
	pages, ok := r.URL.Query()["page"]
	var page int

	if ok {
		page, err = strconv.Atoi(pages[0])
		util.CheckAndResponse(w, err, http.StatusBadRequest, "request's page argument error")
	} else {
		page = 1
	}

	// query the num of count
	stmt, err := util.Db.Prepare("select Count(*) from Article,ATrelation " +
		"where ATrelation.tid=? and ATrelation.aid=Article.id")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")

	row := stmt.QueryRow(id)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database query error")

	err = row.Scan(&count)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database scan error")
	articles.Total = count / PER_PAGE
	articles.CurrentPage = page

	// select corresponding articles
	stmt, err = util.Db.Prepare("select Article.* from Article,ATrelation " +
		"where ATrelation.tid=? and ATrelation.aid=Article.id ORDER BY time DESC limit ?, ?")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")

	rows, err := stmt.Query(id, (page-1)*PER_PAGE, (page)*PER_PAGE)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database query error")
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&article.Id, &article.Title, &article.Intro, &article.Content, &article.Time)
		util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database rows.scan error")
		if i == 0 {
			articles.First = article
		} else if i == 1 {
			articles.Second = article
		} else if i == 2 {
			articles.Third = article
		} else if i == 3 {
			articles.Fouth = article
		} else if i == 4 {
			articles.Fifth = article
		} else {
			panic("something error about rows scan")
		}
		i++
	}

	// set the header, and write the statusOK with body
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.Response{Code:http.StatusOK, Text: "Get articles ok", Body: articles}); err != nil {
		panic(err)
	}
}

func GetArticle(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	// get the page argument for pagination
	id, err := strconv.Atoi(ps.ByName("id"))
	util.CheckAndResponse(w, err, http.StatusBadRequest, "request's id argument error")

	var article util.Article

	// select corresponding article
	stmt, err := util.Db.Prepare("select * from Article where Article.id=?")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")

	row := stmt.QueryRow(id)
	err = row.Scan(&article.Id, &article.Title, &article.Intro, &article.Content, &article.Time)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database query error")

	if err == sql.ErrNoRows {
		// not found, set the header, and write the statusNotFound with body
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(util.Response{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
		return
	} else {
		// other error
		util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database rows.scan error")
		return
	}

	// set the header, and write the statusOK with body
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.Response{Code: http.StatusOK, Text: "Get article ok", Body: article}); err != nil {
		panic(err)
	}
}

func PostArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rb := struct {
		Title string
		Intro string
		Content string
		Tag []int
	}{}

	// read the article structure from request body
	json.NewDecoder(r.Body).Decode(&rb)
	defer r.Body.Close()

	// insert corresponding articles
	stmt, err := util.Db.Prepare("insert Article SET title=?,intro=?,content=?,time=NOW()")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")

	res, err := stmt.Exec(rb.Title, rb.Intro, rb.Content)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database exec error")

	id, err := res.LastInsertId()
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database last insert id error")

	// TODO:
	//      There may be some improvement through
	//		adding cache for tags
	for _, i := range rb.Tag {
		// check whether tags inside, if doesn't exist just log and continue
		stmt, err = util.Db.Prepare("select name from Tag where id=?")
		util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")

		row := stmt.QueryRow(i)
		var tmp string
		err := row.Scan(&tmp)
		if err == sql.ErrNoRows {
			fmt.Fprintf(os.Stdout, "Tag id:%d doestn't exist", i)
			continue
		} else if err != nil {
			util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database row.scan error")
		}

		// insert the relation Table
		stmt, err = util.Db.Prepare("INSERT ATrelation SET aid=?,tid=?")
		util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")

		_, err = stmt.Exec(id, i)
		util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database exec error")
	}

	// set the header, and write the statusOK with body
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.Response{Code: http.StatusOK, Text: "POST SUCCESSFULLY"}); err != nil {
		panic(err)
	}
}

func UpdateArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	util.CheckAndResponse(w, err, http.StatusBadRequest, "request's id argument error")

	// read from request to the postRequest
	postRequest := struct {
		Title string
		Intro string
		Content string
		Tag []int
	}{}
	json.NewDecoder(r.Body).Decode(&postRequest)
	defer r.Body.Close()

	// update article table and delete the relation table
	stmt, err := util.Db.Prepare("update Article set title=?,intro=?,content=?,time=NOW() where id=?")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")

	_, err = stmt.Exec(postRequest.Title, postRequest.Intro, postRequest.Content, id)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database exec error")

	stmt, err = util.Db.Prepare("delete from ATrelation where aid=?")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")

	_, err = stmt.Exec(id)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database exec error")

	// TODO:
	//      There may be some improvement through
	//		adding cache for tags
	for _, i := range postRequest.Tag {
		// check whether tags inside, if doesn't exist just log and continue
		stmt, err = util.Db.Prepare("select name from Tag where id=?")
		util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")

		row := stmt.QueryRow(i)
		var tmp string
		err := row.Scan(&tmp)
		if err == sql.ErrNoRows {
			fmt.Fprintf(os.Stdout, "Tag id:%d doestn't exist", i)
			continue
		} else if err != nil {
			util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database row.scan error")
		}

		// insert the relation Table
		stmt, err = util.Db.Prepare("INSERT ATrelation SET aid=?,tid=?")
		util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")

		_, err = stmt.Exec(id, i)
		util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database exec error")
	}

	// set the header, and write the statusOK with body
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.Response{Code: http.StatusOK, Text: "UPDATE SUCCESSFULLY"}); err != nil {
		panic(err)
	}
}

func DeleteArticle(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	util.CheckAndResponse(w, err, http.StatusBadRequest, "request's id argument error")

	// delete both article and atrelation table
	stmt, err := util.Db.Prepare("delete from Article where id=?")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")

	_, err = stmt.Exec(id)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database exec error")

	stmt, err = util.Db.Prepare("delete from ATrelation where aid=?")
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database prepare error")

	_, err = stmt.Exec(id)
	util.CheckAndResponse(w, err, http.StatusInternalServerError, "Database exec error")

	// insert the relation Table
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(util.Response{Code: http.StatusOK, Text: "DELETE SUCCESSFULLY"}); err != nil {
		panic(err)
	}
}
