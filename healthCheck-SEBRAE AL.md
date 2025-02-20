# SEBRAE AL

## 1. Homolog
### 1.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |      6 |
| Incidentes |      6 |
| Problemas  |    622 |
| Segurança  |    382 |

| VERSÃO  | SUPORTE | FIM DO SUPORTE |
|---------|---------|----------------|
| v1.26.4 | 🟥      | ---            |

> Legenda:
>
> 🟩 - suporte longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.
### 1.2. Informações de recursos
| RECURSOS |  CAPACIDADE  | STATUS |
|----------|--------------|--------|
| CPU      | 24 cores     | 🟩     |
| Memória  | 65304364 Gib | 🟨     |
| PODS     |          660 | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.
### 1.3. Uso de memória dos nodes
|  GRUPO  | QUANTIDADE | STATUS |
|---------|------------|--------|
| X < 65% |          3 | 🟩     |
| X > 65% |          3 | 🟨     |
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
---## 2. Infra
### 2.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |      3 |
| Incidentes |      0 |
| Problemas  |    369 |
| Segurança  |    301 |

|     VERSÃO     | SUPORTE | FIM DO SUPORTE |
|----------------|---------|----------------|
| v1.31.5+rke2r1 | 🟥      | ---            |

> Legenda:
>
> 🟩 - suporte longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.
### 2.2. Informações de recursos
| RECURSOS |  CAPACIDADE  | STATUS |
|----------|--------------|--------|
| CPU      | 12 cores     | 🟩     |
| Memória  | 49131652 Gib | 🟩     |
| PODS     |          330 | 🟩     |

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
| DATA | RECURSO NAMESPACE | NOME | DESCRIÇÃO |
|------|-------------------|------|-----------|

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
---## 3. Produção
### 3.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |     11 |
| Incidentes |     19 |
| Problemas  |    993 |
| Segurança  |    685 |

| VERSÃO  | SUPORTE | FIM DO SUPORTE |
|---------|---------|----------------|
| v1.26.4 | 🟥      | ---            |

> Legenda:
>
> 🟩 - suporte longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.
### 3.2. Informações de recursos
| RECURSOS |  CAPACIDADE   | STATUS |
|----------|---------------|--------|
| CPU      | 44 cores      | 🟩     |
| Memória  | 155455876 Gib | 🟩     |
| PODS     |          1210 | 🟩     |

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
|   DATA   | RECURSO NAMESPACE |         NOME         | DESCRIÇÃO  |
|----------|-------------------|----------------------|------------|
| 15/02/25 | pvclog-stack      | storage-loki-stack-0 | pvc is low |

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