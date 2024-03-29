# [76. minimum-window-substring](https://leetcode.cn/problems/minimum-window-substring/submissions/)

## 题目

给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：
- 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
- 如果 s 中存在这样的子串，我们保证它是唯一的答案。

**示例1：**

~~~
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
~~~

**示例2：**

~~~
输入：s = "a", t = "a"
输出："a"
~~~

**示例3：**

~~~
输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。
~~~

## 题目大意

给定字符串，目标字符串，求解最小覆盖子串

## 解题思路

1. 构建目标子字符串map、窗口map
2. 定义左右指针left=right=0，定义valid=0，定义开始下标start=0
3. 遍历s，从right->len(s)
4. 判断迭代元素是否在指定目标字符串中，对应窗口变化，valid值变化
5. 判断窗口map大小是否等于目标字符串map大小，对应缩小窗口，left变化，start变化，valid变化
6. 返回最新覆盖字符串
