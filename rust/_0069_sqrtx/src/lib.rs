use std::cmp::Ordering;

pub fn my_sqrt(x: i32) -> i32 {
    if x < 2 {
        return x;
    }

    let (mut left, mut right) = (0, x / 2 + 1);
    while left <= right {
        let mid = left + (right - left) / 2;
        match mid.cmp(&(x / mid)) {
            Ordering::Equal => return mid,
            Ordering::Less => left = mid + 1,
            Ordering::Greater => right = mid - 1,
        }
    }

    return right;
}

#[cfg(test)]
mod tests {
    #[test]
    fn my_sqrt() {
        struct Case {
            input: i32,
            want: i32,
        }

        let cases = vec![
            Case { input: 2147395599, want: 46339 },
            Case { input: 16, want: 4 },
            Case { input: 9, want: 3 },
            Case { input: 8, want: 2 },
            Case { input: 7, want: 2 },
            Case { input: 4, want: 2 },
            Case { input: 1, want: 1 },
            Case { input: 0, want: 0 },
        ];

        for case in cases.iter() {
            let got = super::my_sqrt(case.input);
            assert_eq!(case.want, got);
        }
    }
}
