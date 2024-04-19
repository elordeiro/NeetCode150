package main

import (
	"reflect"
	"testing"
)

func Test355DesignTwitter(t *testing.T) {
	tests := []struct {
		ops      []string
		args     [][]int
		expected [][]int
	}{
		{
			[]string{"Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"},
			[][]int{{}, {1, 5}, {1}, {1, 2}, {2, 6}, {1}, {1, 2}, {1}},
			[][]int{{}, {}, {5}, {}, {}, {6, 5}, {}, {5}},
		},
		{
			[]string{"Twitter", "postTweet", "postTweet", "getNewsFeed"},
			[][]int{{}, {1, 5}, {1, 3}, {1}},
			[][]int{{}, {}, {}, {3, 5}},
		},
		{
			[]string{"Twitter", "postTweet", "postTweet", "unfollow", "getNewsFeed"},
			[][]int{{}, {1, 4}, {2, 5}, {1, 2}, {1}},
			[][]int{{}, {}, {}, {}, {4}},
		},
		{
			[]string{"Twitter", "postTweet", "follow", "follow", "getNewsFeed"},
			[][]int{{}, {2, 5}, {1, 2}, {1, 2}, {1}},
			[][]int{{}, {}, {}, {}, {5}},
		},
	}

	for _, test := range tests {
		var obj Twitter
		actual := [][]int{}
		ops, args, expected := test.ops, test.args, test.expected
		for i, op := range ops {
			if op == "Twitter" {
				obj = Constructor()
				actual = append(actual, []int{})
			} else if op == "postTweet" {
				obj.PostTweet(args[i][0], args[i][1])
				actual = append(actual, []int{})
			} else if op == "getNewsFeed" {
				actual = append(actual, obj.GetNewsFeed(args[i][0]))
			} else if op == "follow" {
				obj.Follow(args[i][0], args[i][1])
				actual = append(actual, []int{})
			} else if op == "unfollow" {
				obj.Unfollow(args[i][0], args[i][1])
				actual = append(actual, []int{})
			}
		}

		t.Run("Test355DesignTwitter", func(t *testing.T) {
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v\n", actual, expected)
			}
		})

	}

}
