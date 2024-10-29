
INSERT INTO counter.setting_types (alias, name) VALUES ('buyer', 'Buyer');
INSERT INTO counter.setting_types (alias, name) VALUES ('style', 'Style');
INSERT INTO counter.setting_types (alias, name) VALUES ('color', 'Color');
INSERT INTO counter.setting_types (alias, name) VALUES ('size', 'Size');
INSERT INTO counter.setting_types (alias, name) VALUES ('factory', 'Factory');
INSERT INTO counter.setting_types (alias, name) VALUES ('machine', 'Machine');
;
INSERT INTO counter.settings (id, setting_type_alias, value, parent_id) VALUES (DEFAULT, 'factory', 'Factory 1', 0);
INSERT INTO counter.settings (id, setting_type_alias, value, parent_id) VALUES (DEFAULT, 'machine', 'Machine 1', 1);
INSERT INTO counter.settings (id, setting_type_alias, value, parent_id) VALUES (DEFAULT, 'buyer', 'Buyer Test', 0);
INSERT INTO counter.settings (id, setting_type_alias, value, parent_id) VALUES (DEFAULT, 'style', 'Style Test', 3);
INSERT INTO counter.settings (id, setting_type_alias, value, parent_id) VALUES (DEFAULT, 'color', 'black', 4);
INSERT INTO counter.settings (id, setting_type_alias, value, parent_id) VALUES (DEFAULT, 'size', 'XL', 4);

INSERT INTO counter.items VALUES (123456789, 3, 4, 5, 6);
