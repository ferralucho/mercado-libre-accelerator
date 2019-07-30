
package tickets_test

import (
	"testing"

	"github.com/ferralucho/mercado-libre-accelerator/tickets"
)

func TestGetTotal(t *testing.T) {
	tt := []struct {
		name      string
		price     int
		normales  int
		jubilados int
		invitados int
		expected  int
	}{
		{
			name:      "when price is 100 and [Normales: 0, Jubilados: 0, Invitados: 0] the expected result is 0",
			price:     100,
			normales:  0,
			jubilados: 0,
			invitados: 0,
			expected:  0,
		},
		{
			name:      "when price is 100 and [Normales: 1, Jubilados: 0, Invitados: 0] the expected result is 100",
			price:     100,
			normales:  1,
			jubilados: 0,
			invitados: 0,
			expected:  100,
		},
		{
			name:      "when price is 100 and [Normales: 0, Jubilados: 1, Invitados: 0] the expected result is 50",
			price:     100,
			normales:  0,
			jubilados: 1,
			invitados: 0,
			expected:  50,
		},
		{
			name:      "when price is 100 and [Normales: 0, Jubilados: 0, Invitados: 1] the expected result is 0",
			price:     100,
			normales:  0,
			jubilados: 0,
			invitados: 1,
			expected:  0,
		},
		{
			name:      "when price is 100 and [Normales: 1, Jubilados: 1, Invitados: 1] the expected result is 150",
			price:     100,
			normales:  1,
			jubilados: 1,
			invitados: 1,
			expected:  150,
		},
		{
			name:      "when price is 100, [Normales:2, Jubilados: 3, Invitados: 4] the expected result is 350",
			price:     100,
			normales:  2,
			jubilados: 3,
			invitados: 4,
			expected:  350,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := tickets.GetTotal(tc.price, tc.normales, tc.jubilados, tc.invitados)
			if tc.expected != result {
				t.Errorf("test err: expected %v but got %v", tc.expected, result)
			}
		})
	}
}
