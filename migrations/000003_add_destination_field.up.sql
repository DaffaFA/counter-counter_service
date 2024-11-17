CREATE TABLE counter.styles (
    id          SERIAL  NOT NULL
        CONSTRAINT styles_pk PRIMARY KEY,
    buyer_id    INT     NOT NULL
        CONSTRAINT styles_settings_id_fk REFERENCES counter.settings,
    name        VARCHAR NOT NULL,
    destination VARCHAR NOT NULL DEFAULT 'GLOBAL',
    amount      INT8    NOT NULL DEFAULT 0,
    CONSTRAINT styles_pk_2 UNIQUE (buyer_id, name, destination)
);

ALTER TABLE counter.items
    DROP CONSTRAINT items_settings_id_fk_2;

INSERT INTO counter.styles (id, buyer_id, name, destination, amount)
SELECT 
    settings.id, 
    parent_id, 
    split_part(value, '-', 1), 
    COALESCE(NULLIF(split_part(value, '-', 2), ''), 'GLOBAL'), 
    COUNT(item_scans.time)
FROM counter.settings
LEFT JOIN counter.items ON settings.id = items.style_id
LEFT JOIN counter.item_scans ON items.code = item_scans.qr_code_code
WHERE setting_type_alias = 'style'
GROUP BY 1, 2, 3;

ALTER TABLE counter.items
    ADD CONSTRAINT items_styles_id_fk
        FOREIGN KEY (style_id) REFERENCES counter.styles ON DELETE NO ACTION ON UPDATE NO ACTION;

ALTER TABLE counter.items
    RENAME CONSTRAINT items_settings_id_fk_4 TO items_settings_id_fk_2;

SELECT setval('counter.styles_id_seq', (SELECT MAX(id) FROM counter.styles));
