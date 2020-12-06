package app

import (
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	FlagNameConsulAddress         = "consul.address"
	FlagNameConsulAuthUsername    = "consul.auth.username"
	FlagNameConsulAuthPassword    = "consul.auth.password"
	FlagNameGRPCListenAddr        = "grpc.listen"
	FlagNameHTTPListenAddr        = "http.listen"
	FlagNameHTTPTLSCertFile       = "http.tls.cert"
	FlagNameHTTPTLSKeyFile        = "http.tls.key"
	FlagNameNATSAddresses         = "nats.address"
	FlagNameNATSUsername          = "nats.auth.username"
	FlagNameNATSPassword          = "nats.auth.password"
	FlagNameRedisMasterName       = "redis.master.name"
	FlagNameRedisPassword         = "redis.password"
	FlagNameRedisAddress          = "redis.address"
	FlagNameRedisDB               = "redis.db"
	FlagNameServiceRegisterTarget = "service.register.target"
	FlagNameLogLevel              = "log.level"
)

func init() {
	pflag.String(FlagNameConsulAddress, "localhost:8500", "指定consul的`连接地址`")
	pflag.String(FlagNameConsulAuthUsername, "", "指定连接consul时使用的HTTP基本认证的`用户名`")
	pflag.String(FlagNameConsulAuthPassword, "", "指定连接consul时使用的HTTP基本认证的`密码`")
	pflag.String(FlagNameGRPCListenAddr, ":0", "指定gRPC服务器监听的`IP:Port`")
	pflag.String(FlagNameHTTPListenAddr, ":0", "指定http服务器监听的`IP:Port`")
	pflag.String(FlagNameHTTPTLSCertFile, "", "指定http服务器在开启TLS支持时使用的`证书文件路径`")
	pflag.String(FlagNameHTTPTLSKeyFile, "", "指定http服务器在开启TLS支持时使用的`密钥文件路径`")
	pflag.StringSlice(FlagNameNATSAddresses, []string{"localhost:4222"}, "指定连接nats的`连接地址`，多个地址使用逗号（,）分隔")
	pflag.String(FlagNameNATSUsername, "", "指定连接nats所需认证的`用户名`")
	pflag.String(FlagNameNATSPassword, "", "指定连接nats所需认证的`密码`")
	pflag.String(FlagNameRedisMasterName, "", "指定连接redis的`集群名称`，如果不为空将使用哨兵模式")
	pflag.String(FlagNameRedisPassword, "", "指定连接redis的`密码`")
	pflag.StringSlice(FlagNameRedisAddress, []string{"localhost:6379"}, "指定连接redis的`链接地址`，如果有多个地址将使用集群模式；多个地址使用逗号（,）分隔")
	pflag.Int(FlagNameRedisDB, 0, "指定连接redis的`数据库`编号，默认为0，即默认使用DB0")
	pflag.String(FlagNameServiceRegisterTarget, "", "指定用于服务发现的服务注册地址")
	pflag.Int(FlagNameLogLevel, -1, "日志等级，DebugLevel：-1，InfoLevel：0，WarnLevel：1，ErrorLevel：2，DPanicLevel：3，PanicLevel：4，FatalLevel：5")

	_ = viper.BindPFlags(pflag.CommandLine)

	useHelp := pflag.BoolP("help", "h", false, "帮助")
	pflag.Parse()

	if *useHelp {
		pflag.Usage()
		os.Exit(0)
	}

	logLv := viper.GetInt(FlagNameLogLevel)
	if logLv < -1 || logLv > 5 {
		panic("日志等级错误")
	}

}
