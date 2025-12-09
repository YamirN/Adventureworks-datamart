CREATE TABLE dbo.DIM_Territorio (
    TerritorioKey INT IDENTITY (1, 1) PRIMARY KEY,
    TerritorioID INT NOT NULL,
    NombreTerritorio NVARCHAR (255) NOT NULL,
    CodigoPais NVARCHAR (10) NOT NULL,
    Pais NVARCHAR (255) NOT NULL,
    Continente NVARCHAR (255) NOT NULL,
    FechaCarga DATETIME NOT NULL DEFAULT GETDATE ()
);