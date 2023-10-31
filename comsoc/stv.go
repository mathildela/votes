package comsoc

func STV_SWF(p Profile) (Count, error) {
	p_copy := copyProfile(p)
	count := initCount(p)

	// Pour la majorité, on applique le tiebreak suivant : "l'alternative avec le plus petit indice gagne"
	orderedAlts := make([]Alternative, len(p[0]))
	for i := 0; i < len(p[0]); i++ {
		orderedAlts[i] = Alternative(i + 1)
	}
	tiebreak := TieBreakFactory(orderedAlts)

	for tour := 0; tour < len(p[0])-1; tour++ {
		MajorityWithTieBreak := SWFFactory(MajoritySWF, tiebreak)
		alts, err := MajorityWithTieBreak(p_copy)

		if err != nil {
			return nil, err
		} else {
			p_copy, err = removeAlt(p_copy, alts[len(alts)-1])
			if err != nil {
				return nil, err
			}
			for _, value := range p_copy[0] {
				count[value] += 1
			}
		}
	}

	return count, nil
}

func STV_SCF(p Profile) (bestAlts []Alternative, err error) {
	var count Count
	count, err = STV_SWF(p)
	return maxCount(count), err
}
