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
    return this.getBit(Math.log2(allergenToCode[allergen])) === 1;
  }

  list(): Allergen[] {
    return (Object.keys(allergenToCode) as Allergen[]).filter(this.allergicTo.bind(this));
  }

  private getBit(position: number): number {
    return (this.score >> position) & 1;
  }
}

export default Allergies;
