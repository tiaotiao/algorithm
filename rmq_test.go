package algorithm 

import "testing"

func TestRMQ(t *testing.T) {
    rmq := NewRMQ([]int{0}, true)
    
    rmq = NewRMQ([]int{3,5,2,5,4,3,1,6,3}, false)
    querys := [][]int{
        {0, 2, 5},
        {1, 4, 5},
        {0, 0, 3},
        {0, 8, 6},
        {5, 8, 6},
        {8, 8, 3},
    }
    
    for _, q := range querys {
        i, j, result := q[0], q[1], q[2]
        ans := rmq.Quary(i, j)
        if ans != result {
            t.Error(ans, result, "(", i, j, ")")
        }
    }
}