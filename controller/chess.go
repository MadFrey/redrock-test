/**
 * @Author: lrc
 * @Date: 2022/7/17-9:54
 * @Desc:下棋通信服务
 **/

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"redrock-test/service"
	"redrock-test/util"
	"strconv"
	"strings"
)

var rooms = make(map[string]map[int]*websocket.Conn) //保存每个房间对应的websocket链接
var room = make(map[string]int)                      //判断房间人数
var gameStart = make(map[string]int)                 //判断人数是否满2人准备
var OnlyOnce = make(map[int]bool)                    //用于设置准备和取消准备
var gamer = make(map[int]string)                     //记录每个人所在的房间号
var LocalPlayer [2]int                               //用于记录房间玩家id
var PlayerTurn int                                   //用于判断是当前是谁的回合

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Chess(c *gin.Context) {
	header := c.Request.Header.Get("Cookie")
	roomid := c.Query("roomid") //从请求里获取房价名roomid
	uidd := c.Query("uid")      // 从请求里获取用户id
	uid, _ := strconv.Atoi(uidd)
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
		util.SugarLogger.Error(err)
	}
	defer conn.Close()

	func(conn *websocket.Conn) {
		if room[roomid] == 2 {
			err := conn.WriteMessage(websocket.TextMessage, []byte("房间已满"))
			if err != nil {
				util.SugarLogger.Error(err)
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

	for {
		if gamer[uid] == roomid {
			if !OnlyOnce[uid] {
				err := conn.WriteMessage(websocket.TextMessage, []byte("输入1准备，输入0取消准备"))
				if err != nil {
					util.SugarLogger.Error(err)
				}
			}
		}
		_, data, _ := conn.ReadMessage()
		ok, _ := strconv.Atoi(string(data))
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

		if gameStart[roomid] == 2 {
			err := conn.WriteMessage(websocket.TextMessage, []byte("游戏开始"))
			if err != nil {
				util.SugarLogger.Error(err)
			}
			PlayerTurn = LocalPlayer[1]
			str := service.InitChess()
			if uid == LocalPlayer[0] {
				err := conn.WriteMessage(websocket.TextMessage, []byte(str))
				if err != nil {
					util.SugarLogger.Error(err)
				}
				for {
					if PlayerTurn != uid {
						err = conn.WriteMessage(websocket.TextMessage, []byte("请输入黑方棋子的坐标和移动坐标"))
						if err != nil {
							util.SugarLogger.Error(err)
						}
						_, data1, err := conn.ReadMessage()
						if err != nil {
							util.SugarLogger.Error(err)
						}
						newData := strings.Split(string(data1), " ")
						var temp [4]int //用于记录输入的2个坐标 分别对应 x y tx ty
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
										util.SugarLogger.Error(err)
									}
								}
								return
							}
							for _, v := range rooms[roomid] {
								err := v.WriteMessage(websocket.TextMessage, []byte(str))
								if err != nil {
									util.SugarLogger.Error(err)
								}
							}
							PlayerTurn = uid
						} else {
							err = conn.WriteMessage(websocket.TextMessage, []byte("请重新输入黑方棋子的坐标和移动坐标"))
							if err != nil {
								util.SugarLogger.Error(err)
							}
							continue
						}
					}
				}
			} else {
				err := conn.WriteMessage(websocket.TextMessage, []byte(str))
				if err != nil {
					util.SugarLogger.Error(err)
				}
				for {
					if PlayerTurn != uid {
						err = conn.WriteMessage(websocket.TextMessage, []byte("请输入红方棋子的坐标和移动坐标"))
						if err != nil {
							util.SugarLogger.Error(err)
						}
						_, data, err := conn.ReadMessage()
						if err != nil {
							util.SugarLogger.Error(err)
						}
						newData := strings.Split(string(data), " ")
						var temp [4]int //用于记录输入的2个坐标  分别对应 x y tx ty
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
										util.SugarLogger.Error(err)
									}
								}
								return
							}

							for _, v := range rooms[roomid] {
								err := v.WriteMessage(websocket.TextMessage, []byte(str))
								if err != nil {
									util.SugarLogger.Error(err)
								}
							}
							PlayerTurn = uid
						} else {
							err = conn.WriteMessage(websocket.TextMessage, []byte("请重新输入红方棋子的坐标和移动坐标"))
							if err != nil {
								util.SugarLogger.Error(err)
							}
							continue
						}
					}
				}
			}
		}
	}
}
