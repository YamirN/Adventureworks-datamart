CREATE TABLE dbo.DIM_Producto (
    ProductoKey INT IDENTITY (1, 1) PRIMARY KEY,
    ProductoID INT NOT NULL,
    nombre_producto NVARCHAR (255) NOT NULL,
    costo_estandar DECIMAL(18, 2) NOT NULL,
    categoria NVARCHAR (255),
    subcategoria NVARCHAR (255),
    fecha_carga DATETIME NOT NULL DEFAULT GETDATE ()
);