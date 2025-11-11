// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me_audiobook

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*MeAudiobooksDataSource)(nil)

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
				CustomType:  customfield.NewNestedObjectListType[MeAudiobooksItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"added_at": schema.StringAttribute{
							Description: "The date and time the audiobook was saved\nTimestamps are returned in ISO 8601 format as Coordinated Universal Time (UTC) with a zero offset: YYYY-MM-DDTHH:MM:SSZ.\nIf the time is imprecise (for example, the date/time of an album release), an additional field indicates the precision; see for example, release_date in an album object.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"audiobook": schema.SingleNestedAttribute{
							Description: "Information about the audiobook.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[MeAudiobooksAudiobookDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the audiobook.",
									Computed:    true,
								},
								"authors": schema.ListNestedAttribute{
									Description: "The author(s) for the audiobook.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[MeAudiobooksAudiobookAuthorsDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "The name of the author.",
												Computed:    true,
											},
										},
									},
								},
								"available_markets": schema.ListAttribute{
									Description: "A list of the countries in which the audiobook can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"copyrights": schema.ListNestedAttribute{
									Description: "The copyright statements of the audiobook.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[MeAudiobooksAudiobookCopyrightsDataSourceModel](ctx),
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
									Description: "A description of the audiobook. HTML tags are stripped away from this field, use `html_description` field in case HTML tags are needed.",
									Computed:    true,
								},
								"explicit": schema.BoolAttribute{
									Description: "Whether or not the audiobook has explicit content (true = yes it does; false = no it does not OR unknown).",
									Computed:    true,
								},
								"external_urls": schema.SingleNestedAttribute{
									Description: "External URLs for this audiobook.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[MeAudiobooksAudiobookExternalURLsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"spotify": schema.StringAttribute{
											Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
											Computed:    true,
										},
									},
								},
								"href": schema.StringAttribute{
									Description: "A link to the Web API endpoint providing full details of the audiobook.",
									Computed:    true,
								},
								"html_description": schema.StringAttribute{
									Description: "A description of the audiobook. This field may contain HTML tags.",
									Computed:    true,
								},
								"images": schema.ListNestedAttribute{
									Description: "The cover art for the audiobook in various sizes, widest first.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[MeAudiobooksAudiobookImagesDataSourceModel](ctx),
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
								"languages": schema.ListAttribute{
									Description: "A list of the languages used in the audiobook, identified by their [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"media_type": schema.StringAttribute{
									Description: "The media type of the audiobook.",
									Computed:    true,
								},
								"name": schema.StringAttribute{
									Description: "The name of the audiobook.",
									Computed:    true,
								},
								"narrators": schema.ListNestedAttribute{
									Description: "The narrator(s) for the audiobook.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[MeAudiobooksAudiobookNarratorsDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "The name of the Narrator.",
												Computed:    true,
											},
										},
									},
								},
								"publisher": schema.StringAttribute{
									Description: "The publisher of the audiobook.",
									Computed:    true,
								},
								"total_chapters": schema.Int64Attribute{
									Description: "The number of chapters in this audiobook.",
									Computed:    true,
								},
								"type": schema.StringAttribute{
									Description: "The object type.\nAvailable values: \"audiobook\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("audiobook"),
									},
								},
								"uri": schema.StringAttribute{
									Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the audiobook.",
									Computed:    true,
								},
								"edition": schema.StringAttribute{
									Description: "The edition of the audiobook.",
									Computed:    true,
								},
								"chapters": schema.SingleNestedAttribute{
									Description: "The chapters of the audiobook.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[MeAudiobooksAudiobookChaptersDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"href": schema.StringAttribute{
											Description: "A link to the Web API endpoint returning the full result of the request",
											Computed:    true,
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
										"items": schema.ListNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectListType[MeAudiobooksAudiobookChaptersItemsDataSourceModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the chapter.",
														Computed:    true,
													},
													"audio_preview_url": schema.StringAttribute{
														Description:        "A URL to a 30 second preview (MP3 format) of the chapter. `null` if not available.",
														Computed:           true,
														DeprecationMessage: "This attribute is deprecated.",
													},
													"chapter_number": schema.Int64Attribute{
														Description: "The number of the chapter",
														Computed:    true,
													},
													"description": schema.StringAttribute{
														Description: "A description of the chapter. HTML tags are stripped away from this field, use `html_description` field in case HTML tags are needed.",
														Computed:    true,
													},
													"duration_ms": schema.Int64Attribute{
														Description: "The chapter length in milliseconds.",
														Computed:    true,
													},
													"explicit": schema.BoolAttribute{
														Description: "Whether or not the chapter has explicit content (true = yes it does; false = no it does not OR unknown).",
														Computed:    true,
													},
													"external_urls": schema.SingleNestedAttribute{
														Description: "External URLs for this chapter.",
														Computed:    true,
														CustomType:  customfield.NewNestedObjectType[MeAudiobooksAudiobookChaptersItemsExternalURLsDataSourceModel](ctx),
														Attributes: map[string]schema.Attribute{
															"spotify": schema.StringAttribute{
																Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
																Computed:    true,
															},
														},
													},
													"href": schema.StringAttribute{
														Description: "A link to the Web API endpoint providing full details of the chapter.",
														Computed:    true,
													},
													"html_description": schema.StringAttribute{
														Description: "A description of the chapter. This field may contain HTML tags.",
														Computed:    true,
													},
													"images": schema.ListNestedAttribute{
														Description: "The cover art for the chapter in various sizes, widest first.",
														Computed:    true,
														CustomType:  customfield.NewNestedObjectListType[MeAudiobooksAudiobookChaptersItemsImagesDataSourceModel](ctx),
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
													"is_playable": schema.BoolAttribute{
														Description: "True if the chapter is playable in the given market. Otherwise false.",
														Computed:    true,
													},
													"languages": schema.ListAttribute{
														Description: "A list of the languages used in the chapter, identified by their [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639) code.",
														Computed:    true,
														CustomType:  customfield.NewListType[types.String](ctx),
														ElementType: types.StringType,
													},
													"name": schema.StringAttribute{
														Description: "The name of the chapter.",
														Computed:    true,
													},
													"release_date": schema.StringAttribute{
														Description: "The date the chapter was first released, for example `\"1981-12-15\"`. Depending on the precision, it might be shown as `\"1981\"` or `\"1981-12\"`.",
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
														Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the chapter.",
														Computed:    true,
													},
													"available_markets": schema.ListAttribute{
														Description: "A list of the countries in which the chapter can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.",
														Computed:    true,
														CustomType:  customfield.NewListType[types.String](ctx),
														ElementType: types.StringType,
													},
													"restrictions": schema.SingleNestedAttribute{
														Description: "Included in the response when a content restriction is applied.",
														Computed:    true,
														CustomType:  customfield.NewNestedObjectType[MeAudiobooksAudiobookChaptersItemsRestrictionsDataSourceModel](ctx),
														Attributes: map[string]schema.Attribute{
															"reason": schema.StringAttribute{
																Description: "The reason for the restriction. Supported values:\n- `market` - The content item is not available in the given market.\n- `product` - The content item is not available for the user's subscription type.\n- `explicit` - The content item is explicit and the user's account is set to not play explicit content.\n- `payment_required` - Payment is required to play the content item.\n\nAdditional reasons may be added in the future.\n**Note**: If you use this field, make sure that your application safely handles unknown values.",
																Computed:    true,
															},
														},
													},
													"resume_point": schema.SingleNestedAttribute{
														Description: "The user's most recent position in the chapter. Set if the supplied access token is a user token and has the scope 'user-read-playback-position'.",
														Computed:    true,
														CustomType:  customfield.NewNestedObjectType[MeAudiobooksAudiobookChaptersItemsResumePointDataSourceModel](ctx),
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
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *MeAudiobooksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *MeAudiobooksDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
