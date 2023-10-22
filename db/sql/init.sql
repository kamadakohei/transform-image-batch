DROP TABLE IF EXISTS transform_image_settings;
CREATE TABLE transform_image_settings (
    id         INT AUTO_INCREMENT NOT NULL,
    image_name     VARCHAR(128) NOT NULL,
    output_type    VARCHAR(255) NOT NULL,
    resize_height  DECIMAL(5,2) NOT NULL,
    resize_width   DECIMAL(5,2) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO transform_image_settings
    (image_name, output_type, resize_height, resize_width)
VALUES
    ('example_file1.png', 'png', 50, 50),
    ('example_file2.jpeg', 'jpeg', 10, 50),
    ('example_file3.png', 'png',  50, 50),
    ('example_file4.jpeg', 'jpeg', 30, 50),
    ('example_file5.png', 'png', 50, 50);