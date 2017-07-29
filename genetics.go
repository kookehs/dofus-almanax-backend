package main

import (
	"fmt"
	"sort"
)

// Genes contains functions necessary for crossover, mutation, and selection.
type Genes interface {
	Crossover(genes Genes) (Genes, Genes)
	Fitness(genes Genes) float64
	Mutate(chance float64)
	String() string
}

// Population contains information necessary for simulation.
type Population struct {
	activation float64
	generation int
	members    []Genes
	mutation   float64
	target     Genes
}

// NewPopulation returns an empty Population
// with the given target and mutation chance.
func NewPopulation(activation, mutation float64, target Genes) *Population {
	return &Population{
		activation: activation,
		generation: 0,
		members:    []Genes{},
		mutation:   mutation,
		target:     target,
	}
}

// Add adds the given Genes to the members
func (p *Population) Add(genes Genes) {
	p.members = append(p.members, genes)
}

// Crossover produces children from the given parents.
func (p *Population) Crossover(p1, p2 Genes) (Genes, Genes) {
	return p1.Crossover(p2)
}

// Display prints the generation and each member.
func (p *Population) Display() {
	fmt.Printf("Generation: %d\n", p.generation)

	for _, genes := range p.members {
		fmt.Printf("%g - %s\n", genes.Fitness(p.target), genes.String())
	}
}

// Generation simulates a round of selection, crossover, and mutation.
// Termination is also checked at the end of each round.
func (p *Population) Generation() bool {
	p.Sort()
	p.Display()

	p1, p2 := p.Selection()
	c1, c2 := p.Crossover(p1, p2)

	p.members = p.members[2:]
	p.members = append(p.members, c1, c2)

	p.generation++
	match := p.Termination()

	if match {
		p.Sort()
		p.Display()
	}

	return match
}

// Len returns the number of members.
func (p *Population) Len() int {
	return len(p.members)
}

// Less returns if the value at i is less than the value at j.
func (p *Population) Less(i, j int) bool {
	members := p.members
	target := p.target
	return members[i].Fitness(target) < members[j].Fitness(target)
}

// Mutation randomly mutates members of the Population.
func (p *Population) Mutation() {
	for _, genes := range p.members {
		genes.Mutate(p.mutation)
	}
}

// Selection returns genes for crossover.
func (p *Population) Selection() (Genes, Genes) {
	count := len(p.members)
	return p.members[count-1], p.members[count-2]
}

// Sort sorts members of the Population
func (p *Population) Sort() {
	sort.Sort(p)
}

// Swap swaps the value at i with the value at j.
func (p *Population) Swap(i, j int) {
	members := p.members
	members[i], members[j] = members[j], members[i]
}

// Termination returns if the target gene has been met.
func (p *Population) Termination() bool {
	match := false

	for _, genes := range p.members {
		if genes.Fitness(p.target) > p.activation {
			match = true
		}
	}

	return match
}
