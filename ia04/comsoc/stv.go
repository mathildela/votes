package comsoc

import "fmt"

func STV_SWF(p Profile) (Count, error) {
	tour := 0
	min := 100000
	var alt_min Alternative
	p_copy := copyProfile(p)

	for tour < len(p[0])-1 {
		fmt.Printf("Tour %d\n", tour)
		tour += 1
		min = 100000

		fmt.Printf("profil actuel : %v\n", p_copy)

		c, err := MajoritySWF(p_copy)

		fmt.Printf("count %v\n", c)

		if err != nil {
			return nil, err
		}
		for key, val := range c {
			//fmt.Printf("key : %d, value : %d, min : %d\n", key, val, min)
			if val < min {
				alt_min = key
				min = val
			}
		}

		fmt.Println(alt_min)

		p_copy, err = removeAlt(p_copy, alt_min)
		if err != nil {
			return nil, err
		}
		fmt.Printf("p_copy après remove : %v\n", p_copy)
	}

	fmt.Printf("dernier profile : %v\n", p_copy)
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
