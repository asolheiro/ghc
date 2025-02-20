# SEBRAE AL - 20 de Fevereiro de 2025

## 1. Homolog
### 1.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |      6 |
| Incidentes |      6 |
| Problemas  |    622 |
| Segurança  |    382 |



| VERSÃO  | FIM DO SUPORTE | ÚLTIMA RELEASE DISPONÍVEL | SUPORTE |
|---------|----------------|---------------------------|---------|
| v1.26.4 | 28/02/2024     | 1.26.15                   | 🟥      |

> Legenda:
>
> 🟩 - LTS; 🟨 - menos de 90 dias restantes; 🟧 - menos de 30 dias restantes 🟥 - desatualizado.


### 1.2. Informações de recursos
| RECURSOS | CAPACIDADE | STATUS |
|----------|------------|--------|
| CPU      | 24 cores   | 🟩     |
| Memória  | 62.28 Gib  | 🟨     |
| PODS     |        660 | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.


### 1.3. Uso de memória dos nodes
|  GRUPO  | QUANTIDADE | STATUS |
|---------|------------|--------|
| X < 65% |          1 | 🟩     |
| X > 65% |          5 | 🟨     |
| X > 80% |          0 | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.


### 1.4. Alertas
| DATA | `RECURSO\NAMESPACE` | NOME | DESCRIÇÃO |
|------|---------------------|------|-----------|
| -    | -                   | -    | -         |



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

## 2. Infra
### 2.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |      3 |
| Incidentes |      0 |
| Problemas  |    369 |
| Segurança  |    301 |



|     VERSÃO     | FIM DO SUPORTE |                 ÚLTIMA RELEASE DISPONÍVEL                 |                                                 SUPORTE                                                  |
|----------------|----------------|-----------------------------------------------------------|----------------------------------------------------------------------------------------------------------|
| v1.31.5+rke2r1 | ---            | [RKE2-Releases](https://github.com/rancher/rke2/releases) | [Support-Matrix](https://www.suse.com/pt-br/suse-rke1/support-matrix/all-supported-versions/rke1-v1-31/) |

> Legenda:
>
> 🟩 - LTS; 🟨 - menos de 90 dias restantes; 🟧 - menos de 30 dias restantes 🟥 - desatualizado.


### 2.2. Informações de recursos
| RECURSOS | CAPACIDADE | STATUS |
|----------|------------|--------|
| CPU      | 12 cores   | 🟩     |
| Memória  | 46.86 Gib  | 🟩     |
| PODS     |        330 | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.


### 2.3. Uso de memória dos nodes
|  GRUPO  | QUANTIDADE | STATUS |
|---------|------------|--------|
| X < 65% |          3 | 🟩     |
| X > 65% |          0 | 🟨     |
| X > 80% |          0 | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.


### 2.4. Alertas
| DATA | `RECURSO\NAMESPACE` | NOME | DESCRIÇÃO |
|------|---------------------|------|-----------|
| -    | -                   | -    | -         |



### 2.5. Incidentes
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

## 3. Produção
### 3.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |     11 |
| Incidentes |     19 |
| Problemas  |    993 |
| Segurança  |    685 |



| VERSÃO  | FIM DO SUPORTE | ÚLTIMA RELEASE DISPONÍVEL | SUPORTE |
|---------|----------------|---------------------------|---------|
| v1.26.4 | 28/02/2024     | 1.26.15                   | 🟥      |

> Legenda:
>
> 🟩 - LTS; 🟨 - menos de 90 dias restantes; 🟧 - menos de 30 dias restantes 🟥 - desatualizado.


### 3.2. Informações de recursos
| RECURSOS | CAPACIDADE | STATUS |
|----------|------------|--------|
| CPU      | 44 cores   | 🟩     |
| Memória  | 148.25 Gib | 🟩     |
| PODS     |       1210 | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.


### 3.3. Uso de memória dos nodes
|  GRUPO  | QUANTIDADE | STATUS |
|---------|------------|--------|
| X < 65% |         11 | 🟩     |
| X > 65% |          0 | 🟨     |
| X > 80% |          0 | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.


### 3.4. Alertas
|   DATA   | `RECURSO\NAMESPACE` |         NOME         | DESCRIÇÃO  |
|----------|---------------------|----------------------|------------|
| 15/02/25 | pvc\log-stack       | storage-loki-stack-0 | pvc is low |



### 3.5. Incidentes
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

