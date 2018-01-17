
#include <stdio.h>
#include <stdlib.h>

void _reverse(int* nums, int from, int to) {
    int tmp;
    while (from < to) {
        tmp = nums[from];
        nums[from] = nums[to];
        nums[to] = tmp;
        from += 1;
        to -= 1;
    }
}

void _merge_without_extra_space(int* nums, int left, int mid, int right) {
    int p = left;
    
    while (p <= mid && mid < right) {
        while (p <= mid && nums[p] <= nums[mid+1]) {
            p++;
        }
        if (p > mid) {
            break;
        }
        int offset = 1;
        while (mid + offset + 1 <= right && nums[p] > nums[mid + offset + 1]) {
            offset++;
        }
        // offset -= 1;

        _reverse(nums, p, mid + offset);
        _reverse(nums, p, p + offset - 1);
        _reverse(nums, p + offset, mid+offset);

        p += offset;
        mid = mid + offset;
    }
}

////////////////////////////////////////////
// merge sort

void _merge(int* nums, int left, int mid, int right) {
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

void _merge_sort(int* nums, int left, int right) {
    if (left >= right) {
        return;
    }

    int mid = (left + right) / 2;

    _merge_sort(nums, left, mid);
    _merge_sort(nums, mid+1, right);

    // _merge(nums, left, mid, right);
    _merge_without_extra_space(nums, left, mid, right);
}

void merge_sort(int* nums, int n) {
    _merge_sort(nums, 0, n-1);
}


//////////////////////////////////////////////////////////////////
// testing

void main() {
    
    int n = 1000;
    int a[n];
    for (int i = 0; i < n; i++) {
        a[i] = rand();
    }

    merge_sort(a, n);

    int ok = 1;
    for (int i = 1; i < n; i++) {
        if (a[i-1] > a[i]) {
            ok = 0;
            printf("error: [%d] %d > [%d] %d\n", i-1, a[i-1], i, a[i]);
            break;
        }
    }

    if (ok) {
        printf("ok \n");
    }
}