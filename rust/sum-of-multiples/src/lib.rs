fn remove_duplicates(vec: Vec<u32>) -> Vec<u32> {
    let mut resultat = Vec::new();
    for valeur in vec {
        if !resultat.contains(&valeur) {
            resultat.push(valeur);
        }
    }
    resultat
}

pub fn sum_of_multiples(limit: u32, factors: &[u32]) -> u32 {
    let mut res: Vec<u32> = Vec::new();
    for &f in factors{
        if f == 0 {
            continue;
        }
        for i in (f..limit).step_by(f as usize) {
            res.push(i);
        }
    }
    res = remove_duplicates(res);
    res.iter().sum()
}
