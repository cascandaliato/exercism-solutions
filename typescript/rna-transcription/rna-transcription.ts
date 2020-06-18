type DNANucleotide = 'A' | 'C' | 'G' | 'T';
type RNANucleotide = 'A' | 'C' | 'G' | 'U';

class Transcriptor {
  private readonly dnaToRna: { [key in DNANucleotide]: RNANucleotide } = {
    A: 'U',
    C: 'G',
    G: 'C',
    T: 'A',
  };

  isDNANucleotide(s: string): s is DNANucleotide {
    return s in this.dnaToRna;
  }

  toRna(dnaStrand: string): string {
    let rnaStrand = '';
    for (const letter of dnaStrand) {
      if (!this.isDNANucleotide(letter)) {
        throw new Error('Invalid input DNA.');
      }
      rnaStrand += this.dnaToRna[letter];
    }
    return rnaStrand;
  }
}

export default Transcriptor;
