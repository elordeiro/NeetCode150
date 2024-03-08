class Solution:
    def trap(self, height: list[int]) -> int:
        if (n := len(height)) < 3:
            return 0
        
        total_water = 0
        water = 0
        l, r = 0, 1
        while r < n:
            if height[r] < height[l]:
                water += height[l] - height[r]
            else:
                total_water += water
                water = 0
                l = r
            r += 1
        
        n = l
        r -= 1
        l = r - 1
        while l > n:
            if height[l] < height[r]:
                total_water += height[r] - height[l]
            else:
                r = l
            l -= 1
        
        return total_water


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([0,1,0,2,1,0,1,3,2,1,2,1], 6),
        ([4,2,0,3,2,5], 9)
    ]
    
    passed_all = True
    for i, test in enumerate(tests, 1):
        height, expected = test
        actual = sol.trap(height)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")





# ---------------------------- Other Solutions --------------------------------
def trap(self, height: list[int]) -> int:
        if (n := len(height)) < 3:
            return 0
        
        water = 0

        l = r = 0
        while r < n:
            r = l + 1
            while r < n and height[l] <= height[r]:
                l += 1
                r += 1
            
            bucket = 0
            left_wall = height[l]
            while r < n and left_wall > (right_wall := height[r]):
                bucket += left_wall - right_wall
                r += 1
            
            if r == n:
                break

            water += bucket
            l = r
        
        n = l
        r -= 1
        l = r - 1
        while l > n:
            l = r - 1
            while l > n and height[l] >= height[r]:
                l -= 1
                r -= 1
            
            right_wall = height[r]
            while l > n and right_wall > (left_wall := height[l]):
                water += right_wall - left_wall
                l -= 1

            r = l
        
        return water

def trap(self, height: list[int]) -> int:
        if (n := len(height)) < 3:
            return 0
        
        total_water = 0
        l, r = 0, n - 1
        left_bar, right_bar = height[l], height[r]
        while l <= r:
            if left_bar < right_bar:
                water = left_bar - (bar := height[l])
                l += 1
                if bar > left_bar:
                    left_bar = bar
            else:
                water = right_bar - (bar := height[r])
                r -= 1
                if bar > right_bar:
                    right_bar = bar
            
            if water > 0:
                total_water += water
        
        return total_water