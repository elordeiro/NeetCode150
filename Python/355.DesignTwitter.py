from typing import List
from collections import defaultdict
import heapq


class User:
    def __init__(self) -> None:
        self.following = set()
        self.tweets = []

class Twitter:

    def __init__(self):
        self.users = defaultdict(User)
        self.time = 0

    def postTweet(self, userId: int, tweetId: int) -> None:
        self.users[userId].tweets.append((self.time, tweetId))
        self.time -= 1

    def getNewsFeed(self, userId: int) -> List[int]:
        heap = []
        for friend in self.users[userId].following:
            tweets = self.users[friend].tweets
            for tweet in tweets:
                heapq.heappush(heap, tweet)
        
        for tweet in self.users[userId].tweets:
            heapq.heappush(heap, tweet)
        res = []
        while heap and len(res) < 10:
            res.append(heapq.heappop(heap)[1])
        return res

    def follow(self, followerId: int, followeeId: int) -> None:
        self.users[followerId].following.add(followeeId)

    def unfollow(self, followerId: int, followeeId: int) -> None:
        try:
            self.users[followerId].following.remove(followeeId)
        except:
            pass


null = None
if __name__ == "__main__":
    tests = [
        (
            ["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"], 
            [[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]],
            [null, null, [5], null, null, [6, 5], null, [5]],
        ),
        (
            ["Twitter","postTweet","postTweet","getNewsFeed"],
            [[],[1,5],[1,3],[1]],
            [null,null,null,[3,5]]
        ),
        (
            ["Twitter","postTweet","postTweet","unfollow","getNewsFeed"],
            [[],[1,4],[2,5],[1,2],[1]],
            [None, None, None, None, [4]]
        ),
        (
            ["Twitter","postTweet","follow","follow","getNewsFeed"],
            [[],[2,5],[1,2],[1,2],[1]],
            [null,null,null,null,[5]]
        ),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        
        obj = None
        actual = [None]
        ops, args, expected = test
        
        for j, op in enumerate(ops):
            if op == "Twitter":
                obj = Twitter()
            elif op == "postTweet":
                actual.append(obj.postTweet(*args[j]))
            elif op == "getNewsFeed":
                actual.append(obj.getNewsFeed(*args[j]))
            elif op == "follow":
                actual.append(obj.follow(*args[j]))
            elif op == "unfollow":
                actual.append(obj.unfollow(*args[j]))

        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------