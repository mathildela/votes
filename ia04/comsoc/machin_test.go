package comsoc

import "testing"

//dans ia04/comsoc -> go test

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

func Test_rank(t *testing.T) {
	var alt Alternative = 3
	tab := [...]Alternative{1, 2, 3, 4, 5}
	prefs := tab[:]

	if rank(alt, prefs) != 2 {
		t.Errorf("erreur rank")
	}
}

func Test_isPref(t *testing.T) {
	//isPref(alt1, alt2 Alternative, prefs []Alternative) bool {
	var alt1 Alternative = 3
	var alt2 Alternative = 1
	tab := [...]Alternative{0, 1, 2, 3, 4, 5}
	prefs := tab[:]

	if isPref(alt1, alt2, prefs) != true {
		t.Errorf("erreur isPref")
	}
}

func Test_maxCount(t *testing.T) {
	//func maxCount(count Count) (bestAlts []Alternative)
	//type Count map[Alternative]int

	var count Count = Count{0: 5, 1: 2, 2: 0, 3: 2, 4: 3, 5: 2}
	tab := [...]Alternative{0}
	bestAlts := tab[:]

	res_maxcount := maxCount(count)

	if !isEqualSliceAlt(res_maxcount, bestAlts) {
		t.Errorf("erreur maxCount, res : %v", res_maxcount)
	}

	//to do : tester avec plusieurs max (pas dans un ordre précis)

}
