package comsoc

func CondorcetWinner(p Profile) (bestAlts []Alternative, err error) {
	alts := getAlternatives(p)
	err = checkProfileAlternative(p, alts)
	if err != nil {
		return nil, err
	} else {
		count := initCount(p)
		for _, a1 := range alts {
			for _, a2 := range alts {
				sum1, sum2 := 0, 0
				if a1 != a2 {
					for _, pref := range p {
						if isPref(a1, a2, pref) {
							sum1++
						} else {
							sum2++
						}
					}
					if sum1 > sum2 {
						//vérification des préférences. Comment gérer égalité ?
						//ici on considère que pour qu'un candidat soit gagnant de condorcet, il doit
						//gagner strictement contre tous les candidats
						count[a1]++
					} else {
						count[a2]++
					}
				}
			}
		}
		bestAlts := maxCount(count)
		if count[bestAlts[0]] != (len(alts)-1)*(len(alts)-1) {
			// on multiplie par nb candidats moins 1 car lors du compte chaque candidat a "affronté" tous les autres candidats
			return nil, nil
		} else {
			return bestAlts, nil
		}
	}
}
