package setting

import "time"

type Config struct {
	JwtSecret string
	PageSize  int
	//上传
	PrefixUrl      string
	ImageMaxSize   int
	ImageAllowExts []string
	//http
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	//tcp
	Address string
	//数据库
	Type        string
	User        string
	Password    string
	Host        string
	SqlName     string
	TablePrefix string
}

var (
	RUNTIME_ROOT_PATH = "runtime/"
	IMAGE_SAVE_PATH   = "upload/images/"
	EXPORT_SAVE_PATH  = "export/"
	QR_CODE_SAVE_PATH = "qrcode/"
	FONT_SAVE_PATH    = "fonts/"
	LOG_SAVE_PATH     = "logs/"
	LOG_SAVE_NAME     = "log"
	LOG_FILE_EXT      = "log"
	TIME_FORMAT       = "20060102"
)

var Conf *Config

func Setup(this *Config) {
	//sql
	if this.Type == "" {
		this.Type = "mysql"
	}
	if this.User == "" {
		this.User = "root"
	}
	if this.TablePrefix == "" {
		this.TablePrefix = this.SqlName + "_"
	}
	//server
	if this.ReadTimeout == 0 {
		this.ReadTimeout = 60
	}
	if this.WriteTimeout == 0 {
		this.WriteTimeout = 60
	}
	this.ReadTimeout = this.ReadTimeout * time.Second
	this.WriteTimeout = this.ReadTimeout * time.Second
	//image
	if this.ImageMaxSize == 0 {
		this.ImageMaxSize = 5
	}
	if len(this.ImageAllowExts) == 0 {
		this.ImageAllowExts = []string{".jpg", ".jpeg", ".png"}
	}
	this.ImageMaxSize = this.ImageMaxSize * 1024 * 1024
	//page
	if this.PageSize == 0 {
		this.PageSize = 30
	}
	Conf = this
}