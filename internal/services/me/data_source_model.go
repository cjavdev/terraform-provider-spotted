// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

type MeDataSourceModel struct {
	Country         types.String                                               `tfsdk:"country" json:"country,computed"`
	DisplayName     types.String                                               `tfsdk:"display_name" json:"display_name,computed"`
	Email           types.String                                               `tfsdk:"email" json:"email,computed"`
	Href            types.String                                               `tfsdk:"href" json:"href,computed"`
	ID              types.String                                               `tfsdk:"id" json:"id,computed"`
	Product         types.String                                               `tfsdk:"product" json:"product,computed"`
	Type            types.String                                               `tfsdk:"type" json:"type,computed"`
	Uri             types.String                                               `tfsdk:"uri" json:"uri,computed"`
	ExplicitContent customfield.NestedObject[MeExplicitContentDataSourceModel] `tfsdk:"explicit_content" json:"explicit_content,computed"`
	ExternalURLs    customfield.NestedObject[MeExternalURLsDataSourceModel]    `tfsdk:"external_urls" json:"external_urls,computed"`
	Followers       customfield.NestedObject[MeFollowersDataSourceModel]       `tfsdk:"followers" json:"followers,computed"`
	Images          customfield.NestedObjectList[MeImagesDataSourceModel]      `tfsdk:"images" json:"images,computed"`
}

type MeExplicitContentDataSourceModel struct {
	FilterEnabled types.Bool `tfsdk:"filter_enabled" json:"filter_enabled,computed"`
	FilterLocked  types.Bool `tfsdk:"filter_locked" json:"filter_locked,computed"`
}

type MeExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MeFollowersDataSourceModel struct {
	Href  types.String `tfsdk:"href" json:"href,computed"`
	Total types.Int64  `tfsdk:"total" json:"total,computed"`
}

type MeImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}
