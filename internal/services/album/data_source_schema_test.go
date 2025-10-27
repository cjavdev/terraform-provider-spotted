// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package album_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/spotted-terraform/internal/services/album"
	"github.com/stainless-sdks/spotted-terraform/internal/test_helpers"
)

func TestAlbumDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*album.AlbumDataSourceModel)(nil)
	schema := album.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
