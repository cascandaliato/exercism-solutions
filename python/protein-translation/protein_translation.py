codons = {'Cysteine': ['UGU', 'UGC'],
          'Leucine': ['UUA', 'UUG'],
          'Methionine': ['AUG'],
          'Phenylalanine': ['UUU', 'UUC'],
          'Serine': ['UCU', 'UCC', 'UCA', 'UCG'],
          'Tryptophan': ['UGG'],
          'Tyrosine': ['UAU', 'UAC'],
          'STOP': ['UAA', 'UAG', 'UGA']}

sequences = {sequence: codon for codon in codons for sequence in codons[codon]}


def proteins(strand):
    polypeptide = []
    for sequence in [strand[i:i+3] for i in range(0, len(strand), 3)]:
        if sequences[sequence] == 'STOP':
            break
        else:
            polypeptide.append(sequences[sequence])
    return polypeptide
