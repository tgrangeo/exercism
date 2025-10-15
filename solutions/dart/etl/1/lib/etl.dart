import 'dart:core';

class Etl {
    Map<String,int> transform(Map<String,List<String>> legacy){
    final res = <String,int>{};
    legacy.forEach((k,v){
      v.forEach((s){
         res[s.toLowerCase()] = int.parse(k);
      });
    });
    return res;
  }
}
