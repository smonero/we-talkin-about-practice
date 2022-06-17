package main

// 3 possible states:
// node doesnt have a camera and is not being watched - state 0
//

const (
	NO_CAMERA_AND_NOT_BEING_WATCHED int = 0
	NO_CAMERA_AND_BEING_WATCHED         = 1
	HAS_CAMERA                          = 2
)

func getCameraStatus(node *TreeNode, res *int) int {
	if node == nil {
		return NO_CAMERA_AND_BEING_WATCHED
	}

	left := getCameraStatus(node.Left, res)
	right := getCameraStatus(node.Right, res)

	// if left or right is not being watched, install camera here
	if left == NO_CAMERA_AND_NOT_BEING_WATCHED || right == NO_CAMERA_AND_NOT_BEING_WATCHED {
		*res = *res + 1
		return HAS_CAMERA
	}

	if left == HAS_CAMERA || right == HAS_CAMERA {
		return NO_CAMERA_AND_BEING_WATCHED
	}

	return NO_CAMERA_AND_NOT_BEING_WATCHED
}

func minCameraCover(node *TreeNode) int {
	res := 0

	// need to check for extra camera because root has no parent
	// if left and right are both 1, , but root is 0, then add an extra one
	if getCameraStatus(node, &res) == 0 {
		res++
	}

	return res
}
