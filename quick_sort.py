
def quick_sort(nums, left=None, right=None):
    n = len(nums)
    if left == None: left = 0
    if right == None: right = n-1
    if left >= right: return
    l, r = left, right

    val = nums[l]   # TODO randomly choose pivot
    while l < r:
        while l < r and val <= nums[r]:
            r -= 1
        if l < r:
            nums[l] = nums[r]
            l += 1
        while l < r and nums[l] <= val:
            l += 1
        if l < r:
            nums[r] = nums[l]
            r -= 1
    mid = l
    nums[mid] = val

    quick_sort(nums, left, mid-1)
    quick_sort(nums, mid + 1, right)


def main():
    nums = [3,4,8,5,1,6,0,2,7]
    quick_sort(nums)
    print(nums)

if __name__ == '__main__':
    main()
