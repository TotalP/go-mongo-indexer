package container

type FDPropertyContainer struct {
	Value    string
	TypeUid  string
	Typename string
}

type FIDataSourceDto struct {
	Id                      string
	Class                   string
	IFileMobId              string
	ISourceMobId            string
	IMasterMobId            string
	IChannelId              string
	IMediaFolder            string
	IMediaFileName          string
	IMetadataFolder         string
	IMetadataFileName       string
	IDescriptorResolutionId string
	Properties              map[string]FDPropertyContainer
}
