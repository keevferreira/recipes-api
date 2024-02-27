CREATE TABLE RecipeIngredients (
    ID SERIAL PRIMARY KEY,
    RecipeID INT NOT NULL,
    IngredientID INT NOT NULL,
    Quantity FLOAT,
    Unit VARCHAR(50),
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (RecipeID) REFERENCES Recipes(ID),
    FOREIGN KEY (IngredientID) REFERENCES Ingredients(ID)
);
