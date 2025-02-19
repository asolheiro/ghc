# <ORG_NAME> - <DATE>

## <CLUSTER_INDEX>. alderann - homologação

### <CLUSTER_INDEX>.1. Informações gerais

| Descrição | Número  |
|:----------|:-------:|
| Nodes     | <CLUSTER_NODES>       |
| Incidentes| <CLUSTER_INCIDENTS>       |
| Problemas | <CLUSTER_PROBLEMS>     |
| Segurança | <CLUSTER_SECURITY>     |

| Versão    | Suporte | Fim do suporte |
|:----------|:-------:|:--------------:|
| <CLUSTER_VERSION>   | 🟥      | ---            |

> Legenda:
>
> 🟩 - suporte longo; 🟨 - suporte chegando ao fim; 🟥 - fim do suporte eminente.

### <CLUSTER_INDEX>.2. Informações de recursos

| Recursos | Capacidade    | Status |
|:---------|:--------------|:------:|
| CPU      | <CLUSTER_CPU> cores      | 🟩     |
| Memória  | <CLUSTER_MEMORY> Gb        | 🟩     |
| PODS     | <CLUSTER_MAX_PODS>           | 🟩     |

> Legenda:
>
> 🟩 - nível recomendado ; 🟨 - requer atenção; 🟥 - nível crítico.

### <CLUSTER_INDEX>.3. Uso de memória dos nodes

| Grupo   | Quantidade | Status |
|:--------|:----------:|:------:|
| X < 65% | 1          | 🟩     |
| X > 65% | 4          | 🟨     |
| X > 80% | 2          | 🟥     |

> Legenda:
>
> 🟩 - uso normal; 🟨 - uso grande; 🟥 - uso excessivo.

### <CLUSTER_INDEX>.4. Alertas

| data | Recurso.Namespace | Nome | Descrição | Aberto? |
|--|:------------------|:-----|:----------|:-------:|
| 28/11/24 | pvc.escrita | pvc-smb-escrita | pvc as less than twenty percent available memory | Sim |

### <CLUSTER_INDEX>.5. Incidentes

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