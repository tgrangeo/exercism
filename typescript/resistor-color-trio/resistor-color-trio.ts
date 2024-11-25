export const colorCodes: Record<string, number> = {
  black: 0,
  brown: 1,
  red: 2,
  orange: 3,
  yellow: 4,
  green: 5,
  blue: 6,
  violet: 7,
  grey: 8,
  white: 9,
};

export function decodedResistorValue(colors?: string[]) {
  if (!colors) {
    return undefined;
  }
  let stdColor_0 = colors[0].toLowerCase();
  let stdColor_1 = colors[1].toLowerCase();
  let stdColorZeros = colors[2].toLowerCase();

  let resistanceValueStr =
    colorCodes[stdColor_0].toString() + colorCodes[stdColor_1].toString();
  let numberOfZeros = colorCodes[stdColorZeros];
  let zeros = "0".repeat(numberOfZeros);

  let resistanceValue = +`${resistanceValueStr}${zeros}`;
  let units = "ohms";
  let unitsPrefixes = ["", "kilo", "mega", "giga"];
  let unitsIndex = 0;

  if (resistanceValue) {
    while (resistanceValue % 1000 === 0 && unitsIndex < 3) {
      unitsIndex += 1;
      resistanceValue /= 1000;
    }
  }
  return `${resistanceValue} ${unitsPrefixes[unitsIndex]}${units}`;
}
