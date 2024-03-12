-- Create procedure to update the update_at column
CREATE FUNCTION updated_at_procedure()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';
