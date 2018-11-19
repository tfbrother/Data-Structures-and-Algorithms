# Data-Structures-and-Algorithms
Algorithms & Data Structures in Go

![大纲](https://github.com/tfbrother/Data-Structures-and-Algorithms/blob/master/xmind.jpg?raw=true)

# 数据结构
## 二叉树
### [数组存储](BinaryTree/array.go)
* 左子节点索引：2*(父结点索引)+1
* 右子节点索引：2*(父结点索引)+2
### 链表存储

## [环形队列](Queue/ring.go)
* 核心是注意tail可能比head小
* 先入先出（FIFO）

## [栈](Stack/stack.go)
* 后入先出（LIFO）

## 线性表
### [顺序表](List/sequence.go)
### 链表
#### 静态链表
#### 单链表
#### 循环链表
#### 双向链表

## 图
### [邻接矩阵存储](Map/matrix.go)
* 深度遍历（递归）
* 广度遍历（分层，递归）
* 最小生成树普利姆(Prim)算法
* 最小生成树克鲁斯卡尔(Kruskal)算法

# 算法
## 排序
### 冒泡排序
### 插入排序
### 快速排序
### 合并排序
### 选择排序
### 堆排序

## [一致性hash算法](consistenthash/consistenthash.go)
* 所有的节点以及其虚拟节点形成一个圆环，根据节点hash值顺序排列成圆环
* 对请求的key求hash值，找到在圆环中第一个大于该hash值的节点，就让该结点服务


