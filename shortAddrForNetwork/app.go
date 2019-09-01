package shortAddrForNetwork
import "github.com/gorilla/mux"

type App struct {
	Router *mux.Router
}

//短地址请求
type shortReq struct {
	URL string  //'json:"url" validate:"nonzero" '
	ExpirationInMinutes int64 
}

