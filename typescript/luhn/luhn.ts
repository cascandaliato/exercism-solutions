export default class Luhn {
  private constructor() {}

  static valid(text: string): boolean {
    text = Luhn.stripWhitespace(text);

    if (text.length <= 1 || !Luhn.isNumeric(text)) return false;

    return Luhn.checksum(text) % 10 === 0;
  }

  private static stripWhitespace(text: string): string {
    return text.replace(/\s/g, "");
  }

  private static isNumeric(text: string): boolean {
    return /^\d+$/.test(text);
  }

  private static checksum(text: string): number {
    return text
      .split("")
      .map(Number)
      .map((d, idx) => {
        if (idx % 2 === 1) {
          const doubled = d * 2;
          return doubled > 9 ? doubled - 9 : doubled;
        } else return d;
      })
      .reduce((a, b) => a + b);
  }
}
