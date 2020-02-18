use std::collections::HashSet;

pub fn find(sum: u32) -> HashSet<[u32; 3]> {
    let mut result = HashSet::new();

    for a in 1..sum / 3 {
        for b in a + 1..sum - 2 * a - 1 {
            let c = sum - b - a;
            if a.pow(2) + b.pow(2) == c.pow(2) && a + b + c == sum {
                result.insert([a, b, c]);
            }
        }
    }

    result
}
