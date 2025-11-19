// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audiobook_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-spotted/internal/services/audiobook"
	"github.com/cjavdev/terraform-provider-spotted/internal/test_helpers"
)

func TestAudiobookDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*audiobook.AudiobookDataSourceModel)(nil)
	schema := audiobook.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
