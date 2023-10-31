package comsoc

func CondorcetWinner(p Profile) (bestAlts []Alternative, err error) {
	err = checkProfileAlternative(p, p[0])
	if err != nil {
		return nil, err
	} else {
		count := initCount(p)
		for _, a1 := range p[0] {
			for _, a2 := range p[0] {
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
						count[a1]++
					} else {
						count[a2]++
					}
				}
			}
		}
		bestAlts := maxCount(count)
		if count[bestAlts[0]] != 2*(len(p[0])-1) {
			// on multiplie par nb candidats moins 1 car lors du compte chaque candidat a "affronté" tous les autres candidats
			return nil, nil
		} else {
			return bestAlts, nil
		}
	}
}
