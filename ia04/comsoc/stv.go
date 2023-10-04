package comsoc

func STV_SWF(p Profile) (count Count, err error) {
	alts := getAlternatives(p)
	err = checkProfileAlternative(p, alts)
	if err != nil {
		return nil, err
	} else {
		n_candidat := len(p[0])
		tmp_count := make(Count)
		for i := 0; i < n_candidat; i++ {
			tmp_count, err = MajoritySWF(p)
			//TieBreak := TieBreakFactory()
			//tmp_count, err = SWFFactory(MajoritySWF, TieBreak)
			// Ici utiliser une factory pour éviter les égalités
			if err != nil {
				return nil, err
			}
			worseAlt := minCount(tmp_count)
			p, err = removeAlt(p, worseAlt[0])
			if err != nil {
				return nil, err
			}
		}
		// A la fin il ne doit rester qu'un seul candidat pour chacun des profils
	}
	return nil, nil
}
