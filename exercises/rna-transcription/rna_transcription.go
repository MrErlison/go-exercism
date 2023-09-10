package strand

// ToRNA returns determine the RNA complement of a given DNA sequence
func ToRNA(dna string) string {
	var translate map[rune]string = map[rune]string{'G': "C", 'C': "G", 'T': "A", 'A': "U"}
	var rna string
	for _, v := range dna {
		if value, exist := translate[v]; exist {
			rna += value
		}
	}
	return rna
}
