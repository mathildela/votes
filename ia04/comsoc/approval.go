package comsoc

// func ApprovalSWF(p Profile, thresholds []int) (count Count, err error)
// func ApprovalSCF(p Profile, thresholds []int) (bestAlts []Alternative, err error)

func ApprovalSWF(p Profile, thresholds []int) (count Count, err error) {
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
			if j < thresholds[i] {
				count[p[i][j]] += 1
			}
		}
	}

	return count, err
}

func ApprovalSCF(p Profile, thresholds []int) (bestAlts []Alternative, err error) {
	var count Count
	count, err = ApprovalSWF(p, thresholds)
	return maxCount(count), err
}
