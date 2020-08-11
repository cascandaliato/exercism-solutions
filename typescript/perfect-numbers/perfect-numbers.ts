enum NicomachusCategory {
  Perfect = 'perfect',
  Abudant = 'abundant',
  Deficient = 'deficient',
}

export default class PerfectNumbers {
  static classify(number: number): NicomachusCategory {
    if (number <= 0) {
      throw new Error('Classification is only possible for natural numbers.');
    }

    const factorsSum = PerfectNumbers.factorsSum(number);
    if (factorsSum == number) {
      return NicomachusCategory.Perfect;
    } else if (factorsSum > number) {
      return NicomachusCategory.Abudant;
    } else {
      return NicomachusCategory.Deficient;
    }
  }

  private static factorsSum(number: number): number {
    let sum = 0;
    for (let i = 1; i < number; i++) {
      if (number % i == 0) {
        sum += i;
      }
    }
    return sum;
  }
}
