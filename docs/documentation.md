# Console-Based Library Management System Documentation

## Overview
This Go project implements an in-memory console-based library management system.

## Key Features
1. **Add/Remove Books**: Modify the central in-memory catalog (map key = `Book.ID`).
2. **Borrow/Return Books**: Handle borrowing logic, updating book status ("Available" / "Borrowed"), and keeping track of a member's borrowed list.
3. **Listings**: View available books or member-specific borrowed items.

## Architecture
- `models/`: Domain entities (`Book`, `Member`).
- `services/`: Business logic layer implementing `LibraryManager` interface.
- `controllers/`: Handles CLI parsing (`bufio.Reader`) and user feedback loops.
- `main.go`: Entry point initializing data and bootstrapping the CLI loop.

## How to Run
```bash
go run main.go