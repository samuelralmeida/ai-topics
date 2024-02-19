package frontier

import (
	"testing"

	"github.com/samuelralmeida/ai-topics/search/entity"
	"github.com/stretchr/testify/assert"
)

func Test_queueFrontier_Remove(t *testing.T) {
	type fields struct {
		frontier []entity.Node
	}
	tests := []struct {
		name   string
		fields fields
		want   entity.Node
	}{
		{
			name: "first-in, first-out",
			fields: fields{frontier: []entity.Node{
				{State: entity.Coordinate{Row: 3, Collumn: 1}},
				{State: entity.Coordinate{Row: 4, Collumn: 2}},
				{State: entity.Coordinate{Row: 5, Collumn: 1}},
			}},
			want: entity.Node{State: entity.Coordinate{Row: 3, Collumn: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qf := &queueFrontier{frontier: tt.fields.frontier}

			got := qf.Remove()
			assert.Equal(t, tt.want, got)
			assert.Len(t, qf.frontier, 2)
		})
	}
}

func Test_stackFrontier_Remove(t *testing.T) {
	type fields struct {
		frontier []entity.Node
	}
	tests := []struct {
		name   string
		fields fields
		want   entity.Node
	}{
		{
			name: "last-in, first-out",
			fields: fields{frontier: []entity.Node{
				{State: entity.Coordinate{Row: 3, Collumn: 1}},
				{State: entity.Coordinate{Row: 4, Collumn: 2}},
				{State: entity.Coordinate{Row: 5, Collumn: 1}},
			}},
			want: entity.Node{State: entity.Coordinate{Row: 5, Collumn: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sf := &stackFrontier{frontier: tt.fields.frontier}

			got := sf.Remove()
			assert.Equal(t, tt.want, got)
			assert.Len(t, sf.frontier, 2)
		})
	}
}

func Test_greedyFrontier_Remove(t *testing.T) {
	type fields struct {
		frontier []entity.Node
	}

	tests := []struct {
		name   string
		fields fields
		goal   entity.Coordinate
		want   entity.Node
	}{
		{
			name: "manhatan distance",
			fields: fields{frontier: []entity.Node{
				{State: entity.Coordinate{Row: 4, Collumn: 2}},
				{State: entity.Coordinate{Row: 3, Collumn: 3}},
				{State: entity.Coordinate{Row: 5, Collumn: 3}},
			}},
			goal: entity.Coordinate{Row: 0, Collumn: 11},
			want: entity.Node{State: entity.Coordinate{Row: 3, Collumn: 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gf := &greedyFrontier{frontier: tt.fields.frontier, goal: tt.goal}

			got := gf.Remove()
			assert.Equal(t, tt.want, got)
			assert.Len(t, gf.frontier, 2)
		})
	}
}

func Test_aStarFrontier_Remove(t *testing.T) {
	type fields struct {
		frontier []entity.Node
	}

	tests := []struct {
		name   string
		fields fields
		goal   entity.Coordinate
		want   entity.Node
	}{
		{
			name: "manhatan distance + reach cost",
			fields: fields{frontier: []entity.Node{
				{State: entity.Coordinate{Row: 0, Collumn: 2}, ReachCost: 25},
				{State: entity.Coordinate{Row: 4, Collumn: 2}, ReachCost: 6},
				{State: entity.Coordinate{Row: 4, Collumn: 9}, ReachCost: 15},
			}},
			goal: entity.Coordinate{Row: 0, Collumn: 11},
			want: entity.Node{State: entity.Coordinate{Row: 4, Collumn: 2}, ReachCost: 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gf := &aStarFrontier{frontier: tt.fields.frontier, goal: tt.goal}

			got := gf.Remove()
			assert.Equal(t, tt.want, got)
			assert.Len(t, gf.frontier, 2)
		})
	}
}
