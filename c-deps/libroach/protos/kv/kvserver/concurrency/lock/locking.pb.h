// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: kv/kvserver/concurrency/lock/locking.proto

#ifndef PROTOBUF_INCLUDED_kv_2fkvserver_2fconcurrency_2flock_2flocking_2eproto
#define PROTOBUF_INCLUDED_kv_2fkvserver_2fconcurrency_2flock_2flocking_2eproto

#include <string>

#include <google/protobuf/stubs/common.h>

#if GOOGLE_PROTOBUF_VERSION < 3006000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please update
#error your headers.
#endif
#if 3006000 < GOOGLE_PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata_lite.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/generated_enum_util.h>
// @@protoc_insertion_point(includes)
#define PROTOBUF_INTERNAL_EXPORT_protobuf_kv_2fkvserver_2fconcurrency_2flock_2flocking_2eproto 

namespace protobuf_kv_2fkvserver_2fconcurrency_2flock_2flocking_2eproto {
// Internal implementation detail -- do not use these members.
struct TableStruct {
  static const ::google::protobuf::internal::ParseTableField entries[];
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[];
  static const ::google::protobuf::internal::ParseTable schema[1];
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
}  // namespace protobuf_kv_2fkvserver_2fconcurrency_2flock_2flocking_2eproto
namespace cockroach {
namespace kv {
namespace kvserver {
namespace concurrency {
namespace lock {
}  // namespace lock
}  // namespace concurrency
}  // namespace kvserver
}  // namespace kv
}  // namespace cockroach
namespace cockroach {
namespace kv {
namespace kvserver {
namespace concurrency {
namespace lock {

enum Strength {
  None = 0,
  Shared = 1,
  Upgrade = 2,
  Exclusive = 3,
  Strength_INT_MIN_SENTINEL_DO_NOT_USE_ = ::google::protobuf::kint32min,
  Strength_INT_MAX_SENTINEL_DO_NOT_USE_ = ::google::protobuf::kint32max
};
bool Strength_IsValid(int value);
const Strength Strength_MIN = None;
const Strength Strength_MAX = Exclusive;
const int Strength_ARRAYSIZE = Strength_MAX + 1;

enum Durability {
  Replicated = 0,
  Unreplicated = 1,
  Durability_INT_MIN_SENTINEL_DO_NOT_USE_ = ::google::protobuf::kint32min,
  Durability_INT_MAX_SENTINEL_DO_NOT_USE_ = ::google::protobuf::kint32max
};
bool Durability_IsValid(int value);
const Durability Durability_MIN = Replicated;
const Durability Durability_MAX = Unreplicated;
const int Durability_ARRAYSIZE = Durability_MAX + 1;

// ===================================================================


// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)

}  // namespace lock
}  // namespace concurrency
}  // namespace kvserver
}  // namespace kv
}  // namespace cockroach

namespace google {
namespace protobuf {

template <> struct is_proto_enum< ::cockroach::kv::kvserver::concurrency::lock::Strength> : ::std::true_type {};
template <> struct is_proto_enum< ::cockroach::kv::kvserver::concurrency::lock::Durability> : ::std::true_type {};

}  // namespace protobuf
}  // namespace google

// @@protoc_insertion_point(global_scope)

#endif  // PROTOBUF_INCLUDED_kv_2fkvserver_2fconcurrency_2flock_2flocking_2eproto
