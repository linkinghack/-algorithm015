## 学习笔记

### 递归模板
1. terminator
2. process
3. drill down
4. reverse states

### Python模板:
```python
def recursion(level, param):
    # terminator
    if (level > MAX_LEVEL):
        return
    
    # process current level
    process(level, param)
    edited(param)

    # drill down
    process(level+1, param)

    # reverse state
    reverseState(param)
```

### 分治&回溯
- 本质就是特殊的递归
- 可以递归解决的问题主要是找重复性
- 最近重复性： 分治，回溯
- 最优重复性： 动态规划



#### 分治递归模板:
1. terminator
2. process:  split big problem
3. drill down: subproblem
4. reverse states

python:
```python
def divide_conquer(problem, param1, param2, ...): 
    # recursion terminator 
    if problem is None: 
        print_result 
        return 

    # prepare data 准备切分数据和问题
    data = prepare_data(problem) 
    subproblems = split_problem(problem, data) 

    # conquer subproblems 分别解决切分的问题
    subresult1 = self.divide_conquer(subproblems[0], p1, ...) 
    subresult2 = self.divide_conquer(subproblems[1], p1, ...) 
    subresult3 = self.divide_conquer(subproblems[2], p1, ...) 
    …
    # process and generate the final result  组合子问题结果
    result = process_result(subresult1, subresult2, subresult3, …)
    # revert the current level states
```


#### 回溯
- 一个普通的递归也可以看成是在回溯
- 在回溯过程中提前发现无解状态，尽早的停止其子状态的递归，提升效率
- 体现为在drill down之前进行更加巧妙的判断，决定是否继续深入下一层