package comsoc

func MajoritySWF(p Profile) (count Count, err error) {
	alts := getAlternatives(p)
	err = checkProfileAlternative(p, alts)
	if err != nil {
		return nil, err
	} else {
		// First initialize the map with 0
		count = initCount(p)

		for _, profile := range p {
			count[profile[0]]++
		}
		return count, nil
	}
}

func MajoritySCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := MajoritySWF(p)
	return maxCount(count), err
}
