import (
	"strconv"
	"sync"
)

var mu = &sync.Mutex{}
var myBalance = &balance{amount: 50.00, currency: "GBP"}

type balance struct {
	amount   float64
	currency string
}

func (b *balance) Add(i float64) {
	mu.Lock()
	b.amount += i
	mu.Unlock()
}

func (b *balance) Display() string {
	mu.Lock()
	amt := b.amount
	cur := b.currency
	mu.Unlock()
	return strconv.FormatFloat(amt, 'f', 2, 64) + " " + cur
}

// ---

import (
	"strconv"
	"sync"
)

var myBalance = &balance{amount: 50.00, currency: "GBP"}

type balance struct {
	amount   float64
	currency string
	mu       sync.Mutex
}

func (b *balance) Add(i float64) {
	b.mu.Lock()
	b.amount += i
	b.mu.Unlock()
}

func (b *balance) Display() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return strconv.FormatFloat(b.amount, 'f', 2, 64) + " " + b.currency
}

// ---

import (
	"strconv"
	"sync"
)

var myBalance = &balance{amount: 50.00, currency: "GBP"}

type balance struct {
	amount   float64
	currency string
	mu       sync.RWMutex // more efficient when you have more read than write
}

func (b *balance) Add(i float64) {
	b.mu.Lock()
	b.amount += i
	b.mu.Unlock()
}

func (b *balance) Display() string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return strconv.FormatFloat(b.amount, 'f', 2, 64) + " " + b.currency
}
