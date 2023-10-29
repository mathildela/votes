package comsoc

func STV_SWF(p Profile) (Count, error) {
	p_copy := copyProfile(p)

	// Pour la majorité, on applique le tiebreak suivant : "l'alternative avec le plus petit indice gagne"
	orderedAlts := make([]Alternative, len(p[0]))
	for i := 0; i < len(p[0]); i++ {
		orderedAlts[i] = Alternative(i + 1)
	}
	tiebreak := TieBreakFactory(orderedAlts)

	for tour := 0; tour < len(p[0])-1; tour++ {
		//fmt.Printf("Tour %d\n", tour)
		//fmt.Printf("profil actuel : %v\n", p_copy)
		MajorityWithTieBreak := SWFFactory(MajoritySWF, tiebreak)
		alts, err := MajorityWithTieBreak(p_copy)

		//fmt.Printf("count %v\n", c)

		if err != nil {
			return nil, err
		} else {
			p_copy, err = removeAlt(p_copy, alts[len(alts)-1])
			if err != nil {
				return nil, err
			}
		}

		//fmt.Printf("p_copy après remove : %v\n", p_copy)
	}

	//fmt.Printf("dernier profile : %v\n", p_copy)
	winner := p_copy[0][0]
	final_count := initCount(p)
	final_count[winner] += 1

	return final_count, nil
}

func STV_SCF(p Profile) (bestAlts []Alternative, err error) {
	var count Count
	count, err = STV_SWF(p)
	return maxCount(count), err
}
