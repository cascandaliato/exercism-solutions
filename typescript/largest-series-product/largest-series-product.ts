export default class Series {
  private digits: number[];

  constructor(digits: string) {
    if (!/^\d*$/.test(digits)) throw new Error("Invalid input.");
    this.digits = digits.split("").map(Number);
  }

  largestProduct(len: number): number {
    if (len < 0) throw new Error("Invalid input.");
    else if (len > this.digits.length)
      throw new Error("Slice size is too big.");

    let maxProduct = 0;
    for (let i = 0; i < this.digits.length - len + 1; i++) {
      maxProduct = Math.max(maxProduct, this.multiplyDigits(i, i + len));
    }
    return maxProduct;
  }

  private multiplyDigits(fromIdx: number, toIdx: number): number {
    return this.digits.slice(fromIdx, toIdx).reduce((a, b) => a * b, 1);
  }
}
