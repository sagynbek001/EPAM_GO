package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Request struct {
	Method string `json:"method"`
	Name   string `json:"name"`
}

type Response struct {
	Error  interface{} `json:"error"`
	Result interface{} `json:"result"`
}

var names = map[string]int64{}

func Handle(w http.ResponseWriter, r *http.Request) {
	req, err := parseRequest(r)
	if err != nil {
		returnErr(w, err.Error())
		return
	}
	resp, err := handle(*req)
	if err != nil {
		returnErr(w, err.Error())
		return
	}
	returnOK(w, resp)
}

func parseRequest(r *http.Request) (*Request, error) {
	body := r.Body
	defer body.Close()
	bodyData, err := io.ReadAll(body)
	if err != nil {
		log.Print(err.Error())
		return nil, errors.New("failed to read request body")
	}
	req := &Request{}
	if err = json.Unmarshal(bodyData, req); err != nil {
		log.Print(err.Error())
		return nil, errors.New("request format is wrong")
	}
	return req, nil
}

func handle(req Request) (interface{}, error) {
	var resp interface{}
	if req.Method == "list" {
		values := make([]string, len(names))
		idx := 0
		for value := range names {
			values[idx] = value
			idx++
		}
		resp = values
	} else if req.Method == "register" {
		if names[req.Name] > 0 {
			return nil, fmt.Errorf("this name already exists")
		}
		names[req.Name]++
		resp = fmt.Sprintf("The name was succesfully added, the id is: %d", len(names))
	}
	return resp, nil
}

func returnOK(w http.ResponseWriter, data interface{}) {
	res := Response{
		Error:  nil,
		Result: data,
	}
	res.writeToWeb(w)
}

func returnErr(w http.ResponseWriter, data interface{}) {
	res := Response{
		Error:  data,
		Result: nil,
	}
	res.writeToWeb(w)
}

func (r Response) writeToWeb(w http.ResponseWriter) {
	b, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-Type", "application/json")
	if _, err := w.Write(b); err != nil {
		log.Fatal(err)
	}
}
