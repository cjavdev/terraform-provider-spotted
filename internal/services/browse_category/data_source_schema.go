// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browse_category

import (
	"context"

	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*BrowseCategoryDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"category_id": schema.StringAttribute{
				Description: "The [Spotify category ID](/documentation/web-api/concepts/spotify-uris-ids) for the category.",
				Required:    true,
			},
			"locale": schema.StringAttribute{
				Description: "The desired language, consisting of an [ISO 639-1](http://en.wikipedia.org/wiki/ISO_639-1) language code and an [ISO 3166-1 alpha-2 country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), joined by an underscore. For example: `es_MX`, meaning &quot;Spanish (Mexico)&quot;. Provide this parameter if you want the category strings returned in a particular language.<br/> _**Note**: if `locale` is not supplied, or if the specified language is not available, the category strings returned will be in the Spotify default language (American English)._",
				Optional:    true,
			},
			"href": schema.StringAttribute{
				Description: "A link to the Web API endpoint returning full details of the category.",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "The [Spotify category ID](/documentation/web-api/concepts/spotify-uris-ids) of the category.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the category.",
				Computed:    true,
			},
			"icons": schema.ListNestedAttribute{
				Description: "The category icon, in various sizes.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[BrowseCategoryIconsDataSourceModel](ctx),
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

func (d *BrowseCategoryDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *BrowseCategoryDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
