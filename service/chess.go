/**
 * @Author: lrc
 * @Date: 2022/7/17-9:51
 * @Desc:棋子相关服务
 **/

package service

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"redrock-test/util"
)

var M [10][9]int     //棋盘
var Chess [32]string //棋子

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {

		return true
	},
}

func InitMap() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 9; j++ {
			M[i][j] = -1
		}
	}
}

func InitChess() string {
	InitMap()
	//布置黑方棋子
	Chess[0] = "将" //第0行第4列
	M[0][4] = 0    //m中储存的值对应chess中的棋子名称 如m储存的0，则对应chess[0]中储存的棋子名称 下同
	Chess[1] = "士" //第0行第3列
	M[0][3] = 1
	Chess[2] = "士" //第0行第5列
	M[0][5] = 2
	Chess[3] = "象" //第0行第2列
	M[0][2] = 3
	Chess[4] = "象" //第0行第6列
	M[0][6] = 4
	Chess[5] = "黑马" //第0行第1列
	M[0][1] = 5
	Chess[6] = "黑马" //第0行第7列
	M[0][7] = 6
	Chess[7] = "黑车" //第0行第0列
	M[0][0] = 7
	Chess[8] = "黑车" //第0行第8列
	M[0][8] = 8
	Chess[9] = "黑炮" //第2行第1列
	M[2][1] = 9
	Chess[10] = "黑炮"
	M[2][7] = 10
	for i := 0; i < 5; i++ {
		Chess[11+i] = "卒"
		M[3][i*2] = 11 + i
	}

	//布置红方棋子
	Chess[16] = "帅" //第9行第4列
	M[9][4] = 16
	Chess[17] = "仕" //第9行第3列
	M[9][3] = 17
	Chess[18] = "仕" //第9行第5列
	M[9][5] = 18
	Chess[19] = "相" //第9行第2列
	M[9][2] = 19
	Chess[20] = "相" //第9行第6列
	M[9][6] = 20
	Chess[21] = "马" //第9行第1列
	M[9][1] = 21
	Chess[22] = "马" //第9行第7列
	M[9][7] = 22
	Chess[23] = "车" //第9行第0列
	M[9][0] = 23
	Chess[24] = "车" //第9行第8列
	M[9][8] = 24
	Chess[25] = "炮" //第7行第1列
	M[7][1] = 25
	Chess[26] = "炮" //第7行第7列
	M[7][7] = 26

	for i := 0; i < 5; i++ { //5个红方兵布局
		Chess[27+i] = "兵"
		M[6][i*2] = 27 + i
	}
	str := util.InputChess(M, Chess)
	return str
}

func IsAbleToMove(x int, y int, tx int, ty int) bool {
	fmt.Println(x, y, tx, ty)
	if M[x][y] == 0 || M[x][y] == 16 {
		if M[x][y] == 0 {
			if y == ty && (M[tx][ty] == 0 || M[tx][ty] == 16) {
				for i := tx + 1; i < x; i++ {
					if M[i][ty] != -1 {
						return true
					}
				}
				return true
			}
		} else {
			if y == ty && (M[tx][ty] == 0 || M[tx][ty] == 16) {
				for i := tx - 1; i >= 0; i-- {
					if M[i][ty] != -1 {
						return true
					}
				}
				return true
			}
		}
		fmt.Println(111)
		if y-ty > 1 || ty-y > 1 {
			return false
		}

		if x-tx > 1 || tx-x > 1 { //如果横走或竖走超过一格
			return false
		}
		if M[x][y] == 0 {
			if tx > 2 || ty > 5 || ty < 3 {
				return false
			}
			return true
		} else {
			if tx < 7 || ty > 5 || ty < 3 {
				return false
			}
			return true
		}
		return true
	}
	if M[x][y] == 1 || M[x][y] == 2 || M[x][y] == 17 || M[x][y] == 18 {
		if (tx-x)*(ty-y) == 0 { //如果横走或者竖走
			return false
		}
		if tx-x > 1 || ty-y > 1 || y-ty > 1 || x-tx > 1 { //如果横向或者纵向的位移量大于1，即不是斜走一格
			return false
		}
		if (x > 2 && x < 7) || y < 3 || y > 5 { //如果超出九宫格区域
			return false
		}
		return true
	}
	if M[x][y] == 3 || M[x][y] == 4 || M[x][y] == 19 || M[x][y] == 20 {
		if (tx-x)*(ty-y) == 0 { //如果横走或者竖走
			return false
		}

		if ((tx-x != 2 && tx-x != -2) || (x-tx != 2 && x-tx != -2)) || ((ty-y != 2 && ty-y != -2) || (y-ty != 2 && y-ty != -2)) { //如果横向或者纵向的位移量不同时为2，即不是走田字
			return false
		}
		fmt.Println(111)
		if M[x][y] < 17 && M[x][y] > 0 {
			if tx > 4 { //如果象越过“楚河-汉界”
				return false
			}
		}
		if M[x][y] < 33 && M[x][y] > 16 {
			if tx < 5 { //如果象越过“楚河-汉界”
				return false
			}
		}

		i := 0
		j := 0         //记录象眼位置
		if tx-x == 2 { //象向下跳
			i = x + 1
		}
		if tx-x == -2 { //象向上跳
			i = x - 1
		}
		if ty-y == 2 { //象向右跳
			j = y + 1
		}
		if ty-y == -2 { //象向左跳
			j = y - 1
		}
		if M[i][j] != -1 { //被堵象眼
			return false
		}
		if ty > 9 || tx > 10 || ty < 0 || tx < 0 {
			return false
		}
		return true
	}

	if M[x][y] == 5 || M[x][y] == 6 || M[x][y] == 21 || M[x][y] == 22 {
		if (ty-y)*(tx-x) != 2 && (ty-y)*(tx-x) != -2 { //如果横向位移量乘以竖向位移量不等于2,即如果马不是走日字
			return false
		}
		if tx-x == 2 { //如果马向下跳，并且横向位移量为1，纵向位移量为2
			if M[x+1][y] != -1 { //如果被绊马脚
				return false
			}
		}
		if tx-x == -2 { //如果马向上跳，并且横向位移量为1，纵向位移量为2
			if M[x-1][y] != -1 { //如果被绊马脚
				return false
			}
		}
		if ty-y == 2 { //如果马向右跳，并且横向位移量为2，纵向位移量为1
			if M[x][y+1] != -1 { //如果被绊马脚
				return false
			}
		}

		if ty-y == -2 { //如果马向左跳，并且横向位移量为2，纵向位移量为1
			if M[x][y-1] != -1 { //如果被绊马脚
				return false
			}
		}
		if ty > 9 || tx > 10 || ty < 0 || tx < 0 {
			return false
		}
		return true
	}

	if M[x][y] == 7 || M[x][y] == 8 || M[x][y] == 23 || M[x][y] == 24 {
		if (tx-x)*(ty-y) != 0 { //如果横向位移量和纵向位移量同时都不为0，说明车在斜走，故return false
			return false
		}
		if tx != x { //如果车纵向移动
			if x > tx { //将判断过程简化为纵向从上往下查找中间是否有其他子
				t := tx
				tx = x
				x = t
			}
			for i := x + 1; i < tx; i++ {
				if M[i][ty] != -1 { //如果中间有其他子
					return false
				}
			}
		}

		if ty != y { //如果车横向移动
			if y > ty { //将判断过程简化为横向从左到右查找中间是否有其他子
				t := ty
				ty = y
				y = t
			}
			for i := y + 1; i < ty; i++ {
				if M[x][i] != -1 { //如果中间有其他子
					return false
				}
			}
		}
		if ty > 9 || tx > 10 || ty < 0 || tx < 0 {
			return false
		}
		return true
	}

	if M[x][y] == 9 || M[x][y] == 10 || M[x][y] == 25 || M[x][y] == 26 {
		swapFlagX := false      //记录纵向棋子是否交换过
		swapFlagY := false      //记录横向棋子是否交换过
		if (tx-x)*(ty-y) != 0 { //如果棋子斜走
			return false
		}
		c := 0       //记录两子中间有多少个子
		if tx != x { //如果炮是纵向移动
			if x > tx { //简化后续判断
				t := tx
				tx = x
				x = t
				swapFlagX = true
			}
			for i := x + 1; i < tx; i++ {
				if M[i][y] != -1 { //如果中间有子
					c += 1
				}
			}
		}
		if ty != y { //如果炮是横向 移动
			if y > ty { //简化后续判断
				t := ty
				ty = y
				y = t
				swapFlagY = true
			}
			for i := y + 1; i < ty; i++ {
				if M[x][i] != -1 { //如果中间有子
					c += 1
				}
			}
		}

		if c > 1 { //中间超过一个子
			return false
		}

		if c == 0 { //如果中间没有子
			if swapFlagX == true { //如果之间交换过，需要重新交换回来
				t := tx
				tx = x
				x = t
			}
			if swapFlagY == true {
				t := ty
				ty = y
				y = t
			}
			if M[tx][ty] != -1 { //如果目标处有子存在，则不能移动
				return false
			}
		}

		if c == 1 { //如果中间只有一个子
			if swapFlagX == true { //如果之间交换过，需要重新交换回来
				t := tx
				tx = x
				x = t
			}
			if swapFlagY == true {
				t := ty
				ty = y
				y = t
			}
			if M[tx][ty] == -1 { //如果目标处没有棋子，即不能打空炮
				return false
			}
		}
		if ty > 9 || tx > 10 || ty < 0 || tx < 0 {
			return false
		}
		return true
	}
	if M[x][y] == 11 || M[x][y] == 12 || M[x][y] == 13 || M[x][y] == 14 || M[x][y] == 15 || M[x][y] == 27 || M[x][y] == 28 || M[x][y] == 29 || M[x][y] == 30 || M[x][y] == 31 || M[x][y] == 32 {
		if (tx-x)*(ty-y) != 0 { //如果斜走
			return false
		}
		if tx-x > 1 || ty-y > 1 || x-tx > 1 || y-ty > 1 { //如果一次移动了一格以上
			return false
		}

		if x > 4 && (M[x][y] > 16 && M[x][y] < 33) {
			//如果兵未过河，则只能向上移动,不能左右移动
			if ty-y > 0 || y-ty > 0 { //没过河尝试左右移动
				return false
			}
			if tx-x == 1 { //兵向下移动
				return false
			}
			return true
		} else {
			if x < 5 && (M[x][y] > 16 && M[x][y] < 33) {
				if tx-x == 1 { //兵向下移动
					return false
				}
			} //如果已经过河，可以进行上左右移动，但不能进行向下移动
		}
		if x < 5 && (M[x][y] > 0 && M[x][y] < 17) { //如果兵未过河，则只能向上移动,不能左右移动
			if ty-y > 0 || y-ty > 0 { //没过河尝试左右移动
				return false
			}
			if tx-x == -1 { //兵向下移动
				return false
			}
		} else {
			if x >= 5 {
				if tx-x == -1 { //兵向下移动
					return false
				}
			} //如果已经过河，可以进行上左右移动，但不能进行向下移动
		}
		if ty > 9 || tx > 10 || ty < 0 || tx < 0 {
			return false
		}
		return true
	}
	return true
}

func Move(x int, y int, tx int, ty int) (string, int) {
	flag := 0
	if IsAbleToMove(x, y, tx, ty) {
		if (M[tx][ty] == 0 || M[tx][ty] == 16) && (M[x][y] != 0 && M[x][y] != 16) {
			flag = 1
			str := util.InputChess(M, Chess)
			return str, flag
		}
		if M[x][y] == 0 || M[x][y] == 16 {
			if M[x][y] == 0 {
				if y == ty && (M[tx][ty] == 0 || M[tx][ty] == 16) {
					flag1 := 0
					for i := tx + 1; i < x; i++ {
						if M[i][ty] != -1 {
							flag1 = 1
						}
					}
					if flag1 == 0 {
						flag = 1
						str := util.InputChess(M, Chess)
						return str, flag
					}
				}
			} else {
				if y == ty && (M[tx][ty] == 0 || M[tx][ty] == 16) {
					flag1 := 0
					for i := tx - 1; i >= 0; i-- {
						if M[i][ty] != -1 {
							flag1 = 1
						}
					}
					if flag1 == 0 {
						flag = 1
						str := util.InputChess(M, Chess)
						return str, flag
					}
				}
			}
		}
		if 0 < M[tx][ty] && M[tx][ty] < 33 {
			M[tx][ty] = M[x][y]
			M[x][y] = -1
		} else {
			M[tx][ty] = M[x][y]
			M[x][y] = -1
		}
	}
	str := util.InputChess(M, Chess)
	return str, flag
}
