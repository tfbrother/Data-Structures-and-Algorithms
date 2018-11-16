# Data-Structures-and-Algorithms
Algorithms & Data Structures in Go

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

## [一致性hash算法](ConsistentHash/consistenthash.go)
* 所有的节点以及其虚拟节点形成一个圆环，根据节点hash值顺序排列成圆环
* 对请求的key求hash值，找到在圆环中第一个大于该hash值的节点，就让该结点服务


