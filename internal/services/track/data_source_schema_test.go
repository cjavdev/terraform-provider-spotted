// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package track_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-spotted/internal/services/track"
	"github.com/cjavdev/terraform-provider-spotted/internal/test_helpers"
)

func TestTrackDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*track.TrackDataSourceModel)(nil)
	schema := track.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
