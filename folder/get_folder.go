package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

var (
	ErrInvalidPaths          = errors.New("invalid Paths")
	ErrInvalidOrgID          = errors.New("invalid orgID")
	ErrFolderDoesNotExist    = errors.New("folder does not exist")
	ErrFolderDoesNotExistOrg = errors.New("folder does not exist in the specified organization")
)

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) ([]Folder, error) {

	if orgID == uuid.Nil {
		return nil, ErrInvalidOrgID
	}

	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res, nil

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {

	// Make sure that the orgID is valid
	if orgID == uuid.Nil {
		return nil, ErrInvalidOrgID
	}

	// Get all folders that belong to the orgID
	folders, err := f.GetFoldersByOrgID(orgID)

	// Catch any errors that may have occurred.
	if err != nil {
		return nil, ErrInvalidOrgID
	}

	// Find the parent folder and set it to parentFolder
	var parentFolder *Folder
	for _, folder := range folders {
		if folder.Name == name {
			parentFolder = &folder
			break
		}
	}

	// If the parent folder does not exist, see if it exists in an another organization, else it does not exist
	if parentFolder == nil {
		allFolders := f.folders
		for _, folder := range allFolders {
			if folder.Name == name {
				return nil, ErrFolderDoesNotExistOrg
			}
		}
		return nil, ErrFolderDoesNotExist
	}

	// Find all child folders of the parent folder
	childFolders := []Folder{}
	prefix := parentFolder.Paths + "."

	// Find all child folders of the parent folder
	for _, f := range folders {
		if strings.HasPrefix(f.Paths, prefix) {
			childFolders = append(childFolders, f)
		}
	}

	return childFolders, nil
}
