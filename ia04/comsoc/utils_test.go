package comsoc

import (
	"fmt"
	"testing"
)

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
	var alt1 Alternative = 1
	var alt2 Alternative = 3
	tab := [...]Alternative{0, 1, 2, 3, 4, 5}
	prefs := tab[:]

	if isPref(alt1, alt2, prefs) != true {
		t.Errorf("erreur isPref")
	}
}

func Test_maxCount(t *testing.T) {
	var count Count = Count{0: 5, 1: 2, 2: 0, 3: 2, 4: 3, 5: 2}
	tab := [...]Alternative{0}
	bestAlts := tab[:]

	res_maxcount := maxCount(count)

	if !isEqualSliceAlt(res_maxcount, bestAlts) {
		t.Errorf("erreur maxCount, res : %v", res_maxcount)
	}

}

func Test_maxCount2(t *testing.T) {
	//test si deux alternatives max
	var count Count = Count{0: 5, 1: 2, 2: 5, 3: 2, 4: 3, 5: 2}
	tab1 := [...]Alternative{0, 2}
	tab2 := [...]Alternative{2, 0}

	bestAlts1 := tab1[:]
	bestAlts2 := tab2[:]

	res_maxcount := maxCount(count)

	if !isEqualSliceAlt(res_maxcount, bestAlts1) && !isEqualSliceAlt(res_maxcount, bestAlts2) {
		t.Errorf("erreur maxCount, res : %v", res_maxcount)
	}
}

func Test_contains(t *testing.T) {
	//func contains(alts []Alternative, alt Alternative) bool
	var alt Alternative = 3
	tab := [...]Alternative{0, 1, 2, 3, 4, 5}
	alts := tab[:]

	if contains(alts, alt) != true {
		t.Errorf("erreur fonction contains")
	}

}

func Test_checkProfile(t *testing.T) {
	//func checkProfile(prefs Profile) error

	//type Profile [][]Alternative

	var profile Profile = [][]Alternative{
		{0, 1, 2, 3},
		{1, 3, 2, 4},
		{0, 1, 2, 3},
		{4, 2, 1, 0},
	}

	err := checkProfile(profile)
	if err != nil {
		t.Errorf("Erreur checkProfile")
		fmt.Println(err)
	}
}

func Test_checkProfileAlternative(t *testing.T) {
	//func checkProfileAlternative(prefs Profile, alts []Alternative) error

	var profile Profile = [][]Alternative{
		{0, 1, 2, 3},
		{1, 3, 2, 4},
		{0, 1, 2, 3},
		{4, 2, 1, 0},
	}
	tab := [...]Alternative{0, 1, 2, 3, 4}
	alts := tab[:]

	err := checkProfileAlternative(profile, alts)
	if err != nil {
		t.Errorf("Erreur checkProfile")
		fmt.Println(err)
	}
}
