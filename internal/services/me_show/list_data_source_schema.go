// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me_show

import (
	"context"

	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*MeShowsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"limit": schema.Int64Attribute{
				Description: "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 50),
				},
			},
			"offset": schema.Int64Attribute{
				Description: "The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.",
				Computed:    true,
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[MeShowsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"added_at": schema.StringAttribute{
							Description: "The date and time the show was saved.\nTimestamps are returned in ISO 8601 format as Coordinated Universal Time (UTC) with a zero offset: YYYY-MM-DDTHH:MM:SSZ.\nIf the time is imprecise (for example, the date/time of an album release), an additional field indicates the precision; see for example, release_date in an album object.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"show": schema.SingleNestedAttribute{
							Description: "Information about the show.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[MeShowsShowDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the show.",
									Computed:    true,
								},
								"available_markets": schema.ListAttribute{
									Description: "A list of the countries in which the show can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"copyrights": schema.ListNestedAttribute{
									Description: "The copyright statements of the show.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[MeShowsShowCopyrightsDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"text": schema.StringAttribute{
												Description: "The copyright text for this content.",
												Computed:    true,
											},
											"type": schema.StringAttribute{
												Description: "The type of copyright: `C` = the copyright, `P` = the sound recording (performance) copyright.",
												Computed:    true,
											},
										},
									},
								},
								"description": schema.StringAttribute{
									Description: "A description of the show. HTML tags are stripped away from this field, use `html_description` field in case HTML tags are needed.",
									Computed:    true,
								},
								"explicit": schema.BoolAttribute{
									Description: "Whether or not the show has explicit content (true = yes it does; false = no it does not OR unknown).",
									Computed:    true,
								},
								"external_urls": schema.SingleNestedAttribute{
									Description: "External URLs for this show.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[MeShowsShowExternalURLsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"spotify": schema.StringAttribute{
											Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
											Computed:    true,
										},
									},
								},
								"href": schema.StringAttribute{
									Description: "A link to the Web API endpoint providing full details of the show.",
									Computed:    true,
								},
								"html_description": schema.StringAttribute{
									Description: "A description of the show. This field may contain HTML tags.",
									Computed:    true,
								},
								"images": schema.ListNestedAttribute{
									Description: "The cover art for the show in various sizes, widest first.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[MeShowsShowImagesDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"height": schema.Int64Attribute{
												Description: "The image height in pixels.",
												Computed:    true,
											},
											"url": schema.StringAttribute{
												Description: "The source URL of the image.",
												Computed:    true,
											},
											"width": schema.Int64Attribute{
												Description: "The image width in pixels.",
												Computed:    true,
											},
										},
									},
								},
								"is_externally_hosted": schema.BoolAttribute{
									Description: "True if all of the shows episodes are hosted outside of Spotify's CDN. This field might be `null` in some cases.",
									Computed:    true,
								},
								"languages": schema.ListAttribute{
									Description: "A list of the languages used in the show, identified by their [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"media_type": schema.StringAttribute{
									Description: "The media type of the show.",
									Computed:    true,
								},
								"name": schema.StringAttribute{
									Description: "The name of the episode.",
									Computed:    true,
								},
								"publisher": schema.StringAttribute{
									Description: "The publisher of the show.",
									Computed:    true,
								},
								"total_episodes": schema.Int64Attribute{
									Description: "The total number of episodes in the show.",
									Computed:    true,
								},
								"type": schema.StringAttribute{
									Description: "The object type.\nAvailable values: \"show\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("show"),
									},
								},
								"uri": schema.StringAttribute{
									Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the show.",
									Computed:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *MeShowsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *MeShowsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
