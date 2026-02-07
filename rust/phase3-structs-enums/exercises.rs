// Phase 3 Exercises: Structs and Enums
// Run with: rustc exercises.rs && ./exercises

fn main() {
    exercise_1();
    exercise_2();
    exercise_3();
    exercise_4();
    exercise_5();
}

// Exercise 1: Define a Struct
fn exercise_1() {
    println!("=== Exercise 1: Structs ===");
    
    // TODO: Define a Book struct with title (String), pages (u32), read (bool)
    struct Book {
        title: String,
        pages: u32,
        read: bool,
    }
    
    let book = Book {
        title: String::from("The Rust Book"),
        pages: 500,
        read: false,
    };
    
    println!("{}: {} pages, read: {}", book.title, book.pages, book.read);
}

// Exercise 2: Methods
fn exercise_2() {
    println!("\n=== Exercise 2: Methods ===");
    
    struct Counter {
        count: i32,
    }
    
    impl Counter {
        // TODO: Implement these methods
        
        fn new() -> Self {
            Counter { count: 0 }
        }
        
        fn increment(&mut self) {
            self.count += 1;
        }
        
        fn get(&self) -> i32 {
            self.count
        }
    }
    
    let mut counter = Counter::new();
    counter.increment();
    counter.increment();
    counter.increment();
    println!("Count: {}", counter.get());  // Should print 3
}

// Exercise 3: Enums with Data
fn exercise_3() {
    println!("\n=== Exercise 3: Enums ===");
    
    // TODO: Define a Shape enum with:
    // - Circle(f64) for radius
    // - Rectangle(f64, f64) for width, height
    // - Square(f64) for side
    
    enum Shape {
        Circle(f64),
        Rectangle(f64, f64),
        Square(f64),
    }
    
    impl Shape {
        fn area(&self) -> f64 {
            match self {
                Shape::Circle(r) => 3.14159 * r * r,
                Shape::Rectangle(w, h) => w * h,
                Shape::Square(s) => s * s,
            }
        }
    }
    
    let shapes = vec![
        Shape::Circle(5.0),
        Shape::Rectangle(4.0, 6.0),
        Shape::Square(3.0),
    ];
    
    for shape in &shapes {
        println!("Area: {:.2}", shape.area());
    }
}

// Exercise 4: Option<T>
fn exercise_4() {
    println!("\n=== Exercise 4: Option ===");
    
    fn find_first_even(numbers: &[i32]) -> Option<i32> {
        // TODO: Return Some(n) for first even number, None if no even exists
        for &n in numbers {
            if n % 2 == 0 {
                return Some(n);
            }
        }
        None
    }
    
    let nums1 = [1, 3, 5, 8, 9];
    let nums2 = [1, 3, 5, 7, 9];
    
    match find_first_even(&nums1) {
        Some(n) => println!("First even: {}", n),
        None => println!("No even number"),
    }
    
    // Use unwrap_or for nums2
    let result = find_first_even(&nums2).unwrap_or(-1);
    println!("First even (or -1): {}", result);
}

// Exercise 5: Result<T, E> and Pattern Matching
fn exercise_5() {
    println!("\n=== Exercise 5: Result ===");
    
    #[derive(Debug)]
    enum MathError {
        DivisionByZero,
        NegativeSquareRoot,
    }
    
    fn safe_divide(a: f64, b: f64) -> Result<f64, MathError> {
        if b == 0.0 {
            Err(MathError::DivisionByZero)
        } else {
            Ok(a / b)
        }
    }
    
    fn safe_sqrt(n: f64) -> Result<f64, MathError> {
        if n < 0.0 {
            Err(MathError::NegativeSquareRoot)
        } else {
            Ok(n.sqrt())
        }
    }
    
    // TODO: Use pattern matching to handle results
    let operations = [
        safe_divide(10.0, 2.0),
        safe_divide(10.0, 0.0),
        safe_sqrt(16.0),
        safe_sqrt(-4.0),
    ];
    
    for (i, result) in operations.iter().enumerate() {
        match result {
            Ok(value) => println!("Op {}: {:.2}", i, value),
            Err(MathError::DivisionByZero) => println!("Op {}: division by zero!", i),
            Err(MathError::NegativeSquareRoot) => println!("Op {}: negative sqrt!", i),
        }
    }
}

// ============================================
// CHALLENGES
// ============================================

// Challenge 1: Create a simple state machine using enums
// TrafficLight: Red -> Green -> Yellow -> Red
// Implement a `next()` method that transitions states

// Challenge 2: Implement a Result-based function chain
// parse_int -> double -> to_string
// Use the ? operator to propagate errors
