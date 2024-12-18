package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"sort"
	"time"
)

func main() {
	resp, err := http.Get("https://holland2stay.com/residences?available_to_book%5Bfilter%5D=Available+to+book%2C179&page=1")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.StatusCode, "------------")
	// fmt.Println(string(all))

	file, err := os.OpenFile("./log.html", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	nn, err := writer.Write(all)
	fmt.Println(nn, err)
}

func worker(ctx context.Context) {
	n := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("噢，卖糕的，我死了.....")
			return
		default:
			go func(n int) {
				for {
					fmt.Println("我是：", n)
					time.Sleep(time.Second)
				}
			}(n)
			n++
			fmt.Println("嘿嘿：", n)
			time.Sleep(time.Second)
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var list = new(ListNode)
	current := list
	var high int
	for l1 != nil || l2 != nil {
		sum := high
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		high = sum / 10
		current.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}

		current = current.Next
	}

	if high > 0 {
		current.Next = &ListNode{
			Val:  high,
			Next: nil,
		}
	}

	return list.Next
}

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	length := 1
	for i, v := range s {
		t := 1
		mp := make(map[byte]int)
		mp[byte(v)] = 1
		for k := i + 1; k < n; k++ {
			if mp[s[k]] == 1 {
				break
			}
			mp[s[k]] = 1
			t++
		}

		if t > length {
			length = t
		}
	}

	return length
}

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
// 算法的时间复杂度应该为 O(log (m+n)) 。
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)

	n := len(nums1)
	fmt.Println(n / 2)
	if n%2 == 1 {
		return float64(nums1[n/2])
	} else {
		return float64(nums1[n/2-1]+nums1[n/2]) / 2
	}

}

// 5. 最长回文子串
// 给你一个字符串 s，找到 s 中最长的回文子串。

// abaccabafy
// "cbbd"
func longestPalindrome(s string) string {
	n := len(s)
	l, r := 0, 0
	round := func(left, right int) (int, int) {
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
		}

		return left + 1, right - 1
	}
	for i := 0; i < n; i++ {
		l1, r1 := round(i, i)
		l2, r2 := round(i, i+1)

		if r1-l1 > r-l {
			l, r = l1, r1
		}

		if r2-l2 > r-l {
			l, r = l2, r2
		}

	}

	return s[l : r+1]
}

// 错误
func longestPalindrome2(s string) string {
	n := len(s)
	str := ""
	// fmt.Println(s[0:2])
	for i := 0; i < n; i++ {
		left := i
		right := i

		for left >= 0 && right < n {
			if s[left] != s[right] {
				if len(s[left:right]) > len(str) {
					str = s[left+1 : right-1]
				}
				break
			}

			left--
			right++
		}

		left2 := i
		righ2 := i + 1
		for left2 >= 0 && righ2 < n {
			if s[left2] != s[righ2] {
				if len(s[left2:righ2]) > len(str) {
					fmt.Println(left2, righ2, "-----")
					str = s[left2+1 : righ2-1]
				}
				break
			}

			left2--
			righ2++
		}

	}

	return str
}

// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
//
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
//
// P   A   H   N
// A P L S I I G
// Y   I   R
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
//
// 请你实现这个将字符串进行指定行数变换的函数：
//
// func convert(s string, numRows int) string {
//
// 	n := len(s)
// 	var str = make([]string, n)
//
// 	for k, v := range s {
// 		l := numRows - 2
// 		if l > 0 {
//
// 		}
// 	}
//
// 	return ""
// }

// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//
// 如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
//
// 假设环境不允许存储 64 位整数（有符号或无符号）。

func reverse(x int) int {
	var n int
	for x != 0 {
		n = 10*n + x%10
		x = x / 10
	}

	if n > int(math.Pow(2, 31))-1 || n < int(math.Pow(-2, 31)) {
		n = 0
	}

	return n
}

func reverse2(x int) int {
	var n int
	for x != 0 {
		n = 10*n + x%10
		if n > 2147483647 || n < -2147483648 {
			return 0
		}
		x = x / 10
	}

	return n
}
