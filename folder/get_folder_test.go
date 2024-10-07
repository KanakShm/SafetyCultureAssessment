package folder_test

import (
	"reflect"
	"testing"
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	
	// Sample orgIDs
	orgID1 := uuid.Must(uuid.NewV4())
	orgID2 := uuid.Must(uuid.NewV4())

	tests := [...]struct {
		name    string
		parentName string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:    "Empty folders list",
			orgID:   orgID1,
			folders: []folder.Folder{},
			want:    []folder.Folder{},
		},
		{
			name:  "Single folder matching orgID1",
			orgID: orgID1,
			folders: []folder.Folder{
				{OrgId: orgID1, Paths: "folder1"},
			},
			want: []folder.Folder{
				{OrgId: orgID1, Paths: "folder1"},
			},
		},
		{
			name:  "Multiple folders with some matching orgID1",
			orgID: orgID1,
			folders: []folder.Folder{
				{OrgId: orgID1, Paths: "folder1"},
				{OrgId: orgID2, Paths: "folder2"},
				{OrgId: orgID1, Paths: "folder3"},
			},
			want: []folder.Folder{
				{OrgId: orgID1, Paths: "folder1"},
				{OrgId: orgID1, Paths: "folder3"},
			},
		},
		{
			name:  "No folders match orgID2",
			orgID: orgID2,
			folders: []folder.Folder{
				{OrgId: orgID1, Paths: "folder1"},
				{OrgId: orgID1, Paths: "folder2"},
			},
			want: []folder.Folder{},
		},
		{
			name:  "All folders match orgID2",
			orgID: orgID2,
			folders: []folder.Folder{
				{OrgId: orgID2, Paths: "folder1"},
				{OrgId: orgID2, Paths: "folder2"},
			},
			want: []folder.Folder{
				{OrgId: orgID2, Paths: "folder1"},
				{OrgId: orgID2, Paths: "folder2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)

			if !reflect.DeepEqual(get, tt.want) {
				t.Errorf("GetFoldersByOrgID() = %v, want %v", get, tt.want)
			}
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
    t.Parallel()

    // Generate mock UUIDs
    orgID1 := uuid.Must(uuid.NewV4())
    orgID2 := uuid.Must(uuid.NewV4())

	// Sample folders
    folders := []folder.Folder{
		{OrgId: orgID1, Paths: "parent1.child1"},
        {OrgId: orgID1, Paths: "parent1.child2"},
        {OrgId: orgID1, Paths: "parent1.child2.subChild1"},
        {OrgId: orgID1, Paths: "parent2"},
        {OrgId: orgID2, Paths: "otherParent1.child1"},
        {OrgId: orgID1, Paths: "parent1"},
	}

    tests := [...]struct {
        name       string
        orgID      uuid.UUID
        parentName string
        folders    []folder.Folder
        want       []folder.Folder
    }{
        {
            name:       "Get all child folders for orgID1 with parent1",
            orgID:      orgID1,
            parentName: "parent1",
            folders:    folders,
            want: []folder.Folder{
                {OrgId: orgID1, Paths: "parent1.child1"},
                {OrgId: orgID1, Paths: "parent1.child2"},
                {OrgId: orgID1, Paths: "parent1.child2.subChild1"},
            },
        },
        {
            name:       "Get all child folders for orgID2",
            orgID:      orgID2,
            parentName: "otherParent1",
            folders:    folders,
            want: []folder.Folder{
                {OrgId: orgID2, Paths: "otherParent1.child1"},
            },
        },
        {
            name:       "Get child folders for a parent with no children (parent2)",
            orgID:      orgID1,
            parentName: "parent2",
            folders:    folders,
            want:       []folder.Folder{},
        },
        {
            name:       "Get all child folders for a non-existent orgID",
            orgID:      uuid.Must(uuid.NewV4()),
            parentName: "parent1",
            folders:    folders,
            want:       []folder.Folder{},
        },
        {
            name:       "Get all child folders where parent has only itself (parent1)",
            orgID:      orgID1,
            parentName: "parent1",
            folders:    folders,
            want: []folder.Folder{
                {OrgId: orgID1, Paths: "parent1.child1"},
                {OrgId: orgID1, Paths: "parent1.child2"},
                {OrgId: orgID1, Paths: "parent1.child2.subChild1"},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
            got := f.GetAllChildFolders(tt.orgID, tt.parentName)

            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetAllChildFolders() = %v, want %v", got, tt.want)
            }
        })
    }
}