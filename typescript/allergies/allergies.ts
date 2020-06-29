const allergenToCode = {
  eggs: 1,
  peanuts: 2,
  shellfish: 4,
  strawberries: 8,
  tomatoes: 16,
  chocolate: 32,
  pollen: 64,
  cats: 128,
} as const;

type Allergen = keyof typeof allergenToCode;

class Allergies {
  constructor(private readonly score: number) {}

  allergicTo(allergen: Allergen): boolean {
    return this.list().includes(allergen);
  }

  list(): Allergen[] {
    return (Object.entries(allergenToCode) as [Allergen, number][])
      .filter(([_, c]) => this.getBit(Math.log2(c)) === 1)
      .map(([a, _]) => a);
  }

  private getBit(position: number): number {
    return (this.score >> position) & 1;
  }
}

export default Allergies;
