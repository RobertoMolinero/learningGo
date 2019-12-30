package controllers

import (
	"./models"
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type LineController struct {
	db  *sql.DB
	tpl *template.Template
}

func NewLineController(db *sql.DB, tpl *template.Template) *LineController {
	return &LineController{db: db, tpl: tpl}
}

func (lc LineController) GetLines(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	rows, e := lc.db.Query(`SELECT * FROM line;`)
	checkErrorWithLogFatalln(e)

	defer rows.Close()

	var id int
	var publisher string
	var title string

	var data []models.Line

	for rows.Next() {
		e = rows.Scan(&id, &publisher, &title)
		checkErrorWithLogFatalln(e)

		data = append(data, models.Line{Id: id, Publisher: publisher, Title: title})
	}

	lc.tpl.ExecuteTemplate(res, "line.gohtml", data)
}

func (lc LineController) CreateLine(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	line := models.Line{}
	json.NewDecoder(req.Body).Decode(&line)

	stmtInsert, e := lc.db.Prepare("INSERT INTO line(id, publisher, title) VALUES( ?, ?, ? )")
	checkErrorWithLogFatalln(e)

	defer stmtInsert.Close()

	result, e := stmtInsert.Exec(line.Id, line.Publisher, line.Title)
	checkErrorWithLogFatalln(e)

	lastInsertId, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()

	log.Printf("CREATE LINE | Last Id: %s | Rows Affected: %s", strconv.FormatInt(lastInsertId, 10), strconv.FormatInt(rowsAffected, 10))
}

func (lc LineController) UpdateLine(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	e := req.ParseForm()
	checkErrorWithLogFatalln(e)

	idAsString := req.FormValue("id")
	id, e := strconv.Atoi(idAsString)
	publisher := req.FormValue("publisher")
	title := req.FormValue("title")

	checkErrorWithLogFatalln(e)

	lc.UpdateLineMariaDb(id, publisher, title)

	http.Redirect(res, req, "/line", http.StatusSeeOther)
}

func (lc LineController) UpdateLineMariaDb(id int, publisher string, title string) {
	stmtUpdate, e := lc.db.Prepare("UPDATE line SET publisher = ?, title = ? WHERE id = ?")
	checkErrorWithLogFatalln(e)

	defer stmtUpdate.Close()

	result, e := stmtUpdate.Exec(publisher, title, id)
	checkErrorWithLogFatalln(e)

	lastInsertId, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()

	log.Printf("UPDATE LINE | Last Id: %s | Rows Affected: %s", strconv.FormatInt(lastInsertId, 10), strconv.FormatInt(rowsAffected, 10))
}

func (lc LineController) DeleteLine(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	e := req.ParseForm()
	checkErrorWithLogFatalln(e)

	idAsString := req.FormValue("id")
	id, e := strconv.Atoi(idAsString)
	checkErrorWithLogFatalln(e)

	lc.DeleteLineMariaDb(id)

	http.Redirect(res, req, "/line", http.StatusSeeOther)
}

func (lc LineController) DeleteLineMariaDb(id int) {
	stmtDelete, e := lc.db.Prepare("DELETE FROM line WHERE id = ?")
	checkErrorWithLogFatalln(e)

	defer stmtDelete.Close()

	result, e := stmtDelete.Exec(id)
	checkErrorWithLogFatalln(e)

	lastInsertId, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()

	log.Printf("DELETE LINE | Last Id: %s | Rows Affected: %s", strconv.FormatInt(lastInsertId, 10), strconv.FormatInt(rowsAffected, 10))
}

func checkErrorWithLogFatalln(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
