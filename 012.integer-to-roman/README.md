# Integer to roman
将数字转为罗马数字
```
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
```

**Example 1:**

```
Input: num = 3
Output: "III"
Explanation: 3 is represented as 3 ones.
```

**Example 2:**

```
Input: num = 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.
```

**Example 3:**

```
Input: num = 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
```

 

**Constraints:**

- `1 <= num <= 3999`

## 解题思路

分别将 1-9，10-90，100-900，1000-3000 的罗马字符写下来
```go
var (
	s1    = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	s10   = []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	s100  = []string{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	s1000 = []string{"M", "MM", "MMM"}
)

```

可以看到非常有规律，在 1-3 时为 `I*3`，在到 4 时进位到`IV`，在第 5-8 位置时重复 1-3 的步骤

我们可以将数字拆分，再带入到不同的 slice 中获取对应的值
如
```
1994 = 1000*1 + 100*9 + 10*9 + 4*1
2741 = 1000*2 + 100*7 + 10*4 + 1*1
```

最终超过 95 的人

```go
var (
	s1    = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	s10   = []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	s100  = []string{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	s1000 = []string{"M", "MM", "MMM"}
)

func intToRoman(i int) string {
	var (
		k1000, k100, k10, k int
	)
	k1000 = i / 1000
	result := ""
	if k1000 > 0 {
		result += s1000[k1000-1]
		i = i - k1000*1000
	}
	k100 = i / 100
	if k100 > 0 {
		result += s100[k100-1]
		i = i - k100*100
	}
	k10 = i / 10
	if k10 > 0 {
		result += s10[k10-1]
		i = i - k10*10
	}
	k = i
	if k > 0 {
		result += s1[k-1]
	}
	return result
}

func main() {
	n := intToRoman(18)
	fmt.Println(n)
}

```