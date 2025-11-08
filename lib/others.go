package lib

import (
	"fmt"

	"github.com/tlinden/swayipc"
)

func FindWorkSpace() {
	workSpaces, err := ActiveConnection.GetWorkspaces()
	if err != nil {
		Log.Error.Println(err)
	}
	activeName := ""
	activeId := -1
	rect := swayipc.Rect{}
	for _, prospect := range workSpaces {
		if prospect.Focused {
			activeName = prospect.Name
			activeId = prospect.Id
			rect = prospect.Rect
			fmt.Printf("found '%s' [%d] to have focus.\n", activeName, activeId)
			fmt.Printf("dims: %v\n", prospect.Rect)
			break
		}
	}
	root, _ := ActiveConnection.GetTree()
	fmt.Println("---------------------")
	foundWindows := FindWorkSpaceNodes(root, activeName)
	if len(foundWindows) > 0 {
		ArrangeWindows(activeId, foundWindows, rect)
	}
}

/*Move windows around on the current workspace */
func ArrangeWindows(workspaceId int, windows []*swayipc.Node, rect swayipc.Rect) {
	fmt.Printf("arranging windows in current [%d] workspace\n", workspaceId)
	leafWindows := CollectWindowsFromWorkspace(windows)
	leafCount := len(leafWindows)
	fmt.Printf("move windows [%v] (%d) around.\n", leafWindows, leafCount)

	//first remove all windows so that the containers go away
	for _, win := range leafWindows {
		cmdFloat := fmt.Sprintf("[con_id=%d] floating toggle", win.Id)
		//move windows to the top of the workspace no matter where they are by floating them then unfloating
		ActiveConnection.RunGlobalCommand(cmdFloat) // remove from workspace
	}

	currentPrimaryWidth := PrimaryWidth(rect, leafCount-1)
	currentColumnHeight := ColumnHeight(rect, leafCount-1)
	currentColumnWidth := ColumnWidth(rect, leafCount-1)

	// Now put the windows back in one at a time
	for index, win := range leafWindows {
		ActiveConnection.RunGlobalCommand(fmt.Sprintf("[con_id=%d] floating toggle", win.Id)) //put back on workspace
		ActiveConnection.RunGlobalCommand(fmt.Sprintf("[con_id=%d] focus", win.Id))
		if index == 0 {
			//first window
			ActiveConnection.RunGlobalCommand("split horizontal")
		} else {
			//all other windows
			ActiveConnection.RunGlobalCommand("split vertical")
		}
	}

	// now size everything
	for index, win := range leafWindows {
		if index == 0 {
			cmd := fmt.Sprintf("[con_id=%d] resize set width %dpx", win.Id, currentPrimaryWidth)
			ActiveConnection.RunGlobalCommand(cmd)
		} else {
			cmd := fmt.Sprintf("[con_id=%d] resize set width %dpx height %dpx",
				win.Id, currentColumnWidth, currentColumnHeight)
			ActiveConnection.RunGlobalCommand(cmd)
			cmd = fmt.Sprintf("[con_id=%d] resize set height %dpx", win.Id, currentColumnHeight)
			ActiveConnection.RunGlobalCommand(cmd)
		}
		ActiveConnection.RunGlobalCommand(fmt.Sprintf("[con_id=%d] focus", leafWindows[0].Id))
	}
}
