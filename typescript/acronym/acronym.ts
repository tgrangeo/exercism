export function parse(phrase: string): string {
  let ret = phrase[0];
  for (let i = 0; i < phrase.length; i++) {
    if (phrase[i-1] === ' ' || phrase[i-1] === '-' || ( phrase.charCodeAt(i) <= 90 && phrase.charCodeAt(i) >= 65 && phrase.charCodeAt(i-1) >= 97 && phrase.charCodeAt(i - 1) <= 122))
        ret += phrase[i];
  }
  return ret.toUpperCase()
}
