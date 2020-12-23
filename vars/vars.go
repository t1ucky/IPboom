package vars

import (
	"github.com/patrickmn/go-cache"

	"gopkg.in/cheggaaa/pb.v2"

	"sync"
	"time"
	"strings"
)

var (
	IpList     = "192.168.0.1/24"
	ResultFile = "res.txt"

	UserDict = "user.dic"
	PassDict = "pass.dic"
	DbPortDict = "port.txt"

	TimeOut = 3 * time.Second
	ScanNum = 10

	DebugMode bool

	StartTime time.Time

	ProgressBar      *pb.ProgressBar
	ProcessBarActive *pb.ProgressBar
)

var (
	CacheService *cache.Cache
	Mutex        sync.Mutex
	/*Protocol = []string{//是否是一个动态数组，这是初始化，但是用户输入后就只是用户的值而不应该是叠加
		"SSH","MYSQL","MSSQL","SMB",
	}*/
	Protocol []string//切片
	PortNames = map[int]string{
		21:    "FTP",
		22:    "SSH",
		161:   "SNMP",
		445:   "SMB",
		1433:  "MSSQL",
		3306:  "MYSQL",
		5432:  "POSTGRESQL",
		6379:  "REDIS",
		9200:  "ELASTICSEARCH",
		27017: "MONGODB",
	}
	TCPPortocols = map[string]int{
		"SSH":		22,
		"SMB":		445,
		"MSSQL":	1433,
		"MYSQL":	3306,
	}

	UdpProtocols = map[string]bool{
		"SNMP": true,
	}

	// 标记特定服务的特定用户是否破解成功，成功的话不再尝试破解该用户
	SuccessHash map[string]bool

	SupportProtocols map[string]bool
	DbPort map[string]int
)

func init() {
	SuccessHash = make(map[string]bool)
	CacheService = cache.New(cache.NoExpiration, cache.DefaultExpiration)

	SupportProtocols = make(map[string]bool)
	DbPort = make(map[string]int)


	for _, proto := range PortNames {
		SupportProtocols[strings.ToUpper(proto)] = true
	}

}
