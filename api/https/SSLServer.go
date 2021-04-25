package https

import (
	"InterfaceMockView/models"
	"InterfaceMockView/utils/log"
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"time"
)

var SSLServerMgr = &SSLServer{
	HttpsServer:&http.Server{
		Addr:           ":443",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	},
	Router: nil,
	Certs:  nil,
}

type Certificates struct {
	CertFile string
	KeyFile  string
}

type SSLServer struct {
	HttpsServer *http.Server
	Router      *gin.Engine
	Certs       []Certificates
}

func ListenAndServeTLSSNI(srv *http.Server, Certs []Certificates) error {
	addr := ":https"
	certs := Certs
	config := &tls.Config{}
	if srv.TLSConfig != nil {
		*config = *srv.TLSConfig
	}
	if config.NextProtos == nil {
		config.NextProtos = []string{"http/1.1"}
	}

	var err error

	config.Certificates = make([]tls.Certificate, len(certs))
	for i, v := range certs {
		config.Certificates[i], err = tls.LoadX509KeyPair(v.CertFile, v.KeyFile)
		if err != nil {
			log.Error(err)
		}
	}

	config.BuildNameToCertificate()

	conn, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	tlsListener := tls.NewListener(conn, config)
	err = srv.Serve(tlsListener)
	if err != nil {
		return err
	}
	return nil
}

func (s *SSLServer)GetCerts() {
	var certs []Certificates
	var domains []models.Domain
	if err := models.DB.Find(&domains).Error; err != nil {
		log.Error("certslist Query Fail")
	}
	for _, k := range domains {
		certs = append(certs, Certificates{
			CertFile: "./data/Domain/" + k.CrtFilePath,
			KeyFile:  "./data/Domain/" + k.KeyFilePath,
		})
	}
	if len(certs) == 0 {
		log.Error("Certs len not")
	}
	s.Certs = certs
}

func (s *SSLServer) Start() {
	s.HttpsServer.Handler = s.Router
	if s.GetCerts(); len(s.Certs) != 0 {
		go func() {
			if err := ListenAndServeTLSSNI(s.HttpsServer, s.Certs); err != nil {
				log.Error(err)
			}
		}()
	}
}

func (s *SSLServer) Restart() {
	if err := s.HttpsServer.Shutdown(nil); err != nil {
		s.Start()
	}
}
