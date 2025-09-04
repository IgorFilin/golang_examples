package training

import "testing"

func TestCount(t *testing.T) {
    // Тест 1: пустой slice
    result := CountSymbols([]string{})
    if result != 0 {
        t.Errorf("Expected 0, got %d", result)
    }
	// Тест 2: наполненный символами slice
	result = CountSymbols([]string{"a","b","c"})
    if result != 3 {
        t.Errorf("Expected 3, got %d", result)
    }
	// Тест 2: наполненный словами slice
	result = CountSymbols([]string{"abba","btet","casdasd"})
    if result != 3 {
        t.Errorf("Expected 3, got %d", result)
    }
    // Тест 2: с пустыми символами
	result = CountSymbols([]string{"","btet","casdasd", "test"})
    if result != 3 {
        t.Errorf("Expected 3, got %d", result)
    }
}