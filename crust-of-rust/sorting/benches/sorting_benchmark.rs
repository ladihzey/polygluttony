use crust_of_rust::sorting;
use criterion::{
    black_box,
    criterion_group,
    criterion_main,
    Criterion,
    BenchmarkId,
};


criterion_group!(benches, sorting_benchmarks);
criterion_main!(benches);

fn sorting_benchmarks(c: &mut Criterion) {
    let mut group = c.benchmark_group("sortings");

    for size in [128, 256, 512, 1024, 2048, 4096, 8192, 16384] {
        let vec = init_data(size);

        group.bench_with_input(
            BenchmarkId::new("bubble_sort", size),
            &vec,
            |b, input| b.iter(|| {
                let mut vec = input.to_owned();
                sorting::bubble_sort(&mut vec)
            }),
        );
        group.bench_with_input(
            BenchmarkId::new("insertion_sort", size),
            &vec,
            |b, input| b.iter(|| {
                let mut vec = input.to_owned();
                sorting::bubble_sort(&mut vec)
            }),
        );
        group.bench_with_input(
            BenchmarkId::new("selection_sort", size),
            &vec,
            |b, input| b.iter(|| {
                let mut vec = input.to_owned();
                sorting::bubble_sort(&mut vec)
            }),
        );
        group.bench_with_input(
            BenchmarkId::new("quick_sort", size),
            &vec,
            |b, input| b.iter(|| {
                let mut vec = input.to_owned();
                sorting::bubble_sort(&mut vec)
            }),
        );
        group.bench_with_input(
            BenchmarkId::new("std_unstable_sort", size),
            &vec,
            |b, input| b.iter(|| {
                let mut vec = input.to_owned();
                vec.sort_unstable()
            }),
        );
    }

    group.finish();
}

fn init_data(data_size: usize) -> Vec<i32> {
    let mut vec: Vec<i32> = Vec::with_capacity(data_size);
    for _ in 0..vec.capacity() {
        vec.push(rand::random());
    }
    black_box(vec)
}
