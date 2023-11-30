// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This file is auto-generated from our dictionaries.

package avp

// Diameter AVP types.
const (
	ADCRuleBaseName                            = 1095
	AFChargingIdentifier                       = 505
	AFCorrelationInformation                   = 1276
	AMBR                                       = 1435
	ANGWAddress                                = 1050
	ANID                                       = 1504
	APNAggregateMaxBitrateDL                   = 1040
	APNAggregateMaxBitrateUL                   = 1041
	APNConfiguration                           = 1430
	APNConfigurationProfile                    = 1429
	APNOIReplacement                           = 1427
	ARAPChallengeResponse                      = 84
	ARAPFeatures                               = 71
	ARAPPassword                               = 70
	ARAPSecurity                               = 73
	ARAPSecurityData                           = 74
	ARAPZoneAccess                             = 72
	AUTN                                       = 1449
	AccessNetworkChargingAddress               = 501
	AccessNetworkChargingIdentifierGx          = 1022
	AccessNetworkChargingIdentifierValue       = 503
	AccessNetworkInformation                   = 1263
	AccessRestrictionData                      = 1426
	AccessTransferInformation                  = 2709
	AccessTransferType                         = 2710
	AccountExpiration                          = 2309
	AccountingAuthMethod                       = 406
	AccountingInputOctets                      = 363
	AccountingInputPackets                     = 365
	AccountingOutputOctets                     = 364
	AccountingOutputPackets                    = 366
	AccountingRealtimeRequired                 = 483
	AccountingRecordNumber                     = 485
	AccountingRecordType                       = 480
	AccountingSessionID                        = 44
	AccountingSubSessionID                     = 287
	AcctApplicationID                          = 259
	AcctAuthentic                              = 45
	AcctDelayTime                              = 41
	AcctInterimInterval                        = 85
	AcctLinkCount                              = 51
	AcctMultiSessionID                         = 50
	AcctSessionTime                            = 46
	AcctTunnelConnection                       = 68
	AcctTunnelPacketsLost                      = 86
	AccumulatedCost                            = 2052
	ActiveAPN                                  = 1612
	Adaptations                                = 1217
	AdditionalContentInformation               = 1207
	AdditionalTypeInformation                  = 1205
	AddressData                                = 897
	AddressDomain                              = 898
	AddressType                                = 899
	AddresseeType                              = 1208
	AllAPNConfigurationsIncludedIndicator      = 1428
	AllocationRetentionPriority                = 1034
	AllowedAwfWwsfIdentities                   = 656
	AlternateChargedPartyAddress               = 1280
	AoCCostInformation                         = 2053
	AoCFormat                                  = 2310
	AoCInformation                             = 2054
	AoCRequestType                             = 2055
	AoCService                                 = 2311
	AoCServiceObligatoryType                   = 2312
	AoCServiceType                             = 2313
	AoCSubscriptionInformation                 = 2314
	ApplicID                                   = 1218
	ApplicationPortIdentifer                   = 3010
	ApplicationProvidedCalledPartyAddress      = 837
	ApplicationServer                          = 836
	ApplicationServerID                        = 2101
	ApplicationServerInformation               = 850
	ApplicationServiceProviderIdentity         = 532
	ApplicationSessionID                       = 2103
	AssociatedIdentities                       = 632
	AssociatedPartyAddress                     = 2035
	AssociatedRegisteredIdentities             = 647
	AssociatedURI                              = 856
	AuthApplicationID                          = 258
	AuthGracePeriod                            = 276
	AuthRequestType                            = 274
	AuthSessionState                           = 277
	AuthenticationInfo                         = 1413
	AuthorisedQoS                              = 849
	AuthorizationLifetime                      = 291
	AuxApplicInfo                              = 1219
	BSSID                                      = 2716
	BaseTimeInterval                           = 1265
	BasicServiceCode                           = 3411
	BearerCapability                           = 3412
	BearerIdentifier                           = 1020
	BearerService                              = 854
	BearerUsage                                = 1000
	CCCorrelationID                            = 411
	CCInputOctets                              = 412
	CCMoney                                    = 413
	CCOutputOctets                             = 414
	CCRequestNumber                            = 415
	CCRequestType                              = 416
	CCServiceSpecificUnits                     = 417
	CCSessionFailover                          = 418
	CCSubSessionID                             = 419
	CCTime                                     = 420
	CCTotalOctets                              = 421
	CCUnitType                                 = 454
	CGAddress                                  = 846
	CHAPAlgorithm                              = 403
	CHAPAuth                                   = 402
	CHAPChallenge                              = 60
	CHAPIdent                                  = 404
	CHAPResponse                               = 405
	CLRFlags                                   = 1638
	CNIPMulticastDistribution                  = 921
	CNOperatorSelectionEntity                  = 3421
	CSGAccessMode                              = 2317
	CSGID                                      = 1437
	CSGMembershipIndication                    = 2318
	CSGSubscriptionData                        = 1436
	CUGInformation                             = 2304
	CallBarringInfo                            = 1488
	CallReferenceInfo                          = 720
	CallbackID                                 = 20
	CallbackNumber                             = 19
	CalledAssertedIdentity                     = 1250
	CalledPartyAddress                         = 832
	CalledStationID                            = 30
	CallingPartyAddress                        = 831
	CallingStationID                           = 31
	CancellationType                           = 1420
	CarrierSelectRoutingInformation            = 2023
	CauseCode                                  = 861
	ChangeCondition                            = 2037
	ChangeTime                                 = 2038
	ChargeReasonCode                           = 2118
	ChargedParty                               = 857
	ChargingCharacteristicsSelectionMode       = 2066
	ChargingInformation                        = 618
	ChargingRuleBaseName                       = 1004
	ChargingRuleDefinition                     = 1003
	ChargingRuleInstall                        = 1001
	ChargingRuleName                           = 1005
	ChargingRuleRemove                         = 1002
	CheckBalanceResult                         = 422
	Class                                      = 25
	ClassIdentifier                            = 1214
	ClientAddress                              = 2018
	ClientIdentity                             = 1480
	CompleteDataListIncludedIndicator          = 1468
	ConfidentialityKey                         = 625
	ConfigurationToken                         = 78
	ConnectInfo                                = 77
	ContentClass                               = 1220
	ContentDisposition                         = 828
	ContentID                                  = 2116
	ContentLength                              = 827
	ContentProviderID                          = 2117
	ContentSize                                = 1206
	ContentType                                = 826
	ContextIdentifier                          = 1423
	CostInformation                            = 423
	CostUnit                                   = 424
	CreditControl                              = 426
	CreditControlFailureHandling               = 427
	CurrencyCode                               = 425
	CurrentTariff                              = 2056
	DRMContent                                 = 1221
	DRMP                                       = 301
	DataCodingScheme                           = 2001
	DataReference                              = 703
	DefaultEPSBearerQoS                        = 1049
	DeferredLocationEventType                  = 1230
	DeliveryReportRequested                    = 1216
	DeliveryStatus                             = 2104
	DeregistrationReason                       = 615
	DestinationHost                            = 293
	DestinationInterface                       = 2002
	DestinationRealm                           = 283
	Diagnostics                                = 2039
	DirectDebitingFailureHandling              = 428
	DisconnectCause                            = 273
	DomainName                                 = 1200
	Drmp                                       = 301
	DsaiTag                                    = 711
	DynamicAddressFlag                         = 2051
	DynamicAddressFlagExtension                = 2068
	EPSSubscribedQoSProfile                    = 1431
	EUTRANVector                               = 1414
	EarlyMediaDescription                      = 1272
	Envelope                                   = 1266
	EnvelopeEndTime                            = 1267
	EnvelopeReporting                          = 1268
	EnvelopeStartTime                          = 1269
	ErrorDiagnostic                            = 1614
	ErrorMessage                               = 281
	ErrorReportingHost                         = 294
	Event                                      = 825
	EventChargingTimeStamp                     = 1258
	EventTimestamp                             = 55
	EventTrigger                               = 1006
	EventType                                  = 823
	ExperimentalResult                         = 297
	ExperimentalResultCode                     = 298
	ExpirationDate                             = 1439
	Expires                                    = 888
	ExpiryTime                                 = 709
	Exponent                                   = 429
	ExtPDPAddress                              = 1621
	ExtPDPType                                 = 1620
	ExtendedAPNAMBRDL                          = 2848
	ExtendedAPNAMBRUL                          = 2849
	ExtendedGBRDL                              = 2850
	ExtendedGBRUL                              = 2851
	ExtendedMaxRequestedBWDL                   = 554
	ExtendedMaxRequestedBWUL                   = 555
	ExternalClient                             = 1479
	FailedAVP                                  = 279
	FailedAvp                                  = 279
	FeatureList                                = 630
	FeatureListID                              = 629
	FileRepairSupported                        = 1224
	FilterID                                   = 11
	FinalUnitAction                            = 449
	FinalUnitIndication                        = 430
	FirmwareRevision                           = 267
	FixedUserLocationInfo                      = 2825
	FlowDescription                            = 507
	FlowDirection                              = 1080
	FlowInformation                            = 1058
	FlowLabel                                  = 1057
	Flows                                      = 510
	ForwardingPending                          = 3415
	FramedAppletalkLink                        = 37
	FramedAppletalkNetwork                     = 38
	FramedAppletalkZone                        = 39
	FramedCompression                          = 13
	FramedIPAddress                            = 8
	FramedIPNetmask                            = 9
	FramedIPXNetwork                           = 23
	FramedIPv6Pool                             = 100
	FramedIPv6Prefix                           = 97
	FramedIPv6Route                            = 99
	FramedInterfaceID                          = 96
	FramedMTU                                  = 12
	FramedPool                                 = 88
	FramedProtocol                             = 7
	FramedRoute                                = 22
	FramedRouting                              = 10
	FromAddress                                = 2708
	GERANVector                                = 1416
	GGSNAddress                                = 847
	GMLCAddress                                = 2405
	GMLCNumber                                 = 1474
	GMLCRestriction                            = 1481
	GPRSSubscriptionData                       = 1467
	GSUPoolIdentifier                          = 453
	GSUPoolReference                           = 457
	GrantedServiceUnit                         = 431
	GuaranteedBitrateDL                        = 1025
	GuaranteedBitrateUL                        = 1026
	HPLMNODB                                   = 1418
	HomogeneousSupportofIMSVoiceOverPSSessions = 1493
	HostIPAddress                              = 257
	ICSIndicator                               = 1491
	IMEI                                       = 1402
	IMSApplicationReferenceIdentifier          = 2601
	IMSChargingIdentifier                      = 841
	IMSCommunicationServiceIdentifier          = 1281
	IMSEmergencyIndicator                      = 2322
	IMSIUnauthenticatedFlag                    = 2308
	IMSInformation                             = 876
	IMSVisitedNetworkIdentifier                = 2713
	IPCANType                                  = 1027
	IPRealmDefaultIndication                   = 2603
	ISUPCause                                  = 3416
	ISUPCauseDiagnostics                       = 3422
	ISUPCauseLocation                          = 3423
	ISUPCauseValue                             = 3424
	ISUPLocationNumber                         = 3414
	IdentitySet                                = 708
	IdentityWithEmergencyRegistration          = 651
	IdleTimeout                                = 28
	ImmediateResponsePreferred                 = 1412
	InbandSecurityID                           = 299
	IncomingTrunkGroupID                       = 852
	IncrementalCost                            = 2062
	InitialIMSChargingIdentifier               = 2321
	InstanceID                                 = 3402
	IntegrityKey                               = 626
	InterOperatorIdentifier                    = 838
	InterfaceID                                = 2003
	InterfacePort                              = 2004
	InterfaceText                              = 2005
	InterfaceType                              = 2006
	ItemNumber                                 = 1419
	KASME                                      = 1450
	Kc                                         = 1453
	LCSAPN                                     = 1231
	LCSClientDialedByMS                        = 1233
	LCSClientExternalID                        = 1234
	LCSClientID                                = 1232
	LCSClientName                              = 1235
	LCSClientType                              = 1241
	LCSDataCodingScheme                        = 1236
	LCSFormatIndicator                         = 1237
	LCSInfo                                    = 1473
	LCSInformation                             = 878
	LCSNameString                              = 1238
	LCSPrivacyException                        = 1475
	LCSRequestorID                             = 1239
	LCSRequestorIDString                       = 1240
	LIPAPermission                             = 1618
	LiaFlags                                   = 653
	LocalGWInsertedIndication                  = 2604
	LocalSequenceNumber                        = 2063
	LocalTimeZoneIndication                    = 718
	LocationEstimate                           = 1242
	LocationEstimateType                       = 1243
	LocationType                               = 1244
	LoginIPHost                                = 14
	LoginIPv6Host                              = 98
	LoginLATGroup                              = 36
	LoginLATNode                               = 35
	LoginLATPort                               = 63
	LoginLATService                            = 34
	LoginService                               = 15
	LoginTCPPort                               = 16
	LooseRouteIndication                       = 638
	LowBalanceIndication                       = 2020
	LowPriorityIndicator                       = 2602
	MBMS2G3GIndicator                          = 907
	MBMSChargedParty                           = 2323
	MBMSGWAddress                              = 2307
	MBMSInformation                            = 880
	MBMSServiceArea                            = 903
	MBMSServiceType                            = 906
	MBMSSessionIdentity                        = 908
	MBMSUserServiceType                        = 1225
	MDTConfiguration                           = 1622
	MDTUserConsent                             = 1634
	MIP6AgentInfo                              = 486
	MIP6FeatureVector                          = 124
	MIP6HomeLinkPrefix                         = 125
	MIPHomeAgentAddress                        = 334
	MIPHomeAgentHost                           = 348
	MMBoxStorageRequested                      = 1248
	MMContentType                              = 1203
	MMEName                                    = 2402
	MMENumberforMTSMS                          = 1645
	MMERealm                                   = 2408
	MMSInformation                             = 877
	MMTelInformation                           = 2030
	MMTelSServiceType                          = 2031
	MOLR                                       = 1485
	MPSPriority                                = 1616
	MSCAddress                                 = 3417
	MSISDN                                     = 701
	MTCIWFAddress                              = 3406
	MandatoryCapability                        = 604
	MaxRequestedBandwidthDL                    = 515
	MaxRequestedBandwidthUL                    = 516
	MediaInitiatorFlag                         = 882
	MediaInitiatorParty                        = 1288
	MessageBody                                = 889
	MessageClass                               = 1213
	MessageID                                  = 1210
	MessageSize                                = 1212
	MessageType                                = 1211
	MonitoringKey                              = 1066
	MultiRoundTimeOut                          = 272
	MultipleRegistrationIndication             = 648
	MultipleServicesCreditControl              = 456
	MultipleServicesIndicator                  = 455
	NASFilterRule                              = 400
	NASPort                                    = 5
	NASPortID                                  = 87
	NASPortType                                = 61
	NNIInformation                             = 2703
	NNIType                                    = 2704
	NORFlags                                   = 1443
	NeighbourNodeAddress                       = 2705
	NetworkAccessMode                          = 1417
	NetworkCallReferenceNumber                 = 3418
	NetworkRequestSupport                      = 1024
	NextTariff                                 = 2057
	NodeFunctionality                          = 862
	NodeID                                     = 2064
	Non3GPPIPAccess                            = 1501
	Non3GPPIPAccessAPN                         = 1502
	Non3GPPUserData                            = 1500
	NotificationToUEUser                       = 1478
	NumberOfDiversions                         = 2034
	NumberOfMessagesSent                       = 2019
	NumberOfMessagesSuccessfullyExploded       = 2111
	NumberOfMessagesSuccessfullySent           = 2112
	NumberOfParticipants                       = 885
	NumberOfReceivedTalkBursts                 = 1282
	NumberOfRequestedVectors                   = 1410
	NumberOfTalkBursts                         = 1283
	NumberPortabilityRoutingInformation        = 2024
	OCFeatureVector                            = 622
	OCOLR                                      = 623
	OCReductionPercentage                      = 627
	OCReportType                               = 626
	OCSequenceNumber                           = 624
	OCSupportedFeatures                        = 621
	OCValidityDuration                         = 625
	OMCID                                      = 1466
	OcOlr                                      = 623
	OcSupportedFeatures                        = 621
	Offline                                    = 1008
	OfflineCharging                            = 1278
	OneTimeNotification                        = 712
	Online                                     = 1009
	OnlineChargingFlag                         = 2303
	OperatorDeterminedBarring                  = 1425
	OptionalCapability                         = 605
	OriginHost                                 = 264
	OriginRealm                                = 296
	OriginStateID                              = 278
	OriginatingIOI                             = 839
	OriginatingLineInfo                        = 94
	OriginatingRequest                         = 633
	Originator                                 = 864
	OriginatorAddress                          = 886
	OriginatorInterface                        = 2009
	OriginatorReceivedAddress                  = 2027
	OriginatorSCCPAddress                      = 2008
	OutgoingSessionID                          = 2320
	OutgoingTrunkGroupID                       = 853
	PDNConnectionChargingID                    = 2050
	PDNGWAllocationType                        = 1438
	PDNType                                    = 1456
	PDPAddress                                 = 1227
	PDPAddressPrefixLength                     = 2606
	PDPContext                                 = 1469
	PDPContextType                             = 1247
	PDPType                                    = 1470
	PLMNClient                                 = 1482
	PSAppendFreeFormatData                     = 867
	PSFreeFormatData                           = 866
	PSFurnishChargingInformation               = 865
	PSInformation                              = 874
	PUAFlags                                   = 1442
	PURFlags                                   = 1635
	PacketFilterIdentifier                     = 1060
	PacketFilterUsage                          = 1072
	ParticipantAccessPriority                  = 1259
	ParticipantActionType                      = 2049
	ParticipantGroup                           = 1260
	ParticipantsInvolved                       = 887
	PasswordRetry                              = 75
	PoCChangeCondition                         = 1261
	PoCChangeTime                              = 1262
	PoCControllingAddress                      = 858
	PoCEventType                               = 2025
	PoCGroupName                               = 859
	PoCInformation                             = 879
	PoCServerRole                              = 883
	PoCSessionID                               = 1229
	PoCSessionInitiationtype                   = 1277
	PoCSessionType                             = 884
	PoCUserRole                                = 1252
	PoCUserRoleIDs                             = 1253
	PoCUserRoleinfoUnits                       = 1254
	PortLimit                                  = 62
	PositioningData                            = 1245
	PrePagingSupported                         = 717
	Precedence                                 = 1010
	PreemptionCapability                       = 1047
	PreemptionVulnerability                    = 1048
	PreferredAoCCurrency                       = 2315
	PresenceReportingAreaIdentifier            = 2821
	PresenceReportingAreaInformation           = 2822
	PresenceReportingAreaStatus                = 2823
	Priority                                   = 1209
	PriorityIndication                         = 3006
	PriorityLevel                              = 1046
	PriviledgedSenderIndication                = 652
	ProductName                                = 269
	Prompt                                     = 76
	ProxyHost                                  = 280
	ProxyInfo                                  = 284
	ProxyState                                 = 33
	PublicIdentity                             = 601
	QoSClassIdentifier                         = 1028
	QoSFilterRule                              = 407
	QoSInformation                             = 1016
	QoSSubscribed                              = 1404
	QuotaConsumptionTime                       = 881
	QuotaHoldingTime                           = 871
	RAI                                        = 909
	RAND                                       = 1447
	RATFrequencySelectionPriorityID            = 1440
	RATType                                    = 1032
	RateElement                                = 2058
	RatingGroup                                = 432
	ReAuthRequestType                          = 285
	ReadReplyReportRequested                   = 1222
	RealTimeTariffInformation                  = 2305
	ReasonCode                                 = 616
	ReasonHeader                               = 3401
	ReasonInfo                                 = 617
	ReceivedTalkBurstTime                      = 1284
	ReceivedTalkBurstVolume                    = 1285
	RecipientAddress                           = 1201
	RecipientInfo                              = 2026
	RecipientReceivedAddress                   = 2028
	RecipientSCCPAddress                       = 2010
	RedirectAddressType                        = 433
	RedirectHost                               = 292
	RedirectHostUsage                          = 261
	RedirectInformation                        = 1085
	RedirectMaxCacheTime                       = 262
	RedirectServer                             = 434
	RedirectServerAddress                      = 435
	RedirectSupport                            = 1086
	ReferenceNumber                            = 3007
	RefundInformation                          = 2022
	RegionalSubscriptionZoneCode               = 1446
	RelatedIMSChargingIdentifier               = 2711
	RelatedIMSChargingIdentifierNode           = 2712
	RelationshipMode                           = 2706
	RelayNodeIndicator                         = 1633
	RemainingBalance                           = 2021
	ReplyApplicID                              = 1223
	ReplyMessage                               = 18
	ReplyPathRequested                         = 2011
	ReportingReason                            = 872
	RepositoryDataID                           = 715
	RequestedAction                            = 436
	RequestedDomain                            = 706
	RequestedEUTRANAuthenticationInfo          = 1408
	RequestedNodes                             = 713
	RequestedPartyAddress                      = 1251
	RequestedServiceUnit                       = 437
	RequestedUTRANGERANAuthenticationInfo      = 1409
	RequiredMBMSBearerCapabilities             = 901
	RestrictionFilterRule                      = 438
	ResultCode                                 = 268
	ResynchronizationInfo                      = 1411
	RevalidationTime                           = 1042
	RoamingRestrictedDueToUnsupportedFeature   = 1457
	RoleOfNode                                 = 829
	RouteHeaderReceived                        = 3403
	RouteHeaderTransmitted                     = 3404
	RouteRecord                                = 282
	RuleActivationTime                         = 1043
	RuleDeactivationTime                       = 1044
	SDPAnswerTimestamp                         = 1275
	SDPMediaComponent                          = 843
	SDPMediaDescription                        = 845
	SDPMediaName                               = 844
	SDPOfferTimestamp                          = 1274
	SDPSessionDescription                      = 842
	SDPTimeStamps                              = 1273
	SDPType                                    = 2036
	SGSNAddress                                = 1228
	SGSNNumber                                 = 1489
	SGWAddress                                 = 2067
	SGWChange                                  = 2065
	SIPAuthDataItem                            = 612
	SIPAuthenticate                            = 609
	SIPAuthenticationScheme                    = 608
	SIPAuthorization                           = 610
	SIPItemNumber                              = 613
	SIPMethod                                  = 824
	SIPNumberAuthItems                         = 607
	SIPRequestTimestamp                        = 834
	SIPRequestTimestampFraction                = 2301
	SIPResponseTimestamp                       = 835
	SIPResponseTimestampFraction               = 2302
	SIPTOPermission                            = 1613
	SLRequestType                              = 2904
	SMDeviceTriggerIndicator                   = 3407
	SMDeviceTriggerInformation                 = 3405
	SMDischargeTime                            = 2012
	SMMessageType                              = 2007
	SMProtocolID                               = 2013
	SMSCAddress                                = 2017
	SMSInformation                             = 2000
	SMSNode                                    = 2016
	SMSResult                                  = 3409
	SMSequenceNumber                           = 3408
	SMServiceType                              = 2029
	SMStatus                                   = 2014
	SMUserDataHeader                           = 2015
	SRES                                       = 1454
	SSCode                                     = 1476
	SSID                                       = 1524
	SSStatus                                   = 1477
	STNSR                                      = 1433
	SarFlags                                   = 655
	ScaleFactor                                = 2059
	ScscfRestorationInfo                       = 639
	SecurityParameterIndex                     = 1056
	SendDataIndication                         = 710
	ServedPartyIPAddress                       = 848
	ServerAssignmentType                       = 614
	ServerCapabilities                         = 603
	ServerName                                 = 602
	ServiceContextID                           = 461
	ServiceDataContainer                       = 2040
	ServiceID                                  = 855
	ServiceIdentifier                          = 439
	ServiceIndication                          = 704
	ServiceInformation                         = 873
	ServiceMode                                = 2032
	ServiceParameterInfo                       = 440
	ServiceParameterType                       = 441
	ServiceParameterValue                      = 442
	ServiceSelection                           = 493
	ServiceSpecificData                        = 863
	ServiceSpecificInfo                        = 1249
	ServiceSpecificType                        = 1257
	ServiceType                                = 6
	ServiceTypeIdentity                        = 1484
	ServingNode                                = 2401
	ServingNodeIndication                      = 714
	ServingNodeType                            = 2047
	SessionBinding                             = 270
	SessionDirection                           = 2707
	SessionID                                  = 263
	SessionPriority                            = 650
	SessionServerFailover                      = 271
	SessionTimeout                             = 27
	SipAuthDataItem                            = 612
	SipNumberAuthItems                         = 607
	SoftwareVersion                            = 1403
	SpecificAPNInfo                            = 1472
	SponsorIdentity                            = 531
	StartTime                                  = 2041
	StartofCharging                            = 3419
	StatusASCode                               = 2702
	StopTime                                   = 2042
	SubmissionTime                             = 1202
	SubsReqType                                = 705
	SubscribedPeriodicRAUTAUTimer              = 1619
	SubscribedVSRVCC                           = 1636
	SubscriberRole                             = 2033
	SubscriberStatus                           = 1424
	SubscriptionData                           = 1400
	SubscriptionID                             = 443
	SubscriptionIDData                         = 444
	SubscriptionIDType                         = 450
	SupplementaryService                       = 2048
	SupportedFeatures                          = 628
	SupportedVendorID                          = 265
	TADIdentifier                              = 2717
	TDFIPAddress                               = 1091
	TFTFilter                                  = 1012
	TFTPacketFilterInformation                 = 1013
	TGPPAAAServerName                          = 318
	TGPPChargingCharacteristics                = 13
	TGPPChargingID                             = 2
	TGPPGGSNAddress                            = 7
	TGPPGGSNMCCMNC                             = 9
	TGPPIMSI                                   = 1
	TGPPIMSIMCCMNC                             = 8
	TGPPMSTimeZone                             = 23
	TGPPNSAPI                                  = 10
	TGPPPDPType                                = 3
	TGPPRATType                                = 21
	TGPPSGSNAddress                            = 6
	TGPPSGSNMCCMNC                             = 18
	TGPPSelectionMode                          = 12
	TGPPServiceType                            = 1483
	TGPPSessionStopIndicator                   = 11
	TGPPUserLocationInfo                       = 22
	TMGI                                       = 900
	TSCode                                     = 1487
	TWANUserLocationInfo                       = 2714
	TalkBurstExchange                          = 1255
	TalkBurstTime                              = 1286
	TalkBurstVolume                            = 1287
	TariffChangeUsage                          = 452
	TariffInformation                          = 2060
	TariffTimeChange                           = 451
	TariffXML                                  = 2306
	Teleservice                                = 3413
	TeleserviceList                            = 1486
	TerminalInformation                        = 1401
	TerminatingIOI                             = 840
	TerminationCause                           = 295
	TimeFirstUsage                             = 2043
	TimeLastUsage                              = 2044
	TimeQuotaMechanism                         = 1270
	TimeQuotaThreshold                         = 868
	TimeQuotaType                              = 1271
	TimeStamps                                 = 833
	TimeUsage                                  = 2045
	ToSTrafficClass                            = 1014
	TokenText                                  = 1215
	TotalNumberOfMessagesExploded              = 2113
	TotalNumberOfMessagesSent                  = 2114
	TraceCollectionEntity                      = 1452
	TraceData                                  = 1458
	TraceDepth                                 = 1462
	TraceEventList                             = 1465
	TraceInfo                                  = 1505
	TraceInterfaceList                         = 1464
	TraceNETypeList                            = 1463
	TraceReference                             = 1459
	TrafficDataVolumes                         = 2046
	TranscoderInsertedIndication               = 2605
	TransitIOIList                             = 2701
	Trigger                                    = 1264
	TriggerType                                = 870
	TrunkGroupID                               = 851
	TunnelAssignmentID                         = 82
	TunnelClientAuthID                         = 90
	TunnelClientEndpoint                       = 66
	TunnelMediumType                           = 65
	TunnelPassword                             = 69
	TunnelPreference                           = 83
	TunnelPrivateGroupID                       = 81
	TunnelServerAuthID                         = 91
	TunnelServerEndpoint                       = 67
	TunnelType                                 = 64
	Tunneling                                  = 401
	TypeNumber                                 = 1204
	UESRVCCCapability                          = 1615
	ULAFlags                                   = 1406
	ULRFlags                                   = 1405
	UTRANVector                                = 1415
	UarFlags                                   = 637
	UdrFlags                                   = 719
	UnitCost                                   = 2061
	UnitQuotaThreshold                         = 1226
	UnitValue                                  = 445
	UsageMonitoringInformation                 = 1067
	UsageMonitoringLevel                       = 1068
	UsageMonitoringReport                      = 1069
	UsageMonitoringSupport                     = 1070
	UsedServiceUnit                            = 446
	UserAuthorizationType                      = 623
	UserCSGInformation                         = 2319
	UserDataAlreadyAvailable                   = 624
	UserDataCx                                 = 606
	UserDataSh                                 = 702
	UserEquipmentInfo                          = 458
	UserEquipmentInfoType                      = 459
	UserEquipmentInfoValue                     = 460
	UserID                                     = 1444
	UserIdentity                               = 700
	UserLocationInfoTime                       = 2812
	UserName                                   = 1
	UserParticipatingType                      = 1279
	UserPassword                               = 2
	UserSessionID                              = 830
	VASID                                      = 1102
	VASPID                                     = 1101
	VCSInformation                             = 3410
	VLRNumber                                  = 3420
	VPLMNDynamicAddressAllowed                 = 1432
	VPLMNLIPAAllowed                           = 1617
	ValidityTime                               = 448
	ValueDigits                                = 447
	VendorID                                   = 266
	VendorSpecificApplicationID                = 260
	VisitedNetworkIdentifier                   = 600
	VisitedPLMNID                              = 1407
	VolumeQuotaThreshold                       = 869
	WildcardedImpu                             = 636
	WildcardedPublicIdentity                   = 634
	XRES                                       = 1448
	ePDGAddress                                = 3425
)
