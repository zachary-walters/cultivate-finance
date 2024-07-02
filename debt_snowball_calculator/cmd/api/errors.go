package main

import (
	"encoding/json"
)

func NewError(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

type RequestError struct {
	Err error
	Msg string `json:"msg"`
}

func (r *RequestError) Error() string {
	return r.Err.Error()
}

func reqError(err error) []byte {
	data, _ := json.Marshal(RequestError{
		Err: err,
		Msg: err.Error(),
	})

	return data
}
