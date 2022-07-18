/**
 * @Author: lrc
 * @Date: 2022/7/17-9:51
 * @Desc: token处理
 **/

package service

import (
	"fmt"
	"redrock-test/model"
	"reflect"
	"testing"
)

func TestCreateToken(t *testing.T) {
	tests := []struct {
		name     string
		username string
		want     string
		want1    string
	}{
		{"t", "111111111", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CreateToken(tt.username)
			fmt.Println(got, got1)
			if got == tt.want {
				t.Errorf("CreateToken() got = %v, want %v", got, tt.want)
			}
			if got1 == tt.want1 {
				t.Errorf("CreateToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParseToken(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name               string
		tokenString        string
		refreshTokenString string
		want               *model.MyClaims
		want1              bool
		wantErr            bool
	}{
		{"t", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjExMTExMTExMSIsImV4cCI6MTY1ODExODAyMiwiaXNzIjoiQWxzYWNlIiwibmJmIjoxNjU4MTA3MTYyfQ.r1zGG4zezivrn8rQ9ai0ClwquUj_P0G2ekTS0bSJ50s",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjExMTExMTExMSIsImV4cCI6MTY1ODExMDgyMiwiaXNzIjoiQWxzYWNlIiwibmJmIjoxNjU4MTA3MjIyfQ.RkrL7DMUq_RpOuMI6tBKLmUqV8y6HvEkkxA_2pScNuI",
			nil, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ParseToken(tt.tokenString, tt.refreshTokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
