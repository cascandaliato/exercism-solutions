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
  private readonly allergens: Allergen[] = [];

  constructor(score: number) {
    const codes = Object.values(allergenToCode).sort((a, b) => b - a);

    const allAllergens = codes.reduce((a, b) => a + b, 0);
    score &= allAllergens;

    for (const code of codes) {
      if (score >= code) {
        const allergen = Object.entries(allergenToCode).filter(([, c]) => c === code)[0][0];
        this.allergens.unshift(allergen as Allergen);
        score -= code;
      }
    }
  }

  allergicTo(allergen: Allergen): boolean {
    return this.allergens.includes(allergen);
  }

  list(): Allergen[] {
    return [...this.allergens];
  }
}

export default Allergies;
