// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package artist_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/spotted-terraform/internal/services/artist"
	"github.com/stainless-sdks/spotted-terraform/internal/test_helpers"
)

func TestArtistDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*artist.ArtistDataSourceModel)(nil)
	schema := artist.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
