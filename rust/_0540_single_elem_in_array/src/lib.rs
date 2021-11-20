pub fn single_non_duplicate(nums: Vec<i32>) -> i32 {
    let (mut left, mut right) = (0, nums.len() - 1);
    while left < right {
        let mid = left + (right - left) / 2;
        if nums[mid] == nums[mid ^ 1] {
            left = mid + 1;
        } else {
            right = mid;
        }
    }

    return nums[left];
}

#[cfg(test)]
mod tests {
    #[test]
    fn single_non_duplicate() {
        struct Case {
            input: Vec<i32>,
            want: i32,
        }

        let cases = vec![
            Case { input: vec![1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 6, 7, 7, 8, 8], want: 4 },
            Case { input: vec![0, 0, 3, 3, 7, 7, 10, 10, 11, 11, 17], want: 17 },
            Case { input: vec![0, 1, 1, 2, 2, 3, 3, 4, 4, 8, 8], want: 0 },
            Case { input: vec![0, 0, 3, 3, 7, 7, 10, 11, 11], want: 10 },
            Case { input: vec![1, 1, 2, 3, 3, 4, 4, 8, 8], want: 2 },
            Case { input: vec![22], want: 22 },
        ];

        for case in cases.iter() {
            let got = super::single_non_duplicate(case.input.clone());
            assert_eq!(case.want, got);
        }
    }
}
