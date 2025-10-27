// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audio_analysis

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*AudioAnalysisDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids)\nfor the track.",
				Required:    true,
			},
			"bars": schema.ListNestedAttribute{
				Description: "The time intervals of the bars throughout the track. A bar (or measure) is a segment of time defined as a given number of beats.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AudioAnalysisBarsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"confidence": schema.Float64Attribute{
							Description: "The confidence, from 0.0 to 1.0, of the reliability of the interval.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(0, 1),
							},
						},
						"duration": schema.Float64Attribute{
							Description: "The duration (in seconds) of the time interval.",
							Computed:    true,
						},
						"start": schema.Float64Attribute{
							Description: "The starting point (in seconds) of the time interval.",
							Computed:    true,
						},
					},
				},
			},
			"beats": schema.ListNestedAttribute{
				Description: "The time intervals of beats throughout the track. A beat is the basic time unit of a piece of music; for example, each tick of a metronome. Beats are typically multiples of tatums.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AudioAnalysisBeatsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"confidence": schema.Float64Attribute{
							Description: "The confidence, from 0.0 to 1.0, of the reliability of the interval.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(0, 1),
							},
						},
						"duration": schema.Float64Attribute{
							Description: "The duration (in seconds) of the time interval.",
							Computed:    true,
						},
						"start": schema.Float64Attribute{
							Description: "The starting point (in seconds) of the time interval.",
							Computed:    true,
						},
					},
				},
			},
			"meta": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[AudioAnalysisMetaDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"analysis_time": schema.Float64Attribute{
						Description: "The amount of time taken to analyze this track.",
						Computed:    true,
					},
					"analyzer_version": schema.StringAttribute{
						Description: "The version of the Analyzer used to analyze this track.",
						Computed:    true,
					},
					"detailed_status": schema.StringAttribute{
						Description: "A detailed status code for this track. If analysis data is missing, this code may explain why.",
						Computed:    true,
					},
					"input_process": schema.StringAttribute{
						Description: "The method used to read the track's audio data.",
						Computed:    true,
					},
					"platform": schema.StringAttribute{
						Description: "The platform used to read the track's audio data.",
						Computed:    true,
					},
					"status_code": schema.Int64Attribute{
						Description: "The return code of the analyzer process. 0 if successful, 1 if any errors occurred.",
						Computed:    true,
					},
					"timestamp": schema.Int64Attribute{
						Description: "The Unix timestamp (in seconds) at which this track was analyzed.",
						Computed:    true,
					},
				},
			},
			"sections": schema.ListNestedAttribute{
				Description: "Sections are defined by large variations in rhythm or timbre, e.g. chorus, verse, bridge, guitar solo, etc. Each section contains its own descriptions of tempo, key, mode, time_signature, and loudness.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AudioAnalysisSectionsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"confidence": schema.Float64Attribute{
							Description: `The confidence, from 0.0 to 1.0, of the reliability of the section's "designation".`,
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(0, 1),
							},
						},
						"duration": schema.Float64Attribute{
							Description: "The duration (in seconds) of the section.",
							Computed:    true,
						},
						"key": schema.Int64Attribute{
							Description: "The estimated overall key of the section. The values in this field ranging from 0 to 11 mapping to pitches using standard Pitch Class notation (E.g. 0 = C, 1 = C♯/D♭, 2 = D, and so on). If no key was detected, the value is -1.",
							Computed:    true,
						},
						"key_confidence": schema.Float64Attribute{
							Description: "The confidence, from 0.0 to 1.0, of the reliability of the key. Songs with many key changes may correspond to low values in this field.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(0, 1),
							},
						},
						"loudness": schema.Float64Attribute{
							Description: "The overall loudness of the section in decibels (dB). Loudness values are useful for comparing relative loudness of sections within tracks.",
							Computed:    true,
						},
						"mode": schema.Float64Attribute{
							Description: "Indicates the modality (major or minor) of a section, the type of scale from which its melodic content is derived. This field will contain a 0 for \"minor\", a 1 for \"major\", or a -1 for no result. Note that the major key (e.g. C major) could more likely be confused with the minor key at 3 semitones lower (e.g. A minor) as both keys carry the same pitches.\nAvailable values: -1, 0, 1.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.OneOf(
									-1,
									0,
									1,
								),
							},
						},
						"mode_confidence": schema.Float64Attribute{
							Description: "The confidence, from 0.0 to 1.0, of the reliability of the `mode`.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(0, 1),
							},
						},
						"start": schema.Float64Attribute{
							Description: "The starting point (in seconds) of the section.",
							Computed:    true,
						},
						"tempo": schema.Float64Attribute{
							Description: "The overall estimated tempo of the section in beats per minute (BPM). In musical terminology, tempo is the speed or pace of a given piece and derives directly from the average beat duration.",
							Computed:    true,
						},
						"tempo_confidence": schema.Float64Attribute{
							Description: "The confidence, from 0.0 to 1.0, of the reliability of the tempo. Some tracks contain tempo changes or sounds which don't contain tempo (like pure speech) which would correspond to a low value in this field.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(0, 1),
							},
						},
						"time_signature": schema.Int64Attribute{
							Description: `An estimated time signature. The time signature (meter) is a notational convention to specify how many beats are in each bar (or measure). The time signature ranges from 3 to 7 indicating time signatures of "3/4", to "7/4".`,
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(3, 7),
							},
						},
						"time_signature_confidence": schema.Float64Attribute{
							Description: "The confidence, from 0.0 to 1.0, of the reliability of the `time_signature`. Sections with time signature changes may correspond to low values in this field.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(0, 1),
							},
						},
					},
				},
			},
			"segments": schema.ListNestedAttribute{
				Description: "Each segment contains a roughly conisistent sound throughout its duration.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AudioAnalysisSegmentsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"confidence": schema.Float64Attribute{
							Description: "The confidence, from 0.0 to 1.0, of the reliability of the segmentation. Segments of the song which are difficult to logically segment (e.g: noise) may correspond to low values in this field.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(0, 1),
							},
						},
						"duration": schema.Float64Attribute{
							Description: "The duration (in seconds) of the segment.",
							Computed:    true,
						},
						"loudness_end": schema.Float64Attribute{
							Description: "The offset loudness of the segment in decibels (dB). This value should be equivalent to the loudness_start of the following segment.",
							Computed:    true,
						},
						"loudness_max": schema.Float64Attribute{
							Description: "The peak loudness of the segment in decibels (dB). Combined with `loudness_start` and `loudness_max_time`, these components can be used to describe the \"attack\" of the segment.",
							Computed:    true,
						},
						"loudness_max_time": schema.Float64Attribute{
							Description: "The segment-relative offset of the segment peak loudness in seconds. Combined with `loudness_start` and `loudness_max`, these components can be used to desctibe the \"attack\" of the segment.",
							Computed:    true,
						},
						"loudness_start": schema.Float64Attribute{
							Description: "The onset loudness of the segment in decibels (dB). Combined with `loudness_max` and `loudness_max_time`, these components can be used to describe the \"attack\" of the segment.",
							Computed:    true,
						},
						"pitches": schema.ListAttribute{
							Description: "Pitch content is given by a “chroma” vector, corresponding to the 12 pitch classes C, C#, D to B, with values ranging from 0 to 1 that describe the relative dominance of every pitch in the chromatic scale. For example a C Major chord would likely be represented by large values of C, E and G (i.e. classes 0, 4, and 7).\n\nVectors are normalized to 1 by their strongest dimension, therefore noisy sounds are likely represented by values that are all close to 1, while pure tones are described by one value at 1 (the pitch) and others near 0.\nAs can be seen below, the 12 vector indices are a combination of low-power spectrum values at their respective pitch frequencies.\n![pitch vector](/assets/audio/Pitch_vector.png)",
							Computed:    true,
							Validators: []validator.List{
								listvalidator.ValueFloat64sAre(
									float64validator.Between(0, 1),
								),
							},
							CustomType:  customfield.NewListType[types.Float64](ctx),
							ElementType: types.Float64Type,
						},
						"start": schema.Float64Attribute{
							Description: "The starting point (in seconds) of the segment.",
							Computed:    true,
						},
						"timbre": schema.ListAttribute{
							Description: "Timbre is the quality of a musical note or sound that distinguishes different types of musical instruments, or voices. It is a complex notion also referred to as sound color, texture, or tone quality, and is derived from the shape of a segment’s spectro-temporal surface, independently of pitch and loudness. The timbre feature is a vector that includes 12 unbounded values roughly centered around 0. Those values are high level abstractions of the spectral surface, ordered by degree of importance.\n\nFor completeness however, the first dimension represents the average loudness of the segment; second emphasizes brightness; third is more closely correlated to the flatness of a sound; fourth to sounds with a stronger attack; etc. See an image below representing the 12 basis functions (i.e. template segments).\n![timbre basis functions](/assets/audio/Timbre_basis_functions.png)\n\nThe actual timbre of the segment is best described as a linear combination of these 12 basis functions weighted by the coefficient values: timbre = c1 x b1 + c2 x b2 + ... + c12 x b12, where c1 to c12 represent the 12 coefficients and b1 to b12 the 12 basis functions as displayed below. Timbre vectors are best used in comparison with each other.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.Float64](ctx),
							ElementType: types.Float64Type,
						},
					},
				},
			},
			"tatums": schema.ListNestedAttribute{
				Description: "A tatum represents the lowest regular pulse train that a listener intuitively infers from the timing of perceived musical events (segments).",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AudioAnalysisTatumsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"confidence": schema.Float64Attribute{
							Description: "The confidence, from 0.0 to 1.0, of the reliability of the interval.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(0, 1),
							},
						},
						"duration": schema.Float64Attribute{
							Description: "The duration (in seconds) of the time interval.",
							Computed:    true,
						},
						"start": schema.Float64Attribute{
							Description: "The starting point (in seconds) of the time interval.",
							Computed:    true,
						},
					},
				},
			},
			"track": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[AudioAnalysisTrackDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"analysis_channels": schema.Int64Attribute{
						Description: "The number of channels used for analysis. If 1, all channels are summed together to mono before analysis.",
						Computed:    true,
					},
					"analysis_sample_rate": schema.Int64Attribute{
						Description: "The sample rate used to decode and analyze this track. May differ from the actual sample rate of this track available on Spotify.",
						Computed:    true,
					},
					"code_version": schema.Float64Attribute{
						Description: "A version number for the Echo Nest Musical Fingerprint format used in the codestring field.",
						Computed:    true,
					},
					"codestring": schema.StringAttribute{
						Description: "An [Echo Nest Musical Fingerprint (ENMFP)](https://academiccommons.columbia.edu/doi/10.7916/D8Q248M4) codestring for this track.",
						Computed:    true,
					},
					"duration": schema.Float64Attribute{
						Description: "Length of the track in seconds.",
						Computed:    true,
					},
					"echoprint_version": schema.Float64Attribute{
						Description: "A version number for the EchoPrint format used in the echoprintstring field.",
						Computed:    true,
					},
					"echoprintstring": schema.StringAttribute{
						Description: "An [EchoPrint](https://github.com/spotify/echoprint-codegen) codestring for this track.",
						Computed:    true,
					},
					"end_of_fade_in": schema.Float64Attribute{
						Description: "The time, in seconds, at which the track's fade-in period ends. If the track has no fade-in, this will be 0.0.",
						Computed:    true,
					},
					"key": schema.Int64Attribute{
						Description: "The key the track is in. Integers map to pitches using standard [Pitch Class notation](https://en.wikipedia.org/wiki/Pitch_class). E.g. 0 = C, 1 = C♯/D♭, 2 = D, and so on. If no key was detected, the value is -1.",
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.Between(-1, 11),
						},
					},
					"key_confidence": schema.Float64Attribute{
						Description: "The confidence, from 0.0 to 1.0, of the reliability of the `key`.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 1),
						},
					},
					"loudness": schema.Float64Attribute{
						Description: "The overall loudness of a track in decibels (dB). Loudness values are averaged across the entire track and are useful for comparing relative loudness of tracks. Loudness is the quality of a sound that is the primary psychological correlate of physical strength (amplitude). Values typically range between -60 and 0 db.",
						Computed:    true,
					},
					"mode": schema.Int64Attribute{
						Description: "Mode indicates the modality (major or minor) of a track, the type of scale from which its melodic content is derived. Major is represented by 1 and minor is 0.",
						Computed:    true,
					},
					"mode_confidence": schema.Float64Attribute{
						Description: "The confidence, from 0.0 to 1.0, of the reliability of the `mode`.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 1),
						},
					},
					"num_samples": schema.Int64Attribute{
						Description: "The exact number of audio samples analyzed from this track. See also `analysis_sample_rate`.",
						Computed:    true,
					},
					"offset_seconds": schema.Int64Attribute{
						Description: "An offset to the start of the region of the track that was analyzed. (As the entire track is analyzed, this should always be 0.)",
						Computed:    true,
					},
					"rhythm_version": schema.Float64Attribute{
						Description: "A version number for the Rhythmstring used in the rhythmstring field.",
						Computed:    true,
					},
					"rhythmstring": schema.StringAttribute{
						Description: "A Rhythmstring for this track. The format of this string is similar to the Synchstring.",
						Computed:    true,
					},
					"sample_md5": schema.StringAttribute{
						Description: "This field will always contain the empty string.",
						Computed:    true,
					},
					"start_of_fade_out": schema.Float64Attribute{
						Description: "The time, in seconds, at which the track's fade-out period starts. If the track has no fade-out, this should match the track's length.",
						Computed:    true,
					},
					"synch_version": schema.Float64Attribute{
						Description: "A version number for the Synchstring used in the synchstring field.",
						Computed:    true,
					},
					"synchstring": schema.StringAttribute{
						Description: "A [Synchstring](https://github.com/echonest/synchdata) for this track.",
						Computed:    true,
					},
					"tempo": schema.Float64Attribute{
						Description: "The overall estimated tempo of a track in beats per minute (BPM). In musical terminology, tempo is the speed or pace of a given piece and derives directly from the average beat duration.",
						Computed:    true,
					},
					"tempo_confidence": schema.Float64Attribute{
						Description: "The confidence, from 0.0 to 1.0, of the reliability of the `tempo`.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 1),
						},
					},
					"time_signature": schema.Int64Attribute{
						Description: `An estimated time signature. The time signature (meter) is a notational convention to specify how many beats are in each bar (or measure). The time signature ranges from 3 to 7 indicating time signatures of "3/4", to "7/4".`,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.Between(3, 7),
						},
					},
					"time_signature_confidence": schema.Float64Attribute{
						Description: "The confidence, from 0.0 to 1.0, of the reliability of the `time_signature`.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 1),
						},
					},
					"window_seconds": schema.Int64Attribute{
						Description: "The length of the region of the track was analyzed, if a subset of the track was analyzed. (As the entire track is analyzed, this should always be 0.)",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *AudioAnalysisDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AudioAnalysisDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
