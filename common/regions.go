package common

// Region represents ECS region
type Region string

// Constants of region definition
const (
	Hangzhou        = Region("cn-hangzhou")
	Qingdao         = Region("cn-qingdao")
	Beijing         = Region("cn-beijing")
	BeiJingHpc      = Region("cn-beijing-hpc")
	Hongkong        = Region("cn-hongkong")
	Shenzhen        = Region("cn-shenzhen")
	USWest1         = Region("us-west-1")
	APSouthEast1    = Region("ap-southeast-1")
	APNorthEast1    = Region("ap-northeast-1")
	Shanghai        = Region("cn-shanghai")
	ShenZhenFinance = Region("cn-shenzhen-finance-1")
	EUCentral1      = Region("eu-central-1")   //德国
	MEEast1         = Region("me-east-1")      //中东
	APSouthEast2    = Region("ap-southeast-2") //澳洲
)

var ValidRegions = []Region{Hangzhou, Qingdao, Beijing, BeiJingHpc, Shenzhen, Hongkong,
	Shanghai, USWest1, APSouthEast1, APNorthEast1, ShenZhenFinance,
	EUCentral1, MEEast1, APSouthEast2}
