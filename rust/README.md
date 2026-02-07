# Learning Rust: The Great Ideas

Rust's power comes from a few key innovations. This guide focuses on **what makes Rust special**, not everything in the language.

## The Big Ideas

1. **Ownership** - Memory safety without garbage collection
2. **Borrowing** - Safe references with compile-time checks
3. **Lifetimes** - Explicit scope tracking for references
4. **Pattern Matching** - Exhaustive, expressive control flow
5. **Traits** - Powerful abstraction without runtime cost
6. **Result/Option** - No null, no exceptions, explicit error handling
7. **Fearless Concurrency** - Data races caught at compile time

## Learning Path

### Phase 1: Foundation → `phase1-foundation/`

- [ ] Install Rust via `rustup`
- [ ] `cargo` basics: new, build, run, test
- [ ] Variables, mutability, and shadowing
- [ ] Basic types: integers, floats, bool, char, tuples, arrays

### Phase 2: Ownership → `phase2-ownership/`

- [ ] Ownership rules: one owner, transfer on assignment
- [ ] Move semantics vs Copy
- [ ] Borrowing: `&T` (shared) vs `&mut T` (exclusive)
- [ ] The borrow checker in action

### Phase 3: Structs and Enums → `phase3-structs-enums/`

- [ ] Defining structs, methods with `impl`
- [ ] Enums with data (algebraic data types)
- [ ] `Option<T>` and `Result<T, E>`
- [ ] Pattern matching with `match`

### Phase 4: Traits (TODO)

- [ ] Defining and implementing traits
- [ ] Trait bounds and generics
- [ ] Common traits: `Clone`, `Debug`, `Default`, `From/Into`

### Phase 5: Error Handling (TODO)

- [ ] `Result` and the `?` operator
- [ ] Custom error types
- [ ] `unwrap` vs proper handling

### Phase 6: Lifetimes (TODO)

- [ ] Why lifetimes exist
- [ ] Lifetime annotations `'a`
- [ ] Lifetime elision rules

### Phase 7: Concurrency (TODO)

- [ ] Threads and `move` closures
- [ ] `Arc` and `Mutex`
- [ ] Send and Sync traits
- [ ] Message passing with channels

## Quick Start

```bash
# Install Rust
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# Create first project
cargo new hello-rust
cd hello-rust
cargo run
```

## Resources

- [The Rust Book](https://doc.rust-lang.org/book/) - Official, comprehensive
- [Rustlings](https://github.com/rust-lang/rustlings) - Small exercises
- [Rust by Example](https://doc.rust-lang.org/rust-by-example/) - Learn through code
