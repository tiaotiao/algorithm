
class MonotonicStack:
    def __init__(self, key=None):
        self.stack = []
        if key == None:
            key = lambda x: x
        self.key = key

    def push(self, item):
        out = []
        while len(self.stack) > 0:
            k = self.key(item)
            top = self.key(self.stack[-1])
            if top < k:
                break
            out.append(self.pop())
        self.stack.append(item)
        return out
    
    def pop(self):
        if len(self.stack) == 0:
            return None
        return self.stack.pop()

    def __len__(self):
        return len(self.stack)

    def __getitem__(self, i):
        return self.stack[i]

