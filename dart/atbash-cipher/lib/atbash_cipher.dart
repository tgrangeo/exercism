class AtbashCipher {
  String encode(String str) {
    String res = "";
    for (var i = 0; i < str.length; i++) {
      var char = str[i].toLowerCase();
      int code = char.codeUnitAt(0);
      if (code >= 97 && code <= 122) {
        res += String.fromCharCode(122 - (code - 97));
      } else if (code >= 48 && code <= 57) {
        res += char;
      }
    }
    return res;
  }

  String decode(String str) {
    return encode(str);
  }
}
