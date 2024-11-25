export class Clock {
  private _hour = 0;
  private _min = 0;
  constructor(hour: number, minute?: number) {
    if (hour === 24) this._hour = 0;
    else if (hour >=0 && hour < 24) this._hour = hour;
    else this._hour = hour % 24;
    if (minute === 60 ){
      this._hour += minute / 60;
      this._min = minute % 60;
    } else this._min = minute ?  minute % 60 : 0;
  }

  public toString(): unknown {
   return `${this._hour.toString().padStart(2, '0')}:${this._min.toString().padStart(2, '0')}`;
  }

  public plus(minutes: unknown): Clock {
    throw new Error('Remove this statement and implement this function')
  }

  public minus(minutes: unknown): Clock {
    throw new Error('Remove this statement and implement this function')
  }

  public equals(other: unknown): unknown {
    throw new Error('Remove this statement and implement this function')
  }
}
