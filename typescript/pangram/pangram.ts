export function isPangram(str:string):boolean {
 const alphabet = 'abcdefghijklmnopqrstuvwxyz'
  str = str.toLowerCase();
  for (let letter of alphabet){
    if ( str.indexOf(letter) === -1)
      return false
  }
  return true
}
