class ArmstrongNumbers {
  bool isArmstrongNumber(String str){
    int res = 0;
    for (int i = 0;i < str.length; i++){
      if (str[i] != '0'){
        res += int.parse(str[i]) ^ str.length;
      }
    }
    return res == int.parse(str);
  }
}
