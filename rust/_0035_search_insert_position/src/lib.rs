pub fn search_insert(nums: Vec<i32>, target: i32) -> i32 {
    if nums.len() == 0 {
        return 0;
    }

    let (mut left, mut mid, mut right) = (0, 0, nums.len());
    while left < right {
        mid = left + (right - left) / 2;
        if nums[mid] == target {
            return mid as i32;
        }

        if nums[mid] < target {
            left = mid + 1;
        } else {
            right = mid;
        }
    }

    if left != nums.len() && nums[left] == target {
        return left as i32;
    }

    if nums[mid] <= target {
        mid += 1;
    }

    return mid as i32;
}

#[cfg(test)]
mod tests {
    #[test]
    fn search_insert() {
        struct Case {
            input: Vec<i32>,
            target: i32,
            want: i32,
        }

        let cases = vec![
            Case { input: vec![1, 3, 5, 6, 8], target: 7, want: 4 },
            Case { input: vec![1, 3, 5, 6, 8], target: 5, want: 2 },
            Case { input: vec![-12, -8, -2, 0, 3, 12, 33, 48, 59, 64, 72], target: -5, want: 2 },
            Case { input: vec![1, 3, 5, 6], target: 5, want: 2 },
            Case { input: vec![1, 3, 5, 6], target: 2, want: 1 },
            Case { input: vec![1, 3, 5, 6], target: 7, want: 4 },
            Case { input: vec![1, 3, 5, 6, 8], target: 7, want: 4 },
            Case { input: vec![1, 3, 5, 6], target: 0, want: 0 },
            Case { input: vec![1], target: 0, want: 0 },
            Case { input: vec![], target: 0, want: 0 },
        ];

        for case in cases.iter() {
            let got = super::search_insert(case.input.clone(), case.target);
            assert_eq!(case.want, got);
        }
    }
}
