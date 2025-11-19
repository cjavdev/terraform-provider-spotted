// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_playlist_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-spotted/internal/services/user_playlist"
	"github.com/cjavdev/terraform-provider-spotted/internal/test_helpers"
)

func TestUserPlaylistsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*user_playlist.UserPlaylistsDataSourceModel)(nil)
	schema := user_playlist.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
