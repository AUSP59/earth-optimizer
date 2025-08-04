
package test

import (
    "testing"
    "optiterra/internal/core"
)

func TestOptimizationBasic(t *testing.T) {
    result := core.Optimize(10, 5)
    if result != 15 {
        t.Errorf("Expected 15, got %v", result)
    }
}
