BEGIN;
CREATE INDEX IF NOT EXISTS idx_persons_age ON persons(age);
CREATE INDEX IF NOT EXISTS idx_persons_gender ON persons(gender);
CREATE INDEX IF NOT EXISTS idx_persons_nationality ON persons(nationality);
CREATE INDEX idx_persons_filter_fields ON persons(age, gender, nationality);
COMMIT;

