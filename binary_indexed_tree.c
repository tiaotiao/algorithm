
struct BIT {
    int n;
    int* nums; // index starting from 1
    int* c; // cumulative
};

int lowbit(int x) {
    return x & (-x);
}

struct BIT* BITCreate(int* a, int n) {
    struct BIT* b = malloc(sizeof(struct BIT));
    
    b->n = n;
    int size = n * sizeof(int) + 1;
    b->nums = malloc(size);
    b->c = malloc(size);

    memset((void*)b->nums, 0, size);
    memset((void*)b->c, 0, size);

    if (a == NULL) {
        return b;
    }

    // init
    for (int i = 1; i <= n; i++) {
        int x = a[i-1];
        b->nums[i] = x;
        for (int j = i; j <= n; j += lowbit(j)) {
            b->c[j] += x;
        }
    }

    return b;
}

void BITModify(struct BIT* b, int i, int value) {
    int delta = value - b->nums[i];
    b->nums[i] = value;

    for (int j = i; j <= b->n; j += lowbit(j)) {
        b->c[j] += delta;
    }
}

// sum of nums[1 ~ k]
int BITSum(struct BIT* b, int k) {
    int sum = 0;
    for (int j = k; j > 0; j -= lowbit(j)) {
        sum += b->c[j];
    }
    return sum;
}