package proverb

// Proverb returbs the relevant proverb
func Proverb(rhyme []string) []string {
	if len(rhyme) <= 0 {
		return nil
	}

	var proverbs []string = []string{}
	for i := 0; i < len(rhyme)-1; i++ {
		proverbs = append(proverbs, "For want of a "+rhyme[i]+" the "+rhyme[i+1]+" was lost.")
	}

	proverbs = append(proverbs, "And all for the want of a "+rhyme[0]+".")

	return proverbs
}
