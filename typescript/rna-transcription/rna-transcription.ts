const dnaToRna: { [key: string]: string } = {
  G: 'C',
  C: 'G',
  T: 'A',
  A: 'U',
};

class Transcriptor {
  toRna(dnaStrand: string): string {
    let rnaStrand = '';
    for (const letter of dnaStrand) {
      if (!(letter in dnaToRna)) {
        throw new Error('Invalid input DNA.');
      }
      rnaStrand += dnaToRna[letter];
    }
    return rnaStrand;
  }
}

export default Transcriptor;
