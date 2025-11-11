# State Machines 🤖

## What is a State Machine?

A **state machine** is a simple way to control what your program can do at different times.

Think of it like a **remote control**:
- When TV is OFF → you can only press POWER
- When TV is ON → you can press VOLUME, CHANNEL, POWER
- When TV is MUTED → you can press UNMUTE, POWER

## Real-World Examples

### 🚗 Car
- **PARKED** → can START
- **DRIVING** → can STOP, TURN  
- **STOPPED** → can START, PARK

### 📱 Phone Call
- **IDLE** → can DIAL
- **RINGING** → can ANSWER, DECLINE
- **TALKING** → can HANG UP

### 🎮 Game
- **MENU** → can START GAME, QUIT
- **PLAYING** → can PAUSE, QUIT
- **PAUSED** → can RESUME, QUIT

## Why Use State Machines?

1. **Prevents bugs** - Can't do impossible things (like pause when not playing)
2. **Clear logic** - Easy to understand what's allowed when
3. **Predictable** - Always know what state you're in

## Basic Concept

```
Current State + Action = New State
```

Examples:
- IDLE + START = RUNNING
- RUNNING + PAUSE = PAUSED  
- PAUSED + RESUME = RUNNING

## In Go with Channels

Go uses **nil channels** as a clever trick:
- `nil` channel = action is **disabled**
- Real channel = action is **enabled**

```go
var startCh chan bool    // nil = can't start
var pauseCh chan bool    // nil = can't pause

// In RUNNING state:
startCh = nil                    // Disable start
pauseCh = make(chan bool, 1)     // Enable pause
```

That's it! State machines are just organized ways to control what your program can do. 🎯
