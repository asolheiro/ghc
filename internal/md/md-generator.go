package md

import (
	"fmt"
	"os"

	"github.com/asolheiro/gita-healthcheck/internal/api-calls/alerts"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/auth"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/count"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/incidents"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/metrics"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/problem"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/security"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/version"
	"github.com/asolheiro/gita-healthcheck/internal/maps"
	md "github.com/nao1215/markdown"
)

type FileVars struct {
	Index          int
	Auth           auth.AuthResponse
	OrgName        string
	Cluster        count.Cluster
	Incidents      incidents.Response
	Problem        problem.Response
	Security       security.Response
	ClusterMetrics metrics.TotalMetrics
	NodeMetrics    []metrics.TotalNodeMetrics
	AlertsList     []alerts.Msg
	KubernetesInfo version.KubernetesVersionResponse
}

func GenerateFile(args FileVars) {
	fileName := fmt.Sprintf("%d-gita-report-%s.md", args.Index, args.Cluster.Name)
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m := md.NewMarkdown(f)

	m.H2(fmt.Sprintf("%d. %s", args.Index, args.Cluster.Name))
	m.PlainTextf("")

	m.H3(fmt.Sprintf("%d.1. Informações gerais", args.Index))
	m.PlainTextf("")

	m.Table(md.TableSet{
		Header: []string{"Descrição", "Número"},
		Rows: [][]string{
			{"Nodes", fmt.Sprint(args.Cluster.Node)},
			{"Incidentes", fmt.Sprint(args.Incidents.Total)},
			{"Problemas", fmt.Sprint(args.Problem.Total)},
			{"Segurança", fmt.Sprint(args.Security.Total)},
		},
	})

	m.H3(fmt.Sprintf("%d.2. Versão do Kubernetes", args.Index))
	m.PlainTextf("")

	k8sInfo := []string{
		args.KubernetesInfo.Version,
		args.KubernetesInfo.EndOfSupport,
		args.KubernetesInfo.LastRelease,
		args.KubernetesInfo.SupportStatus,
	}

	m.Table(md.TableSet{
		Header: []string{"Versão", "Fim do suporte", "Última release disponível", "Suporte"},
		Rows:   [][]string{k8sInfo},
	})

	m.PlainText("> Legenda:\n>\n> 🟩 - LTS; 🟨 - menos de 90 dias restantes; 🟧 - menos de 30 dias restantes 🟥 - desatualizado.")

	m.PlainTextf("")

	m.H3(fmt.Sprintf("%d.3. Informações de recursos", args.Index))
	m.PlainTextf("")

	m.Table(md.TableSet{
		Header: []string{"Recursos", "Capacidade", "Status"},
		Rows: [][]string{
			{"CPU", fmt.Sprintf("%d cores", args.ClusterMetrics.TotalCPU), maps.ColorRuleResources(args.ClusterMetrics, "CPU")},
			{"Memória", fmt.Sprintf("%.2f Gib", maps.TotalMemoryToGib(args.ClusterMetrics.TotalMemory)), maps.ColorRuleResources(args.ClusterMetrics, "MEM")},
			{"PODS", fmt.Sprintf("%d", args.ClusterMetrics.TotalPodCapacity), maps.ColorRuleResources(args.ClusterMetrics, "POD")},
		},
	})

	m.PlainText("> Legenda:\n>\n> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.")
	m.PlainTextf("")

	m.H3(fmt.Sprintf("%d.4. Uso de memória dos nodes", args.Index))
	m.PlainTextf("")

	var gt65, gt80, lt65 int
	for _, nodeMetric := range args.NodeMetrics {
		switch {
		case nodeMetric.MemoryUsePercentage >= 80.00:
			gt80++
		case (nodeMetric.MemoryUsePercentage >= 65.00) && (nodeMetric.MemoryUsePercentage < 80.00):
			gt65++
		case nodeMetric.MemoryUsePercentage < 65.00:
			lt65++
		}
	}

	m.Table(md.TableSet{
		Header: []string{"Grupo", "Quantidade", "Status"},
		Rows: [][]string{
			{"X < 65%", fmt.Sprintf("%d", lt65), "🟩"},
			{"X > 65%", fmt.Sprintf("%d", gt65), "🟨"},
			{"X > 80%", fmt.Sprintf("%d", gt80), "🟥"},
		},
	})

	m.PlainText("> Legenda:\n>\n> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.")
	m.PlainTextf("")

	m.H3(fmt.Sprintf("%d.5. Alertas", args.Index))
	m.PlainTextf("")

	var size int
	if len(args.AlertsList) > 0 {
		size = len(args.AlertsList)
	} else {
		size++
	}

	alerts := make([][]string, size)
	if len(args.AlertsList) == 0 {
		alerts[0] = []string{
			"-",
			"-",
			"-",
			"-",
		}
	} else {
		for i, alert := range args.AlertsList {
			alerts[i] = []string{
				alert.DataHora.Date.Format("02/01/06"),
				alert.Kind + "\\" + alert.Namespace,
				alert.Name,
				alert.Msg,
			}
		}
	}

	m.Table(md.TableSet{
		Header: []string{"Data", "`Recurso\\Namespace`", "Nome", "Descrição"},
		Rows:   alerts,
	},
	)

	m.H3(fmt.Sprintf("%d.6. Incidentes", args.Index))
	m.PlainTextf("")

	var (
		loki, monitoring, cattleSystem, gita, kubeSystem, certManager, others int
	)
	if args.Incidents.Total > 0 {
		for _, incident := range args.Incidents.Msg {
			switch incident.Namespace {
			case "loki":
				loki += 1
			case "monitoring":
				monitoring += 1
			case "cattle-monitoring-system":
				monitoring += 1
			case "cattle-system":
				cattleSystem += 1
			case "gita":
				gita += 1
			case "kube-system":
				kubeSystem += 1
			case "cert-manager":
				certManager += 1
			default:
				others += 1
			}
		}
	}

	m.Table(md.TableSet{
		Header: []string{"Stack", "Namespace", "Status"},
		Rows: [][]string{
			{"CerManager", "`cert-manager`", colorRuleIncident(certManager)},
			{"GITA", "`gita`", colorRuleIncident(gita)},
			{"Loki", "`loki`", colorRuleIncident(loki)},
			{"Prometheus", "`monitoring`", colorRuleIncident(monitoring)},
			{"Rancher", "`cattle-system`", colorRuleIncident(cattleSystem)},
			{"Outros", "--", colorRuleIncident(others)},
		},
	})

	m.PlainText("> Legenda:\n>\n> 🟩 - sem incidentes; 🟥 - possui incidentes.")

	m.PlainTextf("")
	m.HorizontalRule()
	m.PlainTextf("\n")

	m.Build()
}

func colorRuleIncident(i int) string {
	if i > 0 {
		return "🟥"
	} else {
		return "🟩"
	}
}
