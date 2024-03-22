-- name: FindAllPets :many
-- Find all pets optionally filtered by category.name
-- :categorieNames - Optional filter by category.name
-- :limit - Limit the number of results
-- :from - from which id to collect the results
SELECT p.id, p.name, c.id AS category_id, c.name AS category_name
FROM pets p
JOIN categories c ON p.category_id = c.id
WHERE
    CASE
        WHEN sqlc.narg('categorieNames')::varchar[] IS NOT NULL THEN
            c.name = ANY(sqlc.narg('categorieNames')::varchar[])
        ELSE
            TRUE
    END
    AND p.id > @startFrom
ORDER BY p.id ASC
LIMIT @maxLimit;


-- name: CreatePet :one
-- Insert a new pet with an existing or new category
-- @name - Name of the pet
-- @categoryName - Name of the category
-- @categoryId - ID of an existing category (optional)
WITH inserted_category AS (
    INSERT INTO categories (name)
    VALUES (@categoryName)
    ON CONFLICT (name) DO UPDATE
    SET name = excluded.name
    RETURNING id, name
)
INSERT INTO pets (name, category_id)
SELECT @name, (SELECT id FROM inserted_category)
RETURNING pets.id AS pet_id, pets.name AS pet_name,
    pets.category_id AS category_id, (SELECT name FROM inserted_category) as category_name;


-- name: UpdatePet :exec
-- Update an existing pet by ID
-- :id - ID of the pet to update
-- :name - New name of the pet
-- :category_name - New category ID of the pet
WITH updated_category AS (
    INSERT INTO categories (name)
    VALUES (@categoryName)
    ON CONFLICT (name) DO UPDATE
    SET name = excluded.name -- Ensure consistent category name
    RETURNING id, name
)
UPDATE pets
SET
    name = COALESCE(@name, name),
    category_id = (SELECT id FROM updated_category)
WHERE pets.id = @petId
    RETURNING id, name, category_id, (SELECT name FROM updated_category) as category_name;

-- name: DeletePet :exec
-- Delete a pet by ID
-- :id - ID of the pet to delete
DELETE FROM pets WHERE id = $1;

-- name: GetPetById :one
-- Get a pet by ID
-- :id - ID of the pet
SELECT p.id, p.name, c.id AS category_id, c.name AS category_name
FROM pets p
JOIN categories c ON p.category_id = c.id
WHERE p.id = $1;