package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"posp_api_go_v2/src/services/dbo/getAgentListById"
	"posp_api_go_v2/src/services/dbo/insertUpdateBmsBookingDetails"
	"posp_api_go_v2/src/services/dbo/insertUpdateBmsLeadDetails"
	"posp_api_go_v2/src/services/dbo/partnerUpgrade"
	"posp_api_go_v2/src/services/exportExcel"
	"posp_api_go_v2/src/services/reports"
	"posp_api_go_v2/src/services/rnd"
)

func initializeRouter() {

	r := mux.NewRouter()

	//Report Package Router
	mount(r, "/api/report", reports.Router())

	//Rnd Package Router
	mount(r, "/api/rnd", rnd.Router())

	//Export Excel Package Router
	mount(r, "/api/ee", exportExcel.Router())

	//getAgentListById Package Router
	//mount(r, "/api/getAgentListById", getAgentListById.Router())
	r.HandleFunc("/api/getAgentListById", getAgentListById.CallAmit).Methods("POST")

	r.HandleFunc("/api/insertUpdateBmsLeadDetails", insertUpdateBmsLeadDetails.HandleRequest).Methods("POST")

	r.HandleFunc("/api/insertUpdateBmsBookingDetails", insertUpdateBmsBookingDetails.HandleRequest).Methods("POST")

	r.HandleFunc("/api/partnerUpgrade", partnerUpgrade.HandleRequest).Methods("POST")

	//Start Server at a port
	RunServer(r)
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}

func RunServer(r *mux.Router) {
	//portNo := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(":9000",
		handlers.CORS(handlers.AllowedHeaders([]string{
			"X-Requested-With",
			"Content-Type",
			"Authorization"}),
			handlers.AllowedMethods([]string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"HEAD",
				"OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(r)))
}
