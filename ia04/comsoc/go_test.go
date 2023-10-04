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

func Test_minCount(t *testing.T) {
	//func minCount(count Count) (worsetAlts []Alternative)
	//type Count map[Alternative]int

	var count_1 Count = Count{0: 4, 1: 2, 2: 0, 3: 2, 4: 5, 5: 2}
	tab_1 := [...]Alternative{2}
	worseAlts_1 := tab_1[:]
	res_1 := minCount(count_1)

	var count_2 Count = Count{0: 4, 1: 2, 2: 1, 3: 2, 4: 4, 5: 1}
	tab_21 := [...]Alternative{2, 5}
	tab_22 := [...]Alternative{5, 2}
	worseAlts_21 := tab_21[:]
	worseAlts_22 := tab_22[:]
	res_2 := minCount(count_2)

	var count_3 Count = Count{}
	tab_3 := [...]Alternative{}
	worseAlts_3 := tab_3[:]
	res_3 := minCount(count_3)

	if !isEqualSliceAlt(res_1, worseAlts_1) {
		t.Errorf("Output incorrect")
	}

	if !isEqualSliceAlt(res_2, worseAlts_21) && !isEqualSliceAlt(res_2, worseAlts_22) {
		t.Errorf("Output incorrect")
	}

	if !isEqualSliceAlt(res_3, worseAlts_3) {
		t.Errorf("Output incorrect")
	}

}

func Test_checkProfile(t *testing.T) {
	pref1 := []Alternative{1, 2, 3}
	pref2 := []Alternative{1, 3}
	pref3 := []Alternative{3, 1, 4}
	pref4 := []Alternative{1, 2, 1}
	alts := []Alternative{1, 2, 3}

	if checkProfile(pref1, alts) != nil {
		t.Errorf("error = %s", checkProfile(pref1, alts))
	}
	if checkProfile(pref2, alts) == nil {
		t.Errorf("error = %s", checkProfile(pref2, alts))
	}
	if checkProfile(pref3, alts) == nil {
		t.Errorf("error = %s", checkProfile(pref3, alts))
	}
	if checkProfile(pref4, alts) == nil {
		t.Errorf("error = %s", checkProfile(pref4, alts))
	}
}

func Test_checkProfileAlternative(t *testing.T) {
	prefs1 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}
	prefs2 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1, 4},
	}
	prefs3 := [][]Alternative{
		{1, 2, 3},
		{1, 4, 3},
		{3, 2, 1},
	}
	prefs4 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 1, 1},
	}
	alts := []Alternative{1, 2, 3}

	if checkProfileAlternative(prefs1, alts) != nil {
		t.Errorf("error = %s", checkProfileAlternative(prefs1, alts))
	}
	if checkProfileAlternative(prefs2, alts) == nil {
		t.Errorf("error = %s", checkProfileAlternative(prefs2, alts))
	}
	if checkProfileAlternative(prefs3, alts) == nil {
		t.Errorf("error = %s", checkProfileAlternative(prefs3, alts))
	}

	if checkProfileAlternative(prefs4, alts) == nil {
		t.Errorf("error = %s", checkProfileAlternative(prefs4, alts))
	}
}

func Test_initCount(t *testing.T) {
	prefs := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}
	res := initCount(prefs)

	if res[1] != 0 {
		t.Errorf("error, result for 1 should be 0, %d computed", res[1])
	}
	if res[2] != 0 {
		t.Errorf("error, result for 2 should be 0, %d computed", res[2])
	}
	if res[3] != 0 {
		t.Errorf("error, result for 3 should be 0, %d computed", res[3])
	}
}

func TestAllDifferent(t *testing.T) {
	count1 := make(Count)
	count1 = Count{1: 2, 2: 1, 3: 3}
	if allDifferentCount(count1) == false {
		t.Errorf("Function should return true")
	}

	count2 := make(Count)
	count2 = Count{1: 2, 2: 1, 3: 2}
	if allDifferentCount(count2) == true {
		t.Errorf("Function should return false")
	}
}

// ---------------------------------------------
// FONCTIONS DE TEST POUR MAJORITY
// ---------------------------------------------

func TestMajoritySWF(t *testing.T) {
	// Cas avec une préférence
	prefs1 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	res1, err1 := MajoritySWF(prefs1)

	if err1 != nil {
		t.Errorf("no error should be returned, %s computed", err1)
	}

	if res1[1] != 2 {
		t.Errorf("error, result for 1 should be 2, %d computed", res1[1])
	}
	if res1[2] != 0 {
		t.Errorf("error, result for 2 should be 0, %d computed", res1[2])
	}
	if res1[3] != 1 {
		t.Errorf("error, result for 3 should be 1, %d computed", res1[3])
	}

	// Cas avec une erreur
	prefs2 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 4},
		{3, 2, 1},
	}

	res2, err2 := MajoritySWF(prefs2)

	if err2 == nil {
		t.Errorf("an error should be returned")
	}
	if res2 != nil {
		t.Errorf("no count result should returned, %T computed", res2)
	}

}

func TestMajoritySCF(t *testing.T) {

	// Cas avec une seule préférence
	prefs1 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	res1, err1 := MajoritySCF(prefs1)

	if err1 != nil {
		t.Error(err1)
	}

	if len(res1) != 1 || res1[0] != 1 {
		t.Errorf("error, 1 should be the only best Alternative")
	}

	// Cas avec plusieurs préférences
	prefs2 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
		{3, 1, 2},
	}

	res2, err2 := MajoritySCF(prefs2)

	if err2 != nil {
		t.Error(err2)
	}

	if len(res2) != 2 {
		t.Errorf("error, there should be two best alternatives")
	}

	if !((res2[0] == 1 && res2[1] == 3) || (res2[0] == 3 && res2[1] == 1)) {
		t.Errorf("error, alternatives returned should be 1 and 3, %T computed", res2)
	}

	// Cas avec une erreur
	prefs3 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1, 4},
	}

	res3, err3 := MajoritySCF(prefs3)

	if err3 == nil {
		t.Errorf("an error should be returned")
	}

	if len(res3) != 0 {
		t.Errorf("no canditate should be returned, %T computed", res3)
	}
}

// ---------------------------------------------
// FONCTIONS DE TEST POUR BORDA
// ---------------------------------------------

func TestBordaSWF(t *testing.T) {
	// Cas avec une préférence
	prefs1 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	res1, err1 := BordaSWF(prefs1)

	if err1 != nil {
		t.Errorf("no error should be returned, %s computed", err1)
	}

	if res1[1] != 4 {
		t.Errorf("error, result for 1 should be 4, %d computed", res1[1])
	}
	if res1[2] != 3 {
		t.Errorf("error, result for 2 should be 3, %d computed", res1[2])
	}
	if res1[3] != 2 {
		t.Errorf("error, result for 3 should be 2, %d computed", res1[3])
	}

	// Cas avec une erreur
	prefs2 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 4},
		{3, 2, 1},
	}

	res2, err2 := BordaSWF(prefs2)

	if err2 == nil {
		t.Errorf("an error should be returned")
	}
	if res2 != nil {
		t.Errorf("no count result should returned, %T computed", res2)
	}
}

func TestBordaSCF(t *testing.T) {
	// Cas avec une préférence
	prefs1 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	res1, err1 := BordaSCF(prefs1)

	if err1 != nil {
		t.Error(err1)
	}

	if len(res1) != 1 || res1[0] != 1 {
		t.Errorf("error, 1 should be the only best Alternative")
	}

	// Cas avec plusieurs préférences
	prefs2 := [][]Alternative{
		{1, 2, 3},
		{1, 3, 2},
		{3, 2, 1},
		{3, 1, 2},
	}

	res2, err2 := BordaSCF(prefs2)

	if err2 != nil {
		t.Error(err2)
	}

	if len(res2) != 2 {
		t.Errorf("error, there should be two best alternatives")
	}

	if !((res2[0] == 1 && res2[1] == 3) || (res2[0] == 3 && res2[1] == 1)) {
		t.Errorf("error, alternatives returned should be 1 and 3, %T computed", res2)
	}

	// Cas avec une erreur
	prefs3 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1, 4},
	}

	res3, err3 := BordaSCF(prefs3)

	if err3 == nil {
		t.Errorf("an error should be returned")
	}

	if len(res3) != 0 {
		t.Errorf("no candidate should be returned, %T computed", res3)
	}
}

// ---------------------------------------------
// FONCTIONS DE TEST POUR VOTE PAR APPROBATION
// ---------------------------------------------

func TestApprovalSWF(t *testing.T) {
	// Cas avec une préférence
	prefs1 := [][]Alternative{
		{1, 2, 3},
		{1, 3, 2},
		{2, 3, 1},
	}
	thresholds1 := []int{2, 1, 2}

	res1, _ := ApprovalSWF(prefs1, thresholds1)

	if res1[1] != 2 {
		t.Errorf("error, result for 1 should be 2, %d computed", res1[1])
	}
	if res1[2] != 2 {
		t.Errorf("error, result for 2 should be 2, %d computed", res1[2])
	}
	if res1[3] != 1 {
		t.Errorf("error, result for 3 should be 1, %d computed", res1[3])
	}

	// Cas d'une liste de seuils erronée
	prefs2 := [][]Alternative{
		{1, 2, 3},
		{1, 3, 2},
		{2, 3, 1},
	}
	thresholds2 := []int{2, 2}

	res2, err2 := ApprovalSWF(prefs2, thresholds2)

	if res2 != nil {
		t.Errorf("error, no count should be returned, computed %T", res2)
	}
	if err2 == nil {
		t.Errorf("an error should be returned")
	}

	// Cas d'une erreur dans les profils
	prefs3 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 4},
		{3, 2, 1},
	}

	thresholds3 := []int{2, 1, 2}

	res3, err3 := ApprovalSWF(prefs3, thresholds3)

	if err3 == nil {
		t.Errorf("an error should be returned")
	}
	if res3 != nil {
		t.Errorf("no count result should returned, %T computed", res2)
	}
}

func TestApprovalSCF(t *testing.T) {
	// Cas avec une préférence
	prefs1 := [][]Alternative{
		{1, 3, 2},
		{1, 2, 3},
		{2, 1, 3},
	}
	thresholds1 := []int{2, 1, 2}

	res1, err1 := ApprovalSCF(prefs1, thresholds1)

	if err1 != nil {
		t.Error(err1)
	}
	if len(res1) != 1 || res1[0] != 1 {
		t.Errorf("error, 1 should be the only best Alternative")
	}

	// Cas avec deux préférences
	prefs2 := [][]Alternative{
		{1, 2, 3},
		{1, 3, 2},
		{2, 3, 1},
	}
	thresholds2 := []int{2, 1, 2}

	res2, err2 := ApprovalSCF(prefs2, thresholds2)

	if err2 != nil {
		t.Error(err2)
	}
	if len(res2) != 2 {
		t.Errorf("error, there should be two best alternatives")
	}

	if !((res2[0] == 1 && res2[1] == 2) || (res2[0] == 2 && res2[1] == 1)) {
		t.Errorf("error, alternatives returned should be 1 and 2, %T computed", res2)
	}

	// Cas avec une erreur
	prefs3 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1, 4},
	}

	res3, err3 := BordaSCF(prefs3)

	if err3 == nil {
		t.Errorf("an error should be returned")
	}

	if len(res3) != 0 {
		t.Errorf("no candidate should be returned, %T computed", res3)
	}
}

// ---------------------------------------------
// FONCTIONS DE TEST POUR TIE_BREAK
// ---------------------------------------------

func TestTieBreakFactory(t *testing.T) {

	TieBreak := TieBreakFactory(nil)
	if TieBreak != nil {
		t.Errorf("No function should be returned because no input")
	}

	orderedAlts := []Alternative{1, 2, 3, 4}
	TieBreak = TieBreakFactory(orderedAlts)
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
		t.Error("No order should be known for the alternatives given")
	}
	if res != -1 {
		t.Errorf("No alternative should be chosen, %d computed", res)
	}

	orderedAlts = []Alternative{1, 2, 3, 4}
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
		{2, 1, 3},
		{2, 3, 1},
	}
	orderedAlts := []Alternative{1, 2, 3}
	TieBreak := TieBreakFactory(orderedAlts)

	SWFMajority := SWFFactory(MajoritySWF, TieBreak)

	res1, err1 := SWFMajority(prefs1)
	if err1 != nil {
		t.Errorf("no error should be returned, %s computed", err1)
	}

	if res1[1] != 4 {
		t.Errorf("error, result for 1 should be 3, %d computed", res1[1])
	}
	if res1[2] != 3 {
		t.Errorf("error, result for 2 should be 0, %d computed", res1[2])
	}
	if res1[3] != 2 {
		t.Errorf("error, result for 3 should be 2, %d computed", res1[3])
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
		t.Errorf("error, result for 1 should be 5, %d computed", res2[1])
	}
	if res2[2] != 2 {
		t.Errorf("error, result for 2 should be 0, %d computed", res2[2])
	}
	if res2[3] != 6 {
		t.Errorf("error, result for 3 should be 2, %d computed", res2[3])
	}

	// La factory n'est pas adaptée au vote par approbation car les inputs de ApprovalSWF
	// ne correspond pas au input de la fonction swf en entrée de SWFFactory

	// Test pour Copeland
	prefs3 := [][]Alternative{
		{1, 2, 3, 4},
		{2, 3, 4, 1},
		{4, 3, 1, 2},
	}

	orderedAlts = []Alternative{3, 2, 4, 1}
	TieBreak = TieBreakFactory(orderedAlts)

	SWFCopeland := SWFFactory(CopelandSWF, TieBreak)
	res3, err3 := SWFCopeland(prefs3)

	if err3 != nil {
		t.Error(err3)
	}
	if res3[1] != -1 {
		t.Errorf("error, result for 1 should be -1, %d computed", res3[1])
	}
	if res3[2] != 1 {
		t.Errorf("error, result for 2 should be 1, %d computed", res3[2])
	}
	if res3[3] != 2 {
		t.Errorf("error, result for 3 should be 2, %d computed", res3[3])
	}
	if res3[4] != 0 {
		t.Errorf("error, result for 4 should be 0, %d computed", res3[4])
	}

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

	// Test pour CondorcetWinner (mais TieBreak ne sera jamais utilisée)
	prefs3 := [][]Alternative{
		{1, 2, 3},
		{1, 3, 2},
		{3, 2, 1},
	}
	orderedAlts = []Alternative{2, 3, 1}
	TieBreak = TieBreakFactory(orderedAlts)

	CondorcetWin := SCFFactory(CondorcetWinner, TieBreak)
	res3, err3 := CondorcetWin(prefs3)

	if err3 != nil {
		t.Error(err3)
	}

	if res3 != 1 {
		t.Errorf("error, result should be 1, %d computed", res3)
	}

	// Test pour Copeland
	prefs4 := [][]Alternative{
		{1, 2, 3, 4},
		{2, 3, 4, 1},
		{4, 3, 1, 2},
	}

	orderedAlts = []Alternative{3, 2, 4, 1, 5}
	TieBreak = TieBreakFactory(orderedAlts)

	SCFCopeland := SCFFactory(CopelandSCF, TieBreak)
	res4, err4 := SCFCopeland(prefs4)

	if err4 != nil {
		t.Error(err4)
	}
	if res4 != 3 {
		t.Errorf("error, result for 1 should be 3, %d computed", res4)
	}

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

// ---------------------------------------------
// FONCTIONS DE TEST POUR CONDORCETWINNER
// ---------------------------------------------

func TestCondorcetWinner(t *testing.T) {
	// Cas avec un gagnant de condorcet
	prefs1 := [][]Alternative{
		{1, 2, 3},
		{1, 3, 2},
		{3, 2, 1},
	}
	// 1 est le gagnant de condorcet

	res1, err1 := CondorcetWinner(prefs1)

	if err1 != nil {
		t.Errorf("no error should be returned, %s computed", err1)
	}
	if len(res1) == 0 || res1[0] != 1 {
		t.Errorf("error, result should be 1, %d computed", res1[0])
	}

	// Cas avec aucun gagnant de condorcet
	prefs2 := [][]Alternative{
		{1, 2, 3},
		{2, 3, 1},
		{3, 1, 2},
	}

	res2, err2 := CondorcetWinner(prefs2)

	if err2 != nil {
		t.Errorf("no error should be returned, %s computed", err2)
	}
	if len(res2) != 0 {
		t.Errorf("there should be no condorcer winner, %T computed", res2)
	}

	// Cas avec une erreur
	prefs3 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 1},
		{3, 2, 1},
	}

	res3, err3 := CondorcetWinner(prefs3)

	if err3 == nil {
		t.Errorf("an error should be returned")
	}
	if res3 != nil {
		t.Errorf("no count result should returned, %T computed", res3)
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

// ---------------------------------------------
// FONCTIONS DE TEST POUR STV
// ---------------------------------------------
