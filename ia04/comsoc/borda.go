package comsoc

//func BordaSWF(p Profile) (count Count, err error)
//func BordaSCF(p Profile) (bestAlts []Alternative, err error)

func BordaSWF(p Profile) (count Count, err error) {
	//vérifications
	err = checkProfileAlternative(p, p[0])
	if err != nil {
		return nil, err
	}

	count = Count{}
	for cpt := 0; cpt < len(p[0]); cpt++ {
		count[p[0][cpt]] = 0 //initialise la map
	}

	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p[0]); j++ {
			count[p[i][j]] += len(p[0]) - j - 1
		}
	}

	return count, err

}

func BordaSCF(p Profile) (bestAlts []Alternative, err error) {
	var count Count
	count, err = BordaSWF(p)
	return maxCount(count), err
}
