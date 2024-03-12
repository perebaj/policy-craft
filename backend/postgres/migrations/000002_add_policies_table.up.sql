CREATE TABLE policies (
  id UUID PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  criteria VARCHAR(255) NOT NULL,
  value INTEGER NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

-- Create a trigger to call the updated_at_procedure function when the policies table is updated
CREATE TRIGGER policies_updated_at_trigger
    BEFORE UPDATE
    ON
        policies
    FOR EACH ROW
EXECUTE PROCEDURE updated_at_procedure();

COMMIT;
