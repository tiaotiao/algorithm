package algorithm

import (
    "testing"
    "reflect"
)

func TestNextPermutation(t *testing.T) {
     var p []int = []int{1, 2, 4, 5, 5, 3}
     var q []int = []int{1, 2, 5, 3, 4, 5}
    
    b := NextPermutation(p)
    if !b {
        t.Fatal("b == false")
    }
    
    if !reflect.DeepEqual(p, q) {
        t.Fatal(p, q)
    }
    
    // a := []int{1, 2, 3, 4}
    // for {
    //     t.Logf("%v\n", a)
    //     b := NextPermutation(a)
    //     if !b {
    //         break
    //     }
    // }
}
