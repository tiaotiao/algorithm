
struct Trie {
    int value;
    struct Trie* child[256];
};

struct Trie* trieNewNode() {
    struct Trie* t = (struct Trie*)malloc(sizeof(struct Trie));
    t->value = -1;
    //t->child = (struct Trie**)malloc(256*sizeof(struct Trie*));
    memset((void*)t->child, 0, 256*sizeof(struct Trie*));
    return t;
}

void TrieInsert(struct Trie* root, char* s, int value) {
    struct Trie* p = root;
    int i = 0;
    while (s[i] > 0) {
        if (p->child[s[i]] == NULL) {
            break;
        }
        p = p->child[s[i]];
        i++;
    }
    while (s[i] > 0) {
        p->child[s[i]] = trieNewNode();
        p = p->child[s[i]];
        i++;
    }
    p->value = value;
}

int TrieFind(struct Trie* root, char *s) {
    struct Trie* p = root;
    int i = 0;
    while (s[i] > 0) {
        if (p->child[s[i]] == NULL) {
            return -1;
        }
        p = p->child[s[i]];
        i++;
    }
    return p->value;
}
