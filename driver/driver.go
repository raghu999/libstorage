package driver

import (
	"github.com/akutz/gofig"
	"golang.org/x/net/context"

	"github.com/emccode/libstorage/model"
)

// NewDriver is a function that constructs a new driver.
type NewDriver func(c *gofig.Config) Driver

// Driver represents a libStorage driver.
type Driver interface {
	// The name of the driver.
	Name() string

	// Init initializes the driver.
	Init() error

	// GetVolumeMapping lists the block devices that are attached to the
	GetVolumeMapping(ctx context.Context) ([]*model.BlockDevice, error)

	// GetInstance retrieves the local instance.
	GetInstance(ctx context.Context) (*model.Instance, error)

	// GetVolume returns all volumes for the instance based on either volumeID
	// or volumeName that are available to the instance.
	GetVolume(
		ctx context.Context,
		volumeID,
		volumeName string) ([]*model.Volume, error)

	// GetVolumeAttach returns the attachment details based on volumeID or
	// volumeName where the volume is currently attached.
	GetVolumeAttach(
		ctx context.Context,
		volumeID string) ([]*model.VolumeAttachment, error)

	// CreateSnapshot is a synch/async operation that returns snapshots that
	// have been performed based on supplying a snapshotName, source volumeID,
	// and optional description.
	CreateSnapshot(
		ctx context.Context,
		snapshotName,
		volumeID,
		description string) ([]*model.Snapshot, error)

	// GetSnapshot returns a list of snapshots for a volume based on volumeID,
	// snapshotID, or snapshotName.
	GetSnapshot(
		ctx context.Context,
		volumeID,
		snapshotID,
		snapshotName string) ([]*model.Snapshot, error)

	// RemoveSnapshot will remove a snapshot based on the snapshotID.
	RemoveSnapshot(ctx context.Context, snapshotID string) error

	// CreateVolume is sync/async and will create an return a new/existing
	// Volume based on volumeID/snapshotID with a name of volumeName and a size
	// in GB.  Optionally based on the storage driver, a volumeType, IOPS, and
	// availabilityZone could be defined.
	CreateVolume(
		ctx context.Context,
		volumeName,
		volumeID,
		snapshotID,
		volumeType string,
		IOPS,
		size int64,
		availabilityZone string) (*model.Volume, error)

	// RemoveVolume will remove a volume based on volumeID.
	RemoveVolume(ctx context.Context, volumeID string) error

	// GetDeviceNextAvailable return a device path that will retrieve the next
	// available disk device that can be used.
	GetDeviceNextAvailable() (string, error)

	// AttachVolume returns a list of VolumeAttachments is sync/async that will
	// attach a volume to an instance based on volumeID and ctx.
	AttachVolume(
		ctx context.Context,
		nextDeviceName,
		volumeID string) ([]*model.VolumeAttachment, error)

	// DetachVolume is sync/async that will detach the volumeID from the local
	// instance or the ctx.
	DetachVolume(
		ctx context.Context,
		volumeID string) error

	// CopySnapshot is a sync/async and returns a snapshot that will copy a
	// snapshot based on volumeID/snapshotID/snapshotName and create a new
	// snapshot of desinationSnapshotName in the destinationRegion location.
	CopySnapshot(
		ctx context.Context,
		volumeID,
		snapshotID,
		snapshotName,
		destinationSnapshotName,
		destinationRegion string) (*model.Snapshot, error)

	// GetClientToolName gets the file name of the tool this driver provides
	// to be executed on the client-side in order to discover a client's
	// instance ID and next, available device name.
	//
	// Use the function GetClientTool to get the actual tool.
	GetClientToolName(ctx context.Context) (string, error)

	// GetClientTool gets the file  for the tool this driver provides
	// to be executed on the client-side in order to discover a client's
	// instance ID and next, available device name.
	//
	// This function returns a byte array that will be either a binary file
	// or a unicode-encoded, plain-text script file. Use the file extension
	// of the client tool's file name to determine the file type.
	//
	// The function GetClientToolName can be used to get the file name.
	GetClientTool(ctx context.Context) ([]byte, error)
}