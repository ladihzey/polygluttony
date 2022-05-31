pub fn sort<T: Ord>(arr: &mut[T]) {
    sort_slice(arr, 0, arr.len().saturating_sub(1));
}

fn sort_slice<T: Ord>(slice: &mut [T], start: usize, end: usize) {
    if start < end {
        let pivot_index = partition(slice, start, end);
        sort_slice(slice, start, pivot_index.saturating_sub(1));
        sort_slice(slice, pivot_index + 1, end);
    }
}

fn partition<T: Ord>(slice: &mut [T], start: usize, end: usize) -> usize {
    let mut index = start;

    for i in start..end {
        if slice[i] <= slice[end] {
            slice.swap(index, i);
            index += 1;
        }
    }

    slice.swap(index, end);
    index
}
