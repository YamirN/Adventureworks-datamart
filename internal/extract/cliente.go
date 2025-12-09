package extract

import (
	"database/sql"

	"ETL_adventure/internal/models/raw"
)

// ExtractProductos extrae productos desde la BD origen
func ExtractClientes(db *sql.DB) ([]raw.ClienteRAW, error) {
	q :=
		`
	
	WITH PrioridadDireccion AS (
    SELECT
        C.CustomerID,
        COALESCE(P.FirstName + ' ' + P.LastName, S.Name) AS NameCustomer,
        P.PersonType,
        A.City,
        SP.Name AS StateProvinceName,
        CR.Name AS CountryRegionName,
        C.PersonID,
        C.StoreID,

        ROW_NUMBER() OVER (
            PARTITION BY C.CustomerID 
            ORDER BY 
                CASE AT.Name 
                    WHEN 'Home' THEN 1      -- Prioridad 1: Domicilio
                    WHEN 'Shipping' THEN 2  -- Prioridad 2: Envío
                    ELSE 3                  -- Prioridad 3: Otros
                END
        ) AS rn
    FROM Sales.Customer AS C
    LEFT JOIN Sales.Store AS S ON C.StoreID = S.BusinessEntityID
    LEFT JOIN Person.Person AS P ON C.PersonID = P.BusinessEntityID
    LEFT JOIN Person.BusinessEntityAddress AS BEA ON C.PersonID = BEA.BusinessEntityID 
    LEFT JOIN Person.AddressType AS AT ON BEA.AddressTypeID = AT.AddressTypeID
    LEFT JOIN Person.Address AS A ON BEA.AddressID = A.AddressID
    LEFT JOIN Person.StateProvince AS SP ON A.StateProvinceID = SP.StateProvinceID
    LEFT JOIN Person.CountryRegion AS CR ON SP.CountryRegionCode = CR.CountryRegionCode
    
    WHERE C.PersonID IS NOT NULL
      
	)
	SELECT 
		T.CustomerID, 
		-- Agregamos la lógica para definir el TipoCliente final basado en las llaves
		
		T.PersonID,
		T.StoreID,
		T.NameCustomer,
		CASE
			WHEN T.StoreID IS NULL THEN 'Persona Pura' -- 18,484
			WHEN T.StoreID IS NOT NULL THEN 'Contacto de Tienda' -- 635
		END AS TipoCliente,
		T.PersonType, 
		T.City, 
		T.StateProvinceName,
		T.CountryRegionName
	FROM PrioridadDireccion AS T
	WHERE rn = 1;
	`

	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []raw.ClienteRAW
	for rows.Next() {
		var r raw.ClienteRAW
		if err := rows.Scan(&r.CustomerID, &r.PersonID, &r.StoreID, &r.NameCustomer, &r.TipoCliente, &r.PersonType, &r.City, &r.StateProvinceName, &r.CountryRegionName); err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, rows.Err()
}
