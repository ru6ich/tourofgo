package practice

import "testing"

func TestPipeline(t *testing.T) {
	nums := Generator(1_000_000)
	primes := Filter(nums)
	sum := Summer(primes)

	expected := 77

	if sum != expected {
		t.Errorf("Ожидали %d, получили %d", expected, sum)
	}
}
