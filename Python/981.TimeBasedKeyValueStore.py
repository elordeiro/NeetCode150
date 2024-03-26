from collections import defaultdict

class TimeMap:
    def __init__(self):
        self.key_to_time = defaultdict(list)
        self.time_to_value = defaultdict(lambda: "")

    def set(self, key: str, value: str, timestamp: int) -> None:
        self.key_to_time[key].append(timestamp)
        self.time_to_value[timestamp] = value

    def get(self, key: str, timestamp: int) -> str:
        if key not in self.key_to_time:
            return ""
        
        tss = self.key_to_time[key]
        n = len(tss)
        
        l, r = 0, n - 1
        while l <= r:
            m = (l + r) // 2
            if tss[m] < timestamp:
                l = m + 1
            elif timestamp < tss[m]:
                r = m - 1
            else:
                return self.time_to_value[timestamp]

        if r < 0:
            return ""
        
        i = min(l, n - 1)
        while tss[i] > timestamp:
            i -= 1
        
        return self.time_to_value[tss[i]]


if __name__ == "__main__":
    tests = []
    
    tm = TimeMap()
    actual = []
    expected = ["bar", "bar", "bar2", "bar2"]
    tm.set("foo", "bar", 1)
    actual.append(tm.get("foo", 1))
    actual.append(tm.get("foo", 3))
    tm.set("foo", "bar2", 4)
    actual.append(tm.get("foo", 4))
    actual.append(tm.get("foo", 5))
    tests.append((actual, expected))
    
    passed_all = True
    for i, test in enumerate(tests, 1):
        actual, expected = test
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------