const ResistanceValues = {
  black: 0,
  brown: 1,
  red: 2,
  orange: 3,
  yellow: 4,
  green: 5,
  blue: 6,
  violet: 7,
  grey: 8,
  white: 9,
};

type Color = keyof typeof ResistanceValues;

export class ResistorColor {
  constructor(private colors: [Color, Color]) {}

  value = (): number => {
    const [firstColor, secondColor] = this.colors;
    const firstDigit = ResistanceValues[firstColor];
    const secondDigit = ResistanceValues[secondColor];
    return firstDigit * 10 + secondDigit;
  };
}
