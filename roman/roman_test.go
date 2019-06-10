package roman

import "testing"

func TestRoman(t *testing.T) {
	t.Run("ip 0 expect ", func(t *testing.T) {
		ip := 0
		r := romanNumber(ip)
		expect := ""
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 1 expect I", func(t *testing.T) {
		ip := 1
		r := romanNumber(ip)
		expect := "I"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 2 expect II", func(t *testing.T) {
		ip := 2
		r := romanNumber(ip)
		expect := "II"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 4 expect IV", func(t *testing.T) {
		ip := 4
		r := romanNumber(ip)
		expect := "IV"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 5 expect V", func(t *testing.T) {
		ip := 5
		r := romanNumber(ip)
		expect := "V"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 6 expect VI", func(t *testing.T) {
		ip := 6
		r := romanNumber(ip)
		expect := "VI"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 9 expect IX", func(t *testing.T) {
		ip := 9
		r := romanNumber(ip)
		expect := "IX"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 10 expect X", func(t *testing.T) {
		ip := 10
		r := romanNumber(ip)
		expect := "X"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 11 expect XI", func(t *testing.T) {
		ip := 11
		r := romanNumber(ip)
		expect := "XI"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 14 expect XIV", func(t *testing.T) {
		ip := 14
		r := romanNumber(ip)
		expect := "XIV"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 15 expect XV", func(t *testing.T) {
		ip := 15
		r := romanNumber(ip)
		expect := "XV"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 20 expect XX", func(t *testing.T) {
		ip := 20
		r := romanNumber(ip)
		expect := "XX"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 21 expect XXI", func(t *testing.T) {
		ip := 21
		r := romanNumber(ip)
		expect := "XXI"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 40 expect XL", func(t *testing.T) {
		ip := 40
		r := romanNumber(ip)
		expect := "XL"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 49 expect XLIX", func(t *testing.T) {
		ip := 49
		r := romanNumber(ip)
		expect := "XLIX"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 50 expect L", func(t *testing.T) {
		ip := 50
		r := romanNumber(ip)
		expect := "L"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 51 expect LI", func(t *testing.T) {
		ip := 51
		r := romanNumber(ip)
		expect := "LI"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 90 expect XC", func(t *testing.T) {
		ip := 90
		r := romanNumber(ip)
		expect := "XC"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 99 expect XCIX", func(t *testing.T) {
		ip := 99
		r := romanNumber(ip)
		expect := "XCIX"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})
	t.Run("ip 100 expect C", func(t *testing.T) {
		ip := 100
		r := romanNumber(ip)
		expect := "C"
		if r != expect {
			t.Errorf("expect % #v but got % #v", expect, r)
		}
	})

}
