package folder

import (
	"errors"
	"strings"
)

var (
	ErrCantMoveFolderToItself = errors.New("cannot move a folder to itself")
	ErrSrcDoesNotExist        = errors.New("source folder does not exist")
	ErrDestDoesNotExist       = errors.New("destination folder does not exist")
	ErrDiffOrg                = errors.New("cannot move a folder to a different organization")
	ErrCantMoveToChild        = errors.New("cannot move a folder to a child of itself")
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	if name == dst {
		return nil, ErrCantMoveFolderToItself
	}

	// Find the source folder and set it to srcFolder
	var srcFolder *Folder
	for _, folder := range f.folders {
		if folder.Paths == name {
			srcFolder = &folder
			break
		}
	}

	// Check if the source folder exists
	if srcFolder == nil {
		return nil, ErrSrcDoesNotExist
	}

	// Find the destination folder and set it to dstFolder
	var dstFolder *Folder
	for _, folder := range f.folders {
		if folder.Paths == dst {
			dstFolder = &folder
			break
		}
	}

	// Check if the destination folder exists
	if dstFolder == nil {
		return nil, ErrDestDoesNotExist
	}

	// Check if the source and destination folders belong to the same organization
	if srcFolder.OrgId != dstFolder.OrgId {
		return nil, ErrDiffOrg
	}

	// Check if the destination folder is a child of the source folder
	if strings.HasPrefix(dstFolder.Paths, srcFolder.Paths+".") {
		return nil, ErrCantMoveToChild
	}

	// Update the paths of the source folder and its children
	oldPrefix := srcFolder.Paths
	newPrefix := dstFolder.Paths + "." + srcFolder.Name

	for i, folder := range f.folders {
		if strings.HasPrefix(folder.Paths, oldPrefix) {
			f.folders[i].Paths = strings.Replace(folder.Paths, oldPrefix, newPrefix, 1)
		}
	}

	for _, folder := range f.folders {
		println(folder.Paths)
	}

	return f.folders, nil
}
