# Data-Structures-and-Algorithms
Algorithms & Data Structures in Go

![大纲](https://github.com/tfbrother/Data-Structures-and-Algorithms/blob/master/xmind.jpg?raw=true)

# 数据结构
## 二叉树
### [数组存储](data-structures/binary-tree/array.go)
* 左子节点索引：2*(父结点索引)+1
* 右子节点索引：2*(父结点索引)+2
### 链表存储

## 二叉搜索树
### [链表存储](data-structures/binary-search-tree/bst.go)
* 主要是理解递归算法
* 添加和遍历都是用的递归

## [环形队列](data-structures/queue/ring.go)
* 核心是注意tail可能比head小
* 先入先出（FIFO）

## [堆](data-structures/heap/heap.go)
* 最大堆/最小堆

## [栈](data-structures/stack/stack.go)
* 后入先出（LIFO）
### 应用
* 编辑器-undo操作
* 操作系统-系统调用栈
* 编译器-括号匹配

## 线性表
### [顺序表](data-structures/list/sequence.go)
### 链表
#### 静态链表
#### 单链表
#### 循环链表
#### 双向链表

## 图
### [邻接矩阵存储](data-structures/graph/matrix.go)
* 深度遍历（递归）
* 广度遍历（分层，递归）
* 最小生成树普利姆(Prim)算法
* 最小生成树克鲁斯卡尔(Kruskal)算法

# 算法
## 排序
### [冒泡排序](algorithms/sorting/bubble.go)
### [插入排序](algorithms/sorting/insertion.go)
### [快速排序](algorithms/sorting/quick.go)
### [合并排序](algorithms/sorting/merge.go)
### [选择排序](algorithms/sorting/selection.go)
### [堆排序](algorithms/sorting/heap.go)

## [一致性hash算法](algorithms/consistenthash/consistenthash.go)
* 所有的节点以及其虚拟节点形成一个圆环，根据节点hash值顺序排列成圆环
* 对请求的key求hash值，找到在圆环中第一个大于该hash值的节点，就让该结点服务


