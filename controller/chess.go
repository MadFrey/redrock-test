package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"redrock-test/service"
	"strconv"
	"strings"
)

var rooms = make(map[string]map[int]*websocket.Conn)
var room = make(map[string]int)
var gameStart = make(map[string]int)
var OnlyOnce = make(map[int]bool) //判断是否已经准备
var gamer = make(map[int]string)
var LocalPlayer [2]int
var PlayerTurn int

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Chess(c *gin.Context) {
	header := c.Request.Header.Get("Cookie")
	fmt.Println(header)
	roomid := c.Query("roomid") //从请求里获取房价名roomid
	uidd := c.Query("uid")      // 从请求里获取用户id
	uid, _ := strconv.Atoi(uidd)
	fmt.Println("roomid:===", roomid)
	fmt.Println("uid:==", uid)
	if header == "" {
		return
	}
	parts := strings.Split(header, " ")
	if !(len(parts) == 2) {
		return
	}
	println(11111)
	_, _, err := service.ParseToken(parts[0], parts[1])
	if err != nil {
		return
	}
	// conn就是建立连接后的连接对象
	conn, err := UP.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	func(conn *websocket.Conn) {
		if room[roomid] == 2 {
			err := conn.WriteMessage(websocket.TextMessage, []byte("房间已满"))
			if err != nil {
				log.Println(err)
			}
			return
		}
		if rooms[roomid] == nil {
			rooms[roomid] = map[int]*websocket.Conn{uid: conn}
			gamer[uid] = roomid
			room[roomid] = 1
			LocalPlayer[0] = uid
		} else {
			rooms[roomid][uid] = conn
			gamer[uid] = roomid
			room[roomid] += 1
			LocalPlayer[1] = uid
		}
	}(conn)
	fmt.Println("rooms:==", rooms)

	for {
		if gamer[uid] == roomid {
			if !OnlyOnce[uid] {
				err := conn.WriteMessage(websocket.TextMessage, []byte("输入1准备，输入0取消准备"))
				if err != nil {
					log.Println("error:==", err)
				}
			}
		}
		_, data, _ := conn.ReadMessage()
		ok, _ := strconv.Atoi(string(data))
		println(ok)
		if ok == 1 {
			if !OnlyOnce[uid] {
				gameStart[roomid]++
				OnlyOnce[uid] = true
			}
		} else {
			if OnlyOnce[uid] {
				gameStart[roomid]--
				OnlyOnce[uid] = false
			}
		}
		fmt.Println(gameStart[roomid])
		if gameStart[roomid] == 2 {
			err := conn.WriteMessage(websocket.TextMessage, []byte("游戏开始"))
			if err != nil {
				log.Println("error:==", err)
			}

			PlayerTurn = LocalPlayer[1]
			str := service.InitChess()
			if uid == LocalPlayer[0] {
				err := conn.WriteMessage(websocket.TextMessage, []byte(str))
				if err != nil {
					log.Println(err)
				}
				for {
					if PlayerTurn != uid {
						err = conn.WriteMessage(websocket.TextMessage, []byte("请输入黑方棋子的坐标和移动坐标"))
						if err != nil {
							log.Println(err)
						}
						_, data1, err := conn.ReadMessage()
						if err != nil {
							log.Println(err)
						}
						newData := strings.Split(string(data1), " ")
						var temp [4]int
						for i, v := range newData {
							IntData, _ := strconv.Atoi(v)
							temp[i] = IntData
						}
						ok := service.IsAbleToMove(temp[0], temp[1], temp[2], temp[3])
						if ok {
							str, flag1 := service.Move(temp[0], temp[1], temp[2], temp[3])
							if flag1 == 1 {
								for _, v := range rooms[roomid] {
									err = v.WriteMessage(websocket.TextMessage, []byte("黑方获胜"))
									if err != nil {
										log.Println(err)
									}
								}
								return
							}
							for _, v := range rooms[roomid] {
								err := v.WriteMessage(websocket.TextMessage, []byte(str))
								if err != nil {
									log.Println("error:==", err)
								}
							}
							PlayerTurn = uid
						} else {
							err = conn.WriteMessage(websocket.TextMessage, []byte("请重新输入黑方棋子的坐标和移动坐标"))
							if err != nil {
								log.Println(err)
							}
							continue
						}
					}
				}
			} else {
				err := conn.WriteMessage(websocket.TextMessage, []byte(str))
				if err != nil {
					log.Println(err)
				}
				for {
					if PlayerTurn != uid {
						err = conn.WriteMessage(websocket.TextMessage, []byte("请输入红方棋子的坐标和移动坐标"))
						if err != nil {
							log.Println(err)
						}
						_, data, err := conn.ReadMessage()
						if err != nil {
							log.Println(err)
						}
						newData := strings.Split(string(data), " ")
						var temp [4]int
						for i, v := range newData {
							IntData, _ := strconv.Atoi(v)
							temp[i] = IntData
						}
						ok := service.IsAbleToMove(temp[0], temp[1], temp[2], temp[3])
						if ok {
							str, flag2 := service.Move(temp[0], temp[1], temp[2], temp[3])

							if flag2 == 1 {
								for _, v := range rooms[roomid] {
									err = v.WriteMessage(websocket.TextMessage, []byte("黑方获胜"))
									if err != nil {
										log.Println(err)
									}
								}
								return
							}
							for _, v := range rooms[roomid] {
								err := v.WriteMessage(websocket.TextMessage, []byte(str))
								if err != nil {
									log.Println("error:==", err)
								}
							}
							PlayerTurn = uid
						} else {
							err = conn.WriteMessage(websocket.TextMessage, []byte("请重新输入红方棋子的坐标和移动坐标"))
							if err != nil {
								log.Println(err)
							}
							continue
						}
					}
				}
			}
		}
	}
}
