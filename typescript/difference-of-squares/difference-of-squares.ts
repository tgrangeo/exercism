export class Squares {
  n:number[] = [];
  constructor(count: number) {
    for(let i = 1; i <= count; i++){
      this.n.push(i)
    }
  }

  get sumOfSquares(): number {
   let sum = 0;
   for(let i of this.n){
    sum  += i * i;
   }
   return sum;
  }

  get squareOfSum(): number {
   let sum = 0;
   for(let i of this.n){
    sum  += i;
   }
   return sum * sum;
  }

  get difference(): number {
    return this.squareOfSum - this.sumOfSquares;
  }
}
