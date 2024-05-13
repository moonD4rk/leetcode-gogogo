## Best Time to Buy and Sell Stock

Say you have an array for which the *i*th element is the price of a given stock on day *i*.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

**Example 1:**

```
Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
```

**Example 2:**

```
Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
```
### 解决办法

1.一次遍历，时间复杂度O(n)，如果当前价钱减去前一天价钱为负，则以当前价钱为最低点，并以当前价钱向后找最大利润，直到遇到另外一个比当前价钱还低的情况，以最低价钱再向后计算最大利润。

```go

func maxProfit(prices []int) int {
	if len(prices) == 0{
		return 0
	}
	var max int
	temp := prices[0]
	for _, v := range prices {
		if (v - temp)<0{ // 如果当前价钱减前一天价钱为负
			temp = v
		}else { // 如果当前价钱减前一天价钱为正
			curr := v - temp // 则计算利润
			max = maxInt(curr, max) // 将最大利润存储起来
		}
	}
	return max
}
func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
```