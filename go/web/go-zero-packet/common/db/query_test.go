package db

import "testing"

func TestDbQueryList_GetWhere(t *testing.T) {
	p := DbQueryList{
		Where: map[string]interface{}{
			"a": 1,
			"b": 0,
			"c": "",
			"d": "123",
		},
	}
	t.Log(p.GetWhere())
}
