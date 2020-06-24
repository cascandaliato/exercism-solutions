enum Planet {
  Earth,
  Jupiter,
  Mars,
  Mercury,
  Neptune,
  Saturn,
  Uranus,
  Venus,
}

class Space {
  private static readonly orbitalPeriods: { [key in Planet]: number } = {
    [Planet.Earth]: 31557600,
    get [Planet.Jupiter]() {
      return 11.862615 * this[Planet.Earth];
    },
    get [Planet.Mars]() {
      return 1.8808158 * this[Planet.Earth];
    },
    get [Planet.Mercury]() {
      return 0.2408467 * this[Planet.Earth];
    },
    get [Planet.Neptune]() {
      return 164.79132 * this[Planet.Earth];
    },
    get [Planet.Saturn]() {
      return 29.447498 * this[Planet.Earth];
    },
    get [Planet.Uranus]() {
      return 84.016846 * this[Planet.Earth];
    },
    get [Planet.Venus]() {
      return 0.61519726 * this[Planet.Earth];
    },
  };

  constructor(public readonly seconds: number) {}

  onEarth = (): number => this.roundedAge(Planet.Earth);
  onJupiter = (): number => this.roundedAge(Planet.Jupiter);
  onMars = (): number => this.roundedAge(Planet.Mars);
  onMercury = (): number => this.roundedAge(Planet.Mercury);
  onNeptune = (): number => this.roundedAge(Planet.Neptune);
  onSaturn = (): number => this.roundedAge(Planet.Saturn);
  onUranus = (): number => this.roundedAge(Planet.Uranus);
  onVenus = (): number => this.roundedAge(Planet.Venus);

  private roundedAge(planet: Planet): number {
    return Math.round((100 * this.seconds) / Space.orbitalPeriods[planet]) / 100;
  }
}

export default Space;
