package comsoc

import "fmt"

//retourne une fonction qui applique l'ordre
//func TieBreakFactory(orderedAlts []Alternative) (func ([]Alternative) (Alternative, error))

func TieBreakFactory(orderedAlts []Alternative) func([]Alternative) (Alternative, error) {

	f := func(bestAlts []Alternative) (Alternative, error) {
		var best Alternative
		var best_rank int = 100000
		if bestAlts == nil {
			err := fmt.Errorf("err : liste d'alternatives vide")
			return -1, err
		}
		check := checkPrefs(bestAlts, bestAlts)
		check2 := checkPrefs(orderedAlts, orderedAlts)
		fmt.Println(check2)
		if check != nil || check2 != nil {
			err := fmt.Errorf("err : listes éléments non unique")
			return -1, err
		}
		for i := 0; i < len(bestAlts); i++ {
			rank := rank(bestAlts[i], orderedAlts)
			//fmt.Printf("[%d] est de rank %d \n", bestAlts[i], rank)
			if rank == -1 {
				err := fmt.Errorf("err : alternative pas dans l'ordre donné")
				return -1, err
			}
			if rank < best_rank {
				//fmt.Printf("%d < %d \n", rank, best_rank)
				best = bestAlts[i]
				best_rank = rank
				//fmt.Printf("best rank : %d, best : %d \n", best_rank, bestAlts[i])
			}
		}
		return best, nil
	}

	return f

}

// SWF doivent renvoyer un ordre total sans égalité
// Les SWF doivent renvoyer des counts à la fin (différence avec le sujet)
func SWFFactory(swf func(p Profile) (Count, error), tiebreak func([]Alternative) (Alternative, error)) func(Profile) (Count, error) {
	if swf == nil || tiebreak == nil {
		return nil
	} else {
		return func(p Profile) (Count, error) {
			count, err := swf(p)
			if err != nil {
				return nil, err
			} else {
				allBestAlts := maxCount(count)
				bestalt, err := tiebreak(allBestAlts)
				if err != nil {
					return nil, err
				}

				for key := range count {
					if bestalt == key {
						count[key]++
					}
				}
				return count, nil
			}
		}
	}
}

// SCF doivent renvoyer un seul élement
//fonctions de Solenn
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
