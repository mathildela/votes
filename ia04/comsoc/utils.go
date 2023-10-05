package comsoc

import (
	"fmt"
)

type Alternative int
type Profile [][]Alternative
type Count map[Alternative]int

// renvoie l'indice oà se trouve alt dans prefs
func rank(alt Alternative, prefs []Alternative) int {
	var ind int = -1
	for i := 0; i < len(prefs); i++ {
		if prefs[i] == alt {
			ind = i
		}
	}
	return ind
}

// renvoie vrai ssi alt1 est préférée à alt2
func isPref(alt1, alt2 Alternative, prefs []Alternative) bool {
	return rank(alt1, prefs) < rank(alt2, prefs)
}

// renvoie les meilleures alternatives pour un décomtpe donné
func maxCount(count Count) (bestAlts []Alternative) {
	var max int = -1
	bestAlts = make([]Alternative, 0)
	for key, val := range count {
		if val > max {
			bestAlts = nil //vider
			bestAlts = append(bestAlts, key)
			max = val
		} else {
			if val == max {
				bestAlts = append(bestAlts, key)
				max = val
			}
		}
	}
	return bestAlts
}

func contains(alts []Alternative, alt Alternative) bool {
	for _, s := range alts {
		if alt == s {
			return true
		}
	}
	return false
}

// vérifie les préférences d'un agent, par ex. qu'ils sont tous complets et que chaque alternative n'apparaît qu'une seule fois

func checkPrefs(prefs []Alternative, alts []Alternative) error {
	//vérification compléture
	if len(prefs) != len(alts) {
		err := fmt.Errorf("err : préférences incomplètes")
		return err
	}

	//vérification unicité alternative
	var verif []Alternative
	var err error
	for i := 0; i < len(prefs); i++ {
		if contains(verif, prefs[i]) {
			err = fmt.Errorf("err : préférences non uniques pour votant %d", i)
		}
		if !contains(alts, prefs[i]) {
			err = fmt.Errorf("err : alternative %d pas dans la liste alts", prefs[i])
		}
		verif = append(verif, prefs[i])
	}
	return err

}

// vérifie le profil donné, par ex. qu'ils sont tous complets et que chaque alternative n'apparaît qu'une seule fois par préférences
func checkProfile(prefs Profile) error {
	//vérification complétude
	lenght := len(prefs[0])
	for i := 0; i < len(prefs); i++ {
		if len(prefs[i]) != lenght {
			err := fmt.Errorf("err : profile incomplet pour %d", i)
			return err
		}
	}

	//vérification unicité alternative
	var verif []Alternative
	var err error
	for i := 0; i < len(prefs); i++ {
		verif = nil
		for j := 0; j < len(prefs[0]); j++ {
			if contains(verif, prefs[i][j]) {
				err = fmt.Errorf("err : préférences non uniques pour votant %d", i)
			}
			verif = append(verif, prefs[i][j])
		}
	}
	return err
}

// vérifie le profil donné, par ex. qu'ils sont tous complets et que chaque alternative de alts apparaît exactement une fois par préférences
func checkProfileAlternative(prefs Profile, alts []Alternative) error {
	//vérification complétude
	lenght := len(prefs[0])
	for i := 0; i < len(prefs); i++ {
		if len(prefs[i]) != lenght {
			err := fmt.Errorf("err : profile incomplet pour %d", i)
			return err
		}
	}

	//vérification unicité alternative
	var verif []Alternative
	var err error
	for i := 0; i < len(prefs); i++ {
		verif = nil
		for j := 0; j < len(prefs[0]); j++ {
			if contains(verif, prefs[i][j]) {
				err = fmt.Errorf("err : préférences non uniques pour votant %d", i)
			}
			if !contains(alts, prefs[i][j]) {
				err = fmt.Errorf("err : alternative %d pas dans la liste alts", prefs[i][j])
			}
			verif = append(verif, prefs[i][j])
		}
	}
	return err
}

func initCount(p Profile) (count Count) {
	count = Count{}
	for cpt := 0; cpt < len(p[0]); cpt++ {
		count[p[0][cpt]] = 0
	}
	return count
}

// fonction de Solenn
// func allDifferentCount(count Count) bool {
// 	for key1, value1 := range count {
// 		for key2, value2 := range count {
// 			if key1 != key2 && value1 == value2 {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

func copyProfile(source Profile) (destination Profile) {
	destination = make([][]Alternative, len(source))
	for i := range destination {
		destination[i] = make([]Alternative, len(source[0]))
	}
	for i := 0; i < len(source); i++ {
		for j := 0; j < len(source[0]); j++ {
			destination[i][j] = source[i][j]
		}
	}
	return destination
}

func removeElement(prefs []Alternative, i Alternative) []Alternative {
	var result []Alternative
	for _, v := range prefs {
		if v != i {
			result = append(result, v)
		}
	}
	return result
}

func removeAlt(p Profile, alt Alternative) (new_p Profile, err error) {
	alts := p[0]
	err = checkProfileAlternative(p, alts)
	if err != nil {
		return nil, err
	} else {
		var new_p Profile
		for _, prefs := range p {
			prefs = removeElement(prefs, alt)
			new_p = append(new_p, prefs)
		}
		err = checkProfileAlternative(new_p, alts)
		if err != nil {
			return nil, err
		} else {
			return new_p, nil
		}
	}
}
