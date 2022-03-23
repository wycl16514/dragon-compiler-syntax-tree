package simple_parser

type NodeInterface interface {
	AddChild(child NodeInterface)
	GetChildren() []NodeInterface
	Attribute() string
}

type SyntaxNode struct {
	T        string
	children []NodeInterface
}

func NewSyntaxNode() *SyntaxNode {
	return &SyntaxNode{
		T: "",
	}
}

func (s *SyntaxNode) AddChild(node NodeInterface) {
	s.children = append(s.children, node)
}

func (s *SyntaxNode) GetChildren() []NodeInterface {
	return s.children
}

func (s *SyntaxNode) Attribute() string {
	if len(s.children) == 0 {
		return s.T
	}

	attribute := ""
	for _, child := range s.children {
		//根据语义规则，父节点的属性是将所有子节点的属性连接后再添加父节点对应的操作符
		attribute = attribute + child.Attribute()
	}

	attribute += s.T

	return attribute
}
