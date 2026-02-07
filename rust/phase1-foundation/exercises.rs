// Phase 1 Exercises: Foundation
// Run with: rustc exercises.rs && ./exercises
// Or create a cargo project and put this in src/main.rs

fn main() {
    exercise_1();
    exercise_2();
    exercise_3();
    exercise_4();
}

// Exercise 1: Variables and Mutability
// Fix this code to make it compile
fn exercise_1() {
    println!("=== Exercise 1: Variables ===");
    
    // TODO: This won't compile. Fix it.
    let mut x = 5;
    x = 10;
    println!("x = {}", x);
}

// Exercise 2: Shadowing
// Use shadowing to transform a string to its length
fn exercise_2() {
    println!("\n=== Exercise 2: Shadowing ===");
    
    let message = "hello rust";
    
    // TODO: Shadow 'message' to be its length (usize)
    // Hint: use message.len()
    let message: usize = message.len();
    
    println!("Length: {}", message);
}

// Exercise 3: Tuples
// Extract values from a tuple
fn exercise_3() {
    println!("\n=== Exercise 3: Tuples ===");
    
    let person: (&str, i32, bool) = ("Alice", 30, true);
    
    // TODO: Destructure the tuple into name, age, is_active
    let (name, age, is_active): (&str, i32, bool) = person;
    
    println!("{} is {} years old, active: {}", name, age, is_active);
}

// Exercise 4: Arrays
// Calculate the sum of an array
fn exercise_4() {
    println!("\n=== Exercise 4: Arrays ===");
    
    let numbers: [i32; 5] = [10, 20, 30, 40, 50];
    
    // TODO: Calculate the sum of all elements
    // Hint: use a for loop or .iter().sum()
    let sum: i32 = numbers.iter().sum();
    
    println!("Sum: {}", sum);  // Should print 150
}

// ============================================
// CHALLENGE: Try these on your own
// ============================================

// Challenge 1: Create a function that takes an array of 3 integers
// and returns a tuple of (min, max, sum)

// Challenge 2: Create a function that swaps two mutable variables
// Hint: Rust has a built-in swap, but try without it first
