package folder

import "github.com/gofrs/uuid"
import "strings"

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res
}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	// Your code here...
	folders := f.GetFoldersByOrgID(orgID)
	res := []Folder{}

	// Check valid orgID
	if !IsValidID(folders, orgID) {
		print("Error: ID does not exist")
		return res
	}

	// Check valid folder name
	if !IsValidString(folders, name) {
		print("Error: Folder does not exist")
		return res
	}
	
	for _, f := range folders {
		// Check if the current path contains the name
		if strings.Contains(f.Paths, name) {
			pathSegments := strings.Split(f.Paths, ".");

			// Check if the name is not the leaf node
			if len(pathSegments) > 1 {
				res = append(res, f)
			}
		}
	}

	return res
}

func IsValidID(folders []Folder, orgID uuid.UUID) bool {
	for _, f := range folders {
		if f.OrgId == orgID {
			return true
		}
	}

	return false
}

func IsValidString(folders []Folder, name string) bool {
	for _, f := range folders {
		if f.Name == name {
			return true
		}
	}

	return false
}
