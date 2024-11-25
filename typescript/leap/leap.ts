export function isLeap(y:number):boolean {
  if (y % 100 === 0 && y % 400 === 0) return true;
  else if (y % 100 !== 0 && y % 4 === 0) return true;
  else return false
}
