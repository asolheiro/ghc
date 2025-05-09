package webhooks

import (
	"fmt"
	"time"
)

// KubernetesReportData contains data for the Kubernetes report card
type KubernetesReportData struct {
	Title          string
	CreatedUtc     time.Time
	NodesCount     int
	IncidentsCount int
	ProblemsCount  int
	SecurityCount  int

	KubernetesVersion         string
	EndOfSupport              string
	LatestRelease             string
	K8sVersionSupportColorMap string

	TotalCPU             int
	ColorMapCPU          string
	TotalMemory          int
	ColorMapMemory       string
	TotalPods            int
	ColorMapPods         string
	CounterLessThan65    int
	CounterGreaterThan65 int
	CounterGreaterThen80 int

	Alerts []Alert

	CertManagerIncidentsColorMap string
	GitaIncidentsColorMap        string
	LokiIncidentsColorMap        string
	MonitoringIncidentsColorMap  string
	RancherIncidentsColorMap     string
	OtherIncidentsColorMap       string
}

type Alert struct {
	AlertName        string
	AlertDate        string
	AlertResource    string
	AlertNamespace   string
	AlertDescription string
}

// SendKubernetesReportCard sends a Kubernetes report as an adaptive card
func SendKubernetesReportCard(webhookURL string, data KubernetesReportData) error {
	// Create the adaptive card that matches the template
	payload := map[string]any{
		"type": "Message",
		"attachments": []map[string]any{
			{
				"contentType": "application/vnd.microsoft.card.adaptive",
				"content": map[string]any{
					"$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
					"type":    "AdaptiveCard",
					"version": "1.5",
					"body": []map[string]any{
						{
							"type":   "TextBlock",
							"size":   "Medium",
							"weight": "Bolder",
							"text":   data.Title,
						},
						{
							"type": "ColumnSet",
							"columns": []map[string]any{
								{
									"type":  "Column",
									"width": "stretch",
									"items": []map[string]any{
										{
											"type":     "TextBlock",
											"spacing":  "None",
											"text":     fmt.Sprintf("Created {{DATE(%s,SHORT)}}", data.CreatedUtc.Format(time.RFC3339)),
											"isSubtle": true,
											"wrap":     true,
										},
										{
											"type": "Container",
											"items": []map[string]any{
												{
													"type": "TextBlock",
													"text": "1. Informações gerais:",
													"wrap": true,
												},
												{
													"type":      "Table",
													"gridStyle": "accent",
													"columns": []map[string]any{
														{"width": 1},
														{"width": 1},
													},
													"rows": []map[string]any{
														{
															"type":  "TableRow",
															"style": "accent",
															"cells": []map[string]any{
																{
																	"type": "TableCell",
																	"items": []map[string]any{
																		{
																			"type":   "TextBlock",
																			"text":   "Métrica",
																			"wrap":   true,
																			"weight": "Bolder",
																		},
																	},
																},
																{
																	"type": "TableCell",
																	"items": []map[string]any{
																		{
																			"type":   "TextBlock",
																			"text":   "Valor",
																			"wrap":   true,
																			"weight": "Bolder",
																		},
																	},
																},
															},
														},
														{
															"type": "TableRow",
															"cells": []map[string]any{
																{
																	"type":  "TableCell",
																	"style": "emphasis",
																	"items": []map[string]any{
																		{
																			"type": "TextBlock",
																			"text": "Nodes",
																			"wrap": true,
																		},
																	},
																},
																{
																	"type":  "TableCell",
																	"style": "emphasis",
																	"items": []map[string]any{
																		{
																			"type": "TextBlock",
																			"text": data.NodesCount,
																			"wrap": true,
																		},
																	},
																},
															},
														},
														{
															"type": "TableRow",
															"cells": []map[string]any{
																{
																	"type":  "TableCell",
																	"style": "emphasis",
																	"items": []map[string]any{
																		{
																			"type": "TextBlock",
																			"text": "Incidentes",
																			"wrap": true,
																		},
																	},
																},
																{
																	"type":  "TableCell",
																	"style": "emphasis",
																	"items": []map[string]any{
																		{
																			"type": "TextBlock",
																			"text": data.IncidentsCount,
																			"wrap": true,
																		},
																	},
																},
															},
														},
														{
															"type": "TableRow",
															"cells": []map[string]any{
																{
																	"type":  "TableCell",
																	"style": "emphasis",
																	"items": []map[string]any{
																		{
																			"type": "TextBlock",
																			"text": "Problemas",
																			"wrap": true,
																		},
																	},
																},
																{
																	"type":  "TableCell",
																	"style": "emphasis",
																	"items": []map[string]any{
																		{
																			"type": "TextBlock",
																			"text": data.ProblemsCount,
																			"wrap": true,
																		},
																	},
																},
															},
														},
														{
															"type": "TableRow",
															"cells": []map[string]any{
																{
																	"type":  "TableCell",
																	"style": "emphasis",
																	"items": []map[string]any{
																		{
																			"type": "TextBlock",
																			"text": "Segurança",
																			"wrap": true,
																		},
																	},
																},
																{
																	"type":  "TableCell",
																	"style": "emphasis",
																	"items": []map[string]any{
																		{
																			"type": "TextBlock",
																			"text": data.SecurityCount,
																			"wrap": true,
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						{
							"type": "Container",
							"items": []map[string]any{
								{
									"type": "TextBlock",
									"text": "2. Versão do Kubernetes",
									"wrap": true,
								},
								{
									"type": "Table",
									"columns": []map[string]any{
										{"width": 1},
										{"width": 1},
										{"width": 1},
										{"width": 1},
									},
									"rows": []map[string]any{
										{
											"type":  "TableRow",
											"style": "accent",
											"cells": []map[string]any{
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type":   "TextBlock",
															"text":   "Versão",
															"wrap":   true,
															"weight": "Bolder",
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type":   "TextBlock",
															"text":   "Fim do suporte",
															"wrap":   true,
															"weight": "Bolder",
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type":   "TextBlock",
															"text":   "Última release disponível",
															"wrap":   true,
															"weight": "Bolder",
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type":   "TextBlock",
															"text":   "Suporte",
															"wrap":   true,
															"weight": "Bolder",
														},
													},
												},
											},
										},
										{
											"type": "TableRow",
											"cells": []map[string]any{
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.KubernetesVersion,
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.EndOfSupport,
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.LatestRelease,
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.K8sVersionSupportColorMap,
															"wrap": true,
														},
													},
												},
											},
										},
									},
								},
							},
						},
						{
							"type": "Container",
							"items": []map[string]any{
								{
									"type": "TextBlock",
									"text": "3. Informações de recursos",
									"wrap": true,
								},
								{
									"type": "Table",
									"columns": []map[string]any{
										{"width": 1},
										{"width": 1},
										{"width": 1},
									},
									"rows": []map[string]any{
										{
											"type":  "TableRow",
											"style": "accent",
											"cells": []map[string]any{
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type":   "TextBlock",
															"text":   "Recursos",
															"wrap":   true,
															"weight": "Bolder",
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type":   "TextBlock",
															"text":   "Capacidade",
															"wrap":   true,
															"weight": "Bolder",
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type":   "TextBlock",
															"text":   "Status",
															"wrap":   true,
															"weight": "Bolder",
														},
													},
												},
											},
										},
										{
											"type": "TableRow",
											"cells": []map[string]any{
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "CPU",
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.TotalCPU,
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.ColorMapCPU,
															"wrap": true,
														},
													},
												},
											},
										},
										{
											"type": "TableRow",
											"cells": []map[string]any{
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "Memória",
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.TotalMemory,
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.ColorMapMemory,
															"wrap": true,
														},
													},
												},
											},
										},
										{
											"type": "TableRow",
											"cells": []map[string]any{
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "PODS",
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.TotalPods,
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.ColorMapPods,
															"wrap": true,
														},
													},
												},
											},
										},
									},
								},
							},
						},
						{
							"type": "Container",
							"items": []map[string]any{
								{
									"type": "TextBlock",
									"text": "4. Uso de memória nos nodes",
									"wrap": true,
								},
								{
									"type": "Table",
									"columns": []map[string]any{
										{"width": 1},
										{"width": 1},
										{"width": 1},
									},
									"rows": []map[string]any{
										{
											"type":  "TableRow",
											"style": "accent",
											"cells": []map[string]any{
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type":   "TextBlock",
															"text":   "Grupo",
															"wrap":   true,
															"weight": "Bolder",
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type":   "TextBlock",
															"text":   "Capacidade",
															"wrap":   true,
															"weight": "Bolder",
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type":   "TextBlock",
															"text":   "Status",
															"wrap":   true,
															"weight": "Bolder",
														},
													},
												},
											},
										},
										{
											"type": "TableRow",
											"cells": []map[string]any{
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "X < 65%",
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.CounterLessThan65,
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "🟩",
															"wrap": true,
														},
													},
												},
											},
										},
										{
											"type": "TableRow",
											"cells": []map[string]any{
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "X > 65%",
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.CounterGreaterThan65,
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "🟨",
															"wrap": true,
														},
													},
												},
											},
										},
										{
											"type": "TableRow",
											"cells": []map[string]any{
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "X > 80%",
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.CounterGreaterThen80,
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "🟥",
															"wrap": true,
														},
													},
												},
											},
										},
									},
								},
							},
						},
						buildAlertsSection(data),
						{
							"type": "Container",
							"items": []map[string]any{
								{
									"type": "TextBlock",
									"text": "6. Incidentes",
									"wrap": true,
								},
								{
									"type": "Table",
									"columns": []map[string]any{
										{"width": 1},
										{"width": 1},
										{"width": 1},
									},
									"rows": []map[string]any{
										{
											"type":  "TableRow",
											"style": "accent",
											"cells": []map[string]any{
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type":   "TextBlock",
															"text":   "Stack",
															"wrap":   true,
															"weight": "Bolder",
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type":   "TextBlock",
															"text":   "Namespace",
															"wrap":   true,
															"weight": "Bolder",
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type":   "TextBlock",
															"text":   "Status",
															"wrap":   true,
															"weight": "Bolder",
														},
													},
												},
											},
										},
										{
											"type": "TableRow",
											"cells": []map[string]any{
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "CertManager",
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "`cert-manager`",
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.CertManagerIncidentsColorMap,
															"wrap": true,
														},
													},
												},
											},
										},
										{
											"type": "TableRow",
											"cells": []map[string]any{
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "GITA",
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "`gita`",
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.GitaIncidentsColorMap,
															"wrap": true,
														},
													},
												},
											},
										},
										{
											"type": "TableRow",
											"cells": []map[string]any{
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "Loki",
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "`loki`",
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.LokiIncidentsColorMap,
															"wrap": true,
														},
													},
												},
											},
										},
										{
											"type": "TableRow",
											"cells": []map[string]any{
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "Prometheus",
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "`monitoring`",
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.MonitoringIncidentsColorMap,
															"wrap": true,
														},
													},
												},
											},
										},
										{
											"type": "TableRow",
											"cells": []map[string]any{
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "Rancher",
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "`cattle-system`",
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.RancherIncidentsColorMap,
															"wrap": true,
														},
													},
												},
											},
										},
										{
											"type": "TableRow",
											"cells": []map[string]any{
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "Outros",
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": "`--`",
															"wrap": true,
														},
													},
												},
												{
													"type": "TableCell",
													"items": []map[string]any{
														{
															"type": "TextBlock",
															"text": data.OtherIncidentsColorMap,
															"wrap": true,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return sendWebhookPayload(webhookURL, payload)
}

func buildAlertsSection(data KubernetesReportData) map[string]any {
	// Create a slice to store all alert rows
	alertRows := []map[string]any{}

	// Iterate through alerts and create a row for each one
	for _, alert := range data.Alerts {
		alertRow := map[string]any{
			"type": "TableRow",
			"cells": []map[string]any{
				{
					"type": "TableCell",
					"items": []map[string]any{
						{
							"type": "TextBlock",
							"text": alert.AlertDate,
							"wrap": true,
						},
					},
				},
				{
					"type": "TableCell",
					"items": []map[string]any{
						{
							"type": "TextBlock",
							"text": alert.AlertResource + "\\" + alert.AlertNamespace,
							"wrap": true,
						},
					},
				},
				{
					"type": "TableCell",
					"items": []map[string]any{
						{
							"type": "TextBlock",
							"text": alert.AlertName,
							"wrap": true,
						},
					},
				},
				{
					"type": "TableCell",
					"items": []map[string]any{
						{
							"type": "TextBlock",
							"text": alert.AlertDescription,
							"wrap": true,
						},
					},
				},
			},
		}
		// Append each alert row to our collection
		alertRows = append(alertRows, alertRow)
	}

	// Create the header row
	headerRow := map[string]any{
		"type":  "TableRow",
		"style": "accent",
		"cells": []map[string]any{
			{
				"type": "TableCell",
				"items": []map[string]any{
					{
						"type":   "TextBlock",
						"text":   "Data",
						"wrap":   true,
						"weight": "Bolder",
					},
				},
			},
			{
				"type": "TableCell",
				"items": []map[string]any{
					{
						"type":   "TextBlock",
						"text":   "Recurso\\Namespace",
						"wrap":   true,
						"weight": "Bolder",
					},
				},
			},
			{
				"type": "TableCell",
				"items": []map[string]any{
					{
						"type":   "TextBlock",
						"text":   "Nome",
						"wrap":   true,
						"weight": "Bolder",
					},
				},
			},
			{
				"type": "TableCell",
				"items": []map[string]any{
					{
						"type":   "TextBlock",
						"text":   "Descrição",
						"wrap":   true,
						"weight": "Bolder",
					},
				},
			},
		},
	}

	// Combine header row with alert rows
	allRows := []map[string]any{headerRow}
	allRows = append(allRows, alertRows...)

	return map[string]any{
		"type": "Container",
		"items": []map[string]any{
			{
				"type": "TextBlock",
				"text": "5. Alertas",
				"wrap": true,
			},
			{
				"type": "Table",
				"columns": []map[string]any{
					{"width": 1},
					{"width": 1},
					{"width": 1},
					{"width": 1},
				},
				"rows": allRows,
			},
		},
	}
}
