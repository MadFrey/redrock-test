package model

import "github.com/gorilla/websocket"

type Message struct {
	Conn          *websocket.Conn
	Uid           int
	RoomId        string
	ForbiddenWord bool
}

type Chess struct {
	ChessName string
	ChessId   int
}
