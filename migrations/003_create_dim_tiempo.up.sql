CREATE TABLE dbo.DIM_Tiempo (
    TiempoKey INT PRIMARY KEY, -- YYYYMMDD
    Fecha DATE NOT NULL,
    Anio INT NOT NULL,
    Mes INT NOT NULL,
    NombreMes NVARCHAR (20) NOT NULL,
    Dia INT NOT NULL,
    DiaSemana INT NOT NULL, -- 1 = Lunes
    NombreDia NVARCHAR (20) NOT NULL,
    Trimestre INT NOT NULL,
    SemanaISO INT NOT NULL,
    EsFinDeSemana BIT NOT NULL,
    FechaCarga DATETIME NOT NULL DEFAULT GETDATE ()
);