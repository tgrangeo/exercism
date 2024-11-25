export class DiffieHellman {
  p!: number;
  g!: number;
  n!: number;
  totient!: number;

  isPrime = (num: number) => {
    for (let i = 2, s = Math.sqrt(num); i <= s; i++) {
      if (num % i === 0) return false;
    }
    return num > 1;
  };

  isPrimitiveRoot = (g: number, p: number) => {
    const requiredSet = new Set<number>();
    for (let x = 1; x < p; x++) {
      requiredSet.add(Math.pow(g, x) % p);
    }
    return requiredSet.size === p - 1;
  };

  constructor(p: number, g: number) {
    if (p < 1 || p > 9999 || g < 1 || g > 9999) throw "Invalid parameters";
    if (!this.isPrime(p) || !this.isPrimitiveRoot(g, p))
      throw "g must be a primitive root modulo p and p must be prime";
    this.p = p;
    this.g = g;
    this.n = p * g;
    this.totient = (p - 1) * (g - 1);
  }

  public getPublicKey(privateKey: number): number {
    if (privateKey < 2 || privateKey >= this.p) {
      throw "Invalid private key";
    }
    return Math.pow(this.g, privateKey) % this.p;
  }

  public getSecret(theirPublicKey: number, myPrivateKey: number): number {
    if (
      typeof theirPublicKey !== "number" ||
      typeof myPrivateKey !== "number"
    ) {
      throw new Error("Public and private keys must be numbers");
    }
    return Math.pow(theirPublicKey, myPrivateKey) % this.p;
  }
}
