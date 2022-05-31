pub fn sort<T: Ord>(arr: &mut [T]) {
    for i in 0..arr.len() {
        let mut min = i;
        for j in i + 1..arr.len() {
            if arr[min] > arr[j] {
                min = j;
            }
        }
        arr.swap(i, min);
    }
}
