// Code generated by triggergen. DO NOT EDIT.

package triggers

import (
	"fmt"

	"github.com/juju/juju/core/database/schema"
)


// ChangeLogTriggersForCloud generates the triggers for the
// cloud table.
func ChangeLogTriggersForCloud(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert trigger for Cloud
CREATE TRIGGER trg_log_cloud_insert
AFTER INSERT ON cloud FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for Cloud
CREATE TRIGGER trg_log_cloud_update
AFTER UPDATE ON cloud FOR EACH ROW
WHEN 
	NEW.name != OLD.name OR
	NEW.cloud_type_id != OLD.cloud_type_id OR
	NEW.endpoint != OLD.endpoint OR
	(NEW.identity_endpoint != OLD.identity_endpoint OR (NEW.identity_endpoint IS NOT NULL AND OLD.identity_endpoint IS NULL) OR (NEW.identity_endpoint IS NULL AND OLD.identity_endpoint IS NOT NULL)) OR
	(NEW.storage_endpoint != OLD.storage_endpoint OR (NEW.storage_endpoint IS NOT NULL AND OLD.storage_endpoint IS NULL) OR (NEW.storage_endpoint IS NULL AND OLD.storage_endpoint IS NOT NULL)) OR
	NEW.skip_tls_verify != OLD.skip_tls_verify 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;

-- delete trigger for Cloud
CREATE TRIGGER trg_log_cloud_delete
AFTER DELETE ON cloud FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}

// ChangeLogTriggersForCloudCredential generates the triggers for the
// cloud_credential table.
func ChangeLogTriggersForCloudCredential(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert trigger for CloudCredential
CREATE TRIGGER trg_log_cloud_credential_insert
AFTER INSERT ON cloud_credential FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for CloudCredential
CREATE TRIGGER trg_log_cloud_credential_update
AFTER UPDATE ON cloud_credential FOR EACH ROW
WHEN 
	NEW.cloud_uuid != OLD.cloud_uuid OR
	NEW.auth_type_id != OLD.auth_type_id OR
	NEW.owner_uuid != OLD.owner_uuid OR
	NEW.name != OLD.name OR
	(NEW.revoked != OLD.revoked OR (NEW.revoked IS NOT NULL AND OLD.revoked IS NULL) OR (NEW.revoked IS NULL AND OLD.revoked IS NOT NULL)) OR
	(NEW.invalid != OLD.invalid OR (NEW.invalid IS NOT NULL AND OLD.invalid IS NULL) OR (NEW.invalid IS NULL AND OLD.invalid IS NOT NULL)) OR
	(NEW.invalid_reason != OLD.invalid_reason OR (NEW.invalid_reason IS NOT NULL AND OLD.invalid_reason IS NULL) OR (NEW.invalid_reason IS NULL AND OLD.invalid_reason IS NOT NULL)) 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;

-- delete trigger for CloudCredential
CREATE TRIGGER trg_log_cloud_credential_delete
AFTER DELETE ON cloud_credential FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}

// ChangeLogTriggersForExternalController generates the triggers for the
// external_controller table.
func ChangeLogTriggersForExternalController(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert trigger for ExternalController
CREATE TRIGGER trg_log_external_controller_insert
AFTER INSERT ON external_controller FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for ExternalController
CREATE TRIGGER trg_log_external_controller_update
AFTER UPDATE ON external_controller FOR EACH ROW
WHEN 
	(NEW.alias != OLD.alias OR (NEW.alias IS NOT NULL AND OLD.alias IS NULL) OR (NEW.alias IS NULL AND OLD.alias IS NOT NULL)) OR
	NEW.ca_cert != OLD.ca_cert 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;

-- delete trigger for ExternalController
CREATE TRIGGER trg_log_external_controller_delete
AFTER DELETE ON external_controller FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}

