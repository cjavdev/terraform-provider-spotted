// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package episode_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-spotted/internal/services/episode"
	"github.com/cjavdev/terraform-provider-spotted/internal/test_helpers"
)

func TestEpisodeDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*episode.EpisodeDataSourceModel)(nil)
	schema := episode.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
