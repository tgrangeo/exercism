export function isArmstrongNumber(number: number | bigint): boolean {
  let str = number.toString();
  let sum: bigint = BigInt(0);
  for (let i = 0; i < str.length; i++) {
    let digit = BigInt(str.charAt(i));
    sum += digit ** BigInt(str.length);
  }
  return sum === BigInt(number);
}
