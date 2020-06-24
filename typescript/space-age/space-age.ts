type Planet = 'Earth' | 'Jupiter' | 'Mars' | 'Mercury' | 'Neptune' | 'Saturn' | 'Uranus' | 'Venus';

class Space {
  private static readonly earthYearSeconds = 31557600;

  private static readonly earthOrbitalPeriods: Record<Planet, number> = {
    Earth: 1.0,
    Jupiter: 11.862615,
    Mars: 1.8808158,
    Mercury: 0.2408467,
    Neptune: 164.79132,
    Saturn: 29.447498,
    Uranus: 84.016846,
    Venus: 0.61519726,
  };

  constructor(public readonly seconds: number) {}

  onEarth = (): number => this.age('Earth');
  onJupiter = (): number => this.age('Jupiter');
  onMars = (): number => this.age('Mars');
  onMercury = (): number => this.age('Mercury');
  onNeptune = (): number => this.age('Neptune');
  onSaturn = (): number => this.age('Saturn');
  onUranus = (): number => this.age('Uranus');
  onVenus = (): number => this.age('Venus');

  private age(planet: Planet): number {
    const secondsInYear = Space.earthOrbitalPeriods[planet] * Space.earthYearSeconds;
    return Math.round((100 * this.seconds) / secondsInYear) / 100;
  }
}

export default Space;
