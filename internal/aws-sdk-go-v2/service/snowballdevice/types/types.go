// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

type Alert struct {
	Arn *string

	Description *string

	DisplayName *string

	LastUpdatedTimestamp *time.Time

	Level AlertLevel

	Name *string

	Namespace *string

	State AlertState

	Subscription *AlertSubscription

	noSmithyDocumentSerde
}

type AlertHistory struct {
	Name *string

	Namespace *string

	Notifications []AlertNotification

	State AlertState

	Timestamp *time.Time

	noSmithyDocumentSerde
}

type AlertNotification struct {
	Channel AlertSubscriptionChannel

	Details map[string]string

	Status AlertNotificationStatus

	Timestamp *time.Time

	noSmithyDocumentSerde
}

type AlertSubscription struct {
	Channels []AlertSubscriptionChannel

	State AlertSubscriptionState

	noSmithyDocumentSerde
}

type AutoStartConfiguration struct {
	AutoStartConfigurationArn *string

	Enabled bool

	IpAddressAssignment IpAddressAssignment

	LaunchTemplateId *string

	LaunchTemplateVersion *string

	PhysicalConnectorType PhysicalConnectorType

	StaticIpAddressConfiguration *StaticIpAddressConfiguration

	noSmithyDocumentSerde
}

type AutoStartConfigurationDetails struct {
	AutoStartConfiguration *AutoStartConfiguration

	AutoStartDetails *AutoStartDetails

	noSmithyDocumentSerde
}

type AutoStartDetails struct {
	InstanceId *string

	InterfaceId *string

	State *string

	StateMessage *string

	noSmithyDocumentSerde
}

type AutoUpdateStrategy struct {
	AutoCheck bool

	AutoCheckFrequency *string

	AutoDownload bool

	AutoDownloadFrequency *string

	AutoInstall bool

	AutoInstallFrequency *string

	AutoReboot bool

	noSmithyDocumentSerde
}

type Capacity struct {
	Available *int64

	Name *string

	Total *int64

	Unit *string

	Used *int64

	noSmithyDocumentSerde
}

type CertificateAssociation struct {

	// This member is required.
	CertificateArn *string

	noSmithyDocumentSerde
}

type CertificateSummary struct {

	// This member is required.
	CertificateArn *string

	SubjectAlternativeNames []string

	noSmithyDocumentSerde
}

type ClusterAssociation struct {

	// This member is required.
	ClusterId *string

	// This member is required.
	State ClusterAssociationState

	ClusteredNetworkInterface *PhysicalNetworkInterface

	noSmithyDocumentSerde
}

type Device struct {
	ActiveNetworkInterface *NetworkInterface

	ClusterAssociation *ClusterAssociation

	DeviceId *string

	NetworkReachability *NetworkReachability

	PhysicalNetworkInterfaces []PhysicalNetworkInterface

	Tags []Tag

	UnlockStatus *UnlockStatus

	noSmithyDocumentSerde
}

type DirectNetworkInterface struct {

	// This member is required.
	DirectNetworkInterfaceArn *string

	// This member is required.
	Driver DirectNetworkDriver

	// This member is required.
	MacAddress *string

	// This member is required.
	PhysicalNetworkInterfaceId *string

	InstanceId *string

	VlanId *int32

	noSmithyDocumentSerde
}

type Endpoint struct {

	// This member is required.
	Host *string

	// This member is required.
	Port int32

	// This member is required.
	Protocol *string

	CertificateAssociation *CertificateAssociation

	Description *string

	DeviceId *string

	Status *ServiceStatus

	noSmithyDocumentSerde
}

type NetworkInterface struct {
	IpAddress *string

	noSmithyDocumentSerde
}

type NetworkReachability struct {

	// This member is required.
	State NetworkReachabilityState

	noSmithyDocumentSerde
}

type PciDevice struct {
	DeviceClass *string

	InstanceId *string

	Name *string

	PciDeviceId *string

	Status PciDeviceStatus

	VendorName *string

	noSmithyDocumentSerde
}

type PhysicalNetworkInterface struct {
	DefaultGateway *string

	IpAddress *string

	IpAddressAssignment IpAddressAssignment

	MacAddress *string

	Netmask *string

	PhysicalConnectorType PhysicalConnectorType

	PhysicalNetworkInterfaceId *string

	noSmithyDocumentSerde
}

type ServiceConfiguration struct {
	AllowedHosts []string

	noSmithyDocumentSerde
}

type ServiceStatus struct {
	Details *string

	State ServiceStatusState

	noSmithyDocumentSerde
}

type ServiceStorage struct {

	// This member is required.
	FreeSpaceBytes int64

	// This member is required.
	TotalSpaceBytes int64

	noSmithyDocumentSerde
}

type StaticIpAddressConfiguration struct {

	// This member is required.
	IpAddress *string

	// This member is required.
	Netmask *string

	DefaultGateway *string

	noSmithyDocumentSerde
}

type Tag struct {
	Key *string

	Value *string

	noSmithyDocumentSerde
}

type TimeSourceStatus struct {

	// This member is required.
	Address *string

	// This member is required.
	State TimeSourceState

	// This member is required.
	Stratum int32

	// This member is required.
	Type TimeSourceType

	noSmithyDocumentSerde
}

type UnlockStatus struct {

	// This member is required.
	State UnlockStatusState

	noSmithyDocumentSerde
}

type VirtualNetworkInterface struct {
	DefaultGateway *string

	IpAddress *string

	IpAddressAssignment IpAddressAssignment

	MacAddress *string

	Netmask *string

	PhysicalNetworkInterfaceId *string

	VirtualNetworkInterfaceArn *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
