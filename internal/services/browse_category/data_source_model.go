// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browse_category

import (
	"context"

	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

type BrowseCategoryDataSourceModel struct {
	CategoryID types.String                                                     `tfsdk:"category_id" path:"category_id,required"`
	Locale     types.String                                                     `tfsdk:"locale" query:"locale,optional"`
	Href       types.String                                                     `tfsdk:"href" json:"href,computed"`
	ID         types.String                                                     `tfsdk:"id" json:"id,computed"`
	Name       types.String                                                     `tfsdk:"name" json:"name,computed"`
	Icons      customfield.NestedObjectList[BrowseCategoryIconsDataSourceModel] `tfsdk:"icons" json:"icons,computed"`
}

func (m *BrowseCategoryDataSourceModel) toReadParams(_ context.Context) (params spotted.BrowseCategoryGetParams, diags diag.Diagnostics) {
	params = spotted.BrowseCategoryGetParams{}

	if !m.Locale.IsNull() {
		params.Locale = param.NewOpt(m.Locale.ValueString())
	}

	return
}

type BrowseCategoryIconsDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}
