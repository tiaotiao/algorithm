
#include <stdlib.h>
#include <string.h>

typedef struct {
    int n;
    int cap;
    int* nums;
} Heap;

void heapDown(Heap* h, int i) {
    int x = h->nums[i];
    int current = i;

    while(current <= h->n/2) {
        int left = current * 2;
        int right = left + 1;
        int child;
        if (right > h->n || h->nums[left] < h->nums[right]) {
            child = left;
        } else {
            child = right;
        }

        if (h->nums[current] <= h->nums[child]) {
            break;
        }

        h->nums[current] = h->nums[child];
        current = child;
    }

    h->nums[current] = x;
}

void heapUp(Heap* h, int i) {
    int x = h->nums[i];
    int current = i;

    while (current > 1) {
        int parent = current / 2;
        if (h->nums[parent] <= h->nums[current]) {
            break;
        }
        h->nums[current] = h->nums[parent];
    }

    h->nums[current] = x;
}

void resize(Heap* h) {
    int newSize = h->cap * 2;
    if (newSize <= 0) {
        newSize = 1;
    }
    h->cap = newSize;
    realloc(h->nums, newSize * sizeof(int) + 1);
}

Heap* Build(int n, int* nums) {
    Heap* h = malloc(sizeof(Heap));
    
    h->n = n;
    h->cap = n;
    h->nums = malloc(n * sizeof(int) + 1);

    if (nums == NULL) {
        memset(h->nums, 0, n * sizeof(int) + 1);
        return h;
    }

    for (int i = h->n / 2; i >= 1; i--) {
        heapDown(h, i);
    }
    return h;
}

void Insert(Heap* h, int value) {
    if (h->n >= h->cap) {
        resize(h);
    }
    h->nums[++h->n] = value;
    heapUp(h, h->n);
}

int Extract(Heap* h) {
    if (h->n <= 0) {
        return -1;
    }
    int value = h->nums[1];
    h->nums[1] = h->nums[h->n--];
    heapDown(h, 1);
}





/////////////////////////////////////////////////////////////////////////////////////////////
// generalized


typedef int (*Comparer)(const void* a, const void* b);  // type of compare function
typedef void (*HeapIndex)(void* v, int idx); // feedback index in heap 

void heapifyDown(void* heap, int i, int n, int size, HeapIndex index, Comparer cmp) {
    void* x = malloc(size);
    memcpy(x, heap+i*size, size);
    
    int j;
    while(i <= n/2) {
        j = i * 2;
        if (j < n && cmp(heap+j*size, heap+(j+1)*size) > 0) {
            j += 1;
        }
        if (cmp(x, heap+j*size) <= 0) {
            break;
        }
        
        memcpy(heap+i*size, heap+j*size, size);
        if (index != NULL) {
            index(heap+i*size, i);
        }
        i = j;
    }

    memcpy(heap+i*size, x, size);
    free(x);
    if (index != NULL) {
        index(heap+i*size, i);
    }
}

void heapifyUp(void* heap, int i, int size, HeapIndex index, Comparer cmp) {
    void* x = malloc(size);
    memcpy(x, heap+i*size, size);

    int j = i;
    while (i > 1) {
        j = i / 2;
        if (cmp(heap+i*size, heap+j*size) >= 0) {
            break;
        }

        memcpy(heap+i*size, heap+j*size, size);
        if (index != NULL) {
            index(heap+i*size, i);
        }
        i = j;
    }

    memcpy(heap+i*size, x, size);
    free(x);
    if (index != NULL) {
        index(heap+i*size, i);
    }
}

void HeapBuild(void* heap, int n, int size, HeapIndex index, Comparer cmp) {
    for (int i = n/2; i >= 1; i--) {
        heapifyDown(heap, i, n, size, index, cmp);
    }
}

void HeapInsert(void* heap, void* v, int* n, int size, HeapIndex index, Comparer cmp) {
    *n += 1;
    memcpy(heap+(*n)*size, v, size);

    if (index != NULL) {
        index(heap+(*n)*size, *n);
    }
    heapifyUp(heap, *n, size, index, cmp);
}

void* HeapExtract(void* heap, int* n, int size, HeapIndex index, Comparer cmp) {
    if (*n < 1) {
        return NULL;
    }
    void* x = malloc(size);
    memcpy(x, heap+1*size, size);

    memcpy(heap+1*size, heap+(*n)*size, size);
    *n -= 1;

    heapifyDown(heap, 1, *n, size, index, cmp);
    return x;
}

int HeapDecrease(void* heap, int i, int size, HeapIndex index, Comparer cmp) {
    heapifyUp(heap, i, size, index, cmp);
    return 0;
}
