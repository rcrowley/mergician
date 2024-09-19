package html

func Attr(n *Node, k string) string {
	for _, attr := range n.Attr {
		if attr.Namespace == "" && attr.Key == k {
			return attr.Val
		}
	}
	return ""
}
