CREATE TABLE dbo.FACT_Ventas (
    FactVentaKey INT IDENTITY (1, 1) PRIMARY KEY,
    ProductoKey INT NOT NULL,
    ClienteKey INT NOT NULL,
    TerritorioKey INT NOT NULL,
    TiempoKey INT NOT NULL,
    OrdenVentaID INT NOT NULL,
    Cantidad INT NOT NULL,
    PrecioUnitario DECIMAL(18, 2) NOT NULL,
    DescuentoPrecioUnitario DECIMAL(18, 2) NOT NULL,
    Total DECIMAL(18, 2) NOT NULL,
    FechaCarga DATETIME NOT NULL DEFAULT GETDATE (),
    FOREIGN KEY (ProductoKey) REFERENCES DIM_Producto (ProductoKey),
    FOREIGN KEY (ClienteKey) REFERENCES DIM_Cliente (ClienteKey),
    FOREIGN KEY (TerritorioKey) REFERENCES DIM_Territorio (TerritorioKey),
    FOREIGN KEY (TiempoKey) REFERENCES DIM_Tiempo (TiempoKey)
);