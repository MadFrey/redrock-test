package util

import "strconv"

func ArrayToString(arr []string) string {
	var result string
	for count, i := range arr {  //遍历数组中所有元素追加成string
		result += i
		if count==len(arr)-1 {
			break
		}
		result+=","
	}
	return result
}

func InputChess(m[10][9]int,Chess[32]string)string  {
	var result string
	for i := 0; i < 10; i++ {
		for j := 0; j < 9; j++ {
			if m[i][j]!=-1 {
				result+=Chess[m[i][j]]+"," +strconv.Itoa(i)+strconv.Itoa(j)+","
			}
		}
	}
	return result
}