// Phase 2 Exercises: Ownership
// These exercises will NOT compile initially.
// Your job: fix them while understanding WHY they fail.
//
// Run with: rustc exercises.rs && ./exercises

fn main() {
    exercise_1();
    exercise_2();
    exercise_3();
    exercise_4();
    exercise_5();
}

// Exercise 1: Move Semantics
// This code won't compile. Fix it WITHOUT using clone().
fn exercise_1() {
    println!("=== Exercise 1: Move ===");
    
    let s1 = String::from("hello");
    let s2 = s1;
    
    // TODO: Fix this - s1 was moved
    // Original broken code:
    // println!("s1 = {}, s2 = {}", s1, s2);
    
    // Solution: only use s2
    println!("s2 = {}", s2);
}

// Exercise 2: Function Ownership
// The function takes ownership. How do you use the string after?
fn exercise_2() {
    println!("\n=== Exercise 2: Function Ownership ===");
    
    let s = String::from("hello");
    
    // Option A: Take and return ownership
    let s = takes_and_returns(s);
    println!("After function: {}", s);
    
    // Option B: Use a reference (better!)
    let s2 = String::from("world");
    borrows(&s2);
    println!("Still valid: {}", s2);
}

fn takes_and_returns(s: String) -> String {
    println!("Got: {}", s);
    s  // return ownership
}

fn borrows(s: &String) {
    println!("Borrowed: {}", s);
}

// Exercise 3: Mutable Borrow
// Fix this to actually modify the string
fn exercise_3() {
    println!("\n=== Exercise 3: Mutable Borrow ===");
    
    let mut s = String::from("hello");
    
    add_world(&mut s);
    
    println!("{}", s);  // Should print "hello world"
}

fn add_world(s: &mut String) {
    s.push_str(" world");
}

// Exercise 4: Borrow Rules
// This violates borrowing rules. Fix it.
fn exercise_4() {
    println!("\n=== Exercise 4: Borrow Rules ===");
    
    let mut data = vec![1, 2, 3];
    
    // Problem: Can't have immutable borrow while mutating
    // Original broken code:
    // let first = &data[0];
    // data.push(4);
    // println!("First: {}", first);
    
    // Solution: Use immutable borrow before mutation
    let first = data[0];  // Copy the value instead of borrow
    data.push(4);
    println!("First was: {}, data is now: {:?}", first, data);
}

// Exercise 5: Returning References
// This won't compile. Why? Fix it.
fn exercise_5() {
    println!("\n=== Exercise 5: Returning References ===");
    
    // This is the safe version
    let result = create_string();
    println!("Created: {}", result);
}

// BAD: Would return reference to dropped value
// fn create_string() -> &String {
//     let s = String::from("hello");
//     &s  // ERROR: s is dropped, reference would dangle
// }

// GOOD: Return owned value
fn create_string() -> String {
    let s = String::from("hello");
    s  // Transfer ownership out
}

// ============================================
// CHALLENGES
// ============================================

// Challenge 1: Write a function that takes a &mut Vec<i32>
// and doubles every element in place

// Challenge 2: Write a function that takes two &str and returns
// the longer one. Hint: you'll need lifetime annotations (Phase 6)
// For now, try returning a String instead.
