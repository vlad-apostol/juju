// Code generated by triggergen. DO NOT EDIT.

package triggers

import (
	"fmt"

	"github.com/juju/juju/core/database/schema"
)


// ChangeLogTriggersForBlockDevice generates the triggers for the
// block_device table.
func ChangeLogTriggersForBlockDevice(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert trigger for BlockDevice
CREATE TRIGGER trg_log_block_device_insert
AFTER INSERT ON block_device FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for BlockDevice
CREATE TRIGGER trg_log_block_device_update
AFTER UPDATE ON block_device FOR EACH ROW
WHEN 
	NEW.machine_uuid != OLD.machine_uuid OR
	NEW.name != OLD.name OR
	(NEW.label != OLD.label OR (NEW.label IS NOT NULL AND OLD.label IS NULL) OR (NEW.label IS NULL AND OLD.label IS NOT NULL)) OR
	(NEW.device_uuid != OLD.device_uuid OR (NEW.device_uuid IS NOT NULL AND OLD.device_uuid IS NULL) OR (NEW.device_uuid IS NULL AND OLD.device_uuid IS NOT NULL)) OR
	(NEW.hardware_id != OLD.hardware_id OR (NEW.hardware_id IS NOT NULL AND OLD.hardware_id IS NULL) OR (NEW.hardware_id IS NULL AND OLD.hardware_id IS NOT NULL)) OR
	(NEW.wwn != OLD.wwn OR (NEW.wwn IS NOT NULL AND OLD.wwn IS NULL) OR (NEW.wwn IS NULL AND OLD.wwn IS NOT NULL)) OR
	(NEW.bus_address != OLD.bus_address OR (NEW.bus_address IS NOT NULL AND OLD.bus_address IS NULL) OR (NEW.bus_address IS NULL AND OLD.bus_address IS NOT NULL)) OR
	(NEW.serial_id != OLD.serial_id OR (NEW.serial_id IS NOT NULL AND OLD.serial_id IS NULL) OR (NEW.serial_id IS NULL AND OLD.serial_id IS NOT NULL)) OR
	(NEW.filesystem_type_id != OLD.filesystem_type_id OR (NEW.filesystem_type_id IS NOT NULL AND OLD.filesystem_type_id IS NULL) OR (NEW.filesystem_type_id IS NULL AND OLD.filesystem_type_id IS NOT NULL)) OR
	(NEW.size_mib != OLD.size_mib OR (NEW.size_mib IS NOT NULL AND OLD.size_mib IS NULL) OR (NEW.size_mib IS NULL AND OLD.size_mib IS NOT NULL)) OR
	(NEW.mount_point != OLD.mount_point OR (NEW.mount_point IS NOT NULL AND OLD.mount_point IS NULL) OR (NEW.mount_point IS NULL AND OLD.mount_point IS NOT NULL)) OR
	(NEW.in_use != OLD.in_use OR (NEW.in_use IS NOT NULL AND OLD.in_use IS NULL) OR (NEW.in_use IS NULL AND OLD.in_use IS NOT NULL)) 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;

-- delete trigger for BlockDevice
CREATE TRIGGER trg_log_block_device_delete
AFTER DELETE ON block_device FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}

// ChangeLogTriggersForStorageAttachment generates the triggers for the
// storage_attachment table.
func ChangeLogTriggersForStorageAttachment(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert trigger for StorageAttachment
CREATE TRIGGER trg_log_storage_attachment_insert
AFTER INSERT ON storage_attachment FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for StorageAttachment
CREATE TRIGGER trg_log_storage_attachment_update
AFTER UPDATE ON storage_attachment FOR EACH ROW
WHEN 
	NEW.unit_uuid != OLD.unit_uuid OR
	NEW.life_id != OLD.life_id 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;

-- delete trigger for StorageAttachment
CREATE TRIGGER trg_log_storage_attachment_delete
AFTER DELETE ON storage_attachment FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}

// ChangeLogTriggersForStorageFilesystem generates the triggers for the
// storage_filesystem table.
func ChangeLogTriggersForStorageFilesystem(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert trigger for StorageFilesystem
CREATE TRIGGER trg_log_storage_filesystem_insert
AFTER INSERT ON storage_filesystem FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for StorageFilesystem
CREATE TRIGGER trg_log_storage_filesystem_update
AFTER UPDATE ON storage_filesystem FOR EACH ROW
WHEN 
	NEW.life_id != OLD.life_id OR
	(NEW.provider_id != OLD.provider_id OR (NEW.provider_id IS NOT NULL AND OLD.provider_id IS NULL) OR (NEW.provider_id IS NULL AND OLD.provider_id IS NOT NULL)) OR
	(NEW.storage_pool_uuid != OLD.storage_pool_uuid OR (NEW.storage_pool_uuid IS NOT NULL AND OLD.storage_pool_uuid IS NULL) OR (NEW.storage_pool_uuid IS NULL AND OLD.storage_pool_uuid IS NOT NULL)) OR
	(NEW.size_mib != OLD.size_mib OR (NEW.size_mib IS NOT NULL AND OLD.size_mib IS NULL) OR (NEW.size_mib IS NULL AND OLD.size_mib IS NOT NULL)) OR
	NEW.provisioning_status_id != OLD.provisioning_status_id 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;

-- delete trigger for StorageFilesystem
CREATE TRIGGER trg_log_storage_filesystem_delete
AFTER DELETE ON storage_filesystem FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}

// ChangeLogTriggersForStorageFilesystemAttachment generates the triggers for the
// storage_filesystem_attachment table.
func ChangeLogTriggersForStorageFilesystemAttachment(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert trigger for StorageFilesystemAttachment
CREATE TRIGGER trg_log_storage_filesystem_attachment_insert
AFTER INSERT ON storage_filesystem_attachment FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for StorageFilesystemAttachment
CREATE TRIGGER trg_log_storage_filesystem_attachment_update
AFTER UPDATE ON storage_filesystem_attachment FOR EACH ROW
WHEN 
	NEW.storage_filesystem_uuid != OLD.storage_filesystem_uuid OR
	NEW.net_node_uuid != OLD.net_node_uuid OR
	NEW.life_id != OLD.life_id OR
	(NEW.mount_point != OLD.mount_point OR (NEW.mount_point IS NOT NULL AND OLD.mount_point IS NULL) OR (NEW.mount_point IS NULL AND OLD.mount_point IS NOT NULL)) OR
	(NEW.read_only != OLD.read_only OR (NEW.read_only IS NOT NULL AND OLD.read_only IS NULL) OR (NEW.read_only IS NULL AND OLD.read_only IS NOT NULL)) OR
	NEW.provisioning_status_id != OLD.provisioning_status_id 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;

-- delete trigger for StorageFilesystemAttachment
CREATE TRIGGER trg_log_storage_filesystem_attachment_delete
AFTER DELETE ON storage_filesystem_attachment FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}

// ChangeLogTriggersForStorageVolume generates the triggers for the
// storage_volume table.
func ChangeLogTriggersForStorageVolume(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert trigger for StorageVolume
CREATE TRIGGER trg_log_storage_volume_insert
AFTER INSERT ON storage_volume FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for StorageVolume
CREATE TRIGGER trg_log_storage_volume_update
AFTER UPDATE ON storage_volume FOR EACH ROW
WHEN 
	NEW.life_id != OLD.life_id OR
	NEW.name != OLD.name OR
	(NEW.provider_id != OLD.provider_id OR (NEW.provider_id IS NOT NULL AND OLD.provider_id IS NULL) OR (NEW.provider_id IS NULL AND OLD.provider_id IS NOT NULL)) OR
	(NEW.storage_pool_uuid != OLD.storage_pool_uuid OR (NEW.storage_pool_uuid IS NOT NULL AND OLD.storage_pool_uuid IS NULL) OR (NEW.storage_pool_uuid IS NULL AND OLD.storage_pool_uuid IS NOT NULL)) OR
	(NEW.size_mib != OLD.size_mib OR (NEW.size_mib IS NOT NULL AND OLD.size_mib IS NULL) OR (NEW.size_mib IS NULL AND OLD.size_mib IS NOT NULL)) OR
	(NEW.hardware_id != OLD.hardware_id OR (NEW.hardware_id IS NOT NULL AND OLD.hardware_id IS NULL) OR (NEW.hardware_id IS NULL AND OLD.hardware_id IS NOT NULL)) OR
	(NEW.wwn != OLD.wwn OR (NEW.wwn IS NOT NULL AND OLD.wwn IS NULL) OR (NEW.wwn IS NULL AND OLD.wwn IS NOT NULL)) OR
	(NEW.persistent != OLD.persistent OR (NEW.persistent IS NOT NULL AND OLD.persistent IS NULL) OR (NEW.persistent IS NULL AND OLD.persistent IS NOT NULL)) OR
	NEW.provisioning_status_id != OLD.provisioning_status_id 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;

-- delete trigger for StorageVolume
CREATE TRIGGER trg_log_storage_volume_delete
AFTER DELETE ON storage_volume FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}

// ChangeLogTriggersForStorageVolumeAttachment generates the triggers for the
// storage_volume_attachment table.
func ChangeLogTriggersForStorageVolumeAttachment(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert trigger for StorageVolumeAttachment
CREATE TRIGGER trg_log_storage_volume_attachment_insert
AFTER INSERT ON storage_volume_attachment FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for StorageVolumeAttachment
CREATE TRIGGER trg_log_storage_volume_attachment_update
AFTER UPDATE ON storage_volume_attachment FOR EACH ROW
WHEN 
	NEW.storage_volume_uuid != OLD.storage_volume_uuid OR
	NEW.net_node_uuid != OLD.net_node_uuid OR
	NEW.life_id != OLD.life_id OR
	(NEW.block_device_uuid != OLD.block_device_uuid OR (NEW.block_device_uuid IS NOT NULL AND OLD.block_device_uuid IS NULL) OR (NEW.block_device_uuid IS NULL AND OLD.block_device_uuid IS NOT NULL)) OR
	(NEW.read_only != OLD.read_only OR (NEW.read_only IS NOT NULL AND OLD.read_only IS NULL) OR (NEW.read_only IS NULL AND OLD.read_only IS NOT NULL)) OR
	NEW.provisioning_status_id != OLD.provisioning_status_id 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;

-- delete trigger for StorageVolumeAttachment
CREATE TRIGGER trg_log_storage_volume_attachment_delete
AFTER DELETE ON storage_volume_attachment FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}

// ChangeLogTriggersForStorageVolumeAttachmentPlan generates the triggers for the
// storage_volume_attachment_plan table.
func ChangeLogTriggersForStorageVolumeAttachmentPlan(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert trigger for StorageVolumeAttachmentPlan
CREATE TRIGGER trg_log_storage_volume_attachment_plan_insert
AFTER INSERT ON storage_volume_attachment_plan FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for StorageVolumeAttachmentPlan
CREATE TRIGGER trg_log_storage_volume_attachment_plan_update
AFTER UPDATE ON storage_volume_attachment_plan FOR EACH ROW
WHEN 
	NEW.storage_volume_uuid != OLD.storage_volume_uuid OR
	NEW.net_node_uuid != OLD.net_node_uuid OR
	NEW.life_id != OLD.life_id OR
	(NEW.device_type_id != OLD.device_type_id OR (NEW.device_type_id IS NOT NULL AND OLD.device_type_id IS NULL) OR (NEW.device_type_id IS NULL AND OLD.device_type_id IS NOT NULL)) OR
	(NEW.block_device_uuid != OLD.block_device_uuid OR (NEW.block_device_uuid IS NOT NULL AND OLD.block_device_uuid IS NULL) OR (NEW.block_device_uuid IS NULL AND OLD.block_device_uuid IS NOT NULL)) OR
	NEW.provisioning_status_id != OLD.provisioning_status_id 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;

-- delete trigger for StorageVolumeAttachmentPlan
CREATE TRIGGER trg_log_storage_volume_attachment_plan_delete
AFTER DELETE ON storage_volume_attachment_plan FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}

