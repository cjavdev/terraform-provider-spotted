// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package playlist_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-spotted/internal/services/playlist"
	"github.com/cjavdev/terraform-provider-spotted/internal/test_helpers"
)

func TestPlaylistDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*playlist.PlaylistDataSourceModel)(nil)
	schema := playlist.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
