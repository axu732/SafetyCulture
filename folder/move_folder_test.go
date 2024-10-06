package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		folders []folder.Folder
		src     string
		dst     string
		want    []folder.Folder
		wantErr error
	}{
		{
			name: "move bravo to delta",
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.FromStringOrNil("org2")},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
			},
			src: "alpha.bravo",
			dst: "alpha.delta",
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "bravo", Paths: "alpha.delta.bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "charlie", Paths: "alpha.delta.bravo.charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.FromStringOrNil("org2")},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
			},
			wantErr: nil,
		},
		{
			name: "move bravo to golf",
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.FromStringOrNil("org2")},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
			},
			src: "alpha.bravo",
			dst: "golf",
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "bravo", Paths: "golf.bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "charlie", Paths: "golf.bravo.charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.FromStringOrNil("org2")},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
			},
			wantErr: nil,
		},
		{
			name: "cannot move a folder to child",
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.FromStringOrNil("org2")},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
			},
			src:     "alpha.bravo",
			dst:     "alpha.bravo.charlie",
			want:    nil,
			wantErr: folder.ErrCantMoveToChild,
		},
		{
			name: "cannot move a folder to itself",
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.FromStringOrNil("org2")},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
			},
			src:     "alpha.bravo",
			dst:     "alpha.bravo",
			want:    nil,
			wantErr: folder.ErrCantMoveFolderToItself,
		},
		{
			name: "cannot move a folder to a different organization",
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.FromStringOrNil("org2")},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
			},
			src:     "alpha.bravo",
			dst:     "foxtrot",
			want:    nil,
			wantErr: folder.ErrDiffOrg,
		},
		{
			name: "source folder does not exist",
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.FromStringOrNil("org2")},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
			},
			src:     "invalid_folder",
			dst:     "alpha.delta",
			want:    nil,
			wantErr: folder.ErrSrcDoesNotExist,
		},
		{
			name: "destination folder does not exist",
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.FromStringOrNil("org2")},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
			},
			src:     "alpha.bravo",
			dst:     "invalid_folder",
			want:    nil,
			wantErr: folder.ErrDestDoesNotExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			got, err := f.MoveFolder(tt.src, tt.dst)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
