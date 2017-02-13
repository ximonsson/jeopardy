package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Player struct {
	ID     int
	Name   string
	Points int
}

var playerAnswering sync.Mutex
var buzzers chan int = make(chan int, 3)
var buzzersActive bool
var currPlayer Player
var players []Player = make([]Player, 3)
var answering chan int = make(chan int)

func NewPlayer(name string) (*Player, int) {
	p := Player{Name: name}
	n := len(players)
	players = append(players, p)
	return &p, n
}

func playerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		_, n := NewPlayer(r.FormValue("name"))
		fmt.Fprintf(w, "%d", n)
	}
}

func buzzerHandler(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("pid")
	pid, _ := strconv.Atoi(v)
	buzzers <- pid
}
