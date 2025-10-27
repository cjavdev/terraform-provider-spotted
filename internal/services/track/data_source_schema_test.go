// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package track_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/spotted-terraform/internal/services/track"
	"github.com/stainless-sdks/spotted-terraform/internal/test_helpers"
)

func TestTrackDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*track.TrackDataSourceModel)(nil)
	schema := track.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
