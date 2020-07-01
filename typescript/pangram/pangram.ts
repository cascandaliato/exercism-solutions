const NUMBER_OF_LETTERS_IN_ALPHABET = 26;
class Pangram {
  constructor(private readonly text: string) {}

  isPangram(): boolean {
    return this.countUniqueLetters() === NUMBER_OF_LETTERS_IN_ALPHABET;
  }

  private countUniqueLetters(): number {
    return new Set(this.text.toLowerCase().match(/[a-z]/g)).size;
  }
}

export default Pangram;
