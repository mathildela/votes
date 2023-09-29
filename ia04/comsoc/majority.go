package comsoc

// Sans tiebreak pour le moment
func MajoritySWF(p Profile) (count Count, err error) {
	// First initialize the map with 0
	count = initCount(p)

	for _, profile := range p {
		count[profile[0]]++
	}
	return count, nil
	// gérer les erreurs (vérifier que le profil est complet en amont)
}

func MajoritySCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := MajoritySWF(p)
	return maxCount(count), err
}
