use super::cursor::Cursor;

pub enum List<T> {
    Link { item: T, next: Box<List<T>> },
    Tail { item: T },
    None,
}

impl<T> List<T> where T: Copy {
    pub fn push(&mut self, x: T) {
        match self {
            Self::None => self.to_tail(x),
            Self::Tail { .. } => self.to_link(x),
            Self::Link { next, .. } => next.push(x),
        }
    }

    pub fn pop(&mut self) -> Option<T> {
        match self {
            Self::None => None,

            Self::Tail { item } => {
                let item = *item;
                self.to_none();
                Some(item)
            },

            Self::Link { item, next } => {
                let mut n = Box::new(Self::None);
                let item = *item;
                std::mem::swap(next, &mut n);
                self.to_next(*n);
                Some(item)
            },
        }
    }

    fn to_next(&mut self, next: List<T>) {
        *self = next;
    }

    fn to_none(&mut self) {
        *self = std::mem::replace(self, List::None);
    }

    fn to_tail(&mut self, x: T) {
        *self = match self {
            Self::None => Self::Tail { item: x },
            Self::Link { .. } => Self::Tail { item: x },
            _ => panic!("Trying to convert tail to tail"),
        }
    }

    fn to_link(&mut self, x: T) {
        *self = match self {
            Self::Tail { item } => {
                Self::Link {
                    item: *item,
                    next: Box::new(Self::Tail { item: x }),
                }
            }
            _ => panic!("Couldn't convert to a link"),
        }
    }

    pub fn new() -> List<T> {
        List::None
    }
}

impl<T> IntoIterator for List<T> where T: Copy {
    type Item = T;
    type IntoIter = Cursor<T>;

    fn into_iter(self) -> Self::IntoIter {
        Cursor {
            curr: self,
        }
    }
}