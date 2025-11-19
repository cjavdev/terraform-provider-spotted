// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audio_analysis

import (
	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AudioAnalysisDataSourceModel struct {
	ID       types.String                                                       `tfsdk:"id" path:"id,required"`
	Bars     customfield.NestedObjectList[AudioAnalysisBarsDataSourceModel]     `tfsdk:"bars" json:"bars,computed"`
	Beats    customfield.NestedObjectList[AudioAnalysisBeatsDataSourceModel]    `tfsdk:"beats" json:"beats,computed"`
	Meta     customfield.NestedObject[AudioAnalysisMetaDataSourceModel]         `tfsdk:"meta" json:"meta,computed"`
	Sections customfield.NestedObjectList[AudioAnalysisSectionsDataSourceModel] `tfsdk:"sections" json:"sections,computed"`
	Segments customfield.NestedObjectList[AudioAnalysisSegmentsDataSourceModel] `tfsdk:"segments" json:"segments,computed"`
	Tatums   customfield.NestedObjectList[AudioAnalysisTatumsDataSourceModel]   `tfsdk:"tatums" json:"tatums,computed"`
	Track    customfield.NestedObject[AudioAnalysisTrackDataSourceModel]        `tfsdk:"track" json:"track,computed"`
}

type AudioAnalysisBarsDataSourceModel struct {
	Confidence types.Float64 `tfsdk:"confidence" json:"confidence,computed"`
	Duration   types.Float64 `tfsdk:"duration" json:"duration,computed"`
	Start      types.Float64 `tfsdk:"start" json:"start,computed"`
}

type AudioAnalysisBeatsDataSourceModel struct {
	Confidence types.Float64 `tfsdk:"confidence" json:"confidence,computed"`
	Duration   types.Float64 `tfsdk:"duration" json:"duration,computed"`
	Start      types.Float64 `tfsdk:"start" json:"start,computed"`
}

type AudioAnalysisMetaDataSourceModel struct {
	AnalysisTime    types.Float64 `tfsdk:"analysis_time" json:"analysis_time,computed"`
	AnalyzerVersion types.String  `tfsdk:"analyzer_version" json:"analyzer_version,computed"`
	DetailedStatus  types.String  `tfsdk:"detailed_status" json:"detailed_status,computed"`
	InputProcess    types.String  `tfsdk:"input_process" json:"input_process,computed"`
	Platform        types.String  `tfsdk:"platform" json:"platform,computed"`
	StatusCode      types.Int64   `tfsdk:"status_code" json:"status_code,computed"`
	Timestamp       types.Int64   `tfsdk:"timestamp" json:"timestamp,computed"`
}

type AudioAnalysisSectionsDataSourceModel struct {
	Confidence              types.Float64 `tfsdk:"confidence" json:"confidence,computed"`
	Duration                types.Float64 `tfsdk:"duration" json:"duration,computed"`
	Key                     types.Int64   `tfsdk:"key" json:"key,computed"`
	KeyConfidence           types.Float64 `tfsdk:"key_confidence" json:"key_confidence,computed"`
	Loudness                types.Float64 `tfsdk:"loudness" json:"loudness,computed"`
	Mode                    types.Float64 `tfsdk:"mode" json:"mode,computed"`
	ModeConfidence          types.Float64 `tfsdk:"mode_confidence" json:"mode_confidence,computed"`
	Start                   types.Float64 `tfsdk:"start" json:"start,computed"`
	Tempo                   types.Float64 `tfsdk:"tempo" json:"tempo,computed"`
	TempoConfidence         types.Float64 `tfsdk:"tempo_confidence" json:"tempo_confidence,computed"`
	TimeSignature           types.Int64   `tfsdk:"time_signature" json:"time_signature,computed"`
	TimeSignatureConfidence types.Float64 `tfsdk:"time_signature_confidence" json:"time_signature_confidence,computed"`
}

type AudioAnalysisSegmentsDataSourceModel struct {
	Confidence      types.Float64                   `tfsdk:"confidence" json:"confidence,computed"`
	Duration        types.Float64                   `tfsdk:"duration" json:"duration,computed"`
	LoudnessEnd     types.Float64                   `tfsdk:"loudness_end" json:"loudness_end,computed"`
	LoudnessMax     types.Float64                   `tfsdk:"loudness_max" json:"loudness_max,computed"`
	LoudnessMaxTime types.Float64                   `tfsdk:"loudness_max_time" json:"loudness_max_time,computed"`
	LoudnessStart   types.Float64                   `tfsdk:"loudness_start" json:"loudness_start,computed"`
	Pitches         customfield.List[types.Float64] `tfsdk:"pitches" json:"pitches,computed"`
	Start           types.Float64                   `tfsdk:"start" json:"start,computed"`
	Timbre          customfield.List[types.Float64] `tfsdk:"timbre" json:"timbre,computed"`
}

type AudioAnalysisTatumsDataSourceModel struct {
	Confidence types.Float64 `tfsdk:"confidence" json:"confidence,computed"`
	Duration   types.Float64 `tfsdk:"duration" json:"duration,computed"`
	Start      types.Float64 `tfsdk:"start" json:"start,computed"`
}

type AudioAnalysisTrackDataSourceModel struct {
	AnalysisChannels        types.Int64   `tfsdk:"analysis_channels" json:"analysis_channels,computed"`
	AnalysisSampleRate      types.Int64   `tfsdk:"analysis_sample_rate" json:"analysis_sample_rate,computed"`
	CodeVersion             types.Float64 `tfsdk:"code_version" json:"code_version,computed"`
	Codestring              types.String  `tfsdk:"codestring" json:"codestring,computed"`
	Duration                types.Float64 `tfsdk:"duration" json:"duration,computed"`
	EchoprintVersion        types.Float64 `tfsdk:"echoprint_version" json:"echoprint_version,computed"`
	Echoprintstring         types.String  `tfsdk:"echoprintstring" json:"echoprintstring,computed"`
	EndOfFadeIn             types.Float64 `tfsdk:"end_of_fade_in" json:"end_of_fade_in,computed"`
	Key                     types.Int64   `tfsdk:"key" json:"key,computed"`
	KeyConfidence           types.Float64 `tfsdk:"key_confidence" json:"key_confidence,computed"`
	Loudness                types.Float64 `tfsdk:"loudness" json:"loudness,computed"`
	Mode                    types.Int64   `tfsdk:"mode" json:"mode,computed"`
	ModeConfidence          types.Float64 `tfsdk:"mode_confidence" json:"mode_confidence,computed"`
	NumSamples              types.Int64   `tfsdk:"num_samples" json:"num_samples,computed"`
	OffsetSeconds           types.Int64   `tfsdk:"offset_seconds" json:"offset_seconds,computed"`
	RhythmVersion           types.Float64 `tfsdk:"rhythm_version" json:"rhythm_version,computed"`
	Rhythmstring            types.String  `tfsdk:"rhythmstring" json:"rhythmstring,computed"`
	SampleMd5               types.String  `tfsdk:"sample_md5" json:"sample_md5,computed"`
	StartOfFadeOut          types.Float64 `tfsdk:"start_of_fade_out" json:"start_of_fade_out,computed"`
	SynchVersion            types.Float64 `tfsdk:"synch_version" json:"synch_version,computed"`
	Synchstring             types.String  `tfsdk:"synchstring" json:"synchstring,computed"`
	Tempo                   types.Float64 `tfsdk:"tempo" json:"tempo,computed"`
	TempoConfidence         types.Float64 `tfsdk:"tempo_confidence" json:"tempo_confidence,computed"`
	TimeSignature           types.Int64   `tfsdk:"time_signature" json:"time_signature,computed"`
	TimeSignatureConfidence types.Float64 `tfsdk:"time_signature_confidence" json:"time_signature_confidence,computed"`
	WindowSeconds           types.Int64   `tfsdk:"window_seconds" json:"window_seconds,computed"`
}
