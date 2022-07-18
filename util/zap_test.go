package util

import (
	"reflect"
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestInitLogger(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"t"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitLogger()
		})
	}
}

func TestGetEncoder(t *testing.T) {
	tests := []struct {
		name string
		want zapcore.Encoder
	}{
		{"t", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEncoder(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLowLogWriter(t *testing.T) {
	tests := []struct {
		name string
		want zapcore.WriteSyncer
	}{
		{"t", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLowLogWriter(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLowLogWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHighLogWriter(t *testing.T) {
	tests := []struct {
		name string
		want zapcore.WriteSyncer
	}{
		{"t", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHighLogWriter(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHighLogWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
