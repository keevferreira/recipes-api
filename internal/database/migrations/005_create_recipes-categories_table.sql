CREATE TABLE RecipeCategories (
    ID SERIAL PRIMARY KEY,
    RecipeID INT NOT NULL,
    CategoryID INT NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (RecipeID) REFERENCES Recipes(ID),
    FOREIGN KEY (CategoryID) REFERENCES Categories(ID)
);