//copied
impl Solution {
    pub fn remove_anagrams(words: Vec<String>) -> Vec<String> {
        let mut res = vec![];
        for w in words {
            let n = res.len();
            if n == 0 {
                res.push(w);
                continue
            }
            let tail = &res[n-1];
            let mut ana: Vec<char> = tail.chars().collect();
            ana.sort();
            let mut w_ana = w.chars().collect::<Vec<char>>();
            w_ana.sort();
            if ana != w_ana {
                res.push(w);
            }
        }
        res
    }
}