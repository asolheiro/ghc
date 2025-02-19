package cmd

import (
	"fmt"
	"os"

	md "github.com/nao1215/markdown"
)

//go:generate go run main.go

func generateFile(organization string) {
	fileName := fmt.Sprintf("gita-report-%s.md", organization)
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m := md.NewMarkdown(f)
	
	m.H1(organization)
	m.H2("1. alderann - homologação")

	m.H3("1.1. Informações gerais")
	m.Table(md.TableSet{
		Header: []string{"Descrição", "Número"},
		Rows: [][]string{
			{"Nodes", "7"},
			{"Incidentes", "7"},
			{"Problemas", "350"},
			{"Segurança", "462"},
		},
	})

	m.Table(md.TableSet{
		Header: []string{"Versão", "Suporte", "Fim do suporte"},
		Rows: [][]string{
			{"1.24.10", "🟥", "---"},
		},
	})

	m.PlainText("> Legenda:\n>\n> 🟩 - suporte longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.")

	m.H3("1.2. Informações de recursos")
	m.Table(md.TableSet{
		Header: []string{"Recursos", "Capacidade", "Status"},
		Rows: [][]string{
			{"CPU", "60 cores", "🟩"},
			{"Memória", "164 Gb", "🟩"},
			{"PODS", "770", "🟩"},
		},
	})

	m.PlainText("> Legenda:\n>\n> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.")

	m.H3("1.3. Uso de memória dos nodes")
	m.Table(md.TableSet{
		Header: []string{"Grupo", "Quantidade", "Status"},
		Rows: [][]string{
			{"X < 65%", "1", "🟩"},
			{"X > 65%", "4", "🟨"},
			{"X > 80%", "2", "🟥"},
		},
	})

	m.PlainText("> Legenda:\n>\n> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.")

	m.H3("1.4. Alertas")
	m.Table(md.TableSet{
		Header: []string{"Data", "Recurso.Namespace", "Nome", "Descrição", "Aberto?"},
		Rows: [][]string{
			{"28/11/24", "pvc.escrita", "pvc-smb-escrita", "pvc as less than twenty percent available memory", "Sim"},
		},
	})

	m.H3("1.5. Incidentes")
	m.Table(md.TableSet{
		Header: []string{"Stack", "Namespace", "Status"},
		Rows: [][]string{
			{"CerManager", "`cert-manager`", "🟩"},
			{"Loki", "`loki`", "🟩"},
			{"Prometheus", "`monitoring`", "🟩"},
			{"Rancher", "`cattle-system`", "🟩"},
			{"Outros", "--", "🟩"},
		},
	})

	m.PlainText("> Legenda:\n>\n> 🟩 - sem incidentes; 🟥 - possui incidentes.")

	m.HorizontalRule()
	m.Build()
}
