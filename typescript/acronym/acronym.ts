export default class Acronym {
  public static parse(phrase: string): string {
    return (
      phrase
        // split by: space, dash, lowercase letter followed by uppercase letter
        .split(/[\s-]|[a-z](?=[A-Z])/)
        .map((word) => word[0])
        .join("")
        .toUpperCase()
    );
  }
}
