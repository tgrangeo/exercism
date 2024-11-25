export class Robot {
  private _name: string = "";
  static oldNames: string[] = [];

  constructor() {
    this.generateName();
  }

  private getRandomBetween(min: number, max: number): number {
    return Math.floor(Math.random() * (max - min) + min);
  }

  private generateName(): void {
    let newName = ""
    do {
      newName = String.fromCharCode(this.getRandomBetween(65, 91)) +          String.fromCharCode(this.getRandomBetween(65, 91)) +     this.getRandomBetween(0, 10).toString() + this.getRandomBetween(0, 10).toString() + this.getRandomBetween(0, 10).toString();
  } while (Robot.oldNames.includes(newName));

     this._name = newName;
    Robot.oldNames.push(this._name);
  }

  public get name(): string {
    if (this._name === "") {
      this.generateName();
    }
    return this._name;
  }

  public resetName(): void {
    this._name = "";
    this.generateName();
  }

  public static releaseNames(): void {
    Robot.oldNames = [];
  }
}
