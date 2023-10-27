package comsoc

import "fmt"

func CopelandSWF(p Profile) (count Count, err error) {
	fmt.Println("copeland")
	err = checkProfileAlternative(p, p[0])
	if err != nil {
		return nil, err
	} else {
		count := initCount(p)
		for _, a1 := range p[0] {
			for _, a2 := range p[0] {
				sum1, sum2 := 0, 0
				if a1 != a2 {
					for _, pref := range p {
						if isPref(a1, a2, pref) {
							sum1++
						} else {
							sum2++
						}
					}
					if sum1 > sum2 {
						//vérification des préférences. Comment gérer égalité ?
						// ici on considère que si égalité, pas de point attribué
						count[a1]++
						count[a2]--
					} else {
						if sum1 < sum2 {
							count[a2]++
							count[a1]--
						}
					}
				}
			}
		}
		fmt.Println("count:", count)
		// on divise les scores par 2 car chaque duo de candidats se sont affrontés 2 fois
		for key, value := range count {
			count[key] = value / 2
		}
		return count, nil
	}
}

func CopelandSCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := CopelandSWF(p)
	return maxCount(count), err
}
