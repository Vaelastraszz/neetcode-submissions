type Tweet struct {
	id        int
	timestamp int
}

type Twitter struct {
	UsersTweets    map[int][]Tweet
	UsersFollowing map[int]map[int]bool
	timer          int
}

func Constructor() Twitter {
	return Twitter{
		UsersTweets:    make(map[int][]Tweet),
		UsersFollowing: make(map[int]map[int]bool),
		timer:          0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	tweet := Tweet{
		id:        tweetId,
		timestamp: this.timer,
	}

	this.UsersTweets[userId] = append(
		this.UsersTweets[userId],
		tweet,
	)

	this.timer++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	h := &MaxHeap{}
	result := []int{}

	myTweets, ok := this.UsersTweets[userId]
	if ok {
		for _, tweet := range myTweets {
			heap.Push(h, tweet)
		}
	}

	following, ok := this.UsersFollowing[userId]
	if ok {
		for followeeId := range following {
			followeeTweets, ok := this.UsersTweets[followeeId]
			if ok {
				for _, tweet := range followeeTweets {
					heap.Push(h, tweet)
				}
			}
		}
	}

	for i := 0; i < 10 && h.Len() > 0; i++ {
		tweet := heap.Pop(h).(Tweet)
		result = append(result, tweet.id)
	}

	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}

	following, ok := this.UsersFollowing[followerId]

	if !ok {
		following = make(map[int]bool)
		this.UsersFollowing[followerId] = following
	}

	following[followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	following, ok := this.UsersFollowing[followerId]
	if !ok {
		return
	}

	delete(following, followeeId)
}

type MaxHeap []Tweet

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].timestamp > h[j].timestamp
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Tweet))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[:n-1]

	return x
}