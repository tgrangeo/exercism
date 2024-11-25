export function age(planet: string, seconds: number): number {
  const earth: number = 31557600;
  let ret = 0;
  switch (planet){
    case 'earth':
      ret = seconds / earth;
      break;
    case 'mercury':
      ret = seconds / (earth * 0.2408467);
      break;
    case 'venus':
      ret = seconds / (earth * 0.61519726);
      break;
    case 'mars':
      ret = seconds / (earth * 1.8808158);
      break;
    case 'jupiter':
      ret = seconds / (earth * 11.862615);
      break;
    case 'saturn':
      ret = seconds / (earth * 29.447498);
      break;
    case 'uranus':
      ret = seconds / (earth *  84.016846);
      break;
    case 'neptune':
      ret = seconds / (earth * 164.79132);
      break;
  }
  return Math.round(ret * 100) / 100;
}
