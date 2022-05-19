package main

import (
	"fmt"
	"strconv"
)

func main() {

	s := "022"
	fmt.Println(strconv.Atoi(s))
}

//func criticalConnections(n int, connections [][]int) [][]int {

// tarjan算法

//}

/*
class Solution {
private:
    int num = 0;
    vector<int> dfn, low;    // 节点的时间戳和追溯值
    vector<vector<int>> res; // 答案
    vector<vector<int>> adj; // 邻接表

    void tarjan(int u, int fa) { // fa 为 u 的父节点
        dfn[u] = low[u] = ++num;
        for (int v : adj[u]) {
            if (v == fa) continue;
            if (!dfn[v]) {
                tarjan(v, u);
                low[u] = min(low[u], low[v]);
                if (dfn[u] < low[v]) res.push_back({u, v});
            } else {
                low[u] = min(low[u], dfn[v]);
            }
        }
    }

public:
    vector<vector<int>> criticalConnections(int n, vector<vector<int>>& connections) {
        dfn.resize(n); // 初始化一下工具人数组们
        low.resize(n);
        adj.resize(n);
        for (auto edge : connections) { // 邻接表存图
            adj[edge[0]].push_back(edge[1]);
            adj[edge[1]].push_back(edge[0]);
        }
        for (int u = 0; u < n; u++) {
            if (!dfn[u]) tarjan(u, -1); // Tarjan 更新以 u 为根的搜索树。-1 表示根节点无父节点
        }
        return res;
    }
};

作者：zealot27
链接：https://leetcode.cn/problems/critical-connections-in-a-network/solution/1192-cha-zhao-ji-qun-nei-de-guan-jian-li-tjb5/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
