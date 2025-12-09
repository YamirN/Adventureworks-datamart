# AdventureWorks Sales Data Mart

Este repositorio contiene un **Data Mart de Ventas** construido a partir
de la base de datos **AdventureWorks 2022**, incluyendo:

- Pipeline **ETL completo en Go (Golang)**
- Modelado en **esquema en estrella**
- Tablas **Dimensión** y **Hecho (FACT_Ventas)**
- Dashboard en **Power BI**
- Scripts SQL usados en el proceso

---

## 📌 Arquitectura del Proyecto

### 1. **Extracción (E)**

Se obtiene la información desde AdventureWorks2022 usando consultas SQL
optimizadas para ventas:

- `Sales.SalesOrderDetail`
- `Sales.SalesOrderHeader`
- `Sales.Store`
- `Sales.SalesTerritory`
- `Sales.Customer`
- `Production.Product`
- `Production.ProductSubcategory`
- `Production.ProductCategory`
- `Person.Person`
- `Person.CountryRegion`
- `Person.BusinessEntityAddress`
- `Person.AddressType`
- `Person.Address`
- `Person.StateProvince`
- `Person.CountryRegion`

### 2. **Transformación (T)**

En esta etapa se construyen:

#### 🟦 Dimensiones:

- `DIM_Tiempo`
- `DIM_Producto`
- `DIM_Cliente`
- `DIM_Territorio`

#### 🟥 Tabla de Hechos:

- `FACT_Ventas`

Con las métricas: - Cantidad - Precio unitario - Descuento - Total -
Claves foráneas de las dimensiones

### 3. **Carga (L)**

Los datos ya transformados se insertan en el Data Mart:

    AdventureWorks_DM
    │── DIM_Producto
    │── DIM_Cliente
    │── DIM_Tiempo
    │── DIM_Territorio
    └── FACT_Ventas

---

## 🛠️ Tecnologías usadas

Componente Tecnología

---

Lenguaje Go (Golang)
Base de datos SQL Server (AdventureWorks2022)
Reportes Power BI
Modelado Esquema en Estrella
Versionamiento Git / GitHub
ETL Scripts propios en Go

---

## 🚀 Ejecución del ETL

1.  Clonar el repositorio\
2.  Configurar el archivo `config.json` con credenciales SQL Server\
3.  Ejecutar:

```bash
go run main.go
```

El proceso:

- Genera las dimensiones\
- Construye el FACT\
- Carga todo al Data Mart

---

## 🧠 Medidas DAX

### Ingresos

```
Ingresos =
 SUMX(
     FACT_Ventas,
     FACT_Ventas[Cantidad] * FACT_Ventas[PrecioUnitario] * (1 - FACT_Ventas[DescuentoPrecioUnitario])
 )
```

### Ganancia

```
Ganancia =
  SUMX(
      FACT_Ventas,
      [Ingresos] - (FACT_Ventas[Cantidad] * RELATED(DIM_Producto[costo_estandar]))
  )
```

### Recuento de pedidos

```
RecuentoOrderID = COUNT(FACT_Ventas[OrdenVentaID])
```

### % Ventas por país

```
%_Ventas_Pais =
DIVIDE(
    [Ventas_totales],
    CALCULATE([Ventas_totales], ALL(DIM_Territorio[NombreTerritorio])),
    0
)
```

### Métricas mensuales

```
IngresosMesAnterior =
VAR MesActual = SELECTEDVALUE(DIM_Tiempo[FechaMes])
VAR MesAnterior = EDATE(MesActual, -1)
RETURN
CALCULATE(
    [Ingresos],
    FILTER(ALL(DIM_Tiempo), DIM_Tiempo[FechaMes] = MesAnterior)
)
```

```
GananciaMesAnterior =
VAR MesActual = SELECTEDVALUE(DIM_Tiempo[FechaMes])
VAR MesAnterior = EDATE(MesActual, -1)
RETURN
CALCULATE(
    [Ganancia],
    FILTER(ALL(DIM_Tiempo), DIM_Tiempo[FechaMes] = MesAnterior)
)
```

```
OrdenesaMesAnterior =
VAR MesActual = SELECTEDVALUE(DIM_Tiempo[FechaMes])
VAR MesAnterior = EDATE(MesActual, -1)
RETURN
CALCULATE(
    [RecuentoOrderID],
    FILTER(ALL(DIM_Tiempo), DIM_Tiempo[FechaMes] = MesAnterior)
)
```

## 📊 Dashboard Power BI

Incluye visualizaciones como:

- KPI de ingresos, ganancia y pedidos\
- Evolución mensual de pedidos, ventas y ganancias\
- Ventas totales de pedidos por mes\
- Porcentaje de ventas por pais\
- Ventas totales por categoria\
- Top 10 productos mas vendidos\
- Top 10 clientes que generan mas ingresos\
- KPI de pedidos, ingresos y ganancia \
- Segmentadores dinámicos por **año**

---

## 🧑‍💻 Autor

Proyecto desarrollado por _Yamir Alex_ como implementación profesional de
un Data Mart con ETL en Go y visualización en Power BI.

---
<img width="1395" height="782" alt="adventureworks-powerbi" src="https://github.com/user-attachments/assets/b197866b-1fa1-4ea1-85d8-dddaf6d02a29" />


