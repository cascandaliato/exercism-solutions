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

const allergens = Object.keys(allergenToCode) as Allergen[];

class Allergies {
  constructor(private readonly score: number) {}

  allergicTo(allergen: Allergen): boolean {
    return !!(this.score & allergenToCode[allergen]);
  }

  list(): Allergen[] {
    return allergens.filter(this.allergicTo.bind(this));
  }
}

export default Allergies;
