# CHESF - 17.02.2025

- [CHESF - 17.02.2025](#chesf---17022025)
  - [1. alderann - homologação](#1-alderann---homologação)
    - [1.1. Informações gerais](#11-informações-gerais)
    - [1.2. Informações de recursos](#12-informações-de-recursos)
    - [1.3. Uso de memória dos nodes](#13-uso-de-memória-dos-nodes)
    - [1.4. Alertas](#14-alertas)
    - [1.5. Incidentes](#15-incidentes)
  - [2. coruscant - produção](#2-coruscant---produção)
    - [2.1 Informações gerais](#21-informações-gerais)
    - [2.2. Informações de recursos](#22-informações-de-recursos)
    - [2.3. Uso de memória dos nodes](#23-uso-de-memória-dos-nodes)
    - [2.4. Alertas](#24-alertas)
    - [2.5. Incidentes](#25-incidentes)
  - [3. dagobah - desenvolvimento](#3-dagobah---desenvolvimento)
    - [3.1 Informações gerais](#31-informações-gerais)
    - [3.2. Informações de recursos](#32-informações-de-recursos)
    - [3.3. Uso de memória dos nodes](#33-uso-de-memória-dos-nodes)
    - [3.4. Alertas](#34-alertas)
    - [3.5. Incidentes](#35-incidentes)
  - [4. deathstar - windows](#4-deathstar---windows)
    - [4.1 Informações gerais](#41-informações-gerais)
    - [4.2. Informações de recursos](#42-informações-de-recursos)
    - [4.3. Uso de memória dos nodes](#43-uso-de-memória-dos-nodes)
    - [4.4. Alertas](#44-alertas)
    - [4.5. Incidentes](#45-incidentes)
  - [5. elet-gke-chesf-dev - desenvolvimento](#5-elet-gke-chesf-dev---desenvolvimento)
    - [5.1 Informações gerais](#51-informações-gerais)
    - [5.2. Informações de recursos](#52-informações-de-recursos)
    - [5.3. Uso de memória dos nodes](#53-uso-de-memória-dos-nodes)
    - [5.4. Alertas](#54-alertas)
    - [5.5. Incidentes](#55-incidentes)
  - [6. elet-gke-chesf-prd - produção](#6-elet-gke-chesf-prd---produção)
    - [6.1 Informações gerais](#61-informações-gerais)
    - [6.2. Informações de recursos](#62-informações-de-recursos)
    - [6.3. Uso de memória dos nodes](#63-uso-de-memória-dos-nodes)
    - [6.4. Alertas](#64-alertas)
    - [6.5. Incidentes](#65-incidentes)
  - [7. elet-gke-chesf-qld](#7-elet-gke-chesf-qld)
    - [7.1. Informações gerais](#71-informações-gerais)
    - [7.2. Informações de recursos](#72-informações-de-recursos)
    - [7.3. Uso de memória dos nodes](#73-uso-de-memória-dos-nodes)
    - [7.4. Alertas](#74-alertas)
    - [7.5. Incidentes](#75-incidentes)
  - [8. local-on-premise](#8-local-on-premise)
    - [8.1 Informações gerais](#81-informações-gerais)
    - [8.2. Informações de recursos](#82-informações-de-recursos)
    - [8.3. Uso de memória dos nodes](#83-uso-de-memória-dos-nodes)
    - [8.4. Alertas](#84-alertas)
    - [8.5. Incidentes](#85-incidentes)
  - [9. tatooine - dgcti](#9-tatooine---dgcti)
    - [9.1 Informações gerais](#91-informações-gerais)
    - [9.2. Informações de recursos](#92-informações-de-recursos)
    - [9.3. Uso de memória dos nodes](#93-uso-de-memória-dos-nodes)
    - [9.4. Alertas](#94-alertas)
    - [9.5. Incidentes](#95-incidentes)

---

## 1. alderann - homologação

### 1.1. Informações gerais

| Descrição | Número  |
|:----------|:-------:|
| Nodes     | 7       |
| Incidentes| 7       |
| Problemas | 350     |
| Segurança | 462     |

| Versão    | Suporte | Fim do suporte |
|:----------|:-------:|:--------------:|
| 1.24.10   | 🟥      | ---            |

> Legenda:
>
> 🟩 - suporte longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.

### 1.2. Informações de recursos

| Recursos | Capacidade    | Status |
|:---------|:--------------|:------:|
| CPU      | 60 cores      | 🟩     |
| Memória  | 164 Gb        | 🟩     |
| PODS     | 770           | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado ; 🟨 - requer atenção; 🟥 - nível crítico.

### 1.3. Uso de memória dos nodes

| Grupo   | Quantidade | Status |
|:--------|:----------:|:------:|
| X < 65% | 1          | 🟩     |
| X > 65% | 4          | 🟨     |
| X > 80% | 2          | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.

### 1.4. Alertas

| data | Recurso.Namespace | Nome | Descrição | Aberto? |
|--|:------------------|:-----|:----------|:-------:|
| 28/11/24 | pvc.escrita | pvc-smb-escrita | pvc as less than twenty percent available memory | Sim |

### 1.5. Incidentes

| Stack      | Namespace       | Status |
|------------|-----------------|:------:|
| CerManager | `cert-manager`  | 🟩     |
| Loki       | `loki`          | 🟩     |
| Prometheus | `monitoring`    | 🟩     |
| Rancher    | `cattle-system` | 🟩     |
| Outros     | --              | 🟩     |

> Legenda:
>
> 🟩 - sem incidentes; 🟥 - possui incidentes.

---

## 2. coruscant - produção

### 2.1 Informações gerais

| Descrição | Número |
|:----------|:------:|
| Nodes     | 11      |
| Problemas | 591    |
| Segurança | 805    |

| Versão    | Suporte | Fim do suporte |
|:----------|:-------:|:--------------:|
| 1.24.6    | 🟥      | ---            |

> Legenda:
>
> 🟩 - suporto longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.

### 2.2. Informações de recursos

| Recursos | Capacidade    | Status |
|:---------|:--------------|:------:|
| CPU      | 128 cores      | 🟩     |
| Memória  | 352 Gb         | 🟩     |
| PODS     | 1210           | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado ; 🟨 - requer atenção; 🟥 - nível crítico.

### 2.3. Uso de memória dos nodes

| Grupo   | Quantidade | Status |
|:--------|:----------:|:------:|
| X < 65% | 4          | 🟩     |
| X > 65% | 6          | 🟨     |
| X > 80% | 1          | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.

### 2.4. Alertas

| data | Recurso.Namespace | Nome | Descrição |
|--|:------------------|:-----|:----------|
| 28/11/24 | pvc.escrita | pvc-smb-escrita | pvc as less than twenty percent available memory |

### 2.5. Incidentes

| Stack      | Namespace                     | Status |
|------------|-------------------------------|:------:|
| CerManager | `cert-manager`                | 🟩     |
| Cluster    | `kube-system`                 | 🟥     |
| Loki       | `loki`                        | 🟩     |
| Prometheus | `cattle-monitoring-system`    | 🟥     |
| Rancher    | `cattle-system`               | 🟥     |
| Gita       | `gita`                        | 🟥     |
| Outros     | --                            | 🟥     |

> Legenda:
>
> 🟩 - sem incidentes; 🟥 - possui incidentes.

Observação: A maioria dos incidentes apresentados está relacionada ao nível de recurso no Node: `The node was low on resource: ephemeral-storage. Container cluster-register was using 494165961, which exceeds its request of 0.`

---

## 3. dagobah - desenvolvimento

### 3.1 Informações gerais

| Descrição | Número |
|:----------|:------:|
| Nodes     | 8      |
| Problemas | 885    |
| Segurança | 54     |

| Versão    | Suporte | Fim do suporte |
|:----------|:-------:|:--------------:|
| 1.24.10   | 🟥      | --- |

> Legenda:
>
> 🟩 - suporto longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.

### 3.2. Informações de recursos

| Recursos | Capacidade    | Status |
|:---------|:--------------|:------:|
| CPU      | 76 cores       | 🟩     |
| Memória  | 195 Gb         | 🟩     |
| PODS     | 236            | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado ; 🟨 - requer atenção; 🟥 - nível crítico.

### 3.3. Uso de memória dos nodes

| Grupo   | Quantidade | Status |
|:--------|:----------:|:------:|
| X < 65% | 2          | 🟩     |
| X > 65% | 4          | 🟨     |
| X > 80% | 2          | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.

### 3.4. Alertas

| data | Recurso.Namespace | Nome | Descrição |
|--|:------------------|:-----|:----------|
| 28/11/24 | pvc.escrita | pvc-smb-escrita | pvc as less than twenty percent available memory |

### 3.5. Incidentes

| Stack      | Namespace                     | Status |
|------------|-------------------------------|:------:|
| CerManager | `cert-manager`                | 🟩     |
| Cluster    | `kube-system`                 | 🟥     |
| Loki       | `loki`                        | 🟩     |
| Prometheus | `cattle-monitoring-system`    | 🟥     |
| Rancher    | `cattle-system`               | 🟥     |
| Outros     | --                            | 🟥     |

Observação: A maioria dos incidentes apresentados está relacionada ao nível de recurso no Node: `The node was low on resource: ephemeral-storage. Container cluster-register was using 494165961, which exceeds its request of 0.`

> Legenda:
>
> 🟩 - sem incidentes; 🟥 - possui incidentes.

---

## 4. deathstar - windows

### 4.1 Informações gerais

| Descrição | Número |
|:----------|:------:|
| Nodes     | 6      |
| Problemas | 257    |
| Segurança | 205    |

| Versão          | Suporte | Fim do suporte |
|:----------------|:-------:|:--------------:|
| 1.23.17+rke2r1  | 🟥      | --- |

> Legenda:
>
> 🟩 - suporto longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.

### 4.2. Informações de recursos

| Recursos | Capacidade    | Status |
|:---------|:--------------|:------:|
| CPU      | 24 cores       | 🟩     |
| Memória  | 47  Gb         | 🟨     |
| PODS     | 68             | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado ; 🟨 - requer atenção; 🟥 - nível crítico.

### 4.3. Uso de memória dos nodes

| Grupo   | Quantidade | Status |
|:--------|:----------:|:------:|
| X < 65% | 2          | 🟩     |
| X > 65% | 4          | 🟨     |
| X > 80% | 0          | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.

### 4.4. Alertas

| data | Recurso.Namespace | Nome | Descrição |
|--|:------------------|:-----|:----------|
|-|-|-|-|

### 4.5. Incidentes

| Stack      | Namespace                     | Status |
|------------|-------------------------------|:------:|
| CerManager | `cert-manager`                | 🟩     |
| Cluster    | `kube-system`                 | 🟩     |
| Loki       | `loki`                        | 🟩     |
| Prometheus | `cattle-monitoring-system`    | 🟩     |
| Rancher    | `cattle-system`               | 🟥     |
| Outros     | --                            | 🟥     |

Observação: A maioria dos incidentes apresentados está relacionada a um erro ao realizar o download da imagem de Node Exporter do Gita.

> Legenda:
>
> 🟩 - sem incidentes; 🟥 - possui incidentes.

---

## 5. elet-gke-chesf-dev - desenvolvimento

### 5.1 Informações gerais

| Descrição | Número |
|:----------|:------:|
| Nodes     | 8      |
| Problemas | 943    |
| Segurança | 707    |

| Versão             | Suporte | Fim do suporte |
|:-------------------|:-------:|:--------------:|
| 1.30.8-gke.1162001 | 🟩      | 30/09/2025     |

> Legenda:
>
> 🟩 - suporto longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.

### 5.2. Informações de recursos

| Recursos | Capacidade    | Status |
|:---------|:--------------|:------:|
| CPU      | 16 cores       | 🟩     |
| Memória  | 62  Gb         | 🟩     |
| PODS     | 880            | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado ; 🟨 - requer atenção; 🟥 - nível crítico.

### 5.3. Uso de memória dos nodes

| Grupo   | Quantidade | Status |
|:--------|:----------:|:------:|
| X < 65% | 3          | 🟩     |
| X > 65% | 3          | 🟨     |
| X > 80% | 2          | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.

### 5.4. Alertas

| data | Recurso.Namespace | Nome | Descrição |
|--|:------------------|:-----|:----------|
|-|-|-|-|

### 5.5. Incidentes

| Stack      | Namespace                     | Status |
|------------|-------------------------------|:------:|
| CerManager | `cert-manager`                | 🟩     |
| Cluster    | `kube-system`                 | 🟩     |
| Loki       | `loki`                        | 🟩     |
| Prometheus | `cattle-monitoring-system`    | 🟩     |
| Rancher    | `cattle-system`               | 🟩     |
| Outros     | --                            | 🟥     |

> Legenda:
>
> 🟩 - sem incidentes; 🟥 - possui incidentes.

---

## 6. elet-gke-chesf-prd - produção

### 6.1 Informações gerais

| Descrição | Número |
|:----------|:------:|
| Nodes     | 9      |
| Problemas | 1548   |
| Segurança | 974    |

| Versão            | Suporte | Fim do suporte |
|:------------------|:-------:|:--------------:|
| 1.30.8-gke1162001 | 🟩      | 30/09/2025     |

> Legenda:
>
> 🟩 - suporto longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.

### 6.2. Informações de recursos

| Recursos | Capacidade    | Status |
|:---------|:--------------|:------:|
| CPU      | 18 cores       | 🟩     |
| Memória  | 78 Gb          | 🟩     |
| PODS     | 990            | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado ; 🟨 - requer atenção; 🟥 - nível crítico.

### 6.3. Uso de memória dos nodes

| Grupo   | Quantidade | Status |
|:--------|:----------:|:------:|
| X < 65% | 7          | 🟩     |
| X > 65% | 1          | 🟨     |
| X > 80% | 1          | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.

### 6.4. Alertas

| data | Recurso.Namespace | Nome | Descrição |
|--|:------------------|:-----|:----------|
|-|-|-|-|

### 6.5. Incidentes

| Stack      | Namespace                     | Status |
|------------|-------------------------------|:------:|
| CerManager | `cert-manager`                | 🟩     |
| Cluster    | `kube-system`                 | 🟩     |
| Loki       | `loki`                        | 🟩     |
| Prometheus | `cattle-monitoring-system`    | 🟩     |
| Rancher    | `cattle-system`               | 🟩     |
| Outros     | --                            | 🟥     |

Observação: A maioria dos incidentes detectados foi relacionada ao nível de memória disponível no node.

> Legenda:
>
> 🟩 - sem incidentes; 🟥 - possui incidentes.

---

## 7. elet-gke-chesf-qld

### 7.1. Informações gerais

| Descrição | Número |
|:----------|:------:|
| Nodes     | 8      |
| Problemas | 523    |
| Segurança | 430    |

| Versão            | Suporte | Fim do suporte |
|:------------------|:-------:|:---------------|
| 1.30.8-gke1162001 | 🟩      | 30/09/2025     |

> Legenda:
>
> 🟩 - suporto longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.

### 7.2. Informações de recursos

| Recursos | Capacidade    | Status |
|:---------|:--------------|:------:|
| CPU      | 14 cores      | 🟩     |
| Memória  | 55 Gb         | 🟩     |
| PODS     | 770           | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado ; 🟨 - requer atenção; 🟥 - nível crítico.

### 7.3. Uso de memória dos nodes

| Grupo   | Quantidade | Status |
|:--------|:----------:|:------:|
| X < 65% | 3          | 🟩     |
| X > 65% | 2          | 🟨     |
| X > 80% | 2          | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.

### 7.4. Alertas

| data | Recurso.Namespace | Nome | Descrição | Aberto? |
|--|:------------------|:-----|:----------|:-------:|
| 27/01/24 | pvc.eletrobras | alien | pvc as less than twenty percent available memory | Sim |

### 7.5. Incidentes

| Stack      | Namespace       | Status |
|------------|-----------------|:------:|
| CerManager | `cert-manager`  | 🟩     |
| Loki       | `loki`          | 🟩     |
| Prometheus | `monitoring`    | 🟩     |
| Rancher    | `cattle-system` | 🟩     |
| Outros     | --              | 🟩     |

> Legenda:
>
> 🟩 - sem incidentes; 🟥 - possui incidentes.

---

## 8. local-on-premise

### 8.1 Informações gerais

| Descrição | Número |
|:----------|:------:|
| Nodes     | 3      |
| Problemas | 137    |
| Segurança | 72     |

| Versão | Suporte | Fim do suporte |
|:-------|:-------:|:--------------:|
| 1.24.6 | 🟥      | ---            |

> Legenda:
>
> 🟩 - suporto longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.

### 8.2. Informações de recursos

| Recursos | Capacidade    | Status |
|:---------|:---------------|:------:|
| CPU      | 18 cores       | 🟩     |
| Memória  | 46  Gb         | 🟩     |
| PODS     | 330            | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado ; 🟨 - requer atenção; 🟥 - nível crítico.

### 8.3. Uso de memória dos nodes

| Grupo   | Quantidade | Status |
|:--------|:----------:|:------:|
| X < 65% | 0          | 🟩     |
| X > 65% | 3          | 🟨     |
| X > 80% | 0          | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.

### 8.4. Alertas

| data | Recurso.Namespace | Nome | Descrição |
|--|:------------------|:-----|:----------|
|-|-|-|-|

### 8.5. Incidentes

| Stack      | Namespace                     | Status |
|------------|-------------------------------|:------:|
| CerManager | `cert-manager`                | 🟩     |
| Cluster    | `kube-system`                 | 🟩     |
| Loki       | `loki`                        | 🟩     |
| Prometheus | `cattle-monitoring-system`    | 🟩     |
| Rancher    | `cattle-system`               | 🟩     |
| Outros     | --                            | 🟩     |

> Legenda:
>
> 🟩 - sem incidentes; 🟥 - possui incidentes.

---

## 9. tatooine - dgcti

### 9.1 Informações gerais

| Descrição | Número |
|:----------|:------:|
| Nodes     | 8      |
| Problemas | 389    |
| Segurança | 266    |

| Versão    | Suporte | Fim do suporte |
|:----------|:-------:|:--------------:|
| 1.24.9    | 🟥      | ---            |

> Legenda:
>
> 🟩 - suporto longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.

### 9.2. Informações de recursos

| Recursos | Capacidade    | Status |
|:---------|:--------------|:------:|
| CPU      | 76 cores       | 🟩     |
| Memória  | 176 Gb         | 🟩     |
| PODS     | 880            | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado ; 🟨 - requer atenção; 🟥 - nível crítico.

### 9.3. Uso de memória dos nodes

| Grupo   | Quantidade | Status |
|:--------|:----------:|:------:|
| X < 65% | 3          | 🟩     |
| X > 65% | 0          | 🟨     |
| X > 80% | 5          | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.

### 9.4. Alertas

| data | Recurso.Namespace | Nome | Descrição |
|--|:------------------|:-----|:----------|
|-|-|-|-|

### 9.5. Incidentes

| Stack      | Namespace                     | Status |
|------------|-------------------------------|:------:|
| CerManager | `cert-manager`                | 🟩     |
| Cluster    | `kube-system`                 | 🟩     |
| Loki       | `loki`                        | 🟩     |
| Prometheus | `cattle-monitoring-system`    | 🟩     |
| Rancher    | `cattle-system`               | 🟩     |
| Outros     | --                            | 🟥     |

> Legenda:
>
> 🟩 - sem incidentes; 🟥 - possui incidentes.

Observação: A maioria dos incidentes está relacionada à quantidade de recursos disponível no node.