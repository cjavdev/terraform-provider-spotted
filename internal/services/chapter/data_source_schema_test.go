// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chapter_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/spotted-terraform/internal/services/chapter"
	"github.com/stainless-sdks/spotted-terraform/internal/test_helpers"
)

func TestChapterDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*chapter.ChapterDataSourceModel)(nil)
	schema := chapter.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
