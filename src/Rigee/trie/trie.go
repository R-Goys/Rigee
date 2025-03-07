package trie

import (
	"strings"
)

//该部分主要为
//实现前缀树的数据结构的包

type Node struct {
	Pattern  string  //待匹配路由的全程
	Part     string  //路由的当前部分
	Children []*Node //儿子节点
	IsWild   bool    //是否精确匹配
}

// MatchPart 只匹配一个
func (node *Node) MatchPart(part string) *Node {
	for _, child := range node.Children {
		if child.IsWild || child.Part == part {
			return child
		}
	}
	return nil
}

// MatchAll 找到所有符合条件的节点
func (node *Node) MatchAll(part string) []*Node {
	nodes := make([]*Node, 0)
	for _, child := range node.Children {
		if child.IsWild || child.Part == part {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// Insert 插入节点
func (node *Node) Insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		node.Pattern = pattern
		return
	}
	part := parts[height]
	child := node.MatchPart(part)
	if child == nil {
		//由于这里不一定是根节点，所以不加pattern字段
		child = &Node{
			Part:   part,
			IsWild: part[0] == '*' || part[0] == ':',
		}
		//将新创建的节点加入子节点
		node.Children = append(node.Children, child)
	}
	child.Insert(pattern, parts, height+1)
}

func (node *Node) Search(parts []string, height int) *Node {
	if len(parts) == height || strings.HasPrefix(node.Part, "*") {
		if node.Pattern == "" {
			return nil
		}
		return node
	}
	part := parts[height]
	children := node.MatchAll(part)
	for _, child := range children {
		if res := child.Search(parts, height+1); res != nil {
			return res
		}
	}
	return nil
}
