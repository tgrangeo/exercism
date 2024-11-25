//
// This is only a SKELETON file for the 'Bank Account' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export class ValueError extends Error {
  constructor() {
    super('Bank account error')
  }
}

export class BankAccount {
  private _balance:number = 0;
  private isOpen:boolean = false;

  constructor() {}

  open(): void {
    if (this.isOpen) throw new ValueError();
    this.isOpen = true;
  }

  close(): void {
    if (!this.isOpen) throw new ValueError();
    this.isOpen = false;
    this._balance = 0;
  }

  deposit(add:number): void {
    if(this.isOpen && add > 0 ) this._balance += add;
    else throw new ValueError();
  }

  withdraw(remove:number): void {
    if (this.isOpen && remove > 0 && this.balance - remove >= 0)                   this._balance -= remove;
    else throw new ValueError();
  }

  get balance(): number {
    if (!this.isOpen) {
      throw new ValueError();
    }
    return this._balance;
  }
}
