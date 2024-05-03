package main

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Event struct {
	Date        time.Time 	`json:"date"`
	Description string    	`json:"description"`
}

type User struct {
	UserID		int `json:"user_id"`
	Events		[]Event
}

type Result struct {
	Events	[]Event
}

// парсинга и валидации параметров POST запросов
func validPostRequest(r *http.Request) (userID int, date string, err error) {
	if r.Method != http.MethodPost {
		err = fmt.Errorf("bad request: bad method %s, method must be %s", r.Method, http.MethodPost)
		return
	}
	if userID, err = strconv.Atoi(r.FormValue("user_id")); err != nil || len(r.FormValue("user_id")) == 0 {
		err = fmt.Errorf("there is no field \"user_id\" in the request body")
		return
	}
	if date = r.FormValue("date"); len(date) == 0 {
		err = fmt.Errorf("there is no field \"user_id\" in the request body")
		return
	}
	return
}

// парсинга и валидации параметров GET запросов
func validGetRequest(r *http.Request) (userID int, date string, err error) {
	if r.Method != http.MethodGet {
		err = fmt.Errorf("bad request: bad method %s, method must be %s", r.Method, http.MethodGet)
		return
	}
	values := r.URL.Query()
	if _, ok := values["user_id"]; ok {
		if userID, err = strconv.Atoi(values.Get("user_id")); err != nil {
			err = fmt.Errorf("incorrect \"user_id\". Please enter a positive integer")
		}
	}
	if _, ok := values["date"]; ok {
		err = fmt.Errorf("there is no field \"user_id\" in the request body")
		return
	}
	return
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Это календарь, для работы с ним...(нужно дополнить)")); err != nil {
		log.Printf("error of home page: %s", err.Error())
	}
	
}


func createEvent(w http.ResponseWriter, r *http.Request) {
	userID, date, err := validPostRequest(r)
	if err != nil {
		ResponseWrapper(w, Result{}, err, http.StatusBadRequest)
	}
	
}

func updateEvent(w http.ResponseWriter, r *http.Request) {


}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	r.Method = "POST"

}

func eventsForDay(w http.ResponseWriter, r *http.Request) {
	r.Method = "GET"

}

func eventsForWeek(w http.ResponseWriter, r *http.Request) {
	r.Method = "GET"

}

func eventsForMonth(w http.ResponseWriter, r *http.Request) {
	r.Method = "GET"

}

// обработчик ответов в формате JSON
func ResponseWrapper(w http.ResponseWriter, result Result, err error, status int) {
	m := make(map[string]string)
	if status != http.StatusOK {
		m["error"] = err.Error() 
	} else if len(result.Events) == 0 {
		m["result"] = "event deleted"
	} else {
		data, err := json.Marshal(result)
		if err != nil {
			err = fmt.Errorf("Internal Server Error: %s", err.Error())
			ResponseWrapper(w, Result{}, err, http.StatusInternalServerError)
			return
		}
		m["result"] = string(data)
	}
	
	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)
	_, err = w.Write(res)
	if err != nil {
		log.Println(err)
	}
}

// логер запросов 
type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s    %s    %v\n", r.Method, r.RequestURI, time.Since(start))
}

func NewLogger(handerToWrap http.Handler) *Logger {
	return &Logger{handler: handerToWrap}
}

// роутер запросов
func CreateRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homePage)						
	mux.HandleFunc("/create_event", createEvent)        //POST
	mux.HandleFunc("/update_event", updateEvent)        //POST
	mux.HandleFunc("/delete_event", deleteEvent)        //POST
	mux.HandleFunc("/events_for_day", eventsForDay)     //GET
	mux.HandleFunc("/events_for_week", eventsForWeek)   //GET
	mux.HandleFunc("/events_for_month", eventsForMonth) //GET

	return mux
}



func main() {

	addr := "localhost:8080"

	router := CreateRouter()
	handler := NewLogger(router)
	server := &http.Server{
		Addr: addr,
		Handler: handler,
	}
	log.Printf("server is listening at %s", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("close server error: %s", err.Error())
	}
}
