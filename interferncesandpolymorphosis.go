//What is an Interface?
//An interface is a collection of method signatures. If a type implements all methods of an interface, it automatically satisfies it — no explicit declaration needed!

//This allows polymorphism: different types can be used interchangeably if they implement the same interface.

//An interface in Go is a type that defines a set of method signatures. If a type (struct) implements those methods, it automatically satisfies the interface.

// No need to explicitly say that a struct implements an interface.

//Syntax
/*
type Device interface {
    Start()
    Stop()
}
	*/

	//Why Interfaces? (Real Benefits)

// Polymorphism: You can use different struct types in a common way.

// Flexibility: Swap implementations without changing the surrounding logic.

// Plug & play behavior: Ideal for APIs, plugins, testing, billing systems, etc.	


/*Subscription Billing System
We’ll simulate a billing system that supports:

MonthlyPlan (e.g., $20/month for 3 months)

YearlyPlan (e.g., $100/year for 2 years)

They both calculate their total cost differently, but are used via one common interface.*/
/*
package main

import (
    "fmt"
)

// Define the interface
type Plan interface {
    Cost() float64
}

// Monthly plan struct and method
type MonthlyPlan struct {
    PricePerMonth float64
    Months        int
}

func (m MonthlyPlan) Cost() float64 {
    return m.PricePerMonth * float64(m.Months)
}

// Yearly plan struct and method
/*
type YearlyPlan struct {
    PricePerYear float64
    Years        int
}

func (y YearlyPlan) Cost() float64 {
    return y.PricePerYear * float64(y.Years)
}

func main() {
    var plans []Plan
    var count int

    fmt.Print("Enter number of plans: ")
    fmt.Scanln(&count)

    for i := 0; i < count; i++ {
        var planType string
        fmt.Printf("Enter plan type for plan %d (monthly/yearly): ", i+1)
        fmt.Scanln(&planType)

        if planType == "monthly" {
            var price float64
            var months int
            fmt.Print("Enter monthly price: ")
            fmt.Scanln(&price)
            fmt.Print("Enter number of months: ")
            fmt.Scanln(&months)
            plans = append(plans, MonthlyPlan{PricePerMonth: price, Months: months})

        } else if planType == "yearly" {
            var price float64
            var years int
            fmt.Print("Enter yearly price: ")
            fmt.Scanln(&price)
            fmt.Print("Enter number of years: ")
            fmt.Scanln(&years)
            plans = append(plans, YearlyPlan{PricePerYear: price, Years: years})

        } else {
            fmt.Println("Invalid plan type. Skipping this one.")
        }
    }

    fmt.Println("\nBilling Summary:")
    total := 0.0
    for i, plan := range plans {
        cost := plan.Cost()
        fmt.Printf("Plan %d: $%.2f\n", i+1, cost)
        total += cost
    }
    fmt.Printf("Total cost for all plans: $%.2f\n", total)
}
*/