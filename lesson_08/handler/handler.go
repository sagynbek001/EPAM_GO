package handler

import (
	"fmt"
	"io/ioutil"
	"main/saver"
	"net/http"
	"time"
)

type Cars struct {
	Id           int    `json:"id"`
	Manufacturer string `json:"manufacturer"`
	Name         string `json:"name"`
	Fuel         string `json:"fuel"`
}

type Handlers struct {
}

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) Configure() *http.ServeMux {
	userMux := http.NewServeMux()
	userMux.HandleFunc("/", mainHandler)
	userMux.HandleFunc("/list", listHandler)
	return userMux
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// send Not found in such case
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		title := "index"
		p, _ := loadPage(title)
		fmt.Fprintf(w, "%s", p.Body)
	case http.MethodPost:
		// get new value for name and address
		name := r.PostFormValue("name")
		address := r.PostFormValue("address")
		// creating a cookie
		cookie := &http.Cookie{
			Name:   "token",
			Value:  name + ":" + address,
			MaxAge: 300,
		}
		// setting the cookie
		http.SetCookie(w, cookie)
		newToken := saver.Token{
			Token:     cookie.Value,
			CreatedAt: time.Now().Local().String(),
			ExpiredAt: time.Now().Local().Add(time.Hour * time.Duration(240)).String(),
		}
		saver.Tokens[int64(len(saver.Tokens))] = newToken
		fmt.Fprintf(w, "The token was saved!")
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/list" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		return
	}
	for i, val := range saver.Tokens {
		fmt.Fprintf(w, "%d \t %s \t %s \t %s \n", i, val.Token, val.CreatedAt, val.ExpiredAt)
	}
}
