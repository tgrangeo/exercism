export function count(subtitle: string): Map<string, number> {
  const words = subtitle.toLowerCase().match(/\b[a-zA-Z0-9]+('[a-zA-Z0-9]+)?\b/g);

  const wordCountMap = new Map<string, number>();
  if (words) {
    for (const word of words) {
      const count = wordCountMap.get(word) || 0;
      wordCountMap.set(word, count + 1);
    }
  }
  return wordCountMap;
}6