package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"https://be-latihan-production-8f9c.up.railway.app",
	"https://my-fe-omega.vercel.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}