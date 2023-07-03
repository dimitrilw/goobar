package alphatrie

import (
	"errors"
)

// This rudimentary AlphaTrie is intended for reuse in code challenges.
// It is not intended for production use. It is not a demonstration of good
// code or optimized performance. But for sparse-tree code challenges,
// it works well enough. And for more complex use cases, there are other
// options, like an adaptive radix tree.

type AlphaTrie struct {
	children [26]*AlphaTrie
	isEnd    bool
}

func NewAlphaTrie() *AlphaTrie {
	return &AlphaTrie{
		children: [26]*AlphaTrie{},
		isEnd:    false,
	}
}

func (at *AlphaTrie) Insert(word string) error {
	n := at
	idxList, err := wordToIndexList(word)
	if err != nil {
		return err
	}

	for _, i := range idxList {
		if n.children[i] == nil {
			n.children[i] = NewAlphaTrie()
		}
		n = n.children[i]
	}
	n.isEnd = true
	return nil
}

func (at *AlphaTrie) walkToEnd(s string) (*AlphaTrie, error) {
	n := at
	idxList, err := wordToIndexList(s)
	if err != nil {
		return nil, err
	}

	for _, i := range idxList {
		if n.children[i] == nil {
			return nil, errors.New("not found")
		}
		n = n.children[i]
	}
	return n, nil
}

func (at *AlphaTrie) Search(word string) bool {
	if n, err := at.walkToEnd(word); err != nil {
		return false
	} else {
		return n.isEnd
	}
}

func (at *AlphaTrie) StartsWith(prefix string) bool {
	_, err := at.walkToEnd(prefix)
	return err == nil
}

// func (at *AlphaTrie) Delete(word string) {} // not yet implemented

func rToIdx(r rune) (int, error) {
	if r < 'a' || 'z' < r {
		return -1, errors.New("invalid character: " + string(r))
	}
	return int(r - 'a'), nil
}

func wordToIndexList(word string) ([]int, error) {
	res := make([]int, len(word))
	for i, r := range word {
		if idx, err := rToIdx(r); err != nil {
			return []int{}, errors.New("error in word: " + word + ", " + err.Error())
		} else {
			res[i] = idx
		}
	}
	return res, nil
}
