/*
    RMQ(range minimum query) solves the problems that finding the minimal value in a sub-array of an array.
    
    https://en.wikipedia.org/wiki/Range_minimum_query
    
    This algorithm is called ST(Sparse Table),
    costs O(nlogn) to preprocess, and O(1) for each query.
*/

package algorithm

type RMQ struct {
    st [][]int
    minimum bool
}

func NewRMQ(a []int, minimum bool) *RMQ {
    n := len(a)
    rmq := new(RMQ)
    rmq.st = make([][]int, n)
    rmq.minimum = minimum
    
    var m = 0
    for int(1 << uint32(m)) <= n {
        m += 1
    }
    
    for i := 0; i < n; i++ {
        rmq.st[i] = make([]int, m)
        rmq.st[i][0] = a[i]
    }
    
    // DP: st[i][p] = min(st[i][p-1], st[i+2^(p-1)][p-1])
    for p := 1; p < m; p++ {
        k := int(1 << uint32(p - 1))
        for i := 0; i < n - k; i++ {
            if minimum {
                rmq.st[i][p] = min(rmq.st[i][p-1], rmq.st[i+k][p-1])
            } else {
                rmq.st[i][p] = max(rmq.st[i][p-1], rmq.st[i+k][p-1])
            }
        }
    }
    
    return rmq
}

func (rmq *RMQ) Quary(i, j int) int {
    var p, k = 0, 1
    for k <= (j-i+1)  {
        p += 1
        k = int(1 << uint32(p))
    }
    p -= 1
    k = int(1 << uint32(p))
    
    if rmq.minimum {
        return min(rmq.st[i][p], rmq.st[j-k+1][p])
    }
    return max(rmq.st[i][p], rmq.st[j-k+1][p])
}
