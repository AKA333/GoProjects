# Ticket Booking System

A small Go example that demonstrates how to use a buffered channel to model a ticket booking system and coordinate concurrent booking attempts from multiple customers.

## Overview

This program creates a fixed pool of available tickets and then simulates many customers trying to book a ticket concurrently.

Instead of protecting a shared counter with a mutex, it uses a buffered channel to represent the available tickets. Customers attempt to receive from the channel and either succeed or immediately receive a "Sold out" result if the channel is empty.

## How it works

- `totalCustomers` defines how many customers try to book tickets.
- `totalTickets` defines how many tickets are available.
- A buffered channel is created with capacity equal to the ticket count.
- The ticket channel is pre-filled with ticket IDs.
- Each customer runs as a goroutine and tries to receive a ticket from the channel.
- If a ticket is available, the customer books it successfully.
- If the channel is empty, the customer receives a "Sold out" message.
- A `sync.WaitGroup` ensures the program waits until all booking attempts are processed before exiting.

## Current configuration

- `totalCustomers = 50`
- `totalTickets = 5`

This means the program prints 5 successful bookings and 45 sold-out messages.

## Run the program

From the `ticket-booking-system` directory:

```bash
go run main.go
```

## Example output

```text
Customer 1, Booked ticket success 1
Customer 2, Booked ticket success 2
Customer 3, Booked ticket success 3
Customer 4, Booked ticket success 4
Customer 5, Booked ticket success 5
Customer 6: Sold out
...
All request processed
```

> Note: The customer order may vary because booking attempts are handled concurrently.

## Customization

To change the booking simulation, update the constants in `main.go`:

- `totalCustomers` to change the number of booking attempts
- `totalTickets` to change the number of available tickets

## What this teaches

- basic Go concurrency with goroutines
- using buffered channels as a shared resource pool
- avoiding mutexes by using channels
- waiting for concurrent work with `sync.WaitGroup`
