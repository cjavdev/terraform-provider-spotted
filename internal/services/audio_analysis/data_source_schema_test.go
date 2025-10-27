// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audio_analysis_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/spotted-terraform/internal/services/audio_analysis"
	"github.com/stainless-sdks/spotted-terraform/internal/test_helpers"
)

func TestAudioAnalysisDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*audio_analysis.AudioAnalysisDataSourceModel)(nil)
	schema := audio_analysis.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
