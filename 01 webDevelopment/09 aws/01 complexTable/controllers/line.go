package controllers

import (
	"./models"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type LineController struct {
	tpl *template.Template
}

var lines []models.Line

func init() {
	lines = []models.Line{
		{Id: uuid.Must(uuid.NewV4()), Text: "Hallo", Number: 145, Date: time.Now()},
		{Id: uuid.Must(uuid.NewV4()), Text: "Hola", Number: 42, Date: time.Now()},
	}
}

func NewLineController(tpl *template.Template) *LineController {
	return &LineController{tpl: tpl}
}

func (lc LineController) GetLines(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var data []models.Line

	for _, v := range lines {
		data = append(data, models.Line{Id: v.Id, Text: v.Text, Number: v.Number, Date: v.Date})
	}

	lc.tpl.ExecuteTemplate(res, "line.gohtml", data)
}

func (lc LineController) CreateLine(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	e := req.ParseForm()
	checkErrorWithLogFatalln(e)

	text := req.FormValue("dialogText")
	number, e := strconv.Atoi(req.FormValue("dialogNumber"))
	checkErrorWithLogFatalln(e)

	date, e := time.Parse("2006-01-02", req.FormValue("dialogDate"))
	checkErrorWithLogFatalln(e)

	lines = append(lines, models.Line{Id: uuid.Must(uuid.NewV4()), Text: text, Number: number, Date: date})

	http.Redirect(res, req, "/line", http.StatusSeeOther)
}

func (lc LineController) UpdateLine(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	e := req.ParseForm()
	checkErrorWithLogFatalln(e)

	idAsString := req.FormValue("id")
	idAsUUID, e := uuid.FromString(idAsString)
	checkErrorWithLogFatalln(e)

	text := req.FormValue("text")
	number, e := strconv.Atoi(req.FormValue("number"))
	checkErrorWithLogFatalln(e)

	date, e := time.Parse("2006-01-02", req.FormValue("date"))
	checkErrorWithLogFatalln(e)

	var newLines []models.Line

	for _, v := range lines {
		if v.Id == idAsUUID {
			newLines = append(newLines, models.Line{Id: idAsUUID, Text: text, Number: number, Date: date})
		} else {
			newLines = append(newLines, v)
		}
	}

	lines = newLines

	http.Redirect(res, req, "/line", http.StatusSeeOther)
}

func (lc LineController) DeleteLine(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	e := req.ParseForm()
	checkErrorWithLogFatalln(e)

	idAsString := req.FormValue("id")
	idAsUUID, e := uuid.FromString(idAsString)
	checkErrorWithLogFatalln(e)

	var newLines []models.Line

	for _, v := range lines {
		if v.Id != idAsUUID {
			newLines = append(newLines, v)
		}
	}

	lines = newLines

	http.Redirect(res, req, "/line", http.StatusSeeOther)
}

func checkErrorWithLogFatalln(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
