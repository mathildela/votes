package comsoc

func CondorcetWinner(p Profile) (bestAlts []Alternative, err error) {
	//deux à deux : on stocke le vainqueur

	err = checkProfileAlternative(p, p[0])
	if err != nil {
		return nil, err
	}

	c := initCount(p)

	var alt1, alt2 Alternative

	for i := 0; i < len(p[0]); i++ {
		for j := i + 1; j < len(p[0]); j++ {
			alt1 = p[0][i]
			alt2 = p[0][j]

			sum1, sum2 := 0, 0

			for idx_pref := 0; idx_pref < len(p); idx_pref++ {
				if isPref(alt1, alt2, p[idx_pref]) {
					//fmt.Printf("In %v, entre %d et %d, le vainqueur est : %d\n", p[idx_pref], alt1, alt2, alt1)
					sum1 += 1
				} else {
					//fmt.Printf("In %v, entre %d et %d, le vainqueur est : %d\n", p[idx_pref], alt1, alt2, alt2)
					sum2 += 1
				}
			}
			if sum1 > sum2 {
				c[alt1] += 1
			} else {
				c[alt2] += 1
			}
		}
	}

	//fmt.Println(c)
	for key, val := range c {
		if val == len(p[0])-1 {
			return []Alternative{key}, nil
		}
	}
	return []Alternative{}, nil

}
