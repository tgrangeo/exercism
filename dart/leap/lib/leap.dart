class Leap {
  bool leapYear(int y){
    if (y % 4 == 0 || (y %100 == 0 && y % 400 == 0)){
      return true;
    } else {
      return false;
    }
  }
}
