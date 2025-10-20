package web

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"text/template"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// Anotation para juntar os arquivos do template em um único embed.FS
// O embed.FS é um recurso do Go para juntar arquivos em um único arquivo
//
//go:embed template/*
var templateContent embed.FS

type Webserver struct {
	TemplateData *TemplateData
}

// NewServer cria uma nova instância do servidor
func NewServer(templateData *TemplateData) *Webserver {
	return &Webserver{
		TemplateData: templateData,
	}
}

// CreateServer cria uma nova instância do servidor com o router go chi
func (we *Webserver) CreateServer() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))
	// promhttp é um handler para o Prometheus
	router.Handle("/metrics", promhttp.Handler())
	router.Get("/", we.HandleRequest)
	return router
}

type TemplateData struct {
	Title              string
	BackgroundColor    string
	ResponseTime       time.Duration
	ExternalCallMethod string
	ExternalCallURL    string
	Content            string
	RequestNameOTEL    string
	OTELTracer         trace.Tracer
}

func (h *Webserver) HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Extrai o carrier do header da requisição
	carrier := propagation.HeaderCarrier(r.Header)
	// Cria um contexto com o carrier
	ctx := r.Context()
	// Extrai o contexto do carrier
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	// Cria um span para a inicialização
	ctx, spanInicial := h.TemplateData.OTELTracer.Start(ctx, "SPAN_INICIAL"+h.TemplateData.RequestNameOTEL)
	time.Sleep(time.Second)
	spanInicial.End()
	
	// Cria um span para a chamada externa
	ctx, span := h.TemplateData.OTELTracer.Start(ctx, "Chama externa"+h.TemplateData.RequestNameOTEL)
	defer span.End()

	// Espera o tempo de resposta
	time.Sleep(time.Millisecond * h.TemplateData.ResponseTime)

	// Se a URL de chamada externa não está vazia, faz a chamada externa
	if h.TemplateData.ExternalCallURL != "" {
		var req *http.Request
		var err error
		// Cria uma nova requisição com o contexto e o método HTTP
		if h.TemplateData.ExternalCallMethod == "GET" {
			// Cria uma nova requisição com o contexto e o método GET
			req, err = http.NewRequestWithContext(ctx, "GET", h.TemplateData.ExternalCallURL, nil)
		} else if h.TemplateData.ExternalCallMethod == "POST" {
			// Cria uma nova requisição com o contexto e o método POST
			req, err = http.NewRequestWithContext(ctx, "POST", h.TemplateData.ExternalCallURL, nil)
		} else {
			// Se o método HTTP não é válido, retorna um erro 500
			http.Error(w, "Invalid ExternalCallMethod", http.StatusInternalServerError)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Injeta o contexto no header da requisição
		otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
		// Faz a chamada externa
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		h.TemplateData.Content = string(bodyBytes)
	}

	tpl := template.Must(template.New("index.html").ParseFS(templateContent, "template/index.html"))
	err := tpl.Execute(w, h.TemplateData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}
