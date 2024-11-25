export class Triangle {
  a!:number;
  b!:number;
  c!:number;
  constructor(...sides:number[]) {
      this.a = sides[0];
      this.b = sides[1];
      this.c = sides[2];
  }

  istriangleValid():boolean{
    if (this.a <= 0 || this.b <= 0 || this.c <= 0)
      return false
     if (!(this.a + this.b >= this.c && this.b + this.c >= this.a && this.a + this.c >= this.b))
      return false
    return true
  }

  get isEquilateral():boolean {
     if (!this.istriangleValid())
      return false
    if (this.a === this.b && this.b === this.c)
      return true;
    else return false
  }

  get isIsosceles():boolean {
    if (!this.istriangleValid())
      return false
    if (this.a === this.b || this.a === this.c || this.b === this.c)
      return true
    else 
      return false
  }

  get isScalene() {
     if (!this.istriangleValid())
      return false
     if (this.a !== this.b && this.a !== this.c && this.b !== this.c)
      return true
    else 
      return false
  }
}
