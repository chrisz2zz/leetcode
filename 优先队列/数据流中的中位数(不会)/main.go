package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	obj.AddNum(12)
	fmt.Println(obj.FindMedian())
	obj.AddNum(10)
	fmt.Println(obj.FindMedian())
	obj.AddNum(13)
	fmt.Println(obj.FindMedian())
	obj.AddNum(11)
	fmt.Println(obj.FindMedian())
	obj.AddNum(5)
	fmt.Println(obj.FindMedian())
	obj.AddNum(15)
	fmt.Println(obj.FindMedian())
	obj.AddNum(6)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())
	obj.AddNum(1)
	fmt.Println(obj.FindMedian())
	obj.AddNum(0)
	fmt.Println(obj.FindMedian())
	obj.AddNum(0)
	fmt.Println(obj.FindMedian())
}

type MedianFinder struct {
	nums []int
}


func Constructor() MedianFinder {

	return MedianFinder{nums: make([]int, 0)}
}


func (this *MedianFinder) AddNum(num int)  {
	if len(this.nums) == 0 {
		this.nums = append(this.nums, num)
		return
	}

	if this.nums[0] > num {
		this.nums = append([]int{num}, this.nums...)
		return
	}

	if this.nums[len(this.nums)-1] < num {
		this.nums = append(this.nums, num)
		return
	}

	for i := 1; i < len(this.nums); i++ {
		if this.nums[i] > num {
			t := make([]int, i)
			copy(t, this.nums[:i])
			t = append(t, num)
			t = append(t, this.nums[i:]...)
			this.nums = t
			break
		}
	}
}


func (this *MedianFinder) FindMedian() float64 {
	length := len(this.nums)
	mid := length/2
	if length % 2 == 0 {
		return float64(this.nums[mid-1] + this.nums[mid]) / 2
	} else {
		return float64(this.nums[mid])
	}
}