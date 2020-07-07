type Nucleotide = 'A' | 'C' | 'G' | 'T';

class NucleotideCount {
  static nucleotideCounts(strand: string) {
    const count: Record<Nucleotide, number> = { A: 0, C: 0, G: 0, T: 0 };

    function isNucleotide(letter: string): letter is Nucleotide {
      return letter in count;
    }

    for (const letter of strand) {
      if (!isNucleotide(letter)) {
        throw new Error('Invalid nucleotide in strand');
      }
      count[letter]++;
    }
    return count;
  }
}

export default NucleotideCount;
