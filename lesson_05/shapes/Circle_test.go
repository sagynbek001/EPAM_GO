package shapes

import (
	"math"
	"testing"
)

func TestCircle_Area(t *testing.T) {
	tests := []struct {
		name    string
		c       Circle
		want    float64
		wantErr bool
	}{
		{
			name:    "no zero",
			c:       Circle{0},
			want:    0,
			wantErr: true,
		},
		{
			name:    "no negative",
			c:       Circle{-1},
			want:    0,
			wantErr: true,
		},
		{
			name:    "regular",
			c:       Circle{2},
			want:    4 * math.Pi,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Circle.Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Circle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Perimeter(t *testing.T) {
	tests := []struct {
		name    string
		c       Circle
		want    float64
		wantErr bool
	}{
		{
			name:    "no zero",
			c:       Circle{0},
			want:    0,
			wantErr: true,
		},
		{
			name:    "no negative",
			c:       Circle{-1},
			want:    0,
			wantErr: true,
		},
		{
			name:    "regular",
			c:       Circle{2},
			want:    4 * math.Pi,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Perimeter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Circle.Perimeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Circle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_String(t *testing.T) {
	tests := []struct {
		name string
		c    Circle
		want string
	}{
		{
			name: "regular",
			c:    Circle{2},
			want: "Circle: radius 2.00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Circle.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
