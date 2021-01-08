class Isogram {
  static isIsogram(text: string): boolean {
    const normalized = Isogram.normalize(text);
    return Isogram.countUniqueChars(normalized) === normalized.length;
  }

  private static normalize(text: string): string {
    return text.toLowerCase().replace(/[^a-z]/g, "");
  }

  private static countUniqueChars(text: string): number {
    return new Set(text).size;
  }
}

export default Isogram;
