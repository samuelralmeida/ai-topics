package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Sentence interface {
	evaluate(model map[string]bool) bool
	formula() string
	symbols() map[string]bool
}

func validate(sentence Sentence) {
	if sentence == nil {
		panic("must be a logical sentence")
	}
}

func parenthesize(s string) string {
	balanced := func(s string) bool {
		count := 0
		for _, c := range s {
			if c == '(' {
				count++
			} else if c == ')' {
				if count <= 0 {
					return false
				}
				count--
			}
		}
		return count == 0
	}

	if len(s) == 0 || isAlpha(s) || (s[0] == '(' && s[len(s)-1] == ')' && balanced(s)) {
		return s
	}

	return fmt.Sprintf("(%s)", s)
}

func isAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

type Symbol struct {
	name string
}

func (s *Symbol) evaluate(model map[string]bool) bool {
	val, ok := model[s.name]
	if !ok {
		panic(fmt.Sprintf("variable %s not in model", s.name))
	}
	return val
}

func (s *Symbol) formula() string {
	return s.name
}

func (s *Symbol) symbols() map[string]bool {
	return map[string]bool{s.name: true}
}

type Not struct {
	operand Sentence
}

func (n *Not) evaluate(model map[string]bool) bool {
	return !n.operand.evaluate(model)
}

func (n *Not) formula() string {
	return "¬" + parenthesize(n.operand.formula())
}

func (n *Not) symbols() map[string]bool {
	return n.operand.symbols()
}

type And struct {
	conjuncts []Sentence
}

func (a *And) evaluate(model map[string]bool) bool {
	for _, conjunct := range a.conjuncts {
		if !conjunct.evaluate(model) {
			return false
		}
	}
	return true
}

func (a *And) formula() string {
	if len(a.conjuncts) == 1 {
		return a.conjuncts[0].formula()
	}
	formulas := make([]string, len(a.conjuncts))
	for i, conjunct := range a.conjuncts {
		formulas[i] = parenthesize(conjunct.formula())
	}
	return strings.Join(formulas, " ∧ ")
}

func (a *And) symbols() map[string]bool {
	symbols := make(map[string]bool)
	for _, conjunct := range a.conjuncts {
		for k, v := range conjunct.symbols() {
			symbols[k] = v
		}
	}
	return symbols
}

type Or struct {
	disjuncts []Sentence
}

func (o *Or) evaluate(model map[string]bool) bool {
	for _, disjunct := range o.disjuncts {
		if disjunct.evaluate(model) {
			return true
		}
	}
	return false
}

func (o *Or) formula() string {
	if len(o.disjuncts) == 1 {
		return o.disjuncts[0].formula()
	}
	formulas := make([]string, len(o.disjuncts))
	for i, disjunct := range o.disjuncts {
		formulas[i] = parenthesize(disjunct.formula())
	}
	return fmt.Sprintf("(%s)", strings.Join(formulas, " ∨ "))
}

func (o *Or) symbols() map[string]bool {
	symbols := make(map[string]bool)
	for _, disjunct := range o.disjuncts {
		for k, v := range disjunct.symbols() {
			symbols[k] = v
		}
	}
	return symbols
}

type Implication struct {
	antecedent Sentence
	consequent Sentence
}

func NewImplication(antecedent, consequent Sentence) *Implication {
	validate(antecedent)
	validate(consequent)
	return &Implication{
		antecedent: antecedent,
		consequent: consequent,
	}
}

func (i *Implication) evaluate(model map[string]bool) bool {
	return (!i.antecedent.evaluate(model)) || i.consequent.evaluate(model)
}

func (i *Implication) formula() string {
	antecedent := parenthesize(i.antecedent.formula())
	consequent := parenthesize(i.consequent.formula())
	return fmt.Sprintf("%s => %s", antecedent, consequent)
}

func (i *Implication) symbols() map[string]bool {
	symbols := make(map[string]bool)
	for k, v := range i.antecedent.symbols() {
		symbols[k] = v
	}
	for k, v := range i.consequent.symbols() {
		symbols[k] = v
	}
	return symbols
}

type Biconditional struct {
	left  Sentence
	right Sentence
}

func NewBiconditional(left, right Sentence) *Biconditional {
	validate(left)
	validate(right)
	return &Biconditional{
		left:  left,
		right: right,
	}
}

func (b *Biconditional) evaluate(model map[string]bool) bool {
	return (b.left.evaluate(model) && b.right.evaluate(model)) ||
		(!b.left.evaluate(model) && !b.right.evaluate(model))
}

func (b *Biconditional) formula() string {
	left := parenthesize(b.left.formula())
	right := parenthesize(b.right.formula())
	return fmt.Sprintf("%s <=> %s", left, right)
}

func (b *Biconditional) symbols() map[string]bool {
	symbols := make(map[string]bool)
	for k, v := range b.left.symbols() {
		symbols[k] = v
	}
	for k, v := range b.right.symbols() {
		symbols[k] = v
	}
	return symbols
}

func modelCheck(knowledge Sentence, query Sentence) bool {
	symbols := unionSets(knowledge.symbols(), query.symbols())
	return checkAll(knowledge, query, symbols, make(map[string]bool))
}

func unionSets(set1 map[string]bool, set2 map[string]bool) map[string]bool {
	// Returns the union of two sets
	unionSet := make(map[string]bool)
	for k, v := range set1 {
		unionSet[k] = v
	}
	for k, v := range set2 {
		unionSet[k] = v
	}
	return unionSet
}

func checkAll(knowledge Sentence, query Sentence, symbols map[string]bool, model map[string]bool) bool {
	if len(symbols) == 0 {
		if knowledge.evaluate(model) {
			return query.evaluate(model)
		}
		return true
	}

	// Choose one of the remaining unused symbols
	remaining := copyMap(symbols)
	var p string
	for k := range remaining {
		p = k
		break
	}
	delete(remaining, p)

	// Create a model where the symbol is true
	modelTrue := copyMap(model)
	modelTrue[p] = true

	// Create a model where the symbol is false
	modelFalse := copyMap(model)
	modelFalse[p] = false

	// Ensure entailment holds in both models
	return checkAll(knowledge, query, remaining, modelTrue) && checkAll(knowledge, query, remaining, modelFalse)
}

func copyMap(m map[string]bool) map[string]bool {
	copyMap := make(map[string]bool)
	for k, v := range m {
		copyMap[k] = v
	}
	return copyMap
}
