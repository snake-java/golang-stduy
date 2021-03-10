package demo7

import "testing"

func TestMapCreate(t *testing.T) {
	a := map[string]int{"你好": 1, "不好": 2}
	t.Log(a["你好"], len(a)) //1 2
	map1 := map[string]int{}
	map1["sutao"] = 2
	t.Log(map1["sutao"], len(map1)) //2 1
	m3 := make(map[int]int, 10)
	t.Log(m3[1], len(m3)) //0 0
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	//go怎么判断是不是空
	if v1, ok := m1[2]; ok {
		t.Logf("key existing %d", v1) // key existing 0
	} else {
		t.Logf("key 3 is not existing")
	}

}

func TestMapRange(t *testing.T) {
	a := map[string]int{"你好": 1, "不好": 2}
	for key, value := range a {
		t.Log(key, value)
	}
}

func TestMapWithValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	t.Log(m[1](1), m[2](2))
}

func TestMapComplierSet(t *testing.T) {
	m := map[int]bool{}
	m[1] = true
	n := 3
	if m[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d isnot existing", n)
	}
	m[2] = true
	t.Log(len(m))
	delete(m, 1)
	if m[1] {
		t.Logf("%d is existing", 1)
	} else {
		t.Logf("%d isnot existing", 1)
	}

}
