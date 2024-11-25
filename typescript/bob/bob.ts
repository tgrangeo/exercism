export function hey(str: string): string {
  if (!str.trim()) {
    return "Fine. Be that way!"
  }else if (str.trim().charAt(str.trim().length - 1) === '?' && str === str.toUpperCase() && /[A-Z]/.test(str)) {
    return "Calm down, I know what I'm doing!"
  } else if (str.trim().charAt(str.trim().length - 1) === '?' ){
    return "Sure."
  } else if (str === str.toUpperCase() && /[A-Z]/.test(str)) {
    return "Whoa, chill out!"
  }
  return "Whatever."
}
