export function score(x: number, y: number): number {
  if (x * x + y * y <= 1) return 10;
  else if (x * x + y * y <= 25) return 5;
  else if (x * x + y * y <= 100) return 1;
  else return 0;
}
