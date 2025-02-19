# Ambiente CIEB
## 1. CIEB Produção
### 1.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |     15 |
| Incidentes |     28 |
| Problemas  |    919 |
| Segurança  |    851 |

|        VERSÃO        | SUPORTE | FIM DO SUPORTE |
|----------------------|---------|----------------|
| v1.29.12-eks-2d5f260 | 🟥      | ---            |

> Legenda:
>
> 🟩 - suporte longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.
### 1.2. Informações de recursos
| RECURSOS |  CAPACIDADE   | STATUS |
|----------|---------------|--------|
| CPU      | 30 cores      | 🟩     |
| Memória  | 118807812 Gib | 🟩     |
| PODS     |           435 | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.
### 1.3. Uso de memória dos nodes
|  GRUPO  | QUANTIDADE | STATUS |
|---------|------------|--------|
| X < 65% |          5 | 🟩     |
| X > 65% |          6 | 🟨     |
| X > 80% |          4 | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.
### 1.4. Alertas
| DATA | RECURSO NAMESPACE | NOME | DESCRIÇÃO |
|------|-------------------|------|-----------|

### 1.5. Incidentes
|   STACK    |    NAMESPACE    | STATUS |
|------------|-----------------|--------|
| CerManager | `cert-manager`  | 🟩     |
| Loki       | `loki`          | 🟩     |
| Prometheus | `monitoring`    | 🟩     |
| Rancher    | `cattle-system` | 🟩     |
| Outros     | --              | 🟥     |

> Legenda:
>
> 🟩 - sem incidentes; 🟥 - possui incidentes.
---