const dnaToRna = {
  G: 'C',
  C: 'G',
  T: 'A',
  A: 'U',
} as const;

type DnaNucleotide = keyof typeof dnaToRna;

const isDnaNucleotide = (s: string): s is DnaNucleotide => s in dnaToRna;

class Transcriptor {
  toRna(dnaStrand: string): string {
    let rnaStrand = '';
    for (const letter of dnaStrand) {
      if (!isDnaNucleotide(letter)) {
        throw new Error('Invalid input DNA.');
      }
      rnaStrand += dnaToRna[letter];
    }
    return rnaStrand;
  }
}

export default Transcriptor;
