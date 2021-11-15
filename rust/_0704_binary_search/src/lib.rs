pub fn search(nums: Vec<i32>, target: i32) -> i32 {
    if nums.len() > 0 {
        let mut left = 0;
        let mut right = (nums.len() - 1) as i32;
        while left <= right {
            let mid = left + (right - left) / 2;
            let mid_idx = mid as usize;
            if nums[mid_idx] == target {
                return mid;
            }

            if nums[mid_idx] < target {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
    }

    return -1;
}

#[cfg(test)]
mod tests {
    #[test]
    fn search() {
        struct Case {
            nums: Vec<i32>,
            target: i32,
            want: i32,
        }

        let cases = vec![
            Case { nums: vec![-1, 0, 3, 5, 9, 12], target: 9, want: 4 },
            Case { nums: vec![-1, 0, 3, 5, 9, 12], target: 2, want: -1 },
            Case { nums: vec![-1, 0, 3, 5, 9, 12], target: 3, want: 2 },
            Case { nums: vec![5], target: -5, want: -1 },
            Case { nums: vec![], target: 3, want: -1 },
        ];

        for case in cases.iter() {
            let got = super::search(case.nums.clone(), case.target);
            assert_eq!(case.want, got);
        }
    }
}
