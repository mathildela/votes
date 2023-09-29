package comsoc

import (
	"testing"
)

// dans ia04/comsoc:
// go test -v 	pour tous les tester
// go test -run TestFunction	pour tester une fonction spécifique

func isEqualSliceAlt(a, b []Alternative) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// ---------------------------------------------
// FONCTIONS DE TESTS POUR UTILS
// ---------------------------------------------

func Test_rank(t *testing.T) {
	// rank(alt Alternative, prefs []Alternative) int
	tab := [...]Alternative{1, 2, 3, 4, 5}
	prefs := tab[:]

	if rank(3, prefs) != 2 {
		t.Errorf("Output expected : 2, Output returned : %d", rank(3, prefs))
	}

	if rank(1, prefs) != 0 {
		t.Errorf("Output expected : 0, Output returned : %d", rank(1, prefs))
	}

	if rank(5, prefs) != 4 {
		t.Errorf("Output expected : 4, Output returned : %d", rank(5, prefs))
	}

	if rank(7, prefs) != -1 {
		t.Errorf("Output expected : -1, Output returned : %d", rank(7, prefs))
	}
}

func Test_isPref(t *testing.T) {
	//isPref(alt1, alt2 Alternative, prefs []Alternative) bool

	tab := [...]Alternative{0, 1, 2, 3, 4, 5}
	prefs := tab[:]

	if isPref(3, 1, prefs) == true {
		t.Errorf("Output expected : false, Output returned : %t", isPref(3, 1, prefs))
	}

	if isPref(0, 5, prefs) == false {
		t.Errorf("Output expected : true, Output returned : %t", isPref(0, 5, prefs))
	}

	if isPref(0, 6, prefs) == true {
		t.Errorf("Output expected : false, Output returned : %t", isPref(0, 6, prefs))
	}

}

func Test_maxCount(t *testing.T) {
	//func maxCount(count Count) (bestAlts []Alternative)
	//type Count map[Alternative]int

	var count_1 Count = Count{0: 4, 1: 2, 2: 0, 3: 2, 4: 5, 5: 2}
	tab_1 := [...]Alternative{4}
	bestAlts_1 := tab_1[:]
	res_1 := maxCount(count_1)

	var count_2 Count = Count{0: 4, 1: 2, 2: 0, 3: 2, 4: 4, 5: 2}
	tab_21 := [...]Alternative{0, 4}
	tab_22 := [...]Alternative{4, 0}
	bestAlts_21 := tab_21[:]
	bestAlts_22 := tab_22[:]
	res_2 := maxCount(count_2)

	var count_3 Count = Count{}
	tab_3 := [...]Alternative{}
	bestAlts_3 := tab_3[:]
	res_3 := maxCount(count_3)

	if !isEqualSliceAlt(res_1, bestAlts_1) {
		t.Errorf("Output incorrect")
	}

	if !isEqualSliceAlt(res_2, bestAlts_21) && !isEqualSliceAlt(res_2, bestAlts_22) {
		t.Errorf("Output incorrect")
	}

	if !isEqualSliceAlt(res_3, bestAlts_3) {
		t.Errorf("Output incorrect")
	}

}

// ---------------------------------------------
// FONCTIONS DE TEST POUR MAJORITY
// ---------------------------------------------

func TestMajoritySWF(t *testing.T) {
	prefs := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	res, _ := MajoritySWF(prefs)

	if res[1] != 2 {
		t.Errorf("error, result for 1 should be 2, %d computed", res[1])
	}
	if res[2] != 0 {
		t.Errorf("error, result for 2 should be 0, %d computed", res[2])
	}
	if res[3] != 1 {
		t.Errorf("error, result for 3 should be 1, %d computed", res[3])
	}
}

func TestMajoritySCF(t *testing.T) {
	prefs := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	res, err := MajoritySCF(prefs)

	if err != nil {
		t.Error(err)
	}

	if len(res) != 1 || res[0] != 1 {
		t.Errorf("error, 1 should be the only best Alternative")
	}
}

// ---------------------------------------------
// FONCTIONS DE TEST POUR BORDA
// ---------------------------------------------

func TestBordaSWF(t *testing.T) {
	prefs := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	res, _ := BordaSWF(prefs)

	if res[1] != 4 {
		t.Errorf("error, result for 1 should be 4, %d computed", res[1])
	}
	if res[2] != 3 {
		t.Errorf("error, result for 2 should be 3, %d computed", res[2])
	}
	if res[3] != 2 {
		t.Errorf("error, result for 3 should be 2, %d computed", res[3])
	}
}

func TestBordaSCF(t *testing.T) {
	prefs := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	res, err := BordaSCF(prefs)

	if err != nil {
		t.Error(err)
	}

	if len(res) != 1 || res[0] != 1 {
		t.Errorf("error, 1 should be the only best Alternative")
	}
}
