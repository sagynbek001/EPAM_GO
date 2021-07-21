package shapes

import (
	"testing"
)

func TestRectangle_Area(t *testing.T) {
	tests := []struct {
		name    string
		r       Rectangle
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "no zero",
			r:       Rectangle{0, 0},
			want:    0,
			wantErr: true,
		},
		{
			name:    "no negative",
			r:       Rectangle{-1, 0},
			want:    0,
			wantErr: true,
		},
		{
			name:    "regular",
			r:       Rectangle{1, 1},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Rectangle.Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rectangle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	tests := []struct {
		name    string
		r       Rectangle
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "no zero",
			r:       Rectangle{0, 0},
			want:    0,
			wantErr: true,
		},
		{
			name:    "no negative",
			r:       Rectangle{-1, 0},
			want:    0,
			wantErr: true,
		},
		{
			name:    "regular",
			r:       Rectangle{1, 1},
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Perimeter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Rectangle.Perimeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rectangle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_String(t *testing.T) {
	tests := []struct {
		name string
		r    Rectangle
		want string
	}{
		// TODO: Add test cases.
		{
			name: "regular",
			r:    Rectangle{1, 1},
			want: "Rectangle with height 1.00 and width 1.00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("Rectangle.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
