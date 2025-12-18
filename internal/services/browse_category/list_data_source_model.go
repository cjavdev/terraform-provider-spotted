// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browse_category

import (
	"context"

	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BrowseCategoriesItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[BrowseCategoriesItemsDataSourceModel] `json:"items,computed"`
}

type BrowseCategoriesDataSourceModel struct {
	Locale   types.String                                                       `tfsdk:"locale" query:"locale,optional"`
	Limit    types.Int64                                                        `tfsdk:"limit" query:"limit,computed_optional"`
	Offset   types.Int64                                                        `tfsdk:"offset" query:"offset,computed_optional"`
	MaxItems types.Int64                                                        `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[BrowseCategoriesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *BrowseCategoriesDataSourceModel) toListParams(_ context.Context) (params spotted.BrowseCategoryListParams, diags diag.Diagnostics) {
	params = spotted.BrowseCategoryListParams{}

	if !m.Limit.IsNull() {
		params.Limit = param.NewOpt(m.Limit.ValueInt64())
	}
	if !m.Locale.IsNull() {
		params.Locale = param.NewOpt(m.Locale.ValueString())
	}
	if !m.Offset.IsNull() {
		params.Offset = param.NewOpt(m.Offset.ValueInt64())
	}

	return
}

type BrowseCategoriesItemsDataSourceModel struct {
	ID        types.String                                                       `tfsdk:"id" json:"id,computed"`
	Href      types.String                                                       `tfsdk:"href" json:"href,computed"`
	Icons     customfield.NestedObjectList[BrowseCategoriesIconsDataSourceModel] `tfsdk:"icons" json:"icons,computed"`
	Name      types.String                                                       `tfsdk:"name" json:"name,computed"`
	Published types.Bool                                                         `tfsdk:"published" json:"published,computed"`
}

type BrowseCategoriesIconsDataSourceModel struct {
	Height    types.Int64  `tfsdk:"height" json:"height,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	Width     types.Int64  `tfsdk:"width" json:"width,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
}
