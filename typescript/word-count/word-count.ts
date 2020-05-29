export default class Words {
  count(text: string): Map<string, number> {
    return text
      .trim()
      .toLowerCase()
      .split(/[ \n\t]+/)
      .reduce((map, word) => map.set(word, (map.get(word) || 0) + 1), new Map<string, number>());
  }
}
