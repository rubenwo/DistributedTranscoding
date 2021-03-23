// Code generated by protoc-gen-go. DO NOT EDIT.
// source: distributor-service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type VideoCodec int32

const (
	VideoCodec_Libx264 VideoCodec = 0
	VideoCodec_Libx265 VideoCodec = 1
	VideoCodec_Vpx     VideoCodec = 2
	VideoCodec_Vp8     VideoCodec = 3
	VideoCodec_Vp9     VideoCodec = 4
)

var VideoCodec_name = map[int32]string{
	0: "Libx264",
	1: "Libx265",
	2: "Vpx",
	3: "Vp8",
	4: "Vp9",
}

var VideoCodec_value = map[string]int32{
	"Libx264": 0,
	"Libx265": 1,
	"Vpx":     2,
	"Vp8":     3,
	"Vp9":     4,
}

func (x VideoCodec) String() string {
	return proto.EnumName(VideoCodec_name, int32(x))
}

func (VideoCodec) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aa095822aff8ecc1, []int{0}
}

type AudioCodec int32

const (
	AudioCodec_Aac  AudioCodec = 0
	AudioCodec_Ac3  AudioCodec = 1
	AudioCodec_Opus AudioCodec = 2
)

var AudioCodec_name = map[int32]string{
	0: "Aac",
	1: "Ac3",
	2: "Opus",
}

var AudioCodec_value = map[string]int32{
	"Aac":  0,
	"Ac3":  1,
	"Opus": 2,
}

func (x AudioCodec) String() string {
	return proto.EnumName(AudioCodec_name, int32(x))
}

func (AudioCodec) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aa095822aff8ecc1, []int{1}
}

type MediaFileType int32

const (
	MediaFileType_Mp4  MediaFileType = 0
	MediaFileType_Mkv  MediaFileType = 1
	MediaFileType_Webm MediaFileType = 2
)

var MediaFileType_name = map[int32]string{
	0: "Mp4",
	1: "Mkv",
	2: "Webm",
}

var MediaFileType_value = map[string]int32{
	"Mp4":  0,
	"Mkv":  1,
	"Webm": 2,
}

func (x MediaFileType) String() string {
	return proto.EnumName(MediaFileType_name, int32(x))
}

func (MediaFileType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aa095822aff8ecc1, []int{2}
}

type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_Ok      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

var UploadStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}

var UploadStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x UploadStatusCode) String() string {
	return proto.EnumName(UploadStatusCode_name, int32(x))
}

func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aa095822aff8ecc1, []int{3}
}

type ProgressState int32

const (
	ProgressState_Distributing ProgressState = 0
	ProgressState_Transcoding  ProgressState = 1
)

var ProgressState_name = map[int32]string{
	0: "Distributing",
	1: "Transcoding",
}

var ProgressState_value = map[string]int32{
	"Distributing": 0,
	"Transcoding":  1,
}

func (x ProgressState) String() string {
	return proto.EnumName(ProgressState_name, int32(x))
}

func (ProgressState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aa095822aff8ecc1, []int{4}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa095822aff8ecc1, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ClusterClientOffer struct {
	MaxConcurrentJobs    int64    `protobuf:"varint,1,opt,name=max_concurrent_jobs,json=maxConcurrentJobs,proto3" json:"max_concurrent_jobs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterClientOffer) Reset()         { *m = ClusterClientOffer{} }
func (m *ClusterClientOffer) String() string { return proto.CompactTextString(m) }
func (*ClusterClientOffer) ProtoMessage()    {}
func (*ClusterClientOffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa095822aff8ecc1, []int{1}
}

func (m *ClusterClientOffer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterClientOffer.Unmarshal(m, b)
}
func (m *ClusterClientOffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterClientOffer.Marshal(b, m, deterministic)
}
func (m *ClusterClientOffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterClientOffer.Merge(m, src)
}
func (m *ClusterClientOffer) XXX_Size() int {
	return xxx_messageInfo_ClusterClientOffer.Size(m)
}
func (m *ClusterClientOffer) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterClientOffer.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterClientOffer proto.InternalMessageInfo

func (m *ClusterClientOffer) GetMaxConcurrentJobs() int64 {
	if m != nil {
		return m.MaxConcurrentJobs
	}
	return 0
}

type VideoSettings struct {
	Codec                VideoCodec `protobuf:"varint,1,opt,name=codec,proto3,enum=VideoCodec" json:"codec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *VideoSettings) Reset()         { *m = VideoSettings{} }
func (m *VideoSettings) String() string { return proto.CompactTextString(m) }
func (*VideoSettings) ProtoMessage()    {}
func (*VideoSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa095822aff8ecc1, []int{2}
}

func (m *VideoSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoSettings.Unmarshal(m, b)
}
func (m *VideoSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoSettings.Marshal(b, m, deterministic)
}
func (m *VideoSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoSettings.Merge(m, src)
}
func (m *VideoSettings) XXX_Size() int {
	return xxx_messageInfo_VideoSettings.Size(m)
}
func (m *VideoSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoSettings.DiscardUnknown(m)
}

var xxx_messageInfo_VideoSettings proto.InternalMessageInfo

func (m *VideoSettings) GetCodec() VideoCodec {
	if m != nil {
		return m.Codec
	}
	return VideoCodec_Libx264
}

type AudioSettings struct {
	Codec                AudioCodec `protobuf:"varint,2,opt,name=codec,proto3,enum=AudioCodec" json:"codec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AudioSettings) Reset()         { *m = AudioSettings{} }
func (m *AudioSettings) String() string { return proto.CompactTextString(m) }
func (*AudioSettings) ProtoMessage()    {}
func (*AudioSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa095822aff8ecc1, []int{3}
}

func (m *AudioSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AudioSettings.Unmarshal(m, b)
}
func (m *AudioSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AudioSettings.Marshal(b, m, deterministic)
}
func (m *AudioSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AudioSettings.Merge(m, src)
}
func (m *AudioSettings) XXX_Size() int {
	return xxx_messageInfo_AudioSettings.Size(m)
}
func (m *AudioSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_AudioSettings.DiscardUnknown(m)
}

var xxx_messageInfo_AudioSettings proto.InternalMessageInfo

func (m *AudioSettings) GetCodec() AudioCodec {
	if m != nil {
		return m.Codec
	}
	return AudioCodec_Aac
}

type TranscodingSettings struct {
	VideoSettings        *VideoSettings `protobuf:"bytes,1,opt,name=video_settings,json=videoSettings,proto3" json:"video_settings,omitempty"`
	AudioSettings        *AudioSettings `protobuf:"bytes,2,opt,name=audio_settings,json=audioSettings,proto3" json:"audio_settings,omitempty"`
	MediaFileType        MediaFileType  `protobuf:"varint,3,opt,name=media_file_type,json=mediaFileType,proto3,enum=MediaFileType" json:"media_file_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TranscodingSettings) Reset()         { *m = TranscodingSettings{} }
func (m *TranscodingSettings) String() string { return proto.CompactTextString(m) }
func (*TranscodingSettings) ProtoMessage()    {}
func (*TranscodingSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa095822aff8ecc1, []int{4}
}

func (m *TranscodingSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranscodingSettings.Unmarshal(m, b)
}
func (m *TranscodingSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranscodingSettings.Marshal(b, m, deterministic)
}
func (m *TranscodingSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranscodingSettings.Merge(m, src)
}
func (m *TranscodingSettings) XXX_Size() int {
	return xxx_messageInfo_TranscodingSettings.Size(m)
}
func (m *TranscodingSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_TranscodingSettings.DiscardUnknown(m)
}

var xxx_messageInfo_TranscodingSettings proto.InternalMessageInfo

func (m *TranscodingSettings) GetVideoSettings() *VideoSettings {
	if m != nil {
		return m.VideoSettings
	}
	return nil
}

func (m *TranscodingSettings) GetAudioSettings() *AudioSettings {
	if m != nil {
		return m.AudioSettings
	}
	return nil
}

func (m *TranscodingSettings) GetMediaFileType() MediaFileType {
	if m != nil {
		return m.MediaFileType
	}
	return MediaFileType_Mp4
}

type Job struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ReferenceNumber      int64                `protobuf:"varint,2,opt,name=reference_number,json=referenceNumber,proto3" json:"reference_number,omitempty"`
	TranscodingSettings  *TranscodingSettings `protobuf:"bytes,3,opt,name=transcoding_settings,json=transcodingSettings,proto3" json:"transcoding_settings,omitempty"`
	InputFileName        string               `protobuf:"bytes,4,opt,name=input_file_name,json=inputFileName,proto3" json:"input_file_name,omitempty"`
	InputFileData        []byte               `protobuf:"bytes,5,opt,name=input_file_data,json=inputFileData,proto3" json:"input_file_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa095822aff8ecc1, []int{5}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Job) GetReferenceNumber() int64 {
	if m != nil {
		return m.ReferenceNumber
	}
	return 0
}

func (m *Job) GetTranscodingSettings() *TranscodingSettings {
	if m != nil {
		return m.TranscodingSettings
	}
	return nil
}

func (m *Job) GetInputFileName() string {
	if m != nil {
		return m.InputFileName
	}
	return ""
}

func (m *Job) GetInputFileData() []byte {
	if m != nil {
		return m.InputFileData
	}
	return nil
}

type Result struct {
	JobId                string           `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	JobReferenceNumber   int64            `protobuf:"varint,2,opt,name=job_reference_number,json=jobReferenceNumber,proto3" json:"job_reference_number,omitempty"`
	StatusCode           UploadStatusCode `protobuf:"varint,3,opt,name=status_code,json=statusCode,proto3,enum=UploadStatusCode" json:"status_code,omitempty"`
	OutputFileData       []byte           `protobuf:"bytes,4,opt,name=output_file_data,json=outputFileData,proto3" json:"output_file_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa095822aff8ecc1, []int{6}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *Result) GetJobReferenceNumber() int64 {
	if m != nil {
		return m.JobReferenceNumber
	}
	return 0
}

func (m *Result) GetStatusCode() UploadStatusCode {
	if m != nil {
		return m.StatusCode
	}
	return UploadStatusCode_Unknown
}

func (m *Result) GetOutputFileData() []byte {
	if m != nil {
		return m.OutputFileData
	}
	return nil
}

type TranscodeRequest struct {
	InputFileName        string               `protobuf:"bytes,4,opt,name=input_file_name,json=inputFileName,proto3" json:"input_file_name,omitempty"`
	InputFileData        []byte               `protobuf:"bytes,5,opt,name=input_file_data,json=inputFileData,proto3" json:"input_file_data,omitempty"`
	TranscodingSettings  *TranscodingSettings `protobuf:"bytes,3,opt,name=transcoding_settings,json=transcodingSettings,proto3" json:"transcoding_settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TranscodeRequest) Reset()         { *m = TranscodeRequest{} }
func (m *TranscodeRequest) String() string { return proto.CompactTextString(m) }
func (*TranscodeRequest) ProtoMessage()    {}
func (*TranscodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa095822aff8ecc1, []int{7}
}

func (m *TranscodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranscodeRequest.Unmarshal(m, b)
}
func (m *TranscodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranscodeRequest.Marshal(b, m, deterministic)
}
func (m *TranscodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranscodeRequest.Merge(m, src)
}
func (m *TranscodeRequest) XXX_Size() int {
	return xxx_messageInfo_TranscodeRequest.Size(m)
}
func (m *TranscodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TranscodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TranscodeRequest proto.InternalMessageInfo

func (m *TranscodeRequest) GetInputFileName() string {
	if m != nil {
		return m.InputFileName
	}
	return ""
}

func (m *TranscodeRequest) GetInputFileData() []byte {
	if m != nil {
		return m.InputFileData
	}
	return nil
}

func (m *TranscodeRequest) GetTranscodingSettings() *TranscodingSettings {
	if m != nil {
		return m.TranscodingSettings
	}
	return nil
}

type TranscodeResponse struct {
	Uuid                 string           `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	StatusCode           UploadStatusCode `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3,enum=UploadStatusCode" json:"status_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TranscodeResponse) Reset()         { *m = TranscodeResponse{} }
func (m *TranscodeResponse) String() string { return proto.CompactTextString(m) }
func (*TranscodeResponse) ProtoMessage()    {}
func (*TranscodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa095822aff8ecc1, []int{8}
}

func (m *TranscodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranscodeResponse.Unmarshal(m, b)
}
func (m *TranscodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranscodeResponse.Marshal(b, m, deterministic)
}
func (m *TranscodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranscodeResponse.Merge(m, src)
}
func (m *TranscodeResponse) XXX_Size() int {
	return xxx_messageInfo_TranscodeResponse.Size(m)
}
func (m *TranscodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TranscodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TranscodeResponse proto.InternalMessageInfo

func (m *TranscodeResponse) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *TranscodeResponse) GetStatusCode() UploadStatusCode {
	if m != nil {
		return m.StatusCode
	}
	return UploadStatusCode_Unknown
}

type ProgressRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProgressRequest) Reset()         { *m = ProgressRequest{} }
func (m *ProgressRequest) String() string { return proto.CompactTextString(m) }
func (*ProgressRequest) ProtoMessage()    {}
func (*ProgressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa095822aff8ecc1, []int{9}
}

func (m *ProgressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProgressRequest.Unmarshal(m, b)
}
func (m *ProgressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProgressRequest.Marshal(b, m, deterministic)
}
func (m *ProgressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProgressRequest.Merge(m, src)
}
func (m *ProgressRequest) XXX_Size() int {
	return xxx_messageInfo_ProgressRequest.Size(m)
}
func (m *ProgressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProgressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProgressRequest proto.InternalMessageInfo

func (m *ProgressRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type ProgressResponse struct {
	State                ProgressState `protobuf:"varint,1,opt,name=state,proto3,enum=ProgressState" json:"state,omitempty"`
	Progress             float32       `protobuf:"fixed32,2,opt,name=progress,proto3" json:"progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ProgressResponse) Reset()         { *m = ProgressResponse{} }
func (m *ProgressResponse) String() string { return proto.CompactTextString(m) }
func (*ProgressResponse) ProtoMessage()    {}
func (*ProgressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa095822aff8ecc1, []int{10}
}

func (m *ProgressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProgressResponse.Unmarshal(m, b)
}
func (m *ProgressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProgressResponse.Marshal(b, m, deterministic)
}
func (m *ProgressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProgressResponse.Merge(m, src)
}
func (m *ProgressResponse) XXX_Size() int {
	return xxx_messageInfo_ProgressResponse.Size(m)
}
func (m *ProgressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProgressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProgressResponse proto.InternalMessageInfo

func (m *ProgressResponse) GetState() ProgressState {
	if m != nil {
		return m.State
	}
	return ProgressState_Distributing
}

func (m *ProgressResponse) GetProgress() float32 {
	if m != nil {
		return m.Progress
	}
	return 0
}

func init() {
	proto.RegisterEnum("VideoCodec", VideoCodec_name, VideoCodec_value)
	proto.RegisterEnum("AudioCodec", AudioCodec_name, AudioCodec_value)
	proto.RegisterEnum("MediaFileType", MediaFileType_name, MediaFileType_value)
	proto.RegisterEnum("UploadStatusCode", UploadStatusCode_name, UploadStatusCode_value)
	proto.RegisterEnum("ProgressState", ProgressState_name, ProgressState_value)
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*ClusterClientOffer)(nil), "ClusterClientOffer")
	proto.RegisterType((*VideoSettings)(nil), "VideoSettings")
	proto.RegisterType((*AudioSettings)(nil), "AudioSettings")
	proto.RegisterType((*TranscodingSettings)(nil), "TranscodingSettings")
	proto.RegisterType((*Job)(nil), "Job")
	proto.RegisterType((*Result)(nil), "Result")
	proto.RegisterType((*TranscodeRequest)(nil), "TranscodeRequest")
	proto.RegisterType((*TranscodeResponse)(nil), "TranscodeResponse")
	proto.RegisterType((*ProgressRequest)(nil), "ProgressRequest")
	proto.RegisterType((*ProgressResponse)(nil), "ProgressResponse")
}

func init() { proto.RegisterFile("distributor-service.proto", fileDescriptor_aa095822aff8ecc1) }

var fileDescriptor_aa095822aff8ecc1 = []byte{
	// 791 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xc7, 0x2d, 0xf9, 0x23, 0xe9, 0x71, 0x6c, 0xd3, 0x4c, 0x06, 0x64, 0xbe, 0x6a, 0x85, 0x6d,
	0xc8, 0xbc, 0x8d, 0xed, 0x9c, 0xa4, 0xd8, 0x2e, 0x33, 0x67, 0x1d, 0x16, 0x2c, 0xcd, 0xa0, 0xa4,
	0x1d, 0xb0, 0x5d, 0x18, 0x94, 0x74, 0x62, 0x30, 0xb1, 0x48, 0x4d, 0xa4, 0xdc, 0xe4, 0x99, 0x76,
	0xb1, 0x07, 0xd8, 0x93, 0xec, 0x6d, 0x06, 0x51, 0xb2, 0x1c, 0xb9, 0x19, 0xd0, 0x8b, 0xde, 0x1d,
	0x9d, 0x0f, 0xf2, 0x77, 0xfe, 0x24, 0x8f, 0xe0, 0xd3, 0x48, 0x68, 0x93, 0x8a, 0x20, 0x33, 0x2a,
	0xfd, 0x46, 0x63, 0xba, 0x14, 0x21, 0xb2, 0x24, 0x55, 0x46, 0x79, 0x5b, 0xd0, 0xfe, 0x31, 0x4e,
	0xcc, 0xbd, 0x77, 0x0a, 0x74, 0xba, 0xc8, 0xb4, 0xc1, 0x74, 0xba, 0x10, 0x28, 0xcd, 0xc5, 0xf5,
	0x35, 0xa6, 0x94, 0xc1, 0x6e, 0xcc, 0xef, 0x66, 0xa1, 0x92, 0x61, 0x96, 0xa6, 0x28, 0xcd, 0xec,
	0x46, 0x05, 0x7a, 0xdf, 0x79, 0xea, 0x1c, 0x34, 0xfd, 0x61, 0xcc, 0xef, 0xa6, 0x55, 0xe4, 0x4c,
	0x05, 0xda, 0x9b, 0x40, 0xef, 0xad, 0x88, 0x50, 0x5d, 0xa2, 0x31, 0x42, 0xce, 0x35, 0x7d, 0x06,
	0xed, 0x50, 0x45, 0x18, 0xda, 0x92, 0xfe, 0xa4, 0xcb, 0x6c, 0x78, 0x9a, 0xbb, 0xfc, 0x22, 0x92,
	0xd7, 0x9c, 0x64, 0x91, 0x78, 0xa4, 0xc6, 0x2d, 0x6b, 0x6c, 0xb8, 0x56, 0xf3, 0x8f, 0x03, 0xbb,
	0x57, 0x29, 0x97, 0x3a, 0x54, 0x91, 0x90, 0xf3, 0xaa, 0xf4, 0x18, 0xfa, 0xcb, 0x7c, 0x83, 0x99,
	0x2e, 0x3d, 0x76, 0xdf, 0xee, 0xa4, 0xcf, 0x6a, 0x58, 0x7e, 0x6f, 0x59, 0xa3, 0x3c, 0x86, 0x3e,
	0xcf, 0xf7, 0x58, 0x97, 0xb9, 0x65, 0x59, 0x8d, 0xcc, 0xef, 0xf1, 0x1a, 0xe8, 0x4b, 0x18, 0xc4,
	0x18, 0x09, 0x3e, 0xbb, 0x16, 0x0b, 0x9c, 0x99, 0xfb, 0x04, 0xf7, 0x9b, 0x16, 0xb9, 0xcf, 0xce,
	0x73, 0xff, 0x2b, 0xb1, 0xc0, 0xab, 0xfb, 0x04, 0xfd, 0x5e, 0xfc, 0xf0, 0xd3, 0xfb, 0xd7, 0x81,
	0xe6, 0x99, 0x0a, 0x68, 0x1f, 0x5c, 0x11, 0x59, 0xc2, 0x27, 0xbe, 0x2b, 0x22, 0xfa, 0x25, 0x90,
	0x14, 0xaf, 0x31, 0x45, 0x19, 0xe2, 0x4c, 0x66, 0x71, 0x80, 0xa9, 0x05, 0x69, 0xfa, 0x83, 0xca,
	0xff, 0xda, 0xba, 0xe9, 0x4f, 0xb0, 0x67, 0xd6, 0xfd, 0xaf, 0xb9, 0x9b, 0x96, 0x7b, 0x8f, 0x3d,
	0x22, 0x8e, 0xbf, 0x6b, 0x1e, 0x51, 0xec, 0x0b, 0x18, 0x08, 0x99, 0x64, 0xa6, 0xe8, 0x41, 0xf2,
	0x18, 0xf7, 0x5b, 0x16, 0xa8, 0x67, 0xdd, 0x39, 0xf3, 0x6b, 0x1e, 0xe3, 0x46, 0x5e, 0xc4, 0x0d,
	0xdf, 0x6f, 0x3f, 0x75, 0x0e, 0x76, 0x1e, 0xe4, 0x9d, 0x72, 0xc3, 0xbd, 0xbf, 0x1d, 0xe8, 0xf8,
	0xa8, 0xb3, 0x85, 0xa1, 0x9f, 0x40, 0xe7, 0x46, 0x05, 0xb3, 0xaa, 0xc5, 0xf6, 0x8d, 0x0a, 0x7e,
	0x8e, 0xe8, 0x0b, 0xd8, 0xcb, 0xdd, 0xff, 0xd3, 0x29, 0xbd, 0x51, 0x81, 0xbf, 0xd1, 0xec, 0x04,
	0xba, 0xda, 0x70, 0x93, 0xe9, 0x59, 0x7e, 0xfa, 0xa5, 0xc6, 0x43, 0xf6, 0x26, 0x59, 0x28, 0x1e,
	0x5d, 0xda, 0x48, 0x7e, 0x3b, 0x7c, 0xd0, 0x95, 0x4d, 0x0f, 0x80, 0xa8, 0xcc, 0xd4, 0x81, 0x5b,
	0x16, 0xb8, 0x5f, 0xf8, 0x2b, 0xe2, 0xbf, 0x1c, 0x20, 0x2b, 0xb9, 0xd0, 0xc7, 0x3f, 0x33, 0xd4,
	0xe6, 0x63, 0xcb, 0xf2, 0xd1, 0xce, 0xcb, 0xfb, 0x03, 0x86, 0x0f, 0x60, 0x75, 0xa2, 0xa4, 0x46,
	0x4a, 0xa1, 0x95, 0x65, 0x95, 0xce, 0xd6, 0xde, 0x14, 0xcd, 0xfd, 0x00, 0xd1, 0xbc, 0xcf, 0x61,
	0xf0, 0x6b, 0xaa, 0xe6, 0x29, 0x6a, 0xbd, 0x12, 0xe2, 0x91, 0xa5, 0xbd, 0x2b, 0x20, 0xeb, 0xb4,
	0x12, 0xe1, 0x33, 0x68, 0xe7, 0x0b, 0x61, 0xf9, 0xd0, 0xfb, 0x6c, 0x95, 0x91, 0x6f, 0x85, 0x7e,
	0x11, 0xa4, 0x23, 0xd8, 0x4e, 0x4a, 0xbf, 0x25, 0x72, 0xfd, 0xea, 0x7b, 0x7c, 0x02, 0xb0, 0x1e,
	0x0e, 0xb4, 0x0b, 0x5b, 0xbf, 0x88, 0xe0, 0x6e, 0xf2, 0xf2, 0x88, 0x34, 0xd6, 0x1f, 0xc7, 0xc4,
	0xa1, 0x5b, 0xd0, 0x7c, 0x9b, 0xdc, 0x11, 0xb7, 0x30, 0xbe, 0x23, 0xcd, 0xc2, 0xf8, 0x9e, 0xb4,
	0xc6, 0x07, 0x00, 0xeb, 0x59, 0x91, 0xbb, 0x4f, 0x78, 0x48, 0x1a, 0xd6, 0x08, 0x0f, 0x89, 0x43,
	0xb7, 0xa1, 0x75, 0x91, 0x64, 0x9a, 0xb8, 0xe3, 0xaf, 0xa0, 0x57, 0x7b, 0xa2, 0x79, 0xce, 0x79,
	0x72, 0x54, 0x24, 0x9f, 0xdf, 0x2e, 0x8b, 0xe4, 0xdf, 0x30, 0x88, 0x89, 0x3b, 0x3e, 0x04, 0xb2,
	0x29, 0x5b, 0x8e, 0xf4, 0x46, 0xde, 0x4a, 0xf5, 0x4e, 0x92, 0x06, 0xed, 0x80, 0x7b, 0x71, 0x4b,
	0x1c, 0x0a, 0xd0, 0x79, 0xc5, 0xc5, 0x02, 0x23, 0xe2, 0x8e, 0x27, 0xd0, 0xab, 0x49, 0x40, 0x09,
	0xec, 0x9c, 0xae, 0xe6, 0xb0, 0x90, 0x73, 0xd2, 0xa0, 0x03, 0xe8, 0x3e, 0x38, 0x77, 0xe2, 0x4c,
	0x10, 0xe8, 0xe9, 0x7a, 0x54, 0x5f, 0x16, 0x93, 0x9a, 0x7e, 0x0d, 0xdd, 0x33, 0x25, 0x64, 0x39,
	0x9e, 0xe9, 0x2e, 0x7b, 0x7f, 0x50, 0x8f, 0x5a, 0xec, 0x4c, 0x05, 0x5e, 0xe3, 0x85, 0x43, 0x9f,
	0xc1, 0x4e, 0x01, 0x5b, 0xbe, 0xc2, 0x2d, 0x56, 0x18, 0xa3, 0x0e, 0x2b, 0x26, 0x7d, 0x63, 0xf2,
	0x0e, 0x60, 0xba, 0x10, 0xab, 0xe5, 0x8f, 0xe0, 0x49, 0x75, 0xa3, 0xe8, 0x90, 0x6d, 0x3e, 0x85,
	0x11, 0x65, 0xef, 0x5d, 0x38, 0xaf, 0x41, 0x0f, 0x61, 0x7b, 0xd5, 0x1e, 0x25, 0x6c, 0xe3, 0xd6,
	0x8c, 0x86, 0x6c, 0xf3, 0x82, 0xe4, 0x6c, 0x3f, 0xec, 0xfc, 0x0e, 0xc9, 0xed, 0xfc, 0x39, 0x4f,
	0xc4, 0xf3, 0xe5, 0xb7, 0x41, 0xc7, 0xfe, 0x82, 0x0e, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x83,
	0x23, 0xca, 0x84, 0x9f, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DistributorServiceClient is the client API for DistributorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DistributorServiceClient interface {
	JoinCluster(ctx context.Context, in *ClusterClientOffer, opts ...grpc.CallOption) (DistributorService_JoinClusterClient, error)
	UploadResult(ctx context.Context, in *Result, opts ...grpc.CallOption) (*Empty, error)
}

type distributorServiceClient struct {
	cc *grpc.ClientConn
}

func NewDistributorServiceClient(cc *grpc.ClientConn) DistributorServiceClient {
	return &distributorServiceClient{cc}
}

func (c *distributorServiceClient) JoinCluster(ctx context.Context, in *ClusterClientOffer, opts ...grpc.CallOption) (DistributorService_JoinClusterClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DistributorService_serviceDesc.Streams[0], "/DistributorService/JoinCluster", opts...)
	if err != nil {
		return nil, err
	}
	x := &distributorServiceJoinClusterClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DistributorService_JoinClusterClient interface {
	Recv() (*Job, error)
	grpc.ClientStream
}

type distributorServiceJoinClusterClient struct {
	grpc.ClientStream
}

func (x *distributorServiceJoinClusterClient) Recv() (*Job, error) {
	m := new(Job)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *distributorServiceClient) UploadResult(ctx context.Context, in *Result, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/DistributorService/UploadResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DistributorServiceServer is the server API for DistributorService service.
type DistributorServiceServer interface {
	JoinCluster(*ClusterClientOffer, DistributorService_JoinClusterServer) error
	UploadResult(context.Context, *Result) (*Empty, error)
}

// UnimplementedDistributorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDistributorServiceServer struct {
}

func (*UnimplementedDistributorServiceServer) JoinCluster(req *ClusterClientOffer, srv DistributorService_JoinClusterServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinCluster not implemented")
}
func (*UnimplementedDistributorServiceServer) UploadResult(ctx context.Context, req *Result) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadResult not implemented")
}

func RegisterDistributorServiceServer(s *grpc.Server, srv DistributorServiceServer) {
	s.RegisterService(&_DistributorService_serviceDesc, srv)
}

func _DistributorService_JoinCluster_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClusterClientOffer)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DistributorServiceServer).JoinCluster(m, &distributorServiceJoinClusterServer{stream})
}

type DistributorService_JoinClusterServer interface {
	Send(*Job) error
	grpc.ServerStream
}

type distributorServiceJoinClusterServer struct {
	grpc.ServerStream
}

func (x *distributorServiceJoinClusterServer) Send(m *Job) error {
	return x.ServerStream.SendMsg(m)
}

func _DistributorService_UploadResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Result)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServiceServer).UploadResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistributorService/UploadResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServiceServer).UploadResult(ctx, req.(*Result))
	}
	return interceptor(ctx, in, info, handler)
}

var _DistributorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DistributorService",
	HandlerType: (*DistributorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadResult",
			Handler:    _DistributorService_UploadResult_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinCluster",
			Handler:       _DistributorService_JoinCluster_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "distributor-service.proto",
}

// CliServiceClient is the client API for CliService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CliServiceClient interface {
	Transcode(ctx context.Context, in *TranscodeRequest, opts ...grpc.CallOption) (*TranscodeResponse, error)
	Progress(ctx context.Context, in *ProgressRequest, opts ...grpc.CallOption) (CliService_ProgressClient, error)
}

type cliServiceClient struct {
	cc *grpc.ClientConn
}

func NewCliServiceClient(cc *grpc.ClientConn) CliServiceClient {
	return &cliServiceClient{cc}
}

func (c *cliServiceClient) Transcode(ctx context.Context, in *TranscodeRequest, opts ...grpc.CallOption) (*TranscodeResponse, error) {
	out := new(TranscodeResponse)
	err := c.cc.Invoke(ctx, "/CliService/Transcode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliServiceClient) Progress(ctx context.Context, in *ProgressRequest, opts ...grpc.CallOption) (CliService_ProgressClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CliService_serviceDesc.Streams[0], "/CliService/Progress", opts...)
	if err != nil {
		return nil, err
	}
	x := &cliServiceProgressClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CliService_ProgressClient interface {
	Recv() (*ProgressResponse, error)
	grpc.ClientStream
}

type cliServiceProgressClient struct {
	grpc.ClientStream
}

func (x *cliServiceProgressClient) Recv() (*ProgressResponse, error) {
	m := new(ProgressResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CliServiceServer is the server API for CliService service.
type CliServiceServer interface {
	Transcode(context.Context, *TranscodeRequest) (*TranscodeResponse, error)
	Progress(*ProgressRequest, CliService_ProgressServer) error
}

// UnimplementedCliServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCliServiceServer struct {
}

func (*UnimplementedCliServiceServer) Transcode(ctx context.Context, req *TranscodeRequest) (*TranscodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transcode not implemented")
}
func (*UnimplementedCliServiceServer) Progress(req *ProgressRequest, srv CliService_ProgressServer) error {
	return status.Errorf(codes.Unimplemented, "method Progress not implemented")
}

func RegisterCliServiceServer(s *grpc.Server, srv CliServiceServer) {
	s.RegisterService(&_CliService_serviceDesc, srv)
}

func _CliService_Transcode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranscodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServiceServer).Transcode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CliService/Transcode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServiceServer).Transcode(ctx, req.(*TranscodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliService_Progress_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProgressRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CliServiceServer).Progress(m, &cliServiceProgressServer{stream})
}

type CliService_ProgressServer interface {
	Send(*ProgressResponse) error
	grpc.ServerStream
}

type cliServiceProgressServer struct {
	grpc.ServerStream
}

func (x *cliServiceProgressServer) Send(m *ProgressResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CliService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CliService",
	HandlerType: (*CliServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transcode",
			Handler:    _CliService_Transcode_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Progress",
			Handler:       _CliService_Progress_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "distributor-service.proto",
}
