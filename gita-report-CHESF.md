# CHESF
## 1. alderann - homologação
### 1.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |      7 |
| Incidentes |      7 |
| Problemas  |    350 |
| Segurança  |    462 |

| VERSÃO  | SUPORTE | FIM DO SUPORTE |
|---------|---------|----------------|
| 1.24.10 | 🟥      | ---            |

> Legenda:
>
> 🟩 - suporte longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.
### 1.2. Informações de recursos
| RECURSOS | CAPACIDADE | STATUS |
|----------|------------|--------|
| CPU      | 60 cores   | 🟩     |
| Memória  | 164 Gb     | 🟩     |
| PODS     |        770 | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.
### 1.3. Uso de memória dos nodes
|  GRUPO  | QUANTIDADE | STATUS |
|---------|------------|--------|
| X < 65% |          1 | 🟩     |
| X > 65% |          4 | 🟨     |
| X > 80% |          2 | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.
### 1.4. Alertas
|   DATA   | RECURSO NAMESPACE |      NOME       |           DESCRIÇÃO            | ABERTO? |
|----------|-------------------|-----------------|--------------------------------|---------|
| 28/11/24 | pvc.escrita       | pvc-smb-escrita | pvc as less than twenty        | Sim     |
|          |                   |                 | percent available memory       |         |

### 1.5. Incidentes
|   STACK    |    NAMESPACE    | STATUS |
|------------|-----------------|--------|
| CerManager | `cert-manager`  | 🟩     |
| Loki       | `loki`          | 🟩     |
| Prometheus | `monitoring`    | 🟩     |
| Rancher    | `cattle-system` | 🟩     |
| Outros     | --              | 🟩     |

> Legenda:
>
> 🟩 - sem incidentes; 🟥 - possui incidentes.
---