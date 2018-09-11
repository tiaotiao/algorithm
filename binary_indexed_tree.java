
import java.util.*;

class BIT {
    int[] sum;
    int[] nums;
    int n;
    
    public BIT(int[] nums) {
        n = nums.length;
        this.nums = nums; 
        this.sum = new int[n+1];
        for (int i = 0; i < n; i++) {
            changeValue(i, nums[i]);
        }
    }
    
    public void update(int i, int val) {
        int delta = val - nums[i];
        nums[i] = val;
        changeValue(i, delta);
    }

    public int sumRange(int i, int j) {
        int s = sumPrefix(j);
        int t = sumPrefix(i-1);
        return s - t;
    }
    
    //////////////////////////////////////////////////
    
    private void changeValue(int i, int delta) {
        int j = i + 1;
        while(j <= n) {
            sum[j] += delta;
            j += lowbit(j);
        }
    }
    
    private int sumPrefix(int i) {
        int s = 0;
        int j = i + 1;
        while(j > 0) {
            s += sum[j];
            j -= lowbit(j);
        }
        return s;
    }
    
    private int lowbit(int x) {
        return x & (-x);
    }
}
