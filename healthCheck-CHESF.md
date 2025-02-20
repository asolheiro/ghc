# CHESF - 20 de Fevereiro de 2025

## 1. alderaan(homologação)
### 1.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |      7 |
| Incidentes |      0 |
| Problemas  |    462 |
| Segurança  |    350 |



|  VERSÃO  | FIM DO SUPORTE | ÚLTIMA RELEASE DISPONÍVEL | SUPORTE |
|----------|----------------|---------------------------|---------|
| v1.24.10 | 28/07/2023     | 1.24.17                   | 🟥      |

> Legenda:
>
> 🟩 - LTS; 🟨 - menos de 90 dias restantes; 🟧 - menos de 30 dias restantes 🟥 - desatualizado.


### 1.2. Informações de recursos
| RECURSOS | CAPACIDADE | STATUS |
|----------|------------|--------|
| CPU      | 60 cores   | 🟩     |
| Memória  | 164.51 Gib | 🟨     |
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
|   DATA   | `RECURSO\NAMESPACE` |      NOME       | DESCRIÇÃO  |
|----------|---------------------|-----------------|------------|
| 28/11/24 | pvc\escrita         | pvc-smb-escrita | pvc is low |



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

## 2. coruscant(Produção)
### 2.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |     11 |
| Incidentes |     40 |
| Problemas  |    805 |
| Segurança  |    591 |



| VERSÃO  | FIM DO SUPORTE | ÚLTIMA RELEASE DISPONÍVEL | SUPORTE |
|---------|----------------|---------------------------|---------|
| v1.24.6 | 28/07/2023     | 1.24.17                   | 🟥      |

> Legenda:
>
> 🟩 - LTS; 🟨 - menos de 90 dias restantes; 🟧 - menos de 30 dias restantes 🟥 - desatualizado.


### 2.2. Informações de recursos
| RECURSOS | CAPACIDADE | STATUS |
|----------|------------|--------|
| CPU      | 128 cores  | 🟩     |
| Memória  | 352.55 Gib | 🟩     |
| PODS     |       1210 | 🟩     |

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
|   DATA   | `RECURSO\NAMESPACE` |      NOME       | DESCRIÇÃO  |
|----------|---------------------|-----------------|------------|
| 28/11/24 | pvc\escrita         | pvc-smb-escrita | pvc is low |



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
| Incidentes |    996 |
| Problemas  |   3596 |
| Segurança  |   1435 |



|  VERSÃO  | FIM DO SUPORTE | ÚLTIMA RELEASE DISPONÍVEL | SUPORTE |
|----------|----------------|---------------------------|---------|
| v1.24.10 | 28/07/2023     | 1.24.17                   | 🟥      |

> Legenda:
>
> 🟩 - LTS; 🟨 - menos de 90 dias restantes; 🟧 - menos de 30 dias restantes 🟥 - desatualizado.


### 3.2. Informações de recursos
| RECURSOS | CAPACIDADE | STATUS |
|----------|------------|--------|
| CPU      | 76 cores   | 🟩     |
| Memória  | 195.90 Gib | 🟩     |
| PODS     |        880 | 🟥     |

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
|   DATA   | `RECURSO\NAMESPACE` |       NOME        |     DESCRIÇÃO     |
|----------|---------------------|-------------------|-------------------|
| 06/04/23 | Node\               | reclnxdagobah-wk1 | disk space is low |
| 28/11/24 | pvc\escrita         | pvc-smb-escrita   | pvc is low        |



### 3.5. Incidentes
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

## 4. deathstar(windows)
### 4.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |      6 |
| Incidentes |      7 |
| Problemas  |    257 |
| Segurança  |    205 |



|     VERSÃO      | FIM DO SUPORTE |                 ÚLTIMA RELEASE DISPONÍVEL                 |                                                 SUPORTE                                                  |
|-----------------|----------------|-----------------------------------------------------------|----------------------------------------------------------------------------------------------------------|
| v1.23.17+rke2r1 | ---            | [RKE2-Releases](https://github.com/rancher/rke2/releases) | [Support-Matrix](https://www.suse.com/pt-br/suse-rke1/support-matrix/all-supported-versions/rke1-v1-31/) |

> Legenda:
>
> 🟩 - LTS; 🟨 - menos de 90 dias restantes; 🟧 - menos de 30 dias restantes 🟥 - desatualizado.


### 4.2. Informações de recursos
| RECURSOS | CAPACIDADE | STATUS |
|----------|------------|--------|
| CPU      | 24 cores   | 🟩     |
| Memória  | 47.25 Gib  | 🟩     |
| PODS     |        660 | 🟩     |

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
| DATA | `RECURSO\NAMESPACE` | NOME | DESCRIÇÃO |
|------|---------------------|------|-----------|
| -    | -                   | -    | -         |



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
| Incidentes |     12 |
| Problemas  |    550 |
| Segurança  |    485 |



|       VERSÃO        | FIM DO SUPORTE | ÚLTIMA RELEASE DISPONÍVEL | SUPORTE |
|---------------------|----------------|---------------------------|---------|
| v1.30.8-gke.1261000 | 30/09/2025     | 1.30.9-gke.1231000        | 🟩      |

> Legenda:
>
> 🟩 - LTS; 🟨 - menos de 90 dias restantes; 🟧 - menos de 30 dias restantes 🟥 - desatualizado.


### 5.2. Informações de recursos
| RECURSOS | CAPACIDADE | STATUS |
|----------|------------|--------|
| CPU      | 14 cores   | 🟩     |
| Memória  | 55.05 Gib  | 🟩     |
| PODS     |        770 | 🟩     |

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
| DATA | `RECURSO\NAMESPACE` | NOME | DESCRIÇÃO |
|------|---------------------|------|-----------|
| -    | -                   | -    | -         |



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
| Problemas  |   1297 |
| Segurança  |    887 |



|       VERSÃO        | FIM DO SUPORTE | ÚLTIMA RELEASE DISPONÍVEL | SUPORTE |
|---------------------|----------------|---------------------------|---------|
| v1.30.8-gke.1261000 | 30/09/2025     | 1.30.9-gke.1231000        | 🟩      |

> Legenda:
>
> 🟩 - LTS; 🟨 - menos de 90 dias restantes; 🟧 - menos de 30 dias restantes 🟥 - desatualizado.


### 6.2. Informações de recursos
| RECURSOS | CAPACIDADE | STATUS |
|----------|------------|--------|
| CPU      | 18 cores   | 🟩     |
| Memória  | 78.44 Gib  | 🟩     |
| PODS     |        990 | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.


### 6.3. Uso de memória dos nodes
|  GRUPO  | QUANTIDADE | STATUS |
|---------|------------|--------|
| X < 65% |          8 | 🟩     |
| X > 65% |          0 | 🟨     |
| X > 80% |          1 | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.


### 6.4. Alertas
| DATA | `RECURSO\NAMESPACE` | NOME | DESCRIÇÃO |
|------|---------------------|------|-----------|
| -    | -                   | -    | -         |



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



|       VERSÃO        | FIM DO SUPORTE | ÚLTIMA RELEASE DISPONÍVEL | SUPORTE |
|---------------------|----------------|---------------------------|---------|
| v1.30.8-gke.1261000 | 30/09/2025     | 1.30.9-gke.1231000        | 🟩      |

> Legenda:
>
> 🟩 - LTS; 🟨 - menos de 90 dias restantes; 🟧 - menos de 30 dias restantes 🟥 - desatualizado.


### 7.2. Informações de recursos
| RECURSOS | CAPACIDADE | STATUS |
|----------|------------|--------|
| CPU      | 18 cores   | 🟩     |
| Memória  | 70.21 Gib  | 🟩     |
| PODS     |        880 | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.


### 7.3. Uso de memória dos nodes
|  GRUPO  | QUANTIDADE | STATUS |
|---------|------------|--------|
| X < 65% |          4 | 🟩     |
| X > 65% |          2 | 🟨     |
| X > 80% |          2 | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.


### 7.4. Alertas
|   DATA   | `RECURSO\NAMESPACE` | NOME  | DESCRIÇÃO  |
|----------|---------------------|-------|------------|
| 28/01/25 | pvc\eletrobras      | alien | pvc is low |



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

## 8. local-on-premise
### 8.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |      3 |
| Incidentes |      0 |
| Problemas  |    137 |
| Segurança  |     72 |



| VERSÃO  | FIM DO SUPORTE | ÚLTIMA RELEASE DISPONÍVEL | SUPORTE |
|---------|----------------|---------------------------|---------|
| v1.24.6 | 28/07/2023     | 1.24.17                   | 🟥      |

> Legenda:
>
> 🟩 - LTS; 🟨 - menos de 90 dias restantes; 🟧 - menos de 30 dias restantes 🟥 - desatualizado.


### 8.2. Informações de recursos
| RECURSOS | CAPACIDADE | STATUS |
|----------|------------|--------|
| CPU      | 18 cores   | 🟩     |
| Memória  | 46.87 Gib  | 🟥     |
| PODS     |        330 | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado; 🟨 - requer atenção; 🟥 - nível crítico.


### 8.3. Uso de memória dos nodes
|  GRUPO  | QUANTIDADE | STATUS |
|---------|------------|--------|
| X < 65% |          0 | 🟩     |
| X > 65% |          0 | 🟨     |
| X > 80% |          3 | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.


### 8.4. Alertas
| DATA | `RECURSO\NAMESPACE` | NOME | DESCRIÇÃO |
|------|---------------------|------|-----------|
| -    | -                   | -    | -         |



### 8.5. Incidentes
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

## 9. tatooine(dgcti)
### 9.1. Informações gerais
| DESCRIÇÃO  | NÚMERO |
|------------|--------|
| Nodes      |      8 |
| Incidentes |     22 |
| Problemas  |    397 |
| Segurança  |    197 |



| VERSÃO  | FIM DO SUPORTE | ÚLTIMA RELEASE DISPONÍVEL | SUPORTE |
|---------|----------------|---------------------------|---------|
| v1.24.9 | 28/07/2023     | 1.24.17                   | 🟥      |

> Legenda:
>
> 🟩 - LTS; 🟨 - menos de 90 dias restantes; 🟧 - menos de 30 dias restantes 🟥 - desatualizado.


### 9.2. Informações de recursos
| RECURSOS | CAPACIDADE | STATUS |
|----------|------------|--------|
| CPU      | 76 cores   | 🟩     |
| Memória  | 176.30 Gib | 🟩     |
| PODS     |        880 | 🟩     |

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
| DATA | `RECURSO\NAMESPACE` | NOME | DESCRIÇÃO |
|------|---------------------|------|-----------|
| -    | -                   | -    | -         |



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

