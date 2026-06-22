package config

var allowedOrigins = []string{
	"http://localhost:5173",
	//"http://be-latihan-production-8f9c.up.railway.app"
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}