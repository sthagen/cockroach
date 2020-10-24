// Code generated by "stringer -type=VersionKey"; DO NOT EDIT.

package clusterversion

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Version19_1-0]
	_ = x[VersionAtomicChangeReplicasTrigger-1]
	_ = x[VersionAtomicChangeReplicas-2]
	_ = x[VersionPartitionedBackup-3]
	_ = x[Version19_2-4]
	_ = x[VersionStart20_1-5]
	_ = x[VersionContainsEstimatesCounter-6]
	_ = x[VersionChangeReplicasDemotion-7]
	_ = x[VersionSecondaryIndexColumnFamilies-8]
	_ = x[VersionNamespaceTableWithSchemas-9]
	_ = x[VersionProtectedTimestamps-10]
	_ = x[VersionPrimaryKeyChanges-11]
	_ = x[VersionAuthLocalAndTrustRejectMethods-12]
	_ = x[VersionPrimaryKeyColumnsOutOfFamilyZero-13]
	_ = x[VersionNoExplicitForeignKeyIndexIDs-14]
	_ = x[VersionHashShardedIndexes-15]
	_ = x[VersionCreateRolePrivilege-16]
	_ = x[VersionStatementDiagnosticsSystemTables-17]
	_ = x[VersionSchemaChangeJob-18]
	_ = x[VersionSavepoints-19]
	_ = x[Version20_1-20]
	_ = x[VersionStart20_2-21]
	_ = x[VersionGeospatialType-22]
	_ = x[VersionEnums-23]
	_ = x[VersionRangefeedLeases-24]
	_ = x[VersionAlterColumnTypeGeneral-25]
	_ = x[VersionAlterSystemJobsAddCreatedByColumns-26]
	_ = x[VersionAddScheduledJobsTable-27]
	_ = x[VersionUserDefinedSchemas-28]
	_ = x[VersionNoOriginFKIndexes-29]
	_ = x[VersionClientRangeInfosOnBatchResponse-30]
	_ = x[VersionNodeMembershipStatus-31]
	_ = x[VersionRangeStatsRespHasDesc-32]
	_ = x[VersionMinPasswordLength-33]
	_ = x[VersionAbortSpanBytes-34]
	_ = x[VersionAlterSystemJobsAddSqllivenessColumnsAddNewSystemSqllivenessTable-35]
	_ = x[VersionMaterializedViews-36]
	_ = x[VersionBox2DType-37]
	_ = x[VersionLeasedDatabaseDescriptors-38]
	_ = x[VersionUpdateScheduledJobsSchema-39]
	_ = x[VersionCreateLoginPrivilege-40]
	_ = x[VersionHBAForNonTLS-41]
	_ = x[Version20_2-42]
	_ = x[VersionStart21_1-43]
}

const _VersionKey_name = "Version19_1VersionAtomicChangeReplicasTriggerVersionAtomicChangeReplicasVersionPartitionedBackupVersion19_2VersionStart20_1VersionContainsEstimatesCounterVersionChangeReplicasDemotionVersionSecondaryIndexColumnFamiliesVersionNamespaceTableWithSchemasVersionProtectedTimestampsVersionPrimaryKeyChangesVersionAuthLocalAndTrustRejectMethodsVersionPrimaryKeyColumnsOutOfFamilyZeroVersionNoExplicitForeignKeyIndexIDsVersionHashShardedIndexesVersionCreateRolePrivilegeVersionStatementDiagnosticsSystemTablesVersionSchemaChangeJobVersionSavepointsVersion20_1VersionStart20_2VersionGeospatialTypeVersionEnumsVersionRangefeedLeasesVersionAlterColumnTypeGeneralVersionAlterSystemJobsAddCreatedByColumnsVersionAddScheduledJobsTableVersionUserDefinedSchemasVersionNoOriginFKIndexesVersionClientRangeInfosOnBatchResponseVersionNodeMembershipStatusVersionRangeStatsRespHasDescVersionMinPasswordLengthVersionAbortSpanBytesVersionAlterSystemJobsAddSqllivenessColumnsAddNewSystemSqllivenessTableVersionMaterializedViewsVersionBox2DTypeVersionLeasedDatabaseDescriptorsVersionUpdateScheduledJobsSchemaVersionCreateLoginPrivilegeVersionHBAForNonTLSVersion20_2VersionStart21_1"

var _VersionKey_index = [...]uint16{0, 11, 45, 72, 96, 107, 123, 154, 183, 218, 250, 276, 300, 337, 376, 411, 436, 462, 501, 523, 540, 551, 567, 588, 600, 622, 651, 692, 720, 745, 769, 807, 834, 862, 886, 907, 978, 1002, 1018, 1050, 1082, 1109, 1128, 1139, 1155}

func (i VersionKey) String() string {
	if i < 0 || i >= VersionKey(len(_VersionKey_index)-1) {
		return "VersionKey(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _VersionKey_name[_VersionKey_index[i]:_VersionKey_index[i+1]]
}
