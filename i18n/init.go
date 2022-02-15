package i18n

var (
	config Config
)

func Init() {
	config = Config{
		PathLangFiles: "-@",
	}
}
