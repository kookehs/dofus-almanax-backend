package main

import (
	"fmt"
	"sort"
)

// Genes contains functions necessary for crossover, mutation, and selection.
type Genes interface {
	Crossover(g Genes)
	Fitness(g Genes) float64
	Mutate()
	String() string
}

// Population contains information necessary for simulation.
type Population struct {
	generation int
	members    []Genes
	target     Genes
}

// NewPopulation returns an empty Population with the given target.
func NewPopulation(t Genes) *Population {
	return &Population{generation: 0, target: t, members: []Genes{}}
}

// Add adds the given Genes to the members
func (p *Population) Add(g Genes) {
	p.members = append(p.members, g)
}

// Display prints the generation and each member.
func (p *Population) Display() {
	fmt.Printf("Generation: %d\n", p.generation)

	for _, v := range p.members {
		fmt.Printf("%g - %s\n", v.Fitness(p.target), v.String())
	}
}

// Generation simulates a round of selection, crossover, and mutation.
func (p *Population) Generation() bool {
	p.Sort()
	p.Display()

	// TODO: Select
	// TODO: Crossover
	// TODO: Mutate

	match := false
	// TODO: Evaluate fitness of genes
	return match
}

// Len returns the number of members.
func (p *Population) Len() int {
	return len(p.members)
}

// Less returns if the value at i is less than the value at j.
func (p *Population) Less(i, j int) bool {
	m := p.members
	t := p.target
	return m[i].Fitness(t) < m[j].Fitness(t)
}

// Sort sorts members of the Population
func (p *Population) Sort() {
	sort.Sort(p)
}

// Swap swaps the value at i with the value at j.
func (p *Population) Swap(i, j int) {
	m := p.members
	m[i], m[j] = m[j], m[i]
}
