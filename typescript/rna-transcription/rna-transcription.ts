export function toRna(dna: string): string {
  let rna = '';
  for (let i = 0; i < dna.length; i++) {
    if (dna.charAt(i) === 'G')
      rna += 'C';
    else if (dna.charAt(i) === 'C')
      rna += 'G';
    else if (dna.charAt(i) === 'T')
      rna += 'A';
    else if (dna.charAt(i) === 'A')
      rna += 'U';
    else
      throw 'Invalid input DNA.';
  }
  return rna;
}
