// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package show

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*ShowDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids)\nfor the show.",
				Required:    true,
			},
			"market": schema.StringAttribute{
				Description: "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the show. HTML tags are stripped away from this field, use `html_description` field in case HTML tags are needed.",
				Computed:    true,
			},
			"explicit": schema.BoolAttribute{
				Description: "Whether or not the show has explicit content (true = yes it does; false = no it does not OR unknown).",
				Computed:    true,
			},
			"href": schema.StringAttribute{
				Description: "A link to the Web API endpoint providing full details of the show.",
				Computed:    true,
			},
			"html_description": schema.StringAttribute{
				Description: "A description of the show. This field may contain HTML tags.",
				Computed:    true,
			},
			"is_externally_hosted": schema.BoolAttribute{
				Description: "True if all of the shows episodes are hosted outside of Spotify's CDN. This field might be `null` in some cases.",
				Computed:    true,
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
			"available_markets": schema.ListAttribute{
				Description: "A list of the countries in which the show can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"languages": schema.ListAttribute{
				Description: "A list of the languages used in the show, identified by their [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"copyrights": schema.ListNestedAttribute{
				Description: "The copyright statements of the show.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ShowCopyrightsDataSourceModel](ctx),
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
			"episodes": schema.SingleNestedAttribute{
				Description: "The episodes of the show.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ShowEpisodesDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"href": schema.StringAttribute{
						Description: "A link to the Web API endpoint returning the full result of the request",
						Computed:    true,
					},
					"items": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[ShowEpisodesItemsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the episode.",
									Computed:    true,
								},
								"audio_preview_url": schema.StringAttribute{
									Description:        "A URL to a 30 second preview (MP3 format) of the episode. `null` if not available.",
									Computed:           true,
									DeprecationMessage: "This attribute is deprecated.",
								},
								"description": schema.StringAttribute{
									Description: "A description of the episode. HTML tags are stripped away from this field, use `html_description` field in case HTML tags are needed.",
									Computed:    true,
								},
								"duration_ms": schema.Int64Attribute{
									Description: "The episode length in milliseconds.",
									Computed:    true,
								},
								"explicit": schema.BoolAttribute{
									Description: "Whether or not the episode has explicit content (true = yes it does; false = no it does not OR unknown).",
									Computed:    true,
								},
								"external_urls": schema.SingleNestedAttribute{
									Description: "External URLs for this episode.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ShowEpisodesItemsExternalURLsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"spotify": schema.StringAttribute{
											Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
											Computed:    true,
										},
									},
								},
								"href": schema.StringAttribute{
									Description: "A link to the Web API endpoint providing full details of the episode.",
									Computed:    true,
								},
								"html_description": schema.StringAttribute{
									Description: "A description of the episode. This field may contain HTML tags.",
									Computed:    true,
								},
								"images": schema.ListNestedAttribute{
									Description: "The cover art for the episode in various sizes, widest first.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[ShowEpisodesItemsImagesDataSourceModel](ctx),
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
									Description: "True if the episode is hosted outside of Spotify's CDN.",
									Computed:    true,
								},
								"is_playable": schema.BoolAttribute{
									Description: "True if the episode is playable in the given market. Otherwise false.",
									Computed:    true,
								},
								"languages": schema.ListAttribute{
									Description: "A list of the languages used in the episode, identified by their [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639) code.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"name": schema.StringAttribute{
									Description: "The name of the episode.",
									Computed:    true,
								},
								"release_date": schema.StringAttribute{
									Description: "The date the episode was first released, for example `\"1981-12-15\"`. Depending on the precision, it might be shown as `\"1981\"` or `\"1981-12\"`.",
									Computed:    true,
								},
								"release_date_precision": schema.StringAttribute{
									Description: "The precision with which `release_date` value is known.\nAvailable values: \"year\", \"month\", \"day\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"year",
											"month",
											"day",
										),
									},
								},
								"type": schema.StringAttribute{
									Description: "The object type.\nAvailable values: \"episode\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("episode"),
									},
								},
								"uri": schema.StringAttribute{
									Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the episode.",
									Computed:    true,
								},
								"language": schema.StringAttribute{
									Description:        "The language used in the episode, identified by a [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code. This field is deprecated and might be removed in the future. Please use the `languages` field instead.",
									Computed:           true,
									DeprecationMessage: "This attribute is deprecated.",
								},
								"restrictions": schema.SingleNestedAttribute{
									Description: "Included in the response when a content restriction is applied.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ShowEpisodesItemsRestrictionsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"reason": schema.StringAttribute{
											Description: "The reason for the restriction. Supported values:\n- `market` - The content item is not available in the given market.\n- `product` - The content item is not available for the user's subscription type.\n- `explicit` - The content item is explicit and the user's account is set to not play explicit content.\n\nAdditional reasons may be added in the future.\n**Note**: If you use this field, make sure that your application safely handles unknown values.",
											Computed:    true,
										},
									},
								},
								"resume_point": schema.SingleNestedAttribute{
									Description: "The user's most recent position in the episode. Set if the supplied access token is a user token and has the scope 'user-read-playback-position'.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ShowEpisodesItemsResumePointDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"fully_played": schema.BoolAttribute{
											Description: "Whether or not the episode has been fully played by the user.",
											Computed:    true,
										},
										"resume_position_ms": schema.Int64Attribute{
											Description: "The user's most recent position in the episode in milliseconds.",
											Computed:    true,
										},
									},
								},
							},
						},
					},
					"limit": schema.Int64Attribute{
						Description: "The maximum number of items in the response (as set in the query or by default).",
						Computed:    true,
					},
					"next": schema.StringAttribute{
						Description: "URL to the next page of items. ( `null` if none)",
						Computed:    true,
					},
					"offset": schema.Int64Attribute{
						Description: "The offset of the items returned (as set in the query or by default)",
						Computed:    true,
					},
					"previous": schema.StringAttribute{
						Description: "URL to the previous page of items. ( `null` if none)",
						Computed:    true,
					},
					"total": schema.Int64Attribute{
						Description: "The total number of items available to return.",
						Computed:    true,
					},
				},
			},
			"external_urls": schema.SingleNestedAttribute{
				Description: "External URLs for this show.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ShowExternalURLsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"spotify": schema.StringAttribute{
						Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
						Computed:    true,
					},
				},
			},
			"images": schema.ListNestedAttribute{
				Description: "The cover art for the show in various sizes, widest first.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ShowImagesDataSourceModel](ctx),
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
		},
	}
}

func (d *ShowDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ShowDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
