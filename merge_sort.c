
void _merge_sort(int* nums, int left, int right) {
    if (left >= right) {
        return;
    }

    int mid = (left + right) / 2;

    _merge_sort(nums, left, mid);
    _merge_sort(nums, mid+1, right);

    int t[right - left + 1], k = 0;

    int p = left, q = mid+1;
    while(p <= mid || q <= right) {
        if (q > right || (p <= mid && nums[p] <= nums[q])) {
            t[k++] = nums[p++];
        } else {
            t[k++] = nums[q++];
        }
    }

    for (int i = left, k = 0; i <= right; i++, k++) {
        nums[i] = t[k];
    }
}

void merge_sort(int* nums, int n) {
    _merge_sort(nums, 0, n-1);
}