package util

import (
	"fmt"
	"testing"
)

func TestPasswordHash(t *testing.T) {
	tests := []struct {
		name    string
		pwd     string
		wantErr bool
	}{
		{"first", "123456", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := PasswordHash(tt.pwd)
			fmt.Println(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("PasswordHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPasswordVerify(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		hashPwd string
	}{
		{"t", "123456", "$2a$10$868zQU/55EFfTOJ2iE2vAemtBjo43nlV5JqsuiXLvUwTTby7SdnE2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !PasswordVerify(tt.want, tt.hashPwd) {
				t.Errorf("PasswordVerify() = %v, want %v", tt.hashPwd, tt.want)
			}
		})
	}
}
