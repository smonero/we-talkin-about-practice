// For each word, count the number of each letter
// compare it to the previous one
// if it doesn't match, put it in the results
// before starting next iteration in loop, save the curr result
// into the prev result
//
// note - because we are case insensitive we can use an array
// instead of a dict

impl Solution {
    pub fn remove_anagrams(words: Vec<String>) -> Vect<String> {
        let mut res = vec![];
        let mut prev_count_dict = [0; 26];

        for word in words {
            let mut current_letters_count_dict = [0; 26];
            word.bytes().for_each(|b| current_letters_count_dict[(b - b'a') as usize] += 1);

            if prev_count_dict != current_letters_count_dict {
                res.push(word);
            }

            prev_count_dict = current_letters_count_dict;
        }
        res
    }
}