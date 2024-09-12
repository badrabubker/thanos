// Copyright (c) The Thanos Authors.
// Licensed under the Apache License 2.0.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.20.0
// source: store/hintspb/hints.proto

package hintspb

import (
	reflect "reflect"
	sync "sync"

	storepb "github.com/thanos-io/thanos/pkg/store/storepb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SeriesRequestHints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// / block_matchers is a list of label matchers that are evaluated against each single block's
	// / labels to filter which blocks get queried. If the list is empty, no per-block filtering
	// / is applied.
	BlockMatchers    []*storepb.LabelMatcher `protobuf:"bytes,1,rep,name=block_matchers,json=blockMatchers,proto3" json:"block_matchers,omitempty"`
	EnableQueryStats bool                    `protobuf:"varint,2,opt,name=enable_query_stats,json=enableQueryStats,proto3" json:"enable_query_stats,omitempty"`
}

func (x *SeriesRequestHints) Reset() {
	*x = SeriesRequestHints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_hintspb_hints_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeriesRequestHints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeriesRequestHints) ProtoMessage() {}

func (x *SeriesRequestHints) ProtoReflect() protoreflect.Message {
	mi := &file_store_hintspb_hints_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeriesRequestHints.ProtoReflect.Descriptor instead.
func (*SeriesRequestHints) Descriptor() ([]byte, []int) {
	return file_store_hintspb_hints_proto_rawDescGZIP(), []int{0}
}

func (x *SeriesRequestHints) GetBlockMatchers() []*storepb.LabelMatcher {
	if x != nil {
		return x.BlockMatchers
	}
	return nil
}

func (x *SeriesRequestHints) GetEnableQueryStats() bool {
	if x != nil {
		return x.EnableQueryStats
	}
	return false
}

type SeriesResponseHints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// / queried_blocks is the list of blocks that have been queried.
	QueriedBlocks []*Block `protobuf:"bytes,1,rep,name=queried_blocks,json=queriedBlocks,proto3" json:"queried_blocks,omitempty"`
	// / query_stats contains statistics of querying store gateway.
	QueryStats *QueryStats `protobuf:"bytes,2,opt,name=query_stats,json=queryStats,proto3" json:"query_stats,omitempty"`
}

func (x *SeriesResponseHints) Reset() {
	*x = SeriesResponseHints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_hintspb_hints_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeriesResponseHints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeriesResponseHints) ProtoMessage() {}

func (x *SeriesResponseHints) ProtoReflect() protoreflect.Message {
	mi := &file_store_hintspb_hints_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeriesResponseHints.ProtoReflect.Descriptor instead.
func (*SeriesResponseHints) Descriptor() ([]byte, []int) {
	return file_store_hintspb_hints_proto_rawDescGZIP(), []int{1}
}

func (x *SeriesResponseHints) GetQueriedBlocks() []*Block {
	if x != nil {
		return x.QueriedBlocks
	}
	return nil
}

func (x *SeriesResponseHints) GetQueryStats() *QueryStats {
	if x != nil {
		return x.QueryStats
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_hintspb_hints_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_store_hintspb_hints_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_store_hintspb_hints_proto_rawDescGZIP(), []int{2}
}

func (x *Block) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type LabelNamesRequestHints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// / block_matchers is a list of label matchers that are evaluated against each single block's
	// / labels to filter which blocks get queried. If the list is empty, no per-block filtering
	// / is applied.
	BlockMatchers []*storepb.LabelMatcher `protobuf:"bytes,1,rep,name=block_matchers,json=blockMatchers,proto3" json:"block_matchers,omitempty"`
}

func (x *LabelNamesRequestHints) Reset() {
	*x = LabelNamesRequestHints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_hintspb_hints_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelNamesRequestHints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelNamesRequestHints) ProtoMessage() {}

func (x *LabelNamesRequestHints) ProtoReflect() protoreflect.Message {
	mi := &file_store_hintspb_hints_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelNamesRequestHints.ProtoReflect.Descriptor instead.
func (*LabelNamesRequestHints) Descriptor() ([]byte, []int) {
	return file_store_hintspb_hints_proto_rawDescGZIP(), []int{3}
}

func (x *LabelNamesRequestHints) GetBlockMatchers() []*storepb.LabelMatcher {
	if x != nil {
		return x.BlockMatchers
	}
	return nil
}

type LabelNamesResponseHints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// / queried_blocks is the list of blocks that have been queried.
	QueriedBlocks []*Block `protobuf:"bytes,1,rep,name=queried_blocks,json=queriedBlocks,proto3" json:"queried_blocks,omitempty"`
}

func (x *LabelNamesResponseHints) Reset() {
	*x = LabelNamesResponseHints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_hintspb_hints_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelNamesResponseHints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelNamesResponseHints) ProtoMessage() {}

func (x *LabelNamesResponseHints) ProtoReflect() protoreflect.Message {
	mi := &file_store_hintspb_hints_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelNamesResponseHints.ProtoReflect.Descriptor instead.
func (*LabelNamesResponseHints) Descriptor() ([]byte, []int) {
	return file_store_hintspb_hints_proto_rawDescGZIP(), []int{4}
}

func (x *LabelNamesResponseHints) GetQueriedBlocks() []*Block {
	if x != nil {
		return x.QueriedBlocks
	}
	return nil
}

type LabelValuesRequestHints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// / block_matchers is a list of label matchers that are evaluated against each single block's
	// / labels to filter which blocks get queried. If the list is empty, no per-block filtering
	// / is applied.
	BlockMatchers []*storepb.LabelMatcher `protobuf:"bytes,1,rep,name=block_matchers,json=blockMatchers,proto3" json:"block_matchers,omitempty"`
}

func (x *LabelValuesRequestHints) Reset() {
	*x = LabelValuesRequestHints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_hintspb_hints_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelValuesRequestHints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelValuesRequestHints) ProtoMessage() {}

func (x *LabelValuesRequestHints) ProtoReflect() protoreflect.Message {
	mi := &file_store_hintspb_hints_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelValuesRequestHints.ProtoReflect.Descriptor instead.
func (*LabelValuesRequestHints) Descriptor() ([]byte, []int) {
	return file_store_hintspb_hints_proto_rawDescGZIP(), []int{5}
}

func (x *LabelValuesRequestHints) GetBlockMatchers() []*storepb.LabelMatcher {
	if x != nil {
		return x.BlockMatchers
	}
	return nil
}

type LabelValuesResponseHints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// / queried_blocks is the list of blocks that have been queried.
	QueriedBlocks []*Block `protobuf:"bytes,1,rep,name=queried_blocks,json=queriedBlocks,proto3" json:"queried_blocks,omitempty"`
}

func (x *LabelValuesResponseHints) Reset() {
	*x = LabelValuesResponseHints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_hintspb_hints_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelValuesResponseHints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelValuesResponseHints) ProtoMessage() {}

func (x *LabelValuesResponseHints) ProtoReflect() protoreflect.Message {
	mi := &file_store_hintspb_hints_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelValuesResponseHints.ProtoReflect.Descriptor instead.
func (*LabelValuesResponseHints) Descriptor() ([]byte, []int) {
	return file_store_hintspb_hints_proto_rawDescGZIP(), []int{6}
}

func (x *LabelValuesResponseHints) GetQueriedBlocks() []*Block {
	if x != nil {
		return x.QueriedBlocks
	}
	return nil
}

// / QueryStats fields are unstable and might change in the future.
type QueryStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlocksQueried          int64                `protobuf:"varint,1,opt,name=blocks_queried,json=blocksQueried,proto3" json:"blocks_queried,omitempty"`
	MergedSeriesCount      int64                `protobuf:"varint,2,opt,name=merged_series_count,json=mergedSeriesCount,proto3" json:"merged_series_count,omitempty"`
	MergedChunksCount      int64                `protobuf:"varint,3,opt,name=merged_chunks_count,json=mergedChunksCount,proto3" json:"merged_chunks_count,omitempty"`
	PostingsTouched        int64                `protobuf:"varint,4,opt,name=postings_touched,json=postingsTouched,proto3" json:"postings_touched,omitempty"`
	PostingsTouchedSizeSum int64                `protobuf:"varint,5,opt,name=postings_touched_size_sum,json=postingsTouchedSizeSum,proto3" json:"postings_touched_size_sum,omitempty"`
	PostingsToFetch        int64                `protobuf:"varint,6,opt,name=postings_to_fetch,json=postingsToFetch,proto3" json:"postings_to_fetch,omitempty"`
	PostingsFetched        int64                `protobuf:"varint,7,opt,name=postings_fetched,json=postingsFetched,proto3" json:"postings_fetched,omitempty"`
	PostingsFetchedSizeSum int64                `protobuf:"varint,8,opt,name=postings_fetched_size_sum,json=postingsFetchedSizeSum,proto3" json:"postings_fetched_size_sum,omitempty"`
	PostingsFetchCount     int64                `protobuf:"varint,9,opt,name=postings_fetch_count,json=postingsFetchCount,proto3" json:"postings_fetch_count,omitempty"`
	SeriesTouched          int64                `protobuf:"varint,10,opt,name=series_touched,json=seriesTouched,proto3" json:"series_touched,omitempty"`
	SeriesTouchedSizeSum   int64                `protobuf:"varint,11,opt,name=series_touched_size_sum,json=seriesTouchedSizeSum,proto3" json:"series_touched_size_sum,omitempty"`
	SeriesFetched          int64                `protobuf:"varint,12,opt,name=series_fetched,json=seriesFetched,proto3" json:"series_fetched,omitempty"`
	SeriesFetchedSizeSum   int64                `protobuf:"varint,13,opt,name=series_fetched_size_sum,json=seriesFetchedSizeSum,proto3" json:"series_fetched_size_sum,omitempty"`
	SeriesFetchCount       int64                `protobuf:"varint,14,opt,name=series_fetch_count,json=seriesFetchCount,proto3" json:"series_fetch_count,omitempty"`
	ChunksTouched          int64                `protobuf:"varint,15,opt,name=chunks_touched,json=chunksTouched,proto3" json:"chunks_touched,omitempty"`
	ChunksTouchedSizeSum   int64                `protobuf:"varint,16,opt,name=chunks_touched_size_sum,json=chunksTouchedSizeSum,proto3" json:"chunks_touched_size_sum,omitempty"`
	ChunksFetched          int64                `protobuf:"varint,17,opt,name=chunks_fetched,json=chunksFetched,proto3" json:"chunks_fetched,omitempty"`
	ChunksFetchedSizeSum   int64                `protobuf:"varint,18,opt,name=chunks_fetched_size_sum,json=chunksFetchedSizeSum,proto3" json:"chunks_fetched_size_sum,omitempty"`
	ChunksFetchCount       int64                `protobuf:"varint,19,opt,name=chunks_fetch_count,json=chunksFetchCount,proto3" json:"chunks_fetch_count,omitempty"`
	DataDownloadedSizeSum  int64                `protobuf:"varint,20,opt,name=data_downloaded_size_sum,json=dataDownloadedSizeSum,proto3" json:"data_downloaded_size_sum,omitempty"`
	GetAllDuration         *durationpb.Duration `protobuf:"bytes,21,opt,name=get_all_duration,json=getAllDuration,proto3" json:"get_all_duration,omitempty"`
	MergeDuration          *durationpb.Duration `protobuf:"bytes,22,opt,name=merge_duration,json=mergeDuration,proto3" json:"merge_duration,omitempty"`
}

func (x *QueryStats) Reset() {
	*x = QueryStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_hintspb_hints_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryStats) ProtoMessage() {}

func (x *QueryStats) ProtoReflect() protoreflect.Message {
	mi := &file_store_hintspb_hints_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryStats.ProtoReflect.Descriptor instead.
func (*QueryStats) Descriptor() ([]byte, []int) {
	return file_store_hintspb_hints_proto_rawDescGZIP(), []int{7}
}

func (x *QueryStats) GetBlocksQueried() int64 {
	if x != nil {
		return x.BlocksQueried
	}
	return 0
}

func (x *QueryStats) GetMergedSeriesCount() int64 {
	if x != nil {
		return x.MergedSeriesCount
	}
	return 0
}

func (x *QueryStats) GetMergedChunksCount() int64 {
	if x != nil {
		return x.MergedChunksCount
	}
	return 0
}

func (x *QueryStats) GetPostingsTouched() int64 {
	if x != nil {
		return x.PostingsTouched
	}
	return 0
}

func (x *QueryStats) GetPostingsTouchedSizeSum() int64 {
	if x != nil {
		return x.PostingsTouchedSizeSum
	}
	return 0
}

func (x *QueryStats) GetPostingsToFetch() int64 {
	if x != nil {
		return x.PostingsToFetch
	}
	return 0
}

func (x *QueryStats) GetPostingsFetched() int64 {
	if x != nil {
		return x.PostingsFetched
	}
	return 0
}

func (x *QueryStats) GetPostingsFetchedSizeSum() int64 {
	if x != nil {
		return x.PostingsFetchedSizeSum
	}
	return 0
}

func (x *QueryStats) GetPostingsFetchCount() int64 {
	if x != nil {
		return x.PostingsFetchCount
	}
	return 0
}

func (x *QueryStats) GetSeriesTouched() int64 {
	if x != nil {
		return x.SeriesTouched
	}
	return 0
}

func (x *QueryStats) GetSeriesTouchedSizeSum() int64 {
	if x != nil {
		return x.SeriesTouchedSizeSum
	}
	return 0
}

func (x *QueryStats) GetSeriesFetched() int64 {
	if x != nil {
		return x.SeriesFetched
	}
	return 0
}

func (x *QueryStats) GetSeriesFetchedSizeSum() int64 {
	if x != nil {
		return x.SeriesFetchedSizeSum
	}
	return 0
}

func (x *QueryStats) GetSeriesFetchCount() int64 {
	if x != nil {
		return x.SeriesFetchCount
	}
	return 0
}

func (x *QueryStats) GetChunksTouched() int64 {
	if x != nil {
		return x.ChunksTouched
	}
	return 0
}

func (x *QueryStats) GetChunksTouchedSizeSum() int64 {
	if x != nil {
		return x.ChunksTouchedSizeSum
	}
	return 0
}

func (x *QueryStats) GetChunksFetched() int64 {
	if x != nil {
		return x.ChunksFetched
	}
	return 0
}

func (x *QueryStats) GetChunksFetchedSizeSum() int64 {
	if x != nil {
		return x.ChunksFetchedSizeSum
	}
	return 0
}

func (x *QueryStats) GetChunksFetchCount() int64 {
	if x != nil {
		return x.ChunksFetchCount
	}
	return 0
}

func (x *QueryStats) GetDataDownloadedSizeSum() int64 {
	if x != nil {
		return x.DataDownloadedSizeSum
	}
	return 0
}

func (x *QueryStats) GetGetAllDuration() *durationpb.Duration {
	if x != nil {
		return x.GetAllDuration
	}
	return nil
}

func (x *QueryStats) GetMergeDuration() *durationpb.Duration {
	if x != nil {
		return x.MergeDuration
	}
	return nil
}

var File_store_hintspb_hints_proto protoreflect.FileDescriptor

var file_store_hintspb_hints_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x69, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2f,
	0x68, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x68, 0x69, 0x6e,
	0x74, 0x73, 0x70, 0x62, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x70, 0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7f, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x22, 0x82, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x35, 0x0a, 0x0e, 0x71, 0x75, 0x65, 0x72,
	0x69, 0x65, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x68, 0x69, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12,
	0x34, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x69, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x22, 0x17, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x55,
	0x0a, 0x16, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x73, 0x22, 0x50, 0x0a, 0x17, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x69, 0x6e, 0x74, 0x73,
	0x12, 0x35, 0x0a, 0x0e, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x68, 0x69, 0x6e, 0x74, 0x73,
	0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65,
	0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x56, 0x0a, 0x17, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x3b, 0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x68, 0x61,
	0x6e, 0x6f, 0x73, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x22,
	0x51, 0x0a, 0x18, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x35, 0x0a, 0x0e, 0x71,
	0x75, 0x65, 0x72, 0x69, 0x65, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x68, 0x69, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x22, 0xd1, 0x08, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x5f, 0x71, 0x75, 0x65, 0x72,
	0x69, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x65, 0x72, 0x67,
	0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x64, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x65, 0x72, 0x67,
	0x65, 0x64, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x64, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x6f, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x5f, 0x74, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x54, 0x6f, 0x75, 0x63,
	0x68, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x19, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5f,
	0x74, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x73, 0x75, 0x6d,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x54, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x53, 0x75, 0x6d, 0x12, 0x2a,
	0x0a, 0x11, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x66, 0x65,
	0x74, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x70, 0x6f, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x54, 0x6f, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x6f,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x19, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x73,
	0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x53, 0x75, 0x6d,
	0x12, 0x30, 0x0a, 0x14, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x66, 0x65, 0x74,
	0x63, 0x68, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12,
	0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x75,
	0x63, 0x68, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x54, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x17, 0x73, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x5f, 0x73, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x73, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x54, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x53, 0x75, 0x6d,
	0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68,
	0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x17, 0x73, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x73,
	0x75, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x53, 0x75, 0x6d, 0x12, 0x2c,
	0x0a, 0x12, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x73, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x5f, 0x74, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x64, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x54, 0x6f, 0x75, 0x63,
	0x68, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x17, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x5f, 0x74, 0x6f,
	0x75, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x73, 0x75, 0x6d, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x54, 0x6f, 0x75, 0x63,
	0x68, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x53, 0x75, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x73, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x46, 0x65, 0x74, 0x63, 0x68, 0x65,
	0x64, 0x12, 0x35, 0x0a, 0x17, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x5f, 0x66, 0x65, 0x74, 0x63,
	0x68, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x73, 0x75, 0x6d, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x14, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x46, 0x65, 0x74, 0x63, 0x68, 0x65,
	0x64, 0x53, 0x69, 0x7a, 0x65, 0x53, 0x75, 0x6d, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x73, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x18, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x73,
	0x75, 0x6d, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x64, 0x61, 0x74, 0x61, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x53, 0x75, 0x6d, 0x12,
	0x43, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0e, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x5f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2d, 0x69, 0x6f, 0x2f, 0x74,
	0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x68, 0x69, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_hintspb_hints_proto_rawDescOnce sync.Once
	file_store_hintspb_hints_proto_rawDescData = file_store_hintspb_hints_proto_rawDesc
)

func file_store_hintspb_hints_proto_rawDescGZIP() []byte {
	file_store_hintspb_hints_proto_rawDescOnce.Do(func() {
		file_store_hintspb_hints_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_hintspb_hints_proto_rawDescData)
	})
	return file_store_hintspb_hints_proto_rawDescData
}

var file_store_hintspb_hints_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_store_hintspb_hints_proto_goTypes = []any{
	(*SeriesRequestHints)(nil),       // 0: hintspb.SeriesRequestHints
	(*SeriesResponseHints)(nil),      // 1: hintspb.SeriesResponseHints
	(*Block)(nil),                    // 2: hintspb.Block
	(*LabelNamesRequestHints)(nil),   // 3: hintspb.LabelNamesRequestHints
	(*LabelNamesResponseHints)(nil),  // 4: hintspb.LabelNamesResponseHints
	(*LabelValuesRequestHints)(nil),  // 5: hintspb.LabelValuesRequestHints
	(*LabelValuesResponseHints)(nil), // 6: hintspb.LabelValuesResponseHints
	(*QueryStats)(nil),               // 7: hintspb.QueryStats
	(*storepb.LabelMatcher)(nil),     // 8: thanos.LabelMatcher
	(*durationpb.Duration)(nil),      // 9: google.protobuf.Duration
}
var file_store_hintspb_hints_proto_depIdxs = []int32{
	8, // 0: hintspb.SeriesRequestHints.block_matchers:type_name -> thanos.LabelMatcher
	2, // 1: hintspb.SeriesResponseHints.queried_blocks:type_name -> hintspb.Block
	7, // 2: hintspb.SeriesResponseHints.query_stats:type_name -> hintspb.QueryStats
	8, // 3: hintspb.LabelNamesRequestHints.block_matchers:type_name -> thanos.LabelMatcher
	2, // 4: hintspb.LabelNamesResponseHints.queried_blocks:type_name -> hintspb.Block
	8, // 5: hintspb.LabelValuesRequestHints.block_matchers:type_name -> thanos.LabelMatcher
	2, // 6: hintspb.LabelValuesResponseHints.queried_blocks:type_name -> hintspb.Block
	9, // 7: hintspb.QueryStats.get_all_duration:type_name -> google.protobuf.Duration
	9, // 8: hintspb.QueryStats.merge_duration:type_name -> google.protobuf.Duration
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_store_hintspb_hints_proto_init() }
func file_store_hintspb_hints_proto_init() {
	if File_store_hintspb_hints_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_hintspb_hints_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SeriesRequestHints); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_store_hintspb_hints_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SeriesResponseHints); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_store_hintspb_hints_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Block); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_store_hintspb_hints_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*LabelNamesRequestHints); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_store_hintspb_hints_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*LabelNamesResponseHints); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_store_hintspb_hints_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*LabelValuesRequestHints); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_store_hintspb_hints_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*LabelValuesResponseHints); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_store_hintspb_hints_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*QueryStats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_hintspb_hints_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_hintspb_hints_proto_goTypes,
		DependencyIndexes: file_store_hintspb_hints_proto_depIdxs,
		MessageInfos:      file_store_hintspb_hints_proto_msgTypes,
	}.Build()
	File_store_hintspb_hints_proto = out.File
	file_store_hintspb_hints_proto_rawDesc = nil
	file_store_hintspb_hints_proto_goTypes = nil
	file_store_hintspb_hints_proto_depIdxs = nil
}
