class MinStack:
    def __init__(self):
        self.stack = [(None, float('inf'))]

    def push(self, val: int) -> None:
        self.stack.append((val, min(val, self.stack[-1][1])))

    def pop(self) -> None:
        self.stack.pop()

    def top(self) -> int:
        return self.stack[-1][0]

    def getMin(self) -> int:
        return self.stack[-1][1]


if __name__ == "__main__":
    tests = []

    actual = []
    ms = MinStack()
    ms.push(-2)
    ms.push(0)
    ms.push(-3)
    actual.append(ms.getMin())
    ms.pop()
    actual.append(ms.top())
    actual.append(ms.getMin())
    expected = [-3,0,-2]
    tests.append((actual, expected))

    actual = []
    ms = MinStack()
    ms.push(0)
    ms.push(1)
    ms.push(0)
    actual.append(ms.getMin())
    ms.pop()
    actual.append(ms.getMin())
    expected = [0,0]
    tests.append((actual, expected))

    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        actual, expected = test
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------