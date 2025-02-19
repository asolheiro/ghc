# CHESF

## 2. coruscant(Produção)
### 2.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |     11 |
| Incidentes |     40 |
| Problemas  |    805 |
| Segurança  |    591 |

| VERSÃO  | SUPORTE | FIM DO SUPORTE |
|---------|---------|----------------|
| v1.24.6 | 🟥      | ---            |

> Legenda:
>
> 🟩 - suporte longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.
### 2.2. Informações de recursos
| RECURSOS |  CAPACIDADE   | STATUS |
|----------|---------------|--------|
| CPU      | 128 cores     | 🟩     |
| Memória  | 369676080 Gib | 🟩     |
| PODS     |          1210 | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.
### 2.3. Uso de memória dos nodes
|  GRUPO  | QUANTIDADE | STATUS |
|---------|------------|--------|
| X < 65% |          4 | 🟩     |
| X > 65% |          5 | 🟨     |
| X > 80% |          2 | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.
### 2.4. Alertas
|   DATA   | RECURSO NAMESPACE |      NOME       | DESCRIÇÃO  |
|----------|-------------------|-----------------|------------|
| 28/11/24 | pvcescrita        | pvc-smb-escrita | pvc is low |

### 2.5. Incidentes
|   STACK    |    NAMESPACE    | STATUS |
|------------|-----------------|--------|
| CerManager | `cert-manager`  | 🟩     |
| Loki       | `loki`          | 🟩     |
| Prometheus | `monitoring`    | 🟥     |
| Rancher    | `cattle-system` | 🟥     |
| Outros     | --              | 🟥     |

> Legenda:
>
> 🟩 - sem incidentes; 🟥 - possui incidentes.
---
## 3. dagobah(desenvolvimento)
### 3.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |      8 |
| Incidentes |     66 |
| Problemas  |    885 |
| Segurança  |    543 |

|  VERSÃO  | SUPORTE | FIM DO SUPORTE |
|----------|---------|----------------|
| v1.24.10 | 🟥      | ---            |

> Legenda:
>
> 🟩 - suporte longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.
### 3.2. Informações de recursos
| RECURSOS |  CAPACIDADE   | STATUS |
|----------|---------------|--------|
| CPU      | 76 cores      | 🟩     |
| Memória  | 205413636 Gib | 🟩     |
| PODS     |           880 | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.
### 3.3. Uso de memória dos nodes
|  GRUPO  | QUANTIDADE | STATUS |
|---------|------------|--------|
| X < 65% |          1 | 🟩     |
| X > 65% |          5 | 🟨     |
| X > 80% |          2 | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.
### 3.4. Alertas
|   DATA   | RECURSO NAMESPACE |      NOME       | DESCRIÇÃO  |
|----------|-------------------|-----------------|------------|
| 28/11/24 | pvcescrita        | pvc-smb-escrita | pvc is low |

### 3.5. Incidentes
|   STACK    |    NAMESPACE    | STATUS |
|------------|-----------------|--------|
| CerManager | `cert-manager`  | 🟩     |
| Loki       | `loki`          | 🟩     |
| Prometheus | `monitoring`    | 🟥     |
| Rancher    | `cattle-system` | 🟥     |
| Outros     | --              | 🟥     |

> Legenda:
>
> 🟩 - sem incidentes; 🟥 - possui incidentes.
---
## 4. deathstar(windows)
### 4.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |      6 |
| Incidentes |      7 |
| Problemas  |    257 |
| Segurança  |    195 |

|     VERSÃO      | SUPORTE | FIM DO SUPORTE |
|-----------------|---------|----------------|
| v1.23.17+rke2r1 | 🟥      | ---            |

> Legenda:
>
> 🟩 - suporte longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.
### 4.2. Informações de recursos
| RECURSOS |  CAPACIDADE  | STATUS |
|----------|--------------|--------|
| CPU      | 24 cores     | 🟩     |
| Memória  | 49546660 Gib | 🟩     |
| PODS     |          660 | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.
### 4.3. Uso de memória dos nodes
|  GRUPO  | QUANTIDADE | STATUS |
|---------|------------|--------|
| X < 65% |          4 | 🟩     |
| X > 65% |          2 | 🟨     |
| X > 80% |          0 | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.
### 4.4. Alertas
| DATA | RECURSO NAMESPACE | NOME | DESCRIÇÃO |
|------|-------------------|------|-----------|

### 4.5. Incidentes
|   STACK    |    NAMESPACE    | STATUS |
|------------|-----------------|--------|
| CerManager | `cert-manager`  | 🟩     |
| Loki       | `loki`          | 🟩     |
| Prometheus | `monitoring`    | 🟩     |
| Rancher    | `cattle-system` | 🟥     |
| Outros     | --              | 🟩     |

> Legenda:
>
> 🟩 - sem incidentes; 🟥 - possui incidentes.
---
## 5. elet-gke-chesf-dev
### 5.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |      7 |
| Incidentes |     11 |
| Problemas  |    543 |
| Segurança  |    478 |

|       VERSÃO        | SUPORTE | FIM DO SUPORTE |
|---------------------|---------|----------------|
| v1.30.8-gke.1162001 | 🟥      | ---            |

> Legenda:
>
> 🟩 - suporte longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.
### 5.2. Informações de recursos
| RECURSOS |  CAPACIDADE  | STATUS |
|----------|--------------|--------|
| CPU      | 14 cores     | 🟩     |
| Memória  | 57721768 Gib | 🟩     |
| PODS     |          770 | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.
### 5.3. Uso de memória dos nodes
|  GRUPO  | QUANTIDADE | STATUS |
|---------|------------|--------|
| X < 65% |          3 | 🟩     |
| X > 65% |          3 | 🟨     |
| X > 80% |          1 | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.
### 5.4. Alertas
| DATA | RECURSO NAMESPACE | NOME | DESCRIÇÃO |
|------|-------------------|------|-----------|

### 5.5. Incidentes
|   STACK    |    NAMESPACE    | STATUS |
|------------|-----------------|--------|
| CerManager | `cert-manager`  | 🟩     |
| Loki       | `loki`          | 🟩     |
| Prometheus | `monitoring`    | 🟥     |
| Rancher    | `cattle-system` | 🟩     |
| Outros     | --              | 🟥     |

> Legenda:
>
> 🟩 - sem incidentes; 🟥 - possui incidentes.
---
## 6. elet-gke-chesf-prd
### 6.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |      9 |
| Incidentes |      9 |
| Problemas  |   1288 |
| Segurança  |    877 |

|       VERSÃO        | SUPORTE | FIM DO SUPORTE |
|---------------------|---------|----------------|
| v1.30.8-gke.1162001 | 🟥      | ---            |

> Legenda:
>
> 🟩 - suporte longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.
### 6.2. Informações de recursos
| RECURSOS |  CAPACIDADE  | STATUS |
|----------|--------------|--------|
| CPU      | 18 cores     | 🟩     |
| Memória  | 82250572 Gib | 🟩     |
| PODS     |          990 | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.
### 6.3. Uso de memória dos nodes
|  GRUPO  | QUANTIDADE | STATUS |
|---------|------------|--------|
| X < 65% |          7 | 🟩     |
| X > 65% |          0 | 🟨     |
| X > 80% |          2 | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.
### 6.4. Alertas
| DATA | RECURSO NAMESPACE | NOME | DESCRIÇÃO |
|------|-------------------|------|-----------|

### 6.5. Incidentes
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
## 7. elet-gke-chesf-qld
### 7.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |      8 |
| Incidentes |      7 |
| Problemas  |    489 |
| Segurança  |    549 |

|       VERSÃO        | SUPORTE | FIM DO SUPORTE |
|---------------------|---------|----------------|
| v1.30.8-gke.1162001 | 🟥      | ---            |

> Legenda:
>
> 🟩 - suporte longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.
### 7.2. Informações de recursos
| RECURSOS |  CAPACIDADE  | STATUS |
|----------|--------------|--------|
| CPU      | 18 cores     | 🟩     |
| Memória  | 73622104 Gib | 🟩     |
| PODS     |          880 | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.
### 7.3. Uso de memória dos nodes
|  GRUPO  | QUANTIDADE | STATUS |
|---------|------------|--------|
| X < 65% |          5 | 🟩     |
| X > 65% |          3 | 🟨     |
| X > 80% |          0 | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.
### 7.4. Alertas
|   DATA   | RECURSO NAMESPACE | NOME  | DESCRIÇÃO  |
|----------|-------------------|-------|------------|
| 28/01/25 | pvceletrobras     | alien | pvc is low |

### 7.5. Incidentes
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

## 9. tatooine(dgcti)
### 9.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |      8 |
| Incidentes |     22 |
| Problemas  |    397 |
| Segurança  |    267 |

| VERSÃO  | SUPORTE | FIM DO SUPORTE |
|---------|---------|----------------|
| v1.24.9 | 🟥      | ---            |

> Legenda:
>
> 🟩 - suporte longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.
### 9.2. Informações de recursos
| RECURSOS |  CAPACIDADE   | STATUS |
|----------|---------------|--------|
| CPU      | 76 cores      | 🟩     |
| Memória  | 184861776 Gib | 🟩     |
| PODS     |           880 | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.
### 9.3. Uso de memória dos nodes
|  GRUPO  | QUANTIDADE | STATUS |
|---------|------------|--------|
| X < 65% |          3 | 🟩     |
| X > 65% |          0 | 🟨     |
| X > 80% |          5 | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.
### 9.4. Alertas
| DATA | RECURSO NAMESPACE | NOME | DESCRIÇÃO |
|------|-------------------|------|-----------|

### 9.5. Incidentes
|   STACK    |    NAMESPACE    | STATUS |
|------------|-----------------|--------|
| CerManager | `cert-manager`  | 🟩     |
| Loki       | `loki`          | 🟩     |
| Prometheus | `monitoring`    | 🟥     |
| Rancher    | `cattle-system` | 🟥     |
| Outros     | --              | 🟥     |

> Legenda:
>
> 🟩 - sem incidentes; 🟥 - possui incidentes.
---