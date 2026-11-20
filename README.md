# Go Bill Manager

A simple command-line application built with **Go** to manage bills, add items, update tips, and save bills as text files.

## Features

- Create a new bill.
- Add items with their prices.
- Add or update a tip.
- Display a formatted bill breakdown.
- Save bills to text files.

## Go Concepts Practiced

- Structs and methods.
- Pointer and value receivers.
- Maps.
- Functions.
- Error handling.
- Reading user input.
- File handling using the `os` package.

## Technologies

- Go (Golang)
- Go Standard Library

## Example

```text
Create a new bill name: Dinner
Created the bill Dinner

Choose option (a - add item, s - save bill, t - add tip): a
Choose item name: Pizza
Choose item price: 25
Item is added - Pizza 25

Choose option: t
Enter tip amount ($): 5
Tip is updated - 5

Choose option: s
Saved bill: Dinner
```