# Phase 1: Foundation

## Setup

```bash
# Install Rust
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# Verify installation
rustc --version
cargo --version

# Create practice project
cargo new practice
cd practice
```

## Cargo Commands

| Command | Purpose |
|---------|---------|
| `cargo new <name>` | Create new project |
| `cargo build` | Compile |
| `cargo run` | Compile and run |
| `cargo test` | Run tests |
| `cargo check` | Fast compile check (no binary) |

## Key Concepts

### Variables & Mutability

```rust
let x = 5;          // immutable by default
let mut y = 10;     // mutable
y = 20;             // ok

let x = "hello";    // shadowing - new binding, can change type
```

### Basic Types

```rust
// Integers: i8, i16, i32, i64, i128, isize
// Unsigned: u8, u16, u32, u64, u128, usize
let a: i32 = 42;
let b: u8 = 255;

// Floats: f32, f64
let pi: f64 = 3.14159;

// Bool & Char
let flag: bool = true;
let c: char = '中';  // 4 bytes, Unicode

// Tuple
let tup: (i32, f64, char) = (500, 6.4, 'a');
let (x, y, z) = tup;  // destructuring
let first = tup.0;    // index access

// Array (fixed size, stack allocated)
let arr: [i32; 5] = [1, 2, 3, 4, 5];
let arr2 = [0; 10];   // [0, 0, 0, ...] 10 times
```

## Exercises

Work through `exercises.rs` in this folder.
