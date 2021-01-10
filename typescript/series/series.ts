class Series {
  readonly digits: number[];

  constructor(input: string) {
    this.digits = input.split("").map(Number);
  }

  slices(len: number): number[][] {
    if (len > this.digits.length)
      throw new Error(
        `"${this.digits} doesn't contain any series of length ${len}`
      );

    const slices: number[][] = [];
    for (let i = 0; i < this.digits.length - len + 1; i++)
      slices.push(this.digits.slice(i, i + len));
    return slices;
  }
}

export default Series;
