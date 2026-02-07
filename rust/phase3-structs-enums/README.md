# Phase 3: Structs and Enums

Rust's type system shines here. Enums + pattern matching = powerful.

## Structs

```rust
// Define
struct User {
    username: String,
    email: String,
    active: bool,
}

// Instantiate
let user = User {
    username: String::from("alice"),
    email: String::from("alice@example.com"),
    active: true,
};

// Access
println!("{}", user.username);

// Shorthand when variable names match fields
let username = String::from("bob");
let email = String::from("bob@example.com");
let user2 = User {
    username,  // same as username: username
    email,
    active: false,
};

// Update syntax
let user3 = User {
    email: String::from("new@example.com"),
    ..user2  // rest from user2 (user2 partially moved!)
};
```

## Methods with impl

```rust
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    // Associated function (like static method)
    fn new(width: u32, height: u32) -> Self {
        Rectangle { width, height }
    }
    
    // Method (takes &self)
    fn area(&self) -> u32 {
        self.width * self.height
    }
    
    // Mutable method
    fn scale(&mut self, factor: u32) {
        self.width *= factor;
        self.height *= factor;
    }
}

let mut rect = Rectangle::new(10, 20);
println!("Area: {}", rect.area());
rect.scale(2);
```

## Enums

Rust enums can hold data - they're algebraic data types.

```rust
enum Message {
    Quit,                       // No data
    Move { x: i32, y: i32 },   // Named fields
    Write(String),              // Single value
    ChangeColor(u8, u8, u8),   // Multiple values
}

let msg = Message::Write(String::from("hello"));
```

## Option<T> - No More Null

```rust
enum Option<T> {
    Some(T),
    None,
}

fn find_user(id: u32) -> Option<String> {
    if id == 1 {
        Some(String::from("Alice"))
    } else {
        None
    }
}

// You MUST handle both cases
match find_user(1) {
    Some(name) => println!("Found: {}", name),
    None => println!("Not found"),
}

// Shortcuts
let name = find_user(1).unwrap_or(String::from("default"));
if let Some(name) = find_user(1) {
    println!("Found: {}", name);
}
```

## Result<T, E> - Explicit Errors

```rust
enum Result<T, E> {
    Ok(T),
    Err(E),
}

fn divide(a: f64, b: f64) -> Result<f64, String> {
    if b == 0.0 {
        Err(String::from("division by zero"))
    } else {
        Ok(a / b)
    }
}

// Handle with match
match divide(10.0, 2.0) {
    Ok(result) => println!("Result: {}", result),
    Err(e) => println!("Error: {}", e),
}

// The ? operator (in functions returning Result)
fn calc() -> Result<f64, String> {
    let x = divide(10.0, 2.0)?;  // Returns early if Err
    let y = divide(x, 2.0)?;
    Ok(y)
}
```

## Pattern Matching

```rust
let x = 5;

match x {
    1 => println!("one"),
    2 | 3 => println!("two or three"),  // multiple
    4..=10 => println!("four to ten"),  // range
    _ => println!("something else"),     // default
}

// Destructuring
let point = (3, 5);
match point {
    (0, 0) => println!("origin"),
    (x, 0) => println!("on x-axis at {}", x),
    (0, y) => println!("on y-axis at {}", y),
    (x, y) => println!("at ({}, {})", x, y),
}

// Guards
match x {
    n if n < 0 => println!("negative"),
    n if n > 0 => println!("positive"),
    _ => println!("zero"),
}
```

## Exercises

Work through `exercises.rs`.
