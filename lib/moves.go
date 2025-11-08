package lib

import "fmt"

func RotateUp() {
	workSpaces, err := ActiveConnection.GetWorkspaces()
	if err != nil {
		Log.Error.Println(err)
	}
	for _, prospect := range workSpaces {
		if prospect.Focused {
			root, err := ActiveConnection.GetTree()
			if err != nil {
				Log.Error.Println(err)
			}
			nodes := FindWorkSpaceNodes(root, prospect.Name)
			windows := CollectWindowsFromWorkspace(nodes)

			for i := len(windows) - 1; i > 1; i-- {
				mcgufin := windows[i]
				other := windows[i-1]
				cmd := fmt.Sprintf("[con_id=%d] focus ; swap container with con_id %d", mcgufin.Id, other.Id)
				ActiveConnection.RunGlobalCommand(cmd)
				Log.Debug.Printf("mv %d->%d\n", mcgufin.Id, other.Id)
			}
			break
		}
	}
}

func RotateDown() {
	workSpaces, err := ActiveConnection.GetWorkspaces()
	if err != nil {
		Log.Error.Println(err)
	}
	for _, prospect := range workSpaces {
		if prospect.Focused {
			root, err := ActiveConnection.GetTree()
			if err != nil {
				Log.Error.Println(err)
			}
			nodes := FindWorkSpaceNodes(root, prospect.Name)
			windows := CollectWindowsFromWorkspace(nodes)

			for i := 1; i < len(windows)-1; i++ {
				mcgufin := windows[i]
				other := windows[i+1]
				cmd := fmt.Sprintf("[con_id=%d] focus ; swap container with con_id %d", mcgufin.Id, other.Id)
				ActiveConnection.RunGlobalCommand(cmd)
				Log.Debug.Printf("mv %d->%d\n", mcgufin.Id, other.Id)
			}
			break
		}
	}
}

/* swap the active with the top */
func SwapLeft() {
	workSpaces, err := ActiveConnection.GetWorkspaces()
	if err != nil {
		Log.Error.Println(err)
	}
	for _, prospect := range workSpaces {
		if prospect.Focused {
			root, err := ActiveConnection.GetTree()
			if err != nil {
				Log.Error.Println(err)
			}
			nodes := FindWorkSpaceNodes(root, prospect.Name)
			windows := CollectWindowsFromWorkspace(nodes)
			if len(windows) > 1 {
				mcgufin := windows[0]
				other := windows[1]
				cmd := fmt.Sprintf("[con_id=%d] focus ; swap container with con_id %d", mcgufin.Id, other.Id)
				ActiveConnection.RunGlobalCommand(cmd)
				Log.Debug.Printf("mv %d->%d\n", mcgufin.Id, other.Id)
			}
			break
		}
	}
}

/* swap the active with the bottom */
func SwapRight() {
	workSpaces, err := ActiveConnection.GetWorkspaces()
	if err != nil {
		Log.Error.Println(err)
	}
	for _, prospect := range workSpaces {
		if prospect.Focused {
			root, err := ActiveConnection.GetTree()
			if err != nil {
				Log.Error.Println(err)
			}
			nodes := FindWorkSpaceNodes(root, prospect.Name)
			windows := CollectWindowsFromWorkspace(nodes)

			if len(windows) > 1 {
				mcgufin := windows[0]
				other := windows[len(windows)-1]
				cmd := fmt.Sprintf("[con_id=%d] focus ; swap container with con_id %d", mcgufin.Id, other.Id)
				ActiveConnection.RunGlobalCommand(cmd)
				Log.Debug.Printf("mv %d->%d\n", mcgufin.Id, other.Id)
			}
			break
		}
	}
}
