package global

const (
	dburi       = "mongodb+srv://sicepot:ekspres@cluster0.qlb5q.mongodb.net/?retryWrites=true&w=majority"
	dbname      = "grpc-blog-app"
	performance = 100
)

var (
	jwtSecret = []byte("blogSecret")
)
