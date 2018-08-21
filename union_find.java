
class UnionFind {
    private int[] s;
    private int cnt;
    public UnionFind(int size) {
        s = new int[size];
        for (int i = 0; i < size; i++) {
            s[i] = -1;
        }
        cnt = size;
    }

    public int union(int a, int b) {
        int ra = find(a);
        int rb = find(b);
        if (ra != rb) {
            s[ra] = rb;
            cnt -= 1;
        }
        return rb;
    }

    public int find(int a) {
        if (s[a] == -1) {
            return a;
        }
        int root = find(s[a]);
        s[a] = root;
        return root;
    }

    public int count() {
        return cnt;
    }
}