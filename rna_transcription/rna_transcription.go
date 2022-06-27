// Package hamming is a solution to the "RNA Transcription" problem
// https://exercism.org/tracks/go/exercises/rna-transcription/
package strand

func ToRNA(dna string) string {
	var rnaStr string
	for i := 0; i < len(dna); i++ {
		switch string(dna[i]) {
		case "G":
			rnaStr += "C"
		case "C":
			rnaStr += "G"
		case "T":
			rnaStr += "A"
		case "A":
			rnaStr += "U"
		}
	}
	return rnaStr
}
