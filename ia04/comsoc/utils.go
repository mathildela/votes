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

func Contains(alts []Alternative, alt Alternative) bool {
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
		return errors.New("Too much or too few candidates in preference")
	}

	var verif_alts []Alternative
	for _, value := range prefs {
		// check if there are duplicated alternatives
		if Contains(verif_alts, value) {
			return errors.New("Alternative not unique")
		} else {
			// check if there is alternative not in alts
			if !Contains(alts, value) {
				return errors.New("Unexpected alternative")
			} else {
				verif_alts = append(verif_alts, value)
			}
		}
	}
	return nil
}

func checkProfileAlternative(prefs Profile, alts []Alternative) error {
	for _, p := range prefs {
		error := checkProfile(p, alts)
		if error != nil {
			return error
		}
	}
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

// TieBreak

func TieBreakFactory(orderedAlts []Alternative) func([]Alternative) (Alternative, error) {
	if len(orderedAlts) == 0 {
		return nil
	} else {
		return func(bestAlts []Alternative) (Alternative, error) {
			if len(bestAlts) == 0 {
				err := errors.New("No alternative")
				return -1, err
			} else {
				// On vérifie que toutes les alternatives de bestAlts sont présentes dans orderedAlts
				for _, alt := range bestAlts {
					if !Contains(orderedAlts, alt) {
						err := errors.New("At least one alternative missing in orderedAlts")
						return -1, err
					}
				}

				// On vérifie qu'il n'y a pas de doublons dans bestAlts
				var verif_alts []Alternative
				for _, value := range orderedAlts {
					if Contains(verif_alts, value) {
						err := errors.New("Alternative not unique")
						return -1, err
					} else {
						verif_alts = append(verif_alts, value)
					}
				}

				for _, alt := range orderedAlts {
					if Contains(bestAlts, alt) {
						return alt, nil
					}
				}
				err := errors.New("No common alternative between orderedAlts and bestAlts")
				return -1, err
			}
		}
	}
}

func allDifferentCount(count Count) bool {
	for key1, value1 := range count {
		for key2, value2 := range count {
			if key1 != key2 && value1 == value2 {
				return false
			}
		}
	}
	return true
}

func sameCount(a Alternative, count Count) []Alternative {
	s := make([]Alternative, 0)
	for key, value := range count {
		if key != a && value == count[a] {
			s = append(s, key)
		}
	}
	return s
}

// SWF doivent renvoyer un ordre total sans égalité
// Les SWF doivent renvoyer des counts à la fin (différence avec le sujet)
func SWFFactory(swf func(p Profile) (Count, error), tiebreak func([]Alternative) (Alternative, error)) func(Profile) ([]Alternative, error) {
	if swf == nil || tiebreak == nil {
		return nil
	} else {
		return func(p Profile) ([]Alternative, error) {
			count, err := swf(p)
			res := make([]Alternative, 0)
			if err != nil {
				return nil, err
			} else {
				for len(count) != 0 {
					//fmt.Println("count og :", count)
					max := maxCount(count)
					//fmt.Println("max = ", max)
					if len(max) == 0 {
						err = errors.New("Erreur dans la fonction maxCount")
						return nil, err
					} else if len(max) == 1 {
						res = append(res, max[0])
						delete(count, max[0])
					} else {
						sameCount := sameCount(max[0], count)
						for sameCount != nil {
							alt, err := tiebreak(sameCount)
							if err != nil {
								return nil, err
							} else {
								res = append(res, alt)
								sameCount = removeElement(sameCount, alt)
								delete(count, alt)
							}
						}
					}
					//fmt.Println("res:", res)
				}
				return res, nil
			}
		}
	}
}

// SCF doivent renvoyer un seul élement
func SCFFactory(scf func(p Profile) ([]Alternative, error), tiebreak func([]Alternative) (Alternative, error)) func(Profile) (Alternative, error) {
	if scf == nil || tiebreak == nil {
		return nil
	} else {
		return func(p Profile) (Alternative, error) {
			bestAlts, err := scf(p)

			if err != nil {
				return -1, err
			} else {
				return tiebreak(bestAlts)
			}
		}
	}
}

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

		err = checkProfileAlternative(new_p, new_p[0])
		if err != nil {
			return nil, err
		} else {
			return new_p, nil
		}
	}
}
