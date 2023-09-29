package comsoc

import (
	"errors"
)

type Alternative int
type Profile [][]Alternative
type Count map[Alternative]int

// renvoie l'indice ou se trouve alt dans prefs
func rank(alt Alternative, prefs []Alternative) int {
	for idx := range prefs {
		if prefs[idx] == alt {
			return idx
		}
	}
	return -1
}

// renvoie vrai ssi alt1 est préférée à alt2
// renvoie faux si alt1 ou alt2 n'est pas une alternative de prefs
func isPref(alt1, alt2 Alternative, prefs []Alternative) bool {
	if rank(alt1, prefs) == -1 || rank(alt2, prefs) == -1 {
		return false
	}

	return rank(alt1, prefs) < rank(alt2, prefs)
}

// renvoie les meilleures alternatives pour un décompte donné
// Vérifier avec plusieurs tests
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
	for _, value := range alts {
		if value == alt {
			return true
		}
	}
	return false
}

// vérifie les préférences d'un agent, par ex. qu'ils sont tous complets et que chaque alternative n'apparaît qu'une seule fois
func checkProfile(prefs []Alternative, alts []Alternative) error {
	// Check if no candidat is missing
	if len(prefs) != len(alts) {
		errors.New("Uncomplete profil")
	}

	// check if all candidates are the one in alts
	var verif_alts []Alternative
	for _, value := range prefs {
		if contains(alts, value) {
			verif_alts = append(verif_alts, value)
		} else {
			errors.New("Unexpected alternative")
		}
	}
	// check if there is duplicates
	return nil
}

// initialise le count, c'est-à-dire créé une clé pour chaque alternative et leur donne 0 comme valeur
func initCount(p Profile) (count Count) {
	count = make(Count)
	for _, profile := range p {
		for i := 0; i < len(profile); i++ {
			_, check := count[profile[i]]
			// check = true if the value already exists in count
			if !check {
				count[profile[i]] = 0
			}
		}
	}
	return count
}
