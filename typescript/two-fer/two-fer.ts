export function twoFer(name:string | undefined): string {
  if (name)
    return `One for ${name}, one for me.`;
  else 
    return `One for you, one for me.`
}
