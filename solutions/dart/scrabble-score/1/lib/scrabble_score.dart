Map<int, String> letter_score = {
    1: "aeioulnrst",
    2: "dg",
    3: "bcmp",
    4: "fhvwy",
    5: "k",
    8: "jx",
    10: "qz",
  };

int score(String str){
  int res = 0;
  for (var c in str.toLowerCase().split('')){
    letter_score.forEach((s,letters) {
      if (letters.contains(c)){
        res += s;
      };
    });
    }
  return res;
}
