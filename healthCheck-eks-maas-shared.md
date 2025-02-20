# eks-maas-shared

## 1. maas-prd
### 1.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |      4 |
| Incidentes |      2 |
| Problemas  |    256 |
| Segurança  |    229 |

|       VERSÃO        | SUPORTE | FIM DO SUPORTE |
|---------------------|---------|----------------|
| v1.30.8-eks-2d5f260 | 🟥      | ---            |

> Legenda:
>
> 🟩 - suporte longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.
### 1.2. Informações de recursos
| RECURSOS |  CAPACIDADE  | STATUS |
|----------|--------------|--------|
| CPU      | 8 cores      | 🟩     |
| Memória  | 31747608 Gib | 🟩     |
| PODS     |          116 | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.
### 1.3. Uso de memória dos nodes
|  GRUPO  | QUANTIDADE | STATUS |
|---------|------------|--------|
| X < 65% |          4 | 🟩     |
| X > 65% |          0 | 🟨     |
| X > 80% |          0 | 🟥     |

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