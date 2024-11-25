interface students{
  [key:number]:string[];
}

export class GradeSchool {
  private _students:students;

  constructor(){
    this._students = {}
  }

  private sort(){
    for (let i in this._students){
      this._students[i].sort();
    }
  }

  private removeOld(name:string){
    for (let i in this._students){
      if (this._students[i].includes(name))
        this._students[i].splice(this._students[i].indexOf(name), 1)
    }
  }
  
  roster() {
    return JSON.parse(JSON.stringify(this._students));
  }

  add(name:string,grade:number) {
     if (!this._students[grade]) {
      this._students[grade] = []; 
    }
    this.removeOld(name);
    this._students[grade].push(name);
    this.sort();
    return this._students[grade];
  }

  grade(grade:number) {
    return this.roster()[grade] ?? [];
  }
}