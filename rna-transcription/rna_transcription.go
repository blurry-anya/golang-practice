package strand

func ToRNA(dna string) string {
	rna := ""

	for _, v := range dna {
		switch v {
		case 'A':
			rna += "U"
		case 'C':
			rna += "G"
		case 'G':
			rna += "C"
		case 'T':
			rna += "A"
		default:
			return ""
		}
	}

	return rna
}
