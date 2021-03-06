package util

import (
	"testing"
)

func Test_SturctToJson(t *testing.T){
	expect := "{\"prod\":\"hk-prod\"}"

	j := map[string]interface{}{
		"prod": "hk-prod",
	}

	s := StructToJson(j)

	if s != expect {
		t.Errorf("expect %s, actual %s", expect, s)
	}
}