class Pangram {
  constructor(private readonly text: string) {}

  isPangram(): boolean {
    return new Set(this.text.toLowerCase().match(/[a-z]/g)).size >= 26;
  }
}

export default Pangram;
