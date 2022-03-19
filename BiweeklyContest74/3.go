package main

/*
import java.util.Comparator;
import java.util.PriorityQueue;

public class Leetcode6022 {

    class Solution {
        public int halveArray(int[] nums) {

            double sum = 0;
            PriorityQueue<Double> pq = new PriorityQueue<>(Comparator.reverseOrder());
            for (int num : nums) {
                sum += num;
                pq.add((double) num);
            }
            double target = sum / 2.0;
            int res = 0;
            while (sum > target) {

                double front = pq.poll();
                System.out.println(front);
                sum -= front / 2.0;
                res++;
                pq.add(front / 2.0);
            }
            return res;
        }
    }
}

*/
