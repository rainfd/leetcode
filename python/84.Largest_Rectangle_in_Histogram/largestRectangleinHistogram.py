
"""
    初始化栈底为[-1]
    将所有元素的序号推进栈，直到a[i-1] > a[i] 数据不再递增
    这时候逐一计算栈内元素的面积
    假设给定的高度为[5,6,1]，现在栈内为[-1, 0(5), 1(6)],i=2
    第一个计算的面积：           1(6)出栈
        宽度 = i - 1 - stack[top-1] = 2 - 1 - 0 = 1
        高度 = a[stack[top]] = 6
    第二个计算的面积:            0(5)出栈
        宽度 = i - 1 - stack[top-1] = 2 - 1 - (-1) = 2
        高度 = a[stack[top]] = 5
"""
class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        maxArea = 0
        stack = [-1]
        for i, h in enumerate(heights):
            while len(stack) > 1 and heights[stack[-1]] > heights[i]:
                area = (i-1-stack[-2]) * heights[stack[-1]]
                maxArea = max(area, maxArea)
                stack.pop()
            stack.append(i)
        while len(stack) > 1:
            area = (len(heights)-1-stack[-2]) * heights[stack[-1]]
            maxArea = max(area, maxArea)
            stack.pop()
        return maxArea
