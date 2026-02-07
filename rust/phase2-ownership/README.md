# Phase 2: Ownership

This is Rust's core innovation. Understand this, and everything else follows.

## The Three Rules

1. Each value has exactly **one owner**
2. When the owner goes out of scope, the value is **dropped**
3. Ownership can be **transferred** (moved) or **borrowed**

## Move Semantics

```rust
let s1 = String::from("hello");
let s2 = s1;        // s1 is MOVED to s2
// println!("{}", s1);  // ERROR: s1 is no longer valid

let s3 = s2.clone();    // Deep copy - both valid
println!("{} {}", s2, s3);  // OK
```

**Why?** String stores data on the heap. Without move semantics, both `s1` and `s2` would point to the same memory, causing double-free on drop.

## Copy vs Move

```rust
// Copy types (stack-only, cheap to copy)
let x = 5;
let y = x;      // Copy, both valid
println!("{} {}", x, y);  // OK

// Move types (heap data or complex)
let v1 = vec![1, 2, 3];
let v2 = v1;    // Move, v1 invalid
// println!("{:?}", v1);  // ERROR
```

**Copy types:** integers, floats, bool, char, tuples of Copy types, arrays of Copy types

## Borrowing

Instead of transferring ownership, **borrow** with references.

```rust
fn main() {
    let s = String::from("hello");
    
    let len = calculate_length(&s);  // Borrow s
    println!("{} has length {}", s, len);  // s still valid
}

fn calculate_length(s: &String) -> usize {
    s.len()
}  // s goes out of scope, but doesn't drop (it's a reference)
```

## Mutable Borrowing

```rust
fn main() {
    let mut s = String::from("hello");
    change(&mut s);
    println!("{}", s);  // "hello, world"
}

fn change(s: &mut String) {
    s.push_str(", world");
}
```

## The Borrowing Rules

1. You can have **either**:
   - One mutable reference (`&mut T`), OR
   - Any number of immutable references (`&T`)
2. References must always be valid (no dangling)

```rust
let mut s = String::from("hello");

let r1 = &s;      // OK
let r2 = &s;      // OK - multiple immutable
// let r3 = &mut s;  // ERROR - can't mix mut and immut

let r4 = &mut s;  // OK - r1, r2 no longer used after this point
```

## Exercises

Work through `exercises.rs` - the borrow checker will teach you!
