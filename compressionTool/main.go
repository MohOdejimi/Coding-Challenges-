package main

import (
	"fmt"
	"os"
	"bufio"
	"container/heap"
)


type TreeNode struct {
	Freq int
	Char rune
	Right *TreeNode
	Left *TreeNode
}

type MinHeap []*TreeNode 

func (mh MinHeap) Len() int {
		return len(mh)
}

func (mh MinHeap) Less(i, j int) bool {
	return mh[i].Freq < mh[j].Freq
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MinHeap) Push(x any) {
	*mh = append(*mh, x.(*TreeNode))
}

func (mh *MinHeap) Pop() any {
	currentHeap := *mh 
	n := len(currentHeap) 

	if n == 0 {
		return nil
	}

	poppedElement := currentHeap[n - 1]
	*mh =  currentHeap[:n-1]
	return poppedElement
}


	
func main() {
	args := os.Args[1:]
	filepath := args[0]

	file, err := os.Open(filepath)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File not found")
		} else if os.IsPermission(err) {
			fmt.Println("Permisiion Denied")
		} else {
			fmt.Println(err.Error())
		}
	}
	defer file.Close() 

	scanner := bufio.NewScanner(file)	

	input := []string{}

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	characterCount := make(map[rune]int) 

	for _, s := range input {
		for _, ch := range s {
			characterCount[ch]++
		}
	}

	mh := MinHeap{}

	for ch, freq := range characterCount {
		node := &TreeNode{
			Freq: freq,
			Char: ch,
		}
		heap.Push(&mh, node)
	}

	
}




	// Steps 

	// 1. Sort the data by frequency 
	// 2. Choose the two smallest two numbers 
	// 3. Merge them(a new number) to form a tree and reorder the list. 
	// 4. Choose the two smallest number
	// 5. Merge them to form a tree and reorder the list.
	// 6. 
