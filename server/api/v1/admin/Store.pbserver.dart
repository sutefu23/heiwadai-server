///
//  Generated code. Do not modify.
//  source: v1/admin/Store.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'Store.pb.dart' as $4;
import '../shared/Store.pb.dart' as $3;
import '../../google/protobuf/empty.pb.dart' as $1;
import 'Store.pbjson.dart';

export 'Store.pb.dart';

abstract class StoreControllerServiceBase extends $pb.GeneratedService {
  $async.Future<$3.Store> getByID($pb.ServerContext ctx, $4.SoreIDRequest request);
  $async.Future<$3.Stores> getAll($pb.ServerContext ctx, $1.Empty request);
  $async.Future<$1.Empty> register($pb.ServerContext ctx, $4.StoreRegisterRequest request);
  $async.Future<$1.Empty> update($pb.ServerContext ctx, $4.StoreUpdateRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'GetByID': return $4.SoreIDRequest();
      case 'GetAll': return $1.Empty();
      case 'Register': return $4.StoreRegisterRequest();
      case 'Update': return $4.StoreUpdateRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'GetByID': return this.getByID(ctx, request as $4.SoreIDRequest);
      case 'GetAll': return this.getAll(ctx, request as $1.Empty);
      case 'Register': return this.register(ctx, request as $4.StoreRegisterRequest);
      case 'Update': return this.update(ctx, request as $4.StoreUpdateRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => StoreControllerServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => StoreControllerServiceBase$messageJson;
}

