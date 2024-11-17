ALTER TABLE counter.items
    DROP CONSTRAINT items_styles_id_fk;

-- Revert the table creation for counter.styles
DROP TABLE IF EXISTS counter.styles;

ALTER TABLE counter.items
    RENAME CONSTRAINT items_settings_id_fk_2 TO items_settings_id_fk_4;

-- Revert the deletion of the constraint from counter.items
ALTER TABLE counter.items
    ADD CONSTRAINT items_settings_id_fk_2
        FOREIGN KEY (style_id) REFERENCES counter.settings;

-- Revert the renaming of the constraint in counter.items
