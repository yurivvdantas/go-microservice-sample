DROP TABLE IF EXISTS CRYPTOS;
CREATE TABLE CRYPTOS (
  id           INT AUTO_INCREMENT NOT NULL,
  name         VARCHAR(255),
  code         VARCHAR(3),
  upvotes      INT NOT NULL,
  downvotes    INT NOT NULL,
  description     VARCHAR(255),
  PRIMARY KEY (`id`)
);

INSERT INTO CRYPTOS
  (name,code,upvotes, downvotes, description)
VALUES
  ('Klever','KLV',0, 0, 'The most amazing crypto');