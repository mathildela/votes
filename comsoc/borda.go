package comsoc

func BordaSWF(p Profile) (count Count, err error) {
	err = checkProfileAlternative(p, p[0])
	if err != nil {
		return nil, err
	} else {
		count = initCount(p)
		for _, profile := range p {
			for idx := 0; idx < len(profile); idx++ {
				count[profile[idx]] = count[profile[idx]] + (len(profile) - 1 - idx)
			}
		}
	}
	return count, nil
}

func BordaSCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := BordaSWF(p)
	return maxCount(count), err
}
