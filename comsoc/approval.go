package comsoc

import (
	"errors"
)

func ApprovalSWF(p Profile, thresholds []int) (count Count, err error) {
	if len(thresholds) == 0 || len(thresholds) != len(p) {
		err := errors.New("Uncomplete threshold list")
		return nil, err
	}
	err = checkProfileAlternative(p, p[0])
	if err != nil {
		return nil, err
	} else {
		count = initCount(p)
		// On parcourt les préférences des votants
		for idx1 := 0; idx1 < len(p); idx1++ {
			// On ajoute 1 pour toutes les préférences sélectionnés par le seuil
			for idx2 := 0; idx2 < thresholds[idx1]; idx2++ {
				count[p[idx1][idx2]] = count[p[idx1][idx2]] + 1
			}
		}
	}
	return count, nil
}

func ApprovalSCF(p Profile, thresholds []int) (bestAlts []Alternative, err error) {
	count, err := ApprovalSWF(p, thresholds)
	return maxCount(count), err
}
