package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
		wantErr error
	}{
		{
			name:    "valid orgID",
			orgID:   uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{{OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)}},
			want:    []folder.Folder{{OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)}},
			wantErr: nil,
		},

		{
			name:    "invalid orgID",
			orgID:   uuid.Nil,
			folders: []folder.Folder{},
			want:    nil,
			wantErr: folder.ErrInvalidOrgID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, get)

		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		parent  string
		want    []folder.Folder
		wantErr error
	}{
		{
			name:  "Valid parent folder alpha",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.delta"},
				{Name: "echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "echo"},
				{Name: "foxtrot", OrgId: uuid.FromStringOrNil("org2"), Paths: "foxtrot"},
			},
			parent: "alpha",
			want: []folder.Folder{
				{Name: "bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.delta"},
			},
			wantErr: nil,
		},
		{
			name:  "Valid parent folder bravo",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.delta"},
				{Name: "echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "echo"},
				{Name: "foxtrot", OrgId: uuid.FromStringOrNil("org2"), Paths: "foxtrot"},
			},
			parent: "bravo",
			want: []folder.Folder{
				{Name: "charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.bravo.charlie"},
			},
			wantErr: nil,
		},
		{
			name:  "Valid parent folder charlie",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.delta"},
				{Name: "echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "echo"},
				{Name: "foxtrot", OrgId: uuid.FromStringOrNil("org2"), Paths: "foxtrot"},
			},
			parent:  "charlie",
			want:    []folder.Folder{},
			wantErr: nil,
		},
		{
			name:  "Valid parent folder echo",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.delta"},
				{Name: "echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "echo"},
				{Name: "foxtrot", OrgId: uuid.FromStringOrNil("org2"), Paths: "foxtrot"},
			},
			parent:  "echo",
			want:    []folder.Folder{},
			wantErr: nil,
		},
		{
			name:  "Invalid parent folder",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.delta"},
				{Name: "echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "echo"},
				{Name: "foxtrot", OrgId: uuid.FromStringOrNil("org2"), Paths: "foxtrot"},
			},
			parent:  "invalid_folder",
			want:    nil,
			wantErr: folder.ErrFolderDoesNotExist,
		},
		{
			name:  "Folder does not exist in the specified organization",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.delta"},
				{Name: "echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "echo"},
				{Name: "foxtrot", OrgId: uuid.FromStringOrNil("org2"), Paths: "foxtrot"},
			},
			parent:  "foxtrot",
			want:    nil,
			wantErr: folder.ErrFolderDoesNotExistOrg,
		},
		{
			name:  "Nil orgID",
			orgID: uuid.Nil,
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.delta"},
				{Name: "echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "echo"},
				{Name: "foxtrot", OrgId: uuid.FromStringOrNil("org2"), Paths: "foxtrot"},
			},
			parent:  "alpha",
			want:    nil,
			wantErr: folder.ErrInvalidOrgID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.GetAllChildFolders(tt.orgID, tt.parent)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, get)
		})
	}
}
