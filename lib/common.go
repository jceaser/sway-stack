/**********************************************************************************************100*/

package lib

import (
	"fmt"

	"github.com/tlinden/swayipc"
)

/***************************************************************************80*/

var (
	ActiveConnection *swayipc.SwayIPC
)

/***************************************************************************80*/
// MARK - Functions

// FindWorkSpaceNodes recursively searches for and returns all nodes within a specific workspace in
// a swayipc Node tree. It takes a pointer to a swayipc Node representing the root node, and a
// string workspace as input. If the current node is the root, output, or workspace type:
//   - If it's a workspace node with a matching name to the input workspace, it prints the workspace
//     name and the number of nodes under it, then returns those nodes.
//
// If the current node is not the target workspace, it continues searching in its child nodes. The
// function recursively calls itself on each child node until it finds the workspace or exhausts all
// nodes. It returns a slice of swayipc Nodes if the workspace is found, otherwise, it returns nil.
func FindWorkSpaceNodes(node *swayipc.Node, workspace string) []*swayipc.Node {
	if node.Type == "root" || node.Type == "output" || node.Type == "workspace" {
		if node.Type == "workspace" && node.Name == workspace {
			fmt.Printf("workspace '%s' has %d nodes.\n", node.Name, len(node.Nodes))
			return node.Nodes
		}
		//keep looking in the child nodes
		for _, innernode := range node.Nodes {
			result := FindWorkSpaceNodes(innernode, workspace)
			if len(result) > 0 {
				return result
			}
		}
	}
	return nil
}

// CollectWindowsFromWorkspace recursively collects leaf nodes (windows without children) from a
// workspace tree. It takes a slice of windows (nodes) as input and returns a slice of leaf nodes
// found within the workspace.
// Parameters:
//   - windows: a slice of pointers to swayipc.Node representing the windows in the workspace tree.
//
// Returns:
//   - a slice of pointers to swayipc.Node representing the leaf nodes (windows without children) in
//     the workspace tree.
func CollectWindowsFromWorkspace(windows []*swayipc.Node) []*swayipc.Node {
	leafNodes := []*swayipc.Node{}

	for _, win := range windows {
		if len(win.Nodes) == 0 {
			leafNodes = append(leafNodes, win)
		} else {
			leafNodes = append(leafNodes, CollectWindowsFromWorkspace(win.Nodes)...)
		}
	}
	return leafNodes
}

func PrimaryWidth(rect swayipc.Rect, count int) int {
	if count < 1 {
		return rect.Width
	}
	return rect.Width - rect.Width/count
}

func ColumnWidth(rect swayipc.Rect, count int) int {
	if count < 1 {
		return rect.Width
	}
	fmt.Printf("columnWidth: %d/%d = %d\n", rect.Width, count, rect.Width/count)
	return rect.Width / count
}

func ColumnHeight(rect swayipc.Rect, count int) int {
	height := rect.Height - rect.Y
	if count < 1 {
		return height
	}
	return height / count
}
