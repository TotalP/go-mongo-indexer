package container

import (
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"strings"
)

func ConvertToDto(input *bson.Raw) FIDataSourceDto {
	var result FIDataSourceDto
	result.Properties = make(map[string]FDPropertyContainer)
	if elements, err := input.Elements(); err == nil {
		for _, element := range elements {
			key := element.Key()
			value := element.Value()
			switch key {
			case "_id":
				result.Id = value.String()
			case "_class":
				continue
			case "iFileMobId":
				result.IFileMobId = value.String()
			case "iSourceMobId":
				result.ISourceMobId = value.String()
			case "iMasterMobId":
				result.IMasterMobId = value.String()
			case "iChannelId":
				result.IChannelId = value.String()
			case "iMediaFolder":
				result.IMediaFolder = value.String()
			case "iMediaFileName":
				result.IMediaFileName = value.String()
			case "iMetadataFolder":
				result.IMetadataFolder = value.String()
			case "iMetadataFileName":
				result.IMetadataFileName = value.String()
			case "iDescriptorResolutionId":
				result.IDescriptorResolutionId = value.String()
			default:
				array := value.Array()
				result.Properties[strings.Replace(key, "#", ".", -1)] = FDPropertyContainer{
					Value:    array.Index(0).Value().String(),
					TypeUid:  array.Index(1).Value().String(),
					Typename: array.Index(2).Value().String(),
				}
			}
		}
	} else {
		log.Fatal("Unable to get the elements from bson.Raw")
	}
	return result
}
