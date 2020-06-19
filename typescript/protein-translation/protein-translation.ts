type Protein =
  | 'Methionine'
  | 'Phenylalanine'
  | 'Leucine'
  | 'Serine'
  | 'Tyrosine'
  | 'Cysteine'
  | 'Tryptophan';

class ProteinTranslation {
  private static readonly codonToProtein: { [codon: string]: Protein } = {
    AUG: 'Methionine',
    UUU: 'Phenylalanine',
    UUC: 'Phenylalanine',
    UUA: 'Leucine',
    UUG: 'Leucine',
    UCU: 'Serine',
    UCC: 'Serine',
    UCA: 'Serine',
    UCG: 'Serine',
    UAU: 'Tyrosine',
    UAC: 'Tyrosine',
    UGU: 'Cysteine',
    UGC: 'Cysteine',
    UGG: 'Tryptophan',
  };

  private static readonly stopCodons = ['UAA', 'UAG', 'UGA'];

  static proteins(rna: string): Protein[] {
    const proteins: Protein[] = [];

    for (const codon of rna.match(/.{1,3}/g) || []) {
      console.log(codon, proteins);
      if (this.stopCodons.includes(codon)) {
        break;
      }
      proteins.push(this.codonToProtein[codon]);
    }

    return proteins;
  }
}

export default ProteinTranslation;
