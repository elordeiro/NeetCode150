package main

import "container/heap"

type Tweet struct {
	time int
	id   int
}

type User struct {
	following map[int]struct{}
	tweets    []Tweet
}

type Twitter struct {
	users map[int]User
	time  int
}

func Constructor() Twitter {
	twitter := Twitter{users: map[int]User{}, time: 0}
	return twitter
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.users[userId]; !ok {
		this.users[userId] = User{following: make(map[int]struct{}), tweets: make([]Tweet, 0)}
	}
	user := this.users[userId]
	tweets := user.tweets
	tweets = append(tweets, Tweet{time: this.time, id: tweetId})
	user.tweets = tweets
	this.users[userId] = user
	this.time += 1
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	hp := TweetHeap{}
	heap.Init(&hp)

	for friend := range this.users[userId].following {
		tweets := this.users[friend].tweets
		for _, tweet := range tweets {
			heap.Push(&hp, tweet)
		}
	}

	for _, tweet := range this.users[userId].tweets {
		heap.Push(&hp, tweet)
	}

	res := []int{}
	for hp.Len() > 0 && len(res) < 10 {
		res = append(res, heap.Pop(&hp).(Tweet).id)
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.users[followerId]; !ok {
		this.users[followerId] = User{following: map[int]struct{}{}, tweets: make([]Tweet, 0)}
	}
	this.users[followerId].following[followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.users[followerId].following, followeeId)
}

type TweetHeap []Tweet

func (heap TweetHeap) Len() int           { return len(heap) }
func (heap TweetHeap) Less(i, j int) bool { return heap[i].time > heap[j].time }
func (heap TweetHeap) Swap(i, j int)      { heap[i], heap[j] = heap[j], heap[i] }
func (heap *TweetHeap) Push(val any)      { *heap = append(*heap, val.(Tweet)) }
func (heap *TweetHeap) Pop() any {
	old := *heap
	n := len(old)
	x := old[n-1]
	*heap = old[:n-1]
	return x
}
