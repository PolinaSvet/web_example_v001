package treeCustom

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Key         int
	Left, Right *Node
	Height      int
}

var initRootStr string

const noKey int = 123456789

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func updateHeight(node *Node) {
	node.Height = 1 + max(height(node.Left), height(node.Right))
}

func rotateRight(y *Node) *Node {
	x := y.Left
	t := x.Right

	x.Right = y
	y.Left = t

	updateHeight(y)
	updateHeight(x)

	return x
}

func rotateLeft(x *Node) *Node {
	y := x.Right
	t := y.Left

	y.Left = x
	x.Right = t

	updateHeight(x)
	updateHeight(y)

	return y
}

func getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

func insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{Key: key, Left: nil, Right: nil, Height: 1}
	}

	if key < root.Key {
		root.Left = insert(root.Left, key)
	} else if key > root.Key {
		root.Right = insert(root.Right, key)
	} else {
		return root
	}

	updateHeight(root)

	balance := getBalance(root)

	// Left Left Case
	if balance > 1 && key < root.Left.Key {
		return rotateRight(root)
	}
	// Right Right Case
	if balance < -1 && key > root.Right.Key {
		return rotateLeft(root)
	}
	// Left Right Case
	if balance > 1 && key > root.Left.Key {
		root.Left = rotateLeft(root.Left)
		return rotateRight(root)
	}
	// Right Left Case
	if balance < -1 && key < root.Right.Key {
		root.Right = rotateRight(root.Right)
		return rotateLeft(root)
	}

	return root
}

func minValueNode(node *Node) *Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func deleteNode(root *Node, key int) *Node {
	if root == nil {
		return root
	}

	if key < root.Key {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil || root.Right == nil {
			var temp *Node
			if root.Left == nil {
				temp = root.Right
			} else {
				temp = root.Left
			}

			if temp == nil {
				temp = root
				root = nil
			} else {
				*root = *temp
			}
		} else {
			temp := minValueNode(root.Right)

			root.Key = temp.Key

			root.Right = deleteNode(root.Right, temp.Key)
		}
	}

	if root == nil {
		return root
	}

	updateHeight(root)

	balance := getBalance(root)

	// Left Left Case
	if balance > 1 && getBalance(root.Left) >= 0 {
		return rotateRight(root)
	}
	// Left Right Case
	if balance > 1 && getBalance(root.Left) < 0 {
		root.Left = rotateLeft(root.Left)
		return rotateRight(root)
	}
	// Right Right Case
	if balance < -1 && getBalance(root.Right) <= 0 {
		return rotateLeft(root)
	}
	// Right Left Case
	if balance < -1 && getBalance(root.Right) > 0 {
		root.Right = rotateRight(root.Right)
		return rotateLeft(root)
	}

	return root
}

func search(root *Node, key int) *Node {
	if root == nil || root.Key == key {
		return root
	}

	if root.Key < key {
		return search(root.Right, key)
	}

	return search(root.Left, key)
}

func printInOrder(root *Node) {
	if root != nil {
		printInOrder(root.Left)
		fmt.Printf("%d ", root.Key)
		printInOrder(root.Right)
	}
}

func printInOrderStr(root *Node, result *string) {
	if root != nil {
		printInOrderStr(root.Left, result)
		*result += fmt.Sprintf("%d ", root.Key)
		printInOrderStr(root.Right, result)
	}
}

func NodeSVG(node *Node, x, y float64, dx float64, nodeStr string, key int, mess string) string {
	if node == nil {
		return ""
	}

	circleColor := "LightSkyBlue"
	if key == node.Key {
		circleColor = "Violet"
	}

	renderedNode := fmt.Sprintf(`<text x="10" y="10" text-anchor="left" fill="black">Исходные данные: %v</text>`, nodeStr)
	renderedNode += fmt.Sprintf(`<text x="10" y="30" text-anchor="left" fill="black">%v</text>`, mess)
	renderedNode += fmt.Sprintf(`<circle cx="%v" cy="%v" r="20" fill="%v" stroke="black" stroke-width="1" /><text x="%v" y="%v" text-anchor="middle" fill="black">%v</text>`, x, y, circleColor, x, y, node.Key)

	if node.Left != nil {
		leftX := x - dx
		leftY := y + 50
		renderedNode += fmt.Sprintf(`<line x1="%v" y1="%v" x2="%v" y2="%v" stroke="black" stroke-width="1" />`, x, y+20, leftX, leftY-20)
		renderedNode += NodeSVG(node.Left, leftX, leftY, dx/2, "", key, "")
	}

	if node.Right != nil {
		rightX := x + dx
		rightY := y + 50
		renderedNode += fmt.Sprintf(`<line x1="%v" y1="%v" x2="%v" y2="%v" stroke="black" stroke-width="1" />`, x, y+20, rightX, rightY-20)
		renderedNode += NodeSVG(node.Right, rightX, rightY, dx/2, "", key, "")
	}

	return renderedNode
}

// RenderTreeSVG возвращает SVG-представление всего дерева
func RenderTreeSVG(root *Node, rootStr string, key int, mess string) (string, error) {
	nodes := NodeSVG(root, 700, 100, 300, rootStr, key, mess)
	output := ""
	for _, node := range nodes {
		output += string(node)
	}

	return output, nil
}

func createRandomBinaryTree(numElements int) (*Node, string) {
	var root *Node
	rand.Seed(time.Now().Unix())
	rootStr := ""
	for i := 0; i < numElements; i++ {
		value := rand.Intn(900)
		rootStr += fmt.Sprintf(`%v `, value)
		root = insert(root, value)
	}

	return root, rootStr
}

func CreateElementRenderTreeSVG(numElements int) (*Node, string) {

	var rootN *Node
	var result string

	rootN, initRootStr = createRandomBinaryTree(numElements)

	printInOrderStr(rootN, &result)
	//fmt.Println(result)
	output, _ := RenderTreeSVG(rootN, initRootStr, noKey, fmt.Sprintf(`Новое дерево: %v`, result))

	return rootN, output
}

func InsertElementRenderTreeSVG(root *Node, key int) (*Node, string) {

	rootN := insert(root, key)
	output, _ := RenderTreeSVG(rootN, initRootStr, key, fmt.Sprintf(`Добавлен элемент: %v`, key))

	return rootN, output
}

func DeleteElementRenderTreeSVG(root *Node, key int) (*Node, string) {

	rootN := deleteNode(root, key)
	output, _ := RenderTreeSVG(rootN, initRootStr, noKey, fmt.Sprintf(`Удален элемент: %v,`, key))

	return rootN, output
}

func FindElementRenderTreeSVG(root *Node, key int) (*Node, string) {

	output := ""
	result := search(root, key)
	if result != nil {
		output, _ = RenderTreeSVG(root, initRootStr, result.Key, fmt.Sprintf(`Найден элемент: %v`, key))
	} else {
		output, _ = RenderTreeSVG(root, initRootStr, noKey, fmt.Sprintf(`Ненайден элемент: %v`, key))
	}

	return root, output
}

func PrintElementRenderTreeSVG(root *Node) (*Node, string) {

	var result string

	printInOrderStr(root, &result)
	//fmt.Println(result)
	output, _ := RenderTreeSVG(root, initRootStr, noKey, fmt.Sprintf(`Текущее дерево: %v`, result))

	return root, output
}
