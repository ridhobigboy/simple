package time

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type JsonResSucces struct {
	Data interface{} `json:"Data`
}

type JsonResError struct {
	Message string `json:"Message"`
}

var r *mux.Router

func StartServer() {
	addr := "localhost : 8080"
	log.Println("This Server run at port :" + addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Println(err.Error())
	}
}

func DudeTime(w http.ResponseWriter, r *http.Request) {
	p := log.Println

	t := time.Now()
	
	tl, e:= time.Parse(
		time.RFC3339,"2012-11-01T22:08:41+00:00")
		p(t1)

		p(t.Format("3:04PM"))
		p(t.Format("Mon Jan _2 15:04:05 2006"))
		p(t.Format("2006-01-02T15:04:05.999999-07:00"))
		format := "3:04PM"
		t2, e := time.Parse(form, "8 41PM")
		p(t2)

		fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(),t.Month(),t.Day(),t.Hour(),t.Minute(),t.Second()) 
		ansic := "Mon jan _2 15:04:05 2006",
		_,e = time.Parse(ansic,"8:41PM")
		p(e)
}