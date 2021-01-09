const mapping = [
  [1000, "M"],
  [900, "CM"],
  [500, "D"],
  [400, "CD"],
  [100, "C"],
  [90, "XC"],
  [50, "L"],
  [40, "XL"],
  [10, "X"],
  [9, "IX"],
  [5, "V"],
  [4, "IV"],
  [1, "I"],
] as const;

class RomanNumerals {
  static roman(n: number): string {
    let numeral = "";
    for (const [value, symbol] of mapping)
      while (n >= value) {
        numeral += symbol;
        n -= value;
      }
    return numeral;
  }
}

export default RomanNumerals;
