package main

type Node struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type List struct {
	Nodes []Node
}

func (list *List) Add(node Node) {
	list.Nodes = append(list.Nodes, node)

}

func (list *List) getList() []Node {
	return list.Nodes
}

func (list *List) Delete(id string) (string, bool) {
	for index, node := range list.Nodes {
		if node.ID == id {
			list.Nodes = append(list.Nodes[:index], list.Nodes[index+1:]...)
			return "Item deleted", true
		}
	}
	return "Item not found", false
}

func (list *List) getNodeByID(id string) (Node, bool) {
	for _, node := range list.Nodes {
		if node.ID == id {
			return node, true
		}
	}
	return Node{}, false
}
