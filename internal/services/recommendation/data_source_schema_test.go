// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package recommendation_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-spotted/internal/services/recommendation"
	"github.com/cjavdev/terraform-provider-spotted/internal/test_helpers"
)

func TestRecommendationDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*recommendation.RecommendationDataSourceModel)(nil)
	schema := recommendation.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
