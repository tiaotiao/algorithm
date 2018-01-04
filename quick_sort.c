
int _partition(int* nums, int left, int right) {
    int x = nums[left]; // TODO switch with a random one
    int tmp;

    // partition
    int p = left, q = right;
    while(p < q) {
        while (p < q && x <= nums[q]) {
            q--;
        }
        nums[p] = nums[q];
        if (p < q) p++;
        while (p < q && nums[p] >= x) {
            p++;
        }
        nums[q] = nums[p];
        if (p < q) q--;
    }

    nums[p] = x;
    return p;
}

void _quick_sort(int* nums, int left, int right) {
    if (left >= right) {
        return;
    }

    int p = _partition(nums, left, right);

    _quick_sort(nums, 0, p-1);
    _quick_sort(nums, p+1, right);
}

void quick_sort(int* nums, int n) {
    _quick_sort(nums, 0, n-1);
}


