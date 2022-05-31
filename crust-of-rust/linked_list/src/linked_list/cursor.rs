use super::list::List;

pub struct Cursor<T> {
    pub curr: List<T>,
}

impl<T> Iterator for Cursor<T> where T: Copy {
    type Item = T;

    fn next(&mut self) -> Option<Self::Item> {
        match self.curr {
            List::None => None,
            List::Tail { item } => {
                self.curr = List::None;
                Some(item)
            },
            List::Link { item, ref mut next } => {
                let mut n = Box::new(List::None);
                std::mem::swap(next, &mut n);
                self.curr = *n;
                Some(item)
            },
        }
    }
}
