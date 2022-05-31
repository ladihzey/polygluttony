mod list;
mod cursor;

pub use list::List;

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn should_push_items() {
        let mut list: List<i32> = List::new();

        list.push(1);
        list.push(2);
        list.push(3);
        list.push(4);

        assert_eq!(list.pop(), Some(1));
        assert_eq!(list.pop(), Some(2));
        assert_eq!(list.pop(), Some(3));
        assert_eq!(list.pop(), Some(4));
        assert_eq!(list.pop(), None);
    }
}
