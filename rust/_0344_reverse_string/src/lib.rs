pub fn reverse_string(s: &mut Vec<char>) {
    if s.len() < 2 {
        return;
    }

    let mut idx = 0;
    loop {
        let pair_idx = s.len() - idx - 1;
        if pair_idx <= idx {
            break;
        }

        let c = s[idx];
        s[idx] = s[pair_idx];
        s[pair_idx] = c;
        idx += 1;
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn reverse_string() {
        struct Case {
            input: Vec<char>,
            want: Vec<char>,
        }

        let cases = vec![
            Case { input: vec!['h', 'e', 'l', 'l', 'o'], want: vec!['o', 'l', 'l', 'e', 'h'] },
            Case { input: vec!['H', 'a', 'n', 'n', 'a', 'h'], want: vec!['h', 'a', 'n', 'n', 'a', 'H'] },
            Case { input: vec!['A'], want: vec!['A'] },
            Case { input: vec![], want: vec![] },
        ];

        for case in cases.iter() {
            let mut got = case.input.clone();
            super::reverse_string(&mut got);
            assert_eq!(case.want, got);
        }
    }
}
