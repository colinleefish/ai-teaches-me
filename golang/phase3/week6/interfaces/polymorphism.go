package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// ===== BASIC POLYMORPHISM =====

// TODO: Define Animal interface with Speak() string and Move() string methods
// YOUR CODE HERE
type Animal interface {
	Speak() string
	Move() string
}

// TODO: Define DogPoly struct with Name and Breed fields (both string)
// YOUR CODE HERE
type DogPoly struct {
	Name  string
	Breed string
}

// TODO: Implement Speak() method for DogPoly - return "[name] says Woof!"
// YOUR CODE HERE
func (d DogPoly) Speak() string {
	return fmt.Sprintf("%s says Woof!", d.Name)
}

// TODO: Implement Move() method for DogPoly - return "[name] runs around"
// YOUR CODE HERE
func (d DogPoly) Move() string {
	return fmt.Sprintf("%s runs around", d.Name)
}

// TODO: Define BirdPoly struct with Name and Species fields (both string)
// YOUR CODE HERE
type BirdPoly struct {
	Name    string
	Species string
}

// TODO: Implement Speak() method for BirdPoly - return "[name] chirps melodiously"
// YOUR CODE HERE
func (b BirdPoly) Speak() string {
	return fmt.Sprintf("%s chirps melodiously", b.Name)
}

// TODO: Implement Move() method for BirdPoly - return "[name] flies gracefully"
// YOUR CODE HERE
func (b BirdPoly) Move() string {
	return fmt.Sprintf("%s flies gracefully", b.Name)
}

// TODO: Define FishPoly struct with Name and Type fields (both string)
// YOUR CODE HERE
type FishPoly struct {
	Name string
	Type string
}

// TODO: Implement Speak() method for FishPoly - return "[name] makes bubbles"
// YOUR CODE HERE
func (f FishPoly) Speak() string {
	return fmt.Sprintf("%s makes bubbles", f.Name)
}

// TODO: Implement Move() method for FishPoly - return "[name] swims smoothly"
// YOUR CODE HERE
func (f FishPoly) Move() string {
	return fmt.Sprintf("%s swims smoothly", f.Name)
}

func demonstrateBasicPolymorphism() {
	fmt.Println("=== Basic Polymorphism ===")

	// TODO: Create a slice of Animal containing different types:
	// DogPoly, BirdPoly, FishPoly instances with sample data
	// YOUR CODE HERE
	animals := []Animal{
		DogPoly{Name: "Rex", Breed: "Golden Retriever"},
		BirdPoly{Name: "Tweety", Species: "Canary"},
		FishPoly{Name: "Nemo", Type: "Clownfish"},
	}

	// TODO: Loop through animals and call Speak() and Move() on each
	// Print the animal number, speak result, move result, and type (%T)
	// YOUR CODE HERE
	for i, animal := range animals {
		fmt.Printf("Animal %d: %T\n", i+1, animal)
		fmt.Println(animal.Speak())
		fmt.Println(animal.Move())
	}

	// TODO: Call makeAnimalShow function with the animals slice
	// YOUR CODE HERE
	makeAnimalShow(animals)
}

// TODO: Implement makeAnimalShow function that takes []Animal
// Print "🎪 Welcome to the Animal Show!"
// Loop through animals and print their Speak() and Move() results
// YOUR CODE HERE

func makeAnimalShow(animals []Animal) {
	fmt.Println("🎪 Welcome to the Animal Show! (Though I think it's a bit weird)")
	for _, animal := range animals {
		fmt.Printf("🐾 Introducing %s: \n", animal)
		fmt.Println(animal.Speak())
		fmt.Println(animal.Move())
		fmt.Println()
	}
}

// ===== STRATEGY PATTERN =====

// TODO: Define SortStrategy interface with:
// - Sort([]int) []int method
// - Name() string method
// YOUR CODE HERE
type SortStrategy interface {
	Sort([]int) []int
	Name() string
}

// TODO: Define BubbleSort struct (empty struct)
// YOUR CODE HERE
type BubbleSort struct{}

// TODO: Implement Sort method for BubbleSort
// - Make a copy of input data
// - Implement bubble sort algorithm (nested loops)
// - Compare adjacent elements and swap if needed
// - Return sorted copy
// YOUR CODE HERE
func (b BubbleSort) Sort(arr []int) []int {
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)

	for i := 0; i < len(copyArr); i++ {
		for j := 0; j < len(copyArr)-i-1; j++ {
			if copyArr[j] > copyArr[j+1] {
				copyArr[j], copyArr[j+1] = copyArr[j+1], copyArr[j]
			}
		}
	}
	return copyArr
}

// TODO: Implement Name method for BubbleSort
// Return "Bubble Sort"
// YOUR CODE HERE
func (b BubbleSort) Name() string {
	return "Bubble Sort"
}

// TODO: Define QuickSort struct (empty struct)
// YOUR CODE HERE
type QuickSort struct{}

// TODO: Implement Sort method for QuickSort
// - Make a copy of input data
// - Call quickSortHelper with the copy
// - Return sorted copy
// YOUR CODE HERE
func (q QuickSort) Sort(arr []int) []int {
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)
	quickSortHelper(copyArr, 0, len(copyArr)-1)
	return copyArr
}

// TODO: Implement Name method for QuickSort
// Return "Quick Sort"
// YOUR CODE HERE
func (q QuickSort) Name() string {
	return "Quick Sort"
}

// TODO: Implement quickSortHelper function
// Parameters: arr []int, low, high int
// Recursive quicksort implementation
// YOUR CODE HERE
func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSortHelper(arr, low, pivot-1)
		quickSortHelper(arr, pivot+1, high)
	}
}

// TODO: Implement partition function
// Parameters: arr []int, low, high int
// Returns pivot index after partitioning
// YOUR CODE HERE
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// TODO: Define DataSorter struct with strategy field of type SortStrategy
// YOUR CODE HERE
type DataSorter struct {
	strategy SortStrategy
}

// TODO: Implement NewDataSorter constructor
// Takes SortStrategy parameter and returns *DataSorter
// YOUR CODE HERE
func NewDataSorter(strategy SortStrategy) *DataSorter {
	return &DataSorter{strategy: strategy}
}

// TODO: Implement SetStrategy method for DataSorter
// Allows changing the sorting strategy at runtime
// YOUR CODE HERE
func (d *DataSorter) SetStrategy(strategy SortStrategy) {
	d.strategy = strategy
}

// TODO: Implement SortData method for DataSorter
// - Print "Using [strategy name]..."
// - Record start time with time.Now()
// - Call strategy.Sort(data)
// - Calculate duration with time.Since(start)
// - Print completion message with duration
// - Return sorted result
// YOUR CODE HERE
func (d *DataSorter) SortData(data []int) []int {
	fmt.Printf("Using %s...\n", d.strategy.Name())
	start := time.Now()
	sorted := d.strategy.Sort(data)
	duration := time.Since(start)
	fmt.Printf("Sorting completed in %s\n", duration)
	return sorted
}

func demonstrateStrategyPattern() {
	fmt.Println("\n=== Strategy Pattern ===")

	// TODO: Create sample data slice to sort
	// YOUR CODE HERE
	data := []int{9, 3, 7, 1, 5, 8, 2, 6, 4}
	// TODO: Create DataSorter with BubbleSort strategy
	// YOUR CODE HERE
	sorter := NewDataSorter(BubbleSort{})

	// TODO: Sort data using bubble sort and print result
	// YOUR CODE HERE
	bubbleSorted := sorter.SortData(data)
	fmt.Printf("Bubble sorted: %v\n", bubbleSorted)

	// TODO: Change strategy to QuickSort using SetStrategy
	// YOUR CODE HERE
	sorter.SetStrategy(QuickSort{})

	// TODO: Sort the same data with new strategy and print result
	// YOUR CODE HERE
	quickSorted := sorter.SortData(data)
	fmt.Printf("Quick sorted: %v\n", quickSorted)
}

// ===== OBSERVER PATTERN =====

// TODO: Define Observer interface with:
// - Update(message string) method
// - GetID() string method
// YOUR CODE HERE
type Observer interface {
	Update(message string)
	GetID() string
}

// TODO: Define Subject interface with:
// - Attach(observer Observer) method
// - Detach(observer Observer) method
// - Notify(message string) method
// YOUR CODE HERE
type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify(message string)
}

// TODO: Define NewsAgency struct with:
// - observers []Observer field
// - news string field
// YOUR CODE HERE
type NewsAgency struct {
	observers []Observer
	news      string
}

// TODO: Implement NewNewsAgency constructor
// Initialize with empty observers slice
// YOUR CODE HERE
func NewNewsAgency() *NewsAgency {
	return &NewsAgency{observers: []Observer{}, news: ""}
}

// TODO: Implement Attach method for NewsAgency
// Add observer to observers slice and print attachment message
// YOUR CODE HERE
func (n *NewsAgency) Attach(ob Observer) {
	n.observers = append(n.observers, ob)
	fmt.Printf("📰 %s subscribed to the news agency\n", ob.GetID())
}

// TODO: Implement Detach method for NewsAgency
// Remove observer from observers slice and print detachment message
// YOUR CODE HERE
func (n *NewsAgency) Detach(ob Observer) {
	for i, observer := range n.observers {
		if observer.GetID() == ob.GetID() {
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
			fmt.Printf("📰 %s unsubscribed from the news agency\n", ob.GetID())
			break
		}
	}
}

// TODO: Implement Notify method for NewsAgency
// Set news field and call Update on all observers
// YOUR CODE HERE
func (n *NewsAgency) Notify(message string) {
	n.news = message
	for _, ob := range n.observers {
		ob.Update(message)
	}
}

// TODO: Define NewsChannel struct with name string field
// YOUR CODE HERE
type NewsChannel struct {
	Name string
}

// TODO: Implement NewNewsChannel constructor
// YOUR CODE HERE
func NewNewsChannel(name string) *NewsChannel {
	return &NewsChannel{Name: name}
}

// TODO: Implement Update method for NewsChannel
// Print "📺 [name] received: [message]"
// YOUR CODE HERE
func (n *NewsChannel) Update(message string) {
	fmt.Printf("📺 %s received: %s\n", n.Name, message)
}

// TODO: Implement GetID method for NewsChannel
// Return the name field
// YOUR CODE HERE
func (n *NewsChannel) GetID() string {
	return n.Name
}

// TODO: Define NewsPaper struct with name string field
// YOUR CODE HERE
type NewsPaper struct {
	Name string
}

// TODO: Implement NewNewsPaper constructor
// YOUR CODE HERE
func NewNewsPaper(name string) *NewsPaper {
	return &NewsPaper{Name: name}
}

// TODO: Implement Update method for NewsPaper
// Print "📰 [name] printing: [message]"
// YOUR CODE HERE
func (n *NewsPaper) Update(message string) {
	fmt.Printf("📰 %s printing: %s\n", n.Name, message)
}

// TODO: Implement GetID method for NewsPaper
// Return the name field
// YOUR CODE HERE
func (n *NewsPaper) GetID() string {
	return n.Name
}

func demonstrateObserverPattern() {
	fmt.Println("\n=== Observer Pattern ===")

	// TODO: Create a NewsAgency instance
	// YOUR CODE HERE
	ap := NewNewsAgency()
	// TODO: Create different observers (NewsChannel and NewsPaper instances)
	// YOUR CODE HERE
	cnn := NewNewsChannel("CNN")
	nyt := NewNewsPaper("New York Times")

	// TODO: Attach all observers to the agency
	// YOUR CODE HERE
	ap.Attach(cnn)
	ap.Attach(nyt)

	// TODO: Broadcast some news using Notify
	// YOUR CODE HERE
	ap.Notify("Breaking news: the sky is falling!")

	// TODO: Detach one observer and broadcast more news
	// YOUR CODE HERE
	ap.Detach(cnn)
	ap.Notify("Breaking news: Trump is back in office!")

}

// ===== FACTORY PATTERN =====

// TODO: Define Transport interface with:
// - Deliver() string method
// - GetType() string method
// YOUR CODE HERE
type Transport interface {
	Deliver() string
	GetType() string
}

// TODO: Define Truck struct with capacity int field
// YOUR CODE HERE
type Truck struct {
	Capacity int
}

// TODO: Implement Deliver() method for Truck
// Return "Delivering by truck (capacity: [capacity] tons)"
// YOUR CODE HERE
func (t Truck) Deliver() string {
	return fmt.Sprintf("Delivering by truck (capacity: %d)", t.Capacity)
}

// TODO: Implement GetType() method for Truck
// Return "Land Transport"
// YOUR CODE HERE
func (t Truck) GetType() string {
	return "Land Transport"
}

// TODO: Define Ship struct with containers int field
// YOUR CODE HERE
type Ship struct {
	Containers int
}

// TODO: Implement Deliver() method for Ship
// Return "Delivering by ship ([containers] containers)"
// YOUR CODE HERE
func (s Ship) Deliver() string {
	return fmt.Sprintf("Delivering by ship (%d containers)", s.Containers)
}

// TODO: Implement GetType() method for Ship
// Return "Sea Transport"
// YOUR CODE HERE
func (s Ship) GetType() string {
	return "Sea Transport"
}

// TODO: Define Plane struct with passengers int field
// YOUR CODE HERE
type Plane struct {
	Passengers int
}

// TODO: Implement Deliver() method for Plane
// Return "Delivering by plane ([passengers] passengers)"
// YOUR CODE HERE
func (p Plane) Deliver() string {
	return fmt.Sprintf("Delivering by plane (%d passengers)", p.Passengers)
}

// TODO: Implement GetType() method for Plane
// Return "Air Transport"
// YOUR CODE HERE
func (p Plane) GetType() string {
	return "Air Transport"
}

// TODO: Define TransportFactory interface with CreateTransport() Transport method
// YOUR CODE HERE
type TransportFactory interface {
	CreateTransport() Transport
}

// TODO: Define LandTransportFactory struct (empty struct)
// YOUR CODE HERE
type LandTransportFactory struct{}

// TODO: Implement CreateTransport() method for LandTransportFactory
// Return Truck with capacity 10
// YOUR CODE HERE
func (l LandTransportFactory) CreateTransport() Transport {
	return Truck{Capacity: 10}
}

// TODO: Define SeaTransportFactory struct (empty struct)
// YOUR CODE HERE
type SeaTransportFactory struct{}

// TODO: Implement CreateTransport() method for SeaTransportFactory
// Return Ship with containers 50
// YOUR CODE HERE
func (s SeaTransportFactory) CreateTransport() Transport {
	return Ship{Containers: 50}
}

// TODO: Define AirTransportFactory struct (empty struct)
// YOUR CODE HERE
type AirTransportFactory struct{}

// TODO: Implement CreateTransport() method for AirTransportFactory
// Return Plane with passengers 200
// YOUR CODE HERE
func (a AirTransportFactory) CreateTransport() Transport {
	return Plane{Passengers: 200}
}

// TODO: Define LogisticsManager struct with factory TransportFactory field
// YOUR CODE HERE
type LogisticsManager struct {
	factory TransportFactory
}

// TODO: Implement NewLogisticsManager constructor
// Takes TransportFactory parameter and returns *LogisticsManager
// YOUR CODE HERE

func NewLogisticsManager(factory TransportFactory) *LogisticsManager {
	return &LogisticsManager{factory: factory}
}

// TODO: Implement PlanDelivery() method for LogisticsManager
// - Call factory.CreateTransport()
// - Print delivery info with "🚚" emoji
// - Print transport type
// YOUR CODE HERE
func (l *LogisticsManager) PlanDelivery() {
	transport := l.factory.CreateTransport()
	fmt.Printf("🚚 %s\n", transport.Deliver())
	fmt.Printf("Type: %s\n", transport.GetType())
}

func demonstrateFactoryPattern() {
	fmt.Println("\n=== Factory Pattern ===")

	// TODO: Create different factories (Land, Sea, Air)
	// YOUR CODE HERE
	landFactory := LandTransportFactory{}
	seaFactory := SeaTransportFactory{}
	airFactory := AirTransportFactory{}

	factories := []TransportFactory{&landFactory, &seaFactory, &airFactory}

	// TODO: Loop through factories and demonstrate each delivery option
	// YOUR CODE HERE
	for _, factory := range factories {
		manager := NewLogisticsManager(factory)
		manager.PlanDelivery()
	}

}

// ===== ADAPTER PATTERN =====

// TODO: Define LegacyPrinter interface with PrintOldFormat(data string) method
// YOUR CODE HERE
type LegacyPrinter interface {
	PrintOldFormat(data string)
}

// TODO: Define OldPrinter struct with name string field
// YOUR CODE HERE
type OldPrinter struct {
	Name string
}

// TODO: Implement PrintOldFormat method for OldPrinter
// Print "🖨️ [Legacy [name]] [UPPERCASE data]"
// Use strings.ToUpper() to convert data to uppercase
// YOUR CODE HERE
func (o *OldPrinter) PrintOldFormat(data string) {
	fmt.Printf("🖨️ %s %s\n", o.Name, strings.ToUpper(data))
}

// TODO: Define ModernPrinter interface with Print(document DocumentAdapter) error method
// YOUR CODE HERE
type ModernPrinter interface {
	Print(document DocumentAdapter) error
}

// TODO: Define DocumentAdapter struct with Title, Content, Author fields (all string)
// YOUR CODE HERE
type DocumentAdapter struct {
	Title   string
	Content string
	Author  string
}

// TODO: Define PrinterAdapter struct with legacyPrinter LegacyPrinter field
// YOUR CODE HERE
type PrinterAdapter struct {
	legacyPrinter LegacyPrinter
}

// TODO: Implement NewPrinterAdapter constructor
// Takes LegacyPrinter parameter and returns *PrinterAdapter
// YOUR CODE HERE
func NewPrinterAdapter(legacyPrinter LegacyPrinter) *PrinterAdapter {
	return &PrinterAdapter{legacyPrinter: legacyPrinter}
}

// TODO: Implement Print method for PrinterAdapter
// - Convert modern document to legacy format: "TITLE: [title] | AUTHOR: [author] | CONTENT: [content]"
// - Call legacyPrinter.PrintOldFormat with the converted format
// - Return nil error
// YOUR CODE HERE
func (p *PrinterAdapter) Print(document DocumentAdapter) error {
	legacyData := fmt.Sprintf("TITLE: %s | AUTHOR: %s | CONTENT: %s", document.Title, document.Author, document.Content)
	p.legacyPrinter.PrintOldFormat(legacyData)
	return nil
}

// TODO: Define AdvancedPrinter struct with model string field
// YOUR CODE HERE
type AdvancedPrinter struct {
	Model string
}

// TODO: Implement Print method for AdvancedPrinter
// - Print "🖨️ [Advanced [model]]"
// - Print Title, Author, Content on separate lines with indentation
// - Return nil error
// YOUR CODE HERE
func (a *AdvancedPrinter) Print(document DocumentAdapter) error {
	fmt.Printf("🖨️ Advanced %s\n", a.Model)
	fmt.Printf("Title: %s\n", document.Title)
	fmt.Printf("Author: %s\n", document.Author)
	fmt.Printf("Content: %s\n", document.Content)
	return nil
}

func demonstrateAdapterPattern() {
	fmt.Println("\n=== Adapter Pattern ===")

	// TODO: Create a DocumentAdapter instance with sample data
	// YOUR CODE HERE
	sampleDoc := DocumentAdapter{
		Title:   "Sample Title",
		Content: "Sample Content",
		Author:  "Sample Author",
	}

	// TODO: Create a slice of ModernPrinter containing:
	// - AdvancedPrinter instances
	// - PrinterAdapter instances wrapping OldPrinter instances
	// YOUR CODE HERE
	advancedPrinter := AdvancedPrinter{Model: "HP-1234"}
	oldPrinter := OldPrinter{Name: "Epson-1234"}

	printerAdapter := NewPrinterAdapter(&oldPrinter)

	printers := []ModernPrinter{&advancedPrinter, printerAdapter}

	// TODO: Loop through printers and call Print on each
	// YOUR CODE HERE
	for _, printer := range printers {
		err := printer.Print(sampleDoc)
		if err != nil {
			fmt.Printf("Error printing with %T: %v\n", printer, err)
		}
	}
}

// ===== DECORATOR PATTERN =====

// TODO: Define Coffee interface with:
// - Cost() float64 method
// - Description() string method
// YOUR CODE HERE
type Coffee interface {
	Cost() float64
	Description() string
}

// TODO: Define SimpleCoffee struct (empty struct)
// YOUR CODE HERE
type SimpleCoffee struct{}

// TODO: Implement Cost() method for SimpleCoffee
// Return 2.0
// YOUR CODE HERE
func (sc SimpleCoffee) Cost() float64 {
	return 2.0
}

// TODO: Implement Description() method for SimpleCoffee
// Return "Simple coffee"
// YOUR CODE HERE
func (sc SimpleCoffee) Description() string {
	return "Simple Coffee"
}

// TODO: Define CoffeeDecorator struct with coffee Coffee field
// YOUR CODE HERE
type CoffeeDecorator struct {
	coffee Coffee
}

// TODO: Define MilkDecorator struct that embeds CoffeeDecorator
// YOUR CODE HERE
type MilkDecorator struct {
	CoffeeDecorator
}

// TODO: Implement NewMilkDecorator constructor
// Takes Coffee parameter and returns *MilkDecorator
// YOUR CODE HERE
func NewMilkDecorator(c Coffee) *MilkDecorator {
	return &MilkDecorator{CoffeeDecorator: CoffeeDecorator{coffee: c}}
}

// TODO: Implement Cost() method for MilkDecorator
// Return coffee.Cost() + 0.5
// YOUR CODE HERE
func (mk *MilkDecorator) Cost() float64 {
	return mk.coffee.Cost() + 0.5
}

// TODO: Implement Description() method for MilkDecorator
// Return coffee.Description() + ", milk"
// YOUR CODE HERE
func (mk *MilkDecorator) Description() string {
	return mk.coffee.Description() + ", milk"
}

// TODO: Define SugarDecorator struct that embeds CoffeeDecorator
// YOUR CODE HERE
type SugarDecorator struct {
	CoffeeDecorator
}

// TODO: Implement NewSugarDecorator constructor
// YOUR CODE HERE
func NewSugarDecorator(c Coffee) *SugarDecorator {
	return &SugarDecorator{CoffeeDecorator: CoffeeDecorator{coffee: c}}
}

// TODO: Implement Cost() method for SugarDecorator
// Return coffee.Cost() + 0.2
// YOUR CODE HERE
func (s *SugarDecorator) Cost() float64 {
	return s.coffee.Cost() + 0.2
}

// TODO: Implement Description() method for SugarDecorator
// Return coffee.Description() + ", sugar"
// YOUR CODE HERE
func (s *SugarDecorator) Description() string {
	return s.coffee.Description() + ", sugar"
}

// TODO: Define WhipDecorator struct that embeds CoffeeDecorator
// YOUR CODE HERE
type WhipDecorator struct {
	CoffeeDecorator
}

// TODO: Implement NewWhipDecorator constructor
// YOUR CODE HERE
func NewWhipDecorator(c Coffee) *WhipDecorator {
	return &WhipDecorator{CoffeeDecorator: CoffeeDecorator{coffee: c}}
}

// TODO: Implement Cost() method for WhipDecorator
// Return coffee.Cost() + 0.7
// YOUR CODE HERE
func (w *WhipDecorator) Cost() float64 {
	return w.coffee.Cost() + 0.7
}

// TODO: Implement Description() method for WhipDecorator
// Return coffee.Description() + ", whipped cream"
// YOUR CODE HERE
func (w *WhipDecorator) Description() string {
	return w.coffee.Description() + ", whipped cream"
}

func demonstrateDecoratorPattern() {
	fmt.Println("\n=== Decorator Pattern ===")

	// TODO: Start with simple coffee and print its cost and description
	// YOUR CODE HERE
	sc := SimpleCoffee{}
	fmt.Printf("Simple Coffee: %s ($%.2f)\n", sc.Description(), sc.Cost())

	// TODO: Add milk decorator and print the result
	// YOUR CODE HERE
	milk := NewMilkDecorator(&sc)
	fmt.Printf("Milk Coffee: %s ($%.2f)\n", milk.Description(), milk.Cost())

	// TODO: Add sugar decorator to the milk coffee and print
	// YOUR CODE HERE
	sugar := NewSugarDecorator(milk)
	fmt.Printf("Sugar Coffee: %s ($%.2f)\n", sugar.Description(), sugar.Cost())

	// TODO: Add whipped cream decorator and print the final result
	// YOUR CODE HERE
	whip := NewWhipDecorator(sugar)
	fmt.Printf("Whipped Cream Coffee: %s ($%.2f)\n", whip.Description(), whip.Cost())

	// TODO: Create different coffee combinations and print them
	// YOUR CODE HERE

	fmt.Println("Different coffee combinations:")

	latte := NewMilkDecorator(&SimpleCoffee{})
	fmt.Printf("Latte: %s ($%.2f)\n", latte.Description(), latte.Cost())

	mocha := NewWhipDecorator(NewSugarDecorator(NewMilkDecorator(&SimpleCoffee{})))
	fmt.Printf("Mocha: %s ($%.2f)\n", mocha.Description(), mocha.Cost())

	americano := NewSugarDecorator(NewMilkDecorator(&SimpleCoffee{}))
	fmt.Printf("Americano: %s ($%.2f)\n", americano.Description(), americano.Cost())

}

// ===== POLYMORPHIC BEHAVIOR WITH RANDOM SELECTION =====

// TODO: Define GameCharacter interface with:
// - Attack() string method
// - Defend() string method
// - GetHealth() int method
// - TakeDamage(damage int) method
// YOUR CODE HERE
type GameCharacter interface {
	Attack() string
	Defend() string
	GetHealth() int
	TakeDamage(damage int)
}

// TODO: Define Warrior struct with name (string), health (int), armor (int) fields
// YOUR CODE HERE
type Warrior struct {
	Name   string
	Health int
	Armor  int
}

// TODO: Implement NewWarrior constructor
// Initialize with health 100, armor 20
// YOUR CODE HERE
func NewWarrior(name string) *Warrior {
	return &Warrior{
		Name:   name,
		Health: 100,
		Armor:  20,
	}
}

// TODO: Implement Attack() method for Warrior
// - Generate random damage: 25 + rand.Intn(10)
// - Return "[name] swings sword for [damage] damage!"
// YOUR CODE HERE
func (w *Warrior) Attack() string {
	damage := 25 + rand.Intn(10)
	return fmt.Sprintf("%s swings sword for %d damage!", w.Name, damage)
}

// TODO: Implement Defend() method for Warrior
// Return "[name] raises shield (armor: [armor])"
// YOUR CODE HERE
func (w *Warrior) Defend() string {
	return fmt.Sprintf("%s raises shield (armor: %d)", w.Name, w.Armor)
}

// TODO: Implement GetHealth() method for Warrior
// Return health field
// YOUR CODE HERE
func (w *Warrior) GetHealth() int {
	return w.Health
}

// TODO: Implement TakeDamage(damage int) method for Warrior
// - Calculate actual damage: damage - armor/2 (minimum 0)
// - Reduce health by actual damage
// - Ensure health doesn't go below 0
// YOUR CODE HERE
func (w *Warrior) TakeDamage(damage int) {
	actualDamage := damage - w.Armor/2
	if actualDamage < 0 {
		actualDamage = 0
	}
	w.Health -= actualDamage
	if w.Health < 0 {
		w.Health = 0
	}
}

// TODO: Define Mage struct with name (string), health (int), mana (int) fields
// YOUR CODE HERE
type Mage struct {
	Name   string
	Health int
	Mana   int
}

// TODO: Implement NewMage constructor
// Initialize with health 70, mana 100
// YOUR CODE HERE

func NewMage(name string) *Mage {
	return &Mage{
		Name:   name,
		Health: 70,
		Mana:   100,
	}
}

// TODO: Implement Attack() method for Mage
// - If mana >= 20: reduce mana by 20, generate damage 30 + rand.Intn(15)
// - Return "[name] casts fireball for [damage] damage! (mana: [mana])"
// - If mana < 20: return "[name] has no mana for spells!"
// YOUR CODE HERE

func (m *Mage) Attack() string {
	if m.Mana >= 20 {
		damage := 30 + rand.Intn(15)
		m.Mana -= 20
		return fmt.Sprintf("%s casts fireball for %d damage! (mana: %d)", m.Name, damage, m.Mana)
	}
	return fmt.Sprintf("%s has no mana for spells!", m.Name)
}

// TODO: Implement Defend() method for Mage
// - If mana >= 10: reduce mana by 10, return "[name] casts magic shield (mana: [mana])"
// - If mana < 10: return "[name] dodges attack"
// YOUR CODE HERE
func (m *Mage) Defend() string {
	if m.Mana >= 10 {
		m.Mana -= 10
		return fmt.Sprintf("%s casts magic shield (mana: %d)", m.Name, m.Mana)
	}
	return fmt.Sprintf("%s dodges attack", m.Name)
}

// TODO: Implement GetHealth() method for Mage
// Return health field
// YOUR CODE HERE
func (m *Mage) GetHealth() int {
	return m.Health
}

// TODO: Implement TakeDamage(damage int) method for Mage
// - Reduce health by damage
// - Ensure health doesn't go below 0
// YOUR CODE HERE
func (m *Mage) TakeDamage(damage int) {
	m.Health -= damage
	if m.Health < 0 {
		m.Health = 0
	}
}

func demonstrateGamePolymorphism() {
	fmt.Println("\n=== Game Character Polymorphism ===")

	// TODO: Create a slice of GameCharacter with Warrior and Mage instances
	// YOUR CODE HERE
	garen := NewWarrior("Garen")
	zoe := NewMage("Zoe")
	characters := []GameCharacter{garen, zoe}

	// TODO: Print "🗡️ Battle Simulation!"
	// YOUR CODE HERE
	fmt.Println("🗡️ Battle Simulation!")
	// TODO: Simulate 3 rounds of battle
	// For each round:
	// - Print round number
	// - For each character with health > 0: randomly attack or defend
	// - Apply random damage to all characters
	// YOUR CODE HERE
	for round := 1; round <= 3; round++ {
		fmt.Printf("\n🔄 Round %d\n", round)
		for _, character := range characters {
			if character.GetHealth() > 0 {
				action := rand.Intn(2)
				if action == 0 {
					fmt.Printf("%s\n", character.Attack())
				} else {
					fmt.Printf("%s\n", character.Defend())
				}
			}
			damage := rand.Intn(10) + 1
			character.TakeDamage(damage)
		}
	}

	// TODO: Print final status of all characters
	// Show health and alive/KO status
	// YOUR CODE HERE
	fmt.Println("\n🏁 Final Status:")
	for _, character := range characters {
		fmt.Printf("%+v: Health: %d, Alive: %v\n", character, character.GetHealth(), character.GetHealth() > 0)
	}
}

// ===== MAIN DEMO FUNCTION =====

func main() {
	fmt.Println("🔄 Go Polymorphism and Design Patterns Practice")
	fmt.Println("===============================================")

	rand.Seed(time.Now().UnixNano())

	// demonstrateBasicPolymorphism()
	// demonstrateStrategyPattern()
	// TODO: Uncomment these as you implement them
	// demonstrateObserverPattern()
	// demonstrateFactoryPattern()
	// demonstrateAdapterPattern()
	// demonstrateDecoratorPattern()
	demonstrateGamePolymorphism()

	fmt.Println("\n✅ Polymorphism practice completed!")
	fmt.Println("\n🎯 Learning Goals:")
	fmt.Println("- Understand polymorphism through interfaces")
	fmt.Println("- Implement Strategy pattern for algorithm selection")
	fmt.Println("- Apply Observer pattern for event notification")
	fmt.Println("- Use Factory pattern for object creation")
	fmt.Println("- Implement Adapter pattern for interface compatibility")
	fmt.Println("- Apply Decorator pattern for behavior extension")
}
