package main

import "fmt"

var (
	AKnight = Symbol{"A is a Knight"}
	AKnave  = Symbol{"A is a Knave"}
	BKnight = Symbol{"B is a Knight"}
	BKnave  = Symbol{"B is a Knave"}
	CKnight = Symbol{"C is a Knight"}
	CKnave  = Symbol{"C is a Knave"}
)

// Puzzle 0
// A says "I am both a knight and a knave."
var knowledge0 = And{[]Sentence{
	&Or{[]Sentence{
		&And{[]Sentence{&AKnight, &Not{operand: &AKnave}}},
		&And{[]Sentence{&Not{&AKnight}, &AKnave}},
	}},
	&Implication{
		antecedent: &AKnight,
		consequent: &And{[]Sentence{&AKnight, &AKnave}},
	},
	&Implication{
		antecedent: &AKnave,
		consequent: &Not{&And{[]Sentence{&AKnight, &AKnave}}},
	},
}}

// Puzzle 1
// A says "We are both knaves."
// B says nothing.
var knowledge1 = And{[]Sentence{
	&Or{[]Sentence{
		&And{[]Sentence{&AKnight, &Not{&AKnave}}},
		&And{[]Sentence{&Not{&AKnight}, &AKnave}},
	}},
	&Or{[]Sentence{
		&And{[]Sentence{&BKnight, &Not{&BKnave}}},
		&And{[]Sentence{&Not{&BKnight}, &BKnave}},
	}},
	&Implication{
		antecedent: &AKnight,
		consequent: &And{[]Sentence{&AKnave, &BKnave}},
	},
	&Implication{
		antecedent: &AKnave,
		consequent: &Not{&And{[]Sentence{&AKnave, &BKnave}}},
	},
}}

// Puzzle 2
// A says "We are the same kind."
// B says "We are of different kinds."
var knowledge2 = And{[]Sentence{
	&Or{[]Sentence{
		&And{[]Sentence{&AKnight, &Not{&AKnave}}},
		&And{[]Sentence{&Not{&AKnight}, &AKnave}},
	}},
	&Or{[]Sentence{
		&And{[]Sentence{&BKnight, &Not{&BKnave}}},
		&And{[]Sentence{&Not{&BKnight}, &BKnave}},
	}},
	&Implication{
		antecedent: &AKnight,
		consequent: &And{[]Sentence{&AKnight, &BKnight}},
	},
	&Implication{
		antecedent: &AKnave,
		consequent: &Not{&And{[]Sentence{&AKnave, &BKnave}}},
	},
	&Implication{
		antecedent: &BKnight,
		consequent: &And{[]Sentence{&BKnight, &AKnave}},
	},
	&Implication{
		antecedent: &BKnave,
		consequent: &And{[]Sentence{&BKnave, &AKnave}},
	},
}}

// Puzzle 3
// A says either "I am a knight." or "I am a knave.", but you don't know which.
// B says "A said 'I am a knave'."
// B says "C is a knave."
// C says "A is a knight."
var knowledge3 = And{[]Sentence{
	&Or{[]Sentence{
		&And{[]Sentence{&AKnight, &Not{&AKnave}}},
		&And{[]Sentence{&Not{&AKnight}, &AKnave}},
	}},
	&Or{[]Sentence{
		&And{[]Sentence{&BKnight, &Not{&BKnave}}},
		&And{[]Sentence{&Not{&BKnight}, &BKnave}},
	}},
	&Or{[]Sentence{
		&And{[]Sentence{&CKnight, &Not{&CKnave}}},
		&And{[]Sentence{&Not{&CKnight}, &CKnave}},
	}},
	&Implication{
		antecedent: &AKnight,
		consequent: &Or{[]Sentence{&AKnight, &AKnave}},
	},
	&Implication{
		antecedent: &AKnave,
		consequent: &Not{&And{[]Sentence{&AKnight, &AKnave}}},
	},
	&Implication{
		antecedent: &BKnight,
		consequent: &And{[]Sentence{
			&Implication{&AKnight, &AKnave},
			&Implication{&AKnave, &Not{&AKnave}},
		}},
	},
	&Implication{
		antecedent: &BKnave,
		consequent: &And{[]Sentence{
			&Implication{&AKnight, &AKnight},
			&Implication{&AKnave, &Not{&AKnight}},
		}},
	},
	&Implication{&BKnight, &CKnave},
	&Implication{&BKnave, &CKnight},
	&Implication{&CKnight, &AKnight},
	&Implication{&CKnave, &AKnave},
}}

func main() {
	symbols := []Symbol{AKnight, AKnave, BKnight, BKnave, CKnight, CKnave}

	puzzles := map[string]And{
		"Puzzle 0": knowledge0,
		"Puzzle 1": knowledge1,
		"Puzzle 2": knowledge2,
		"Puzzle 3": knowledge3,
	}

	for puzzle, knowledge := range puzzles {
		fmt.Println(puzzle)
		if len(knowledge.conjuncts) == 0 {
			fmt.Println("    Not yet implemented.")
		} else {
			for _, symbol := range symbols {
				if modelCheck(&knowledge, &symbol) {
					fmt.Printf("    %s\n", symbol.name)
				}
			}
		}
	}
}
