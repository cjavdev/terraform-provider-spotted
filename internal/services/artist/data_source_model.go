// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package artist

import (
	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ArtistDataSourceModel struct {
	ID           types.String                                                `tfsdk:"id" path:"id,required"`
	Href         types.String                                                `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                `tfsdk:"name" json:"name,computed"`
	Popularity   types.Int64                                                 `tfsdk:"popularity" json:"popularity,computed"`
	Type         types.String                                                `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                `tfsdk:"uri" json:"uri,computed"`
	Genres       customfield.List[types.String]                              `tfsdk:"genres" json:"genres,computed"`
	ExternalURLs customfield.NestedObject[ArtistExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Followers    customfield.NestedObject[ArtistFollowersDataSourceModel]    `tfsdk:"followers" json:"followers,computed"`
	Images       customfield.NestedObjectList[ArtistImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
}

type ArtistExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type ArtistFollowersDataSourceModel struct {
	Href  types.String `tfsdk:"href" json:"href,computed"`
	Total types.Int64  `tfsdk:"total" json:"total,computed"`
}

type ArtistImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}
