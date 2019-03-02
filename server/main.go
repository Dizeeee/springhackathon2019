package main
import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var TruckLoads = map[string]int{"truck1": 0, "truck2": 0}
var funcMap = map[string]interface{}{}

var indexTemplate = template.Must(template.New("index").Funcs(funcMap).ParseFiles("../website/index.html"))
var fileMux sync.Mutex

func include(filename string) (string, error) {
	fileMux.Lock()
	file, err := ioutil.ReadFile(filename)
	fileMux.Unlock()
	if err != nil {
		return "", err
	}
	s := fmt.Sprintf("%s", file)
	return s, nil
}
func includeHTML(filename string) (template.HTML, error) {
	text, err := include(filename)
	return template.HTML(text), err
}

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func init() {
	// TruckLoads = make(map[string]int)

	http.Handle("/", appHandler(root))
	http.Handle("/getValues", appHandler(getValues))
	http.Handle("/setValue", appHandler(setValue))
	http.Handle("/main.css", appHandler(getCss))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func root(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	r.ParseForm()

	//type data struct{ Counter int }
	err := indexTemplate.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		return err
	}
	return nil
}

func getValues(w http.ResponseWriter, r *http.Request) error {
	type truckData struct {
		Id    string
		Value int
	}

	jTrucks, err := json.Marshal(TruckLoads)
	if err != nil {
		return err
	}

	s := string(jTrucks[:])
	fmt.Fprintln(w, s)
	return nil
}

func setValue(w http.ResponseWriter, r *http.Request) error {
	var variables struct {
		Id    string
		Value int
	}

	r.ParseForm()
	variables.Id = r.FormValue("Id")
	valInt64, err := strconv.ParseInt(r.FormValue("Value"), 0, 0)
	if err != nil {
		return err
	}
	variables.Value = int(valInt64)

	if _, ok := TruckLoads[variables.Id]; ok {
		TruckLoads[variables.Id] = variables.Value
	} else {
		return errors.New("Id not found")
	}
	return nil
}

func getCss(w http.ResponseWriter, r *http.Request) error {
	file, err := include("../website/main.css")
	if err != nil {
		return nil
	}
	fmt.Fprintln(w, file)
	return nil
}