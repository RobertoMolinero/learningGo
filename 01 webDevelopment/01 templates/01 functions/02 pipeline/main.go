package main

import (
	"html/template"
	"log"
	"math"
	"os"
)

func double(x int) int {
	return 2 * x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func squareRoot(x float64) float64 {
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"d": double,
	"sq": square,
	"sqRt": squareRoot,
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("../templates/pipeline.gohtml"))
}

func main() {
	e := tpl.ExecuteTemplate(os.Stdout, "pipeline.gohtml", 3)

	if e != nil {
		log.Fatalln(e)
	}
}
