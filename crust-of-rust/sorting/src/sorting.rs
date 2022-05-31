mod bubble_sort;
mod selection_sort;
mod insertion_sort;
mod quick_sort;

pub use bubble_sort::sort as bubble_sort;
pub use selection_sort::sort as selection_sort;
pub use insertion_sort::sort as insertion_sort;
pub use quick_sort::sort as quick_sort;

#[cfg(test)]
mod tests {
    use pretty_assertions::{assert_eq};
    use test_case::test_case;
    use super::*;

    type SortFunction = fn(&mut [i32]);

    #[test_case(bubble_sort)]
    #[test_case(selection_sort)]
    #[test_case(insertion_sort)]
    #[test_case(quick_sort)]
    fn should_tolerate_empty_arr(sort: SortFunction) {
        let mut arr = [];
        sort(&mut arr);
        assert_eq!(arr, []);
    }

    #[test_case(bubble_sort)]
    #[test_case(selection_sort)]
    #[test_case(insertion_sort)]
    #[test_case(quick_sort)]
    fn should_tolerate_single_item_arr(sort: SortFunction) {
        let mut arr = [1];
        sort(&mut arr);
        assert_eq!(arr, [1]);
    }

    #[test_case(bubble_sort)]
    #[test_case(selection_sort)]
    #[test_case(insertion_sort)]
    #[test_case(quick_sort)]
    fn should_sort_arr(sort: SortFunction) {
        let mut arr = [123, 0, -100, 12, 25, 257, 1];
        sort(&mut arr);
        assert_eq!(arr, [-100, 0, 1, 12, 25, 123, 257]);
    }

    #[test_case(bubble_sort)]
    #[test_case(selection_sort)]
    #[test_case(insertion_sort)]
    #[test_case(quick_sort)]
    fn should_sort_as_std_library(sort: SortFunction) {
        let mut vec1: Vec<i32> = Vec::with_capacity(1000);
        for _ in 0..vec1.capacity() {
            vec1.push(rand::random());
        }
        let mut vec2 = vec1.to_owned();

        sort(&mut vec1);
        vec2.sort_unstable();

        assert_eq!(vec1, vec2);
    }
}
