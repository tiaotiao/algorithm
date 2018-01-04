

int* InitUnionFind(int n) {
    int* set = (int*)malloc(n*sizeof(int));
    for (int i = 0; i < n; i++) {
        set[i] = i;
    }
    return set;
}

int Find(int* set, int p) {
    if (set[p] == p) {
        return p;
    }
    int root = Find(set, set[p]);
    set[p] = root;
    return root;
}

int Union(int* set, int p, int q) {
    int root_p = Find(set, p);
    int root_q = Find(set, q);
    set[root_p] = root_q;
    return root_q;
}
