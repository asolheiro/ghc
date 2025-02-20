# Ambiente CIEB - 20 de Fevereiro de 2025

## 1. CIEB Produção
### 1.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |     15 |
| Incidentes |     21 |
| Problemas  |    924 |
| Segurança  |    838 |



|        VERSÃO        | FIM DO SUPORTE | ÚLTIMA RELEASE DISPONÍVEL | SUPORTE |
|----------------------|----------------|---------------------------|---------|
| v1.29.12-eks-2d5f260 | 23/03/2025     | 1-29-eks-31               | 🟧      |

> Legenda:
>
> 🟩 - LTS; 🟨 - menos de 90 dias restantes; 🟧 - menos de 30 dias restantes 🟥 - desatualizado.


### 1.2. Informações de recursos
| RECURSOS | CAPACIDADE | STATUS |
|----------|------------|--------|
| CPU      | 30 cores   | 🟩     |
| Memória  | 113.30 Gib | 🟨     |
| PODS     |        435 | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.


### 1.3. Uso de memória dos nodes
|  GRUPO  | QUANTIDADE | STATUS |
|---------|------------|--------|
| X < 65% |          6 | 🟩     |
| X > 65% |          5 | 🟨     |
| X > 80% |          4 | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.


### 1.4. Alertas
|   DATA   | `RECURSO\NAMESPACE` |                     NOME                      |   DESCRIÇÃO   |
|----------|---------------------|-----------------------------------------------|---------------|
| 20/02/25 | Node\               | ip-192-168-123-201.us-east-2.compute.internal | memory is low |



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

