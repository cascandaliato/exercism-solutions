export default class Squares {
  squareOfSum: number = 0;
  sumOfSquares: number = 0;
  difference: number = 0;

  constructor(private num: number) {
    if (num > 0) {
      this.squareOfSum = ((num * (num + 1)) / 2) ** 2;
      this.sumOfSquares = (num * (num + 1) * (2 * num + 1)) / 6;
      this.difference = this.squareOfSum - this.sumOfSquares;
    }
  }
}
