package model

import "github.com/gorilla/websocket"

type Message struct {
	Conn  *websocket.Conn
	Uid  int
	Roomid string
	Forbiddenword bool
}

type Chess struct {
	ChessName string
	ChessId int
}