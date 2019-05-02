package models

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
)

func TestGetAllSuccess(t *testing.T) {
	result, err := GetAll()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
