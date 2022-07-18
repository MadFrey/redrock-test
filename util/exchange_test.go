package util

import (
	"testing"
)

func TestArrayToString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"t", "1,2,3"},
	}
	arr := []string{"1", "2", "3"}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayToString(arr); got != tt.want {
				t.Errorf("ArrayToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInputChess(t *testing.T) {
	var M [10][9]int //棋盘
	var Chess [32]string
	tests := []struct {
		name string
		want string
	}{
		{"t", "黑车,00,黑马,01,象,02,士,03,将,04,士,05,象,06,黑马,07,黑车,08,黑炮,21,黑炮,27,卒,30,卒,32,卒,34,卒,36,卒,38,兵,60,兵,62,兵,64,兵,66,兵,68,炮,71,炮,77,车,90,马,91,相,92,仕,93,帅,94,仕,95,相,96,马,97,车,98,"},
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 9; j++ {
			M[i][j] = -1
		}
	}
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InputChess(M, Chess); got != tt.want {
				t.Errorf("InputChess() = %v, want %v", got, tt.want)
			}
		})
	}
}
