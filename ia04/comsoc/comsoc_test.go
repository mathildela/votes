package comsoc

import (
	"fmt"
	"reflect"
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

func Test_checkPrefs(t *testing.T) {
	//func checkPrefs(prefs []Alternative, alts []Alternative) error
	tab1 := [...]Alternative{1, 2, 3, 4, 0}
	prefs := tab1[:]
	tab := [...]Alternative{0, 1, 2, 3, 4}
	alts := tab[:]

	err := checkPrefs(prefs, alts)
	if err != nil {
		t.Errorf("Erreur checkProfile")
		fmt.Println(err)
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

func Test_majoritySWF(t *testing.T) {
	//func MajoritySWF(p Profile) (count Count, err error)

	var profile Profile = [][]Alternative{
		{0, 1, 2, 3},
		{1, 3, 2, 0},
		{0, 1, 2, 3},
		{3, 2, 1, 0},
	}

	var count_test Count = Count{0: 2, 1: 1, 2: 0, 3: 1}

	count, err := MajoritySWF(profile)

	if err != nil {
		t.Errorf("Erreur dans le profile passé en paramètre pour MajoritySWF")
		fmt.Println(err)
	} else {
		if !reflect.DeepEqual(count, count_test) {
			t.Errorf("erreur dans le comptage retourné / résultat : %v", count)
		}
	}

}

// func MajoritySCF(p Profile) (bestAlts []Alternative, err error)
func Test_majoritySCF(t *testing.T) {

	var profile Profile = [][]Alternative{
		{0, 1, 2, 3},
		{1, 3, 2, 0},
		{0, 1, 2, 3},
		{3, 2, 1, 0},
	}

	tab := [...]Alternative{0}
	bestAlts_test := tab[:]

	bestAlts, err := MajoritySCF(profile)

	if err != nil {
		t.Errorf("Erreur dans le profile passé en paramètres pour MajoritySCF")
		fmt.Println(err)
	}

	if !isEqualSliceAlt(bestAlts_test, bestAlts) {
		t.Errorf("erreur dans bestAlts de MajoriySCF, res : %v", bestAlts)
	}

}

// TESTS DE MOODLE

// version 2.0.0

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

func TestApprovalSWF(t *testing.T) {
	prefs := [][]Alternative{
		{1, 2, 3},
		{1, 3, 2},
		{2, 3, 1},
	}
	thresholds := []int{2, 1, 2}

	res, _ := ApprovalSWF(prefs, thresholds)

	if res[1] != 2 {
		t.Errorf("error, result for 1 should be 2, %d computed", res[1])
	}
	if res[2] != 2 {
		t.Errorf("error, result for 2 should be 2, %d computed", res[2])
	}
	if res[3] != 1 {
		t.Errorf("error, result for 3 should be 1, %d computed", res[3])
	}
}

func TestApprovalSCF(t *testing.T) {
	prefs := [][]Alternative{
		{1, 3, 2},
		{1, 2, 3},
		{2, 1, 3},
	}
	thresholds := []int{2, 1, 2}

	res, err := ApprovalSCF(prefs, thresholds)

	if err != nil {
		t.Error(err)
	}
	if len(res) != 1 || res[0] != 1 {
		t.Errorf("error, 1 should be the only best Alternative")
	}
}

// ---------------------------------------------
// FONCTIONS DE TEST POUR TIE_BREAK (SOLENN)
// ---------------------------------------------

func TestTieBreakFactory(t *testing.T) {

	// TieBreak := TieBreakFactory(nil)
	// if TieBreak != nil {
	// 	t.Errorf("No function should be returned because no input")
	// }

	orderedAlts := []Alternative{1, 2, 3, 4}
	TieBreak := TieBreakFactory(orderedAlts)
	alts := []Alternative{4, 2, 1}
	res, err := TieBreak(alts)
	if res != 1 {
		t.Errorf("the best Alternative should be 1, %d computed", res)
	}
	if err != nil {
		t.Errorf("there should be no error, %s computed", err)
	}

	orderedAlts = []Alternative{1, 2, 1, 3, 4}
	TieBreak = TieBreakFactory(orderedAlts)
	alts = []Alternative{4, 2, 1}
	res, err = TieBreak(alts)
	if err == nil {
		t.Error("No order should be known for the alternatives given here")
	}
	if res != -1 {
		t.Errorf("No alternative should be chosen, %d computed", res)
	}

	alts = []Alternative{5, 6, 7}
	res, err = TieBreak(alts)
	if err == nil {
		t.Error("No order should be known for the alternatives given")
	}
	if res != -1 {
		t.Errorf("No alternative should be chosen, %d computed", res)
	}

	alts = nil
	res, err = TieBreak(alts)
	if err == nil {
		t.Error("No alternative was given")
	}
	if res != -1 {
		t.Errorf("No alternative should be chosen, %d computed", res)
	}

}

// ---------------------------------------------
// FONCTIONS DE TEST POUR SWFFACTORY
// ---------------------------------------------

func TestSWFFactory(t *testing.T) {
	// Test pour Majority
	prefs1 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
		{3, 1, 2},
	}
	orderedAlts := []Alternative{1, 2, 3}
	TieBreak := TieBreakFactory(orderedAlts)

	SWFMajority := SWFFactory(MajoritySWF, TieBreak)

	res1, err1 := SWFMajority(prefs1)
	if err1 != nil {
		t.Errorf("no error should be returned, %s computed", err1)
	}

	if res1[1] != 3 {
		t.Errorf("error majority, result for 1 should be 3, %d computed", res1[1])
	}
	if res1[2] != 0 {
		t.Errorf("error majority, result for 2 should be 0, %d computed", res1[2])
	}
	if res1[3] != 2 {
		t.Errorf("error majority, result for 3 should be 2, %d computed", res1[3])
	}

	// Test pour Majority avec 3 valeurs tie-break
	prefs5 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
		{3, 1, 2},
		{2, 1, 3},
		{2, 1, 3},
	}
	orderedAlts = []Alternative{1, 2, 3}
	TieBreak = TieBreakFactory(orderedAlts)

	SWFMajority = SWFFactory(MajoritySWF, TieBreak)

	res5, err5 := SWFMajority(prefs5)
	if err5 != nil {
		t.Errorf("no error should be returned, %s computed", err1)
	}

	if res5[1] != 3 {
		t.Errorf("error majority, result for 1 should be 3, %d computed", res1[1])
	}
	if res5[2] != 2 {
		t.Errorf("error majority, result for 2 should be 2, %d computed", res1[2])
	}
	if res1[3] != 2 {
		t.Errorf("error majority, result for 3 should be 2, %d computed", res1[3])
	}

	// Test pour Borda
	prefs2 := [][]Alternative{
		{1, 2, 3},
		{1, 3, 2},
		{3, 2, 1},
		{3, 1, 2},
	}

	orderedAlts = []Alternative{2, 3, 1}
	TieBreak = TieBreakFactory(orderedAlts)

	SWFBorda := SWFFactory(BordaSWF, TieBreak)
	res2, err2 := SWFBorda(prefs2)

	if err2 != nil {
		t.Error(err2)
	}

	if res2[1] != 5 {
		t.Errorf("error borda, result for 1 should be 5, %d computed", res2[1])
	}
	if res2[2] != 2 {
		t.Errorf("error borda, result for 2 should be 2, %d computed", res2[2])
	}
	if res2[3] != 6 {
		t.Errorf("error borda, result for 3 should be 6, %d computed", res2[3])
	}

	// La factory n'est pas adaptée au vote par approbation car les inputs de ApprovalSWF
	// ne correspond pas au input de la fonction swf en entrée de SWFFactory

	// // Test pour Copeland
	// prefs3 := [][]Alternative{
	// 	{1, 2, 3, 4},
	// 	{2, 3, 4, 1},
	// 	{4, 3, 1, 2},
	// }

	// orderedAlts = []Alternative{3, 2, 4, 1}
	// TieBreak = TieBreakFactory(orderedAlts)

	// SWFCopeland := SWFFactory(CopelandSWF, TieBreak)
	// res3, err3 := SWFCopeland(prefs3)

	// if err3 != nil {
	// 	t.Error(err3)
	// }
	// if res3[1] != -1 {
	// 	t.Errorf("error, result for 1 should be -1, %d computed", res3[1])
	// }
	// if res3[2] != 1 {
	// 	t.Errorf("error, result for 2 should be 1, %d computed", res3[2])
	// }
	// if res3[3] != 2 {
	// 	t.Errorf("error, result for 3 should be 2, %d computed", res3[3])
	// }
	// if res3[4] != 0 {
	// 	t.Errorf("error, result for 4 should be 0, %d computed", res3[4])
	// }

	// Test pour STV
	// Les TieBreak doivent être gérés dans la fonction. Pas possible avec la Factory
	// (ou faire un cas à part si on détecte le nom de la fonction ?)

	// Cas avec des erreurs
	prefs4 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
		{3, 1, 2},
	}
	orderedAlts = []Alternative{1, 2}
	TieBreak = TieBreakFactory(orderedAlts)

	SWFMajority = SWFFactory(MajoritySWF, TieBreak)

	res4, err4 := SWFMajority(prefs4)

	if err4 == nil {
		t.Errorf("an error should be returned")
	}
	if res4 != nil {
		t.Errorf("no count result should returned, %T computed", res4)
	}
}

// ---------------------------------------------
// FONCTIONS DE TEST POUR SCFFACTORY
// ---------------------------------------------

func TestSCFFactory(t *testing.T) {
	// Test pour Majority
	prefs1 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
		{3, 1, 2},
	}
	orderedAlts := []Alternative{1, 2, 3}
	TieBreak := TieBreakFactory(orderedAlts)

	SCFMajority := SCFFactory(MajoritySCF, TieBreak)

	res1, err1 := SCFMajority(prefs1)
	if err1 != nil {
		t.Errorf("no error should be returned, %s computed", err1)
	}

	if res1 != 1 {
		t.Errorf("error, result should be 1, %d computed", res1)
	}

	// Test pour Borda
	prefs2 := [][]Alternative{
		{1, 2, 3},
		{1, 3, 2},
		{3, 2, 1},
		{3, 1, 2},
	}

	orderedAlts = []Alternative{2, 3, 1}
	TieBreak = TieBreakFactory(orderedAlts)

	SCFBorda := SCFFactory(BordaSCF, TieBreak)
	res2, err2 := SCFBorda(prefs2)

	if err2 != nil {
		t.Error(err2)
	}

	if res2 != 3 {
		t.Errorf("error, result should be 3, %d computed", res2)
	}

	// Pas possible de tester pour ApprovalSCF car inputs non compatibles

	// // Test pour CondorcetWinner (mais TieBreak ne sera jamais utilisée)
	// prefs3 := [][]Alternative{
	// 	{1, 2, 3},
	// 	{1, 3, 2},
	// 	{3, 2, 1},
	// }
	// orderedAlts = []Alternative{2, 3, 1}
	// TieBreak = TieBreakFactory(orderedAlts)

	// CondorcetWin := SCFFactory(CondorcetWinner, TieBreak)
	// res3, err3 := CondorcetWin(prefs3)

	// if err3 != nil {
	// 	t.Error(err3)
	// }

	// if res3 != 1 {
	// 	t.Errorf("error, result should be 1, %d computed", res3)
	// }

	// // Test pour Copeland
	// prefs4 := [][]Alternative{
	// 	{1, 2, 3, 4},
	// 	{2, 3, 4, 1},
	// 	{4, 3, 1, 2},
	// }

	// orderedAlts = []Alternative{3, 2, 4, 1, 5}
	// TieBreak = TieBreakFactory(orderedAlts)

	// SCFCopeland := SCFFactory(CopelandSCF, TieBreak)
	// res4, err4 := SCFCopeland(prefs4)

	// if err4 != nil {
	// 	t.Error(err4)
	// }
	// if res4 != 3 {
	// 	t.Errorf("error, result for 1 should be 3, %d computed", res4)
	// }

	// Test pour STV
	// Les TieBreak doivent être gérés dans la fonction. Pas possible avec la Factory
	// (ou faire un cas à part si on détecte le nom de la fonction ?)

	// Cas avec des erreurs
	prefs5 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
		{3, 1, 2},
	}
	orderedAlts = []Alternative{1, 2, 1, 3}
	TieBreak = TieBreakFactory(orderedAlts)

	SCFMajority = SCFFactory(MajoritySCF, TieBreak)

	res5, err5 := SCFMajority(prefs5)

	if err5 == nil {
		t.Errorf("an error should be returned")
	}
	if res5 != -1 {
		t.Errorf("no count result should returned, %T computed", res5)
	}

}

func TestCondorcetWinner(t *testing.T) {
	prefs1 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	prefs2 := [][]Alternative{
		{1, 2, 3},
		{2, 3, 1},
		{3, 1, 2},
	}

	res1, _ := CondorcetWinner(prefs1)
	res2, _ := CondorcetWinner(prefs2)

	if len(res1) == 0 || res1[0] != 1 {
		t.Errorf("error, 1 should be the only best alternative for prefs1 but %d computed", res1[0])
	}
	if len(res2) != 0 {
		t.Errorf("no best alternative for prefs2")
	}
}

// ---------------------------------------------
// FONCTIONS DE TEST POUR COPELAND
// ---------------------------------------------

func TestCopelandSWF(t *testing.T) {
	// Cas avec une préférence
	prefs1 := [][]Alternative{
		{1, 2, 3, 4},
		{2, 3, 4, 1},
		{4, 3, 1, 2},
	}

	res1, err1 := CopelandSWF(prefs1)

	if err1 != nil {
		t.Errorf("no error should be returned, %s computed", err1)
	}

	if res1[1] != -1 {
		t.Errorf("error, result for 1 should be -1, %d computed", res1[1])
	}
	if res1[2] != 1 {
		t.Errorf("error, result for 2 should be 1, %d computed", res1[2])
	}
	if res1[3] != 1 {
		t.Errorf("error, result for 3 should be 1, %d computed", res1[3])
	}
	if res1[4] != -1 {
		t.Errorf("error, result for 4 should be -1, %d computed", res1[4])
	}

	// Cas avec une erreur
	prefs2 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 4},
		{3, 2, 1},
	}

	res2, err2 := CopelandSWF(prefs2)

	if err2 == nil {
		t.Errorf("an error should be returned")
	}
	if res2 != nil {
		t.Errorf("no count result should returned, %T computed", res2)
	}
}

func TestCopelandSCF(t *testing.T) {
	// Cas avec une préférence
	prefs1 := [][]Alternative{
		{1, 2, 3},
		{3, 1, 2},
		{3, 1, 2},
	}

	res1, err1 := CopelandSCF(prefs1)

	if err1 != nil {
		t.Errorf("no error should be returned, %s computed", err1)
	}
	if res1[0] != 3 {
		t.Errorf("the best Alternative should be 3, %d computed", res1[0])
	}

	// Cas avec deux préférences
	prefs2 := [][]Alternative{
		{1, 2, 3, 4},
		{2, 3, 4, 1},
		{4, 3, 1, 2},
	}

	res2, err2 := CopelandSCF(prefs2)

	if err2 != nil {
		t.Errorf("no error should be returned, %s computed", err2)
	}
	if !((res2[0] == 3 && res2[1] == 2) || (res2[0] == 2 && res2[1] == 3)) {
		t.Errorf("error, alternatives returned should be 2 and 3, %T computed", res2)
	}

	// Cas avec une erreur
	prefs3 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 1},
		{3, 2, 1},
	}

	res3, err3 := CopelandSCF(prefs3)

	if err3 == nil {
		t.Errorf("an error should be returned")
	}
	if len(res3) != 0 {
		t.Errorf("no count result should returned, %T computed", res3)
	}
}

func TestSTV_SWF(t *testing.T) {
	prefs1 := [][]Alternative{
		{1, 2, 3},
		{3, 1, 2},
		{3, 1, 2},
	}

	res, err := STV_SWF(prefs1)

	if res[3] != 1 {
		t.Errorf("the winner should be 3")
	}

	if err != nil {
		t.Errorf("no error should be returned")
	}

}
