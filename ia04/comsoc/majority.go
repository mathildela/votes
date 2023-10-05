package comsoc

/*
func MajoritySWF(p Profile) (count Count, err error)
func MajoritySCF(p Profile) (bestAlts []Alternative, err error)
*/

func MajoritySWF(p Profile) (count Count, err error) {
	//vérifications
	err = checkProfileAlternative(p, p[0])
	if err != nil {
		return nil, err
	}

	count = Count{}
	for cpt := 0; cpt < len(p[0]); cpt++ {
		count[p[0][cpt]] = 0
	}

	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p[0]); j++ {
			if j == 0 {
				count[p[i][j]] += 1
			}
		}
	}

	return count, err
}

func MajoritySCF(p Profile) (bestAlts []Alternative, err error) {
	var count Count
	count, err = MajoritySWF(p)
	return maxCount(count), err
}
