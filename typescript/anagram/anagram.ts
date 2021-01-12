export default class Anagram {
  constructor(private word: string) {}

  matches = (...candidates: string[]): string[] =>
    candidates.filter(this.isAnagram);

  private isAnagram = (candidate: string): boolean =>
    candidate.toLowerCase() !== this.word.toLowerCase() &&
    Anagram.normalize(candidate) === Anagram.normalize(this.word);

  private static normalize = (word: string): string =>
    [...word.toLowerCase()].sort().join("");
}
