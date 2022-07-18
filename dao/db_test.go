/**
 * @Author: lrc
 * @Date: 2022/7/18-10:36
 * @Desc:
 **/

package dao

import (
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestInit(t *testing.T) {
	dns := "root:sjk123456@tcp(localhost:3306)/chess"
	_, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	err = Init(dns)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
}
