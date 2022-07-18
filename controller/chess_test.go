/**
 * @Author: lrc
 * @Date: 2022/7/17-9:54
 * @Desc:下棋通信服务
 **/

package controller

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"redrock-test/util"
	"strings"
	"testing"
)

func TestChess(t *testing.T) {
	util.InitLogger()
	r := CreateRouter()
	req := httptest.NewRequest("POST",
		"/user/register",
		strings.NewReader(`{"username": "liwenzhou","password":"123456","rePassword":"123456"}`))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
