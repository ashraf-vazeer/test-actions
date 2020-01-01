package calculation

// Sub is for finding sub
func Sub(a, b int) int {
        return a - b
}

// Sum is for finding sum
func Sum(a, b int) int {
        return a + b
}
nfvadmin@ilos:~/opensource/src/ci-cd-test/cicd/src/calculation$ cat calc_test.go
package calculation

import "testing"

func TestSub(t *testing.T) {
        total := Sub(5, 5)
        if total != 0 {
                t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 0)
        }
}

func TestSum(t *testing.T) {
        total := Sum(5, 5)
        if total != 10 {
                t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
        }
}
