package main

import (
	"fmt"
	"sort"
)

//func sortJumbled(mapping []int, nums []int) []int {
//
//}

type customizedSort []int

func (s customizedSort) Len() int {
	return len(s)
}
func (s customizedSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s customizedSort) Less(i, j int) bool {

	return s[i] < s[j]
}

func main() {
	fruits := []int{11, 33, 4}
	sort.Sort(customizedSort(fruits))
	fmt.Println(fruits)
}

//import java.util.Arrays;
//
//public class LeetcodeM {
//
//public int[] sortJumbled(int[] mapping, int[] nums) {
//
//int n = nums.length;
//Integer[] indexes = new Integer[n];
//for (int i = 0; i < indexes.length; i++) {
//indexes[i] = i;
//}
//Arrays.sort(
//indexes,
//(i, j) -> {
//int x = 0, y = 0;
//int tx = nums[i], ty = nums[j];
//int pro = 1;
//if (tx == 0) {
//x = mapping[0];
//} else {
//while (tx > 0) {
//x = mapping[(tx % 10)] * pro + x;
//tx /= 10;
//pro *= 10;
//}
//}
//
//pro = 1;
//if (ty == 0) {
//y = mapping[0];
//} else {
//while (ty > 0) {
//y = mapping[(ty % 10)] * pro + y;
//ty /= 10;
//pro *= 10;
//}
//}
//
//if (x == y) {
//return i - j;
//} else {
//return x - y;
//}
//});
//int[] res = new int[n];
//for (int i = 0; i < res.length; i++) {
//res[i] = nums[indexes[i]];
//}
//return res;
//}
//
//public static void main(String[] args) {
//
//LeetcodeM ins = new LeetcodeM();
//ins.sortJumbled(
//new int[] {8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, new int[] {0, 1, 2, 3, 4, 5, 6, 7, 8, 9});
//}
//}
