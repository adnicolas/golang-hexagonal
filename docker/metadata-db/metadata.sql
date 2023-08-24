BEGIN;

SET CLIENT_ENCODING TO 'UTF8';

CREATE TABLE public."role" (
	id varchar NOT NULL,
	CONSTRAINT "PK_b36bcfe02fc8de3c57a8b2391c2" PRIMARY KEY (id)
);

CREATE TABLE public."user" (
	id serial4 NOT NULL,
	uuid uuid NOT NULL,
	"name" varchar NOT NULL,
	surname varchar NOT NULL,
	email varchar NOT NULL,
	"password" varchar NOT NULL,
	"roleId" varchar NULL,
	CONSTRAINT "PK_cace4a159ff9f2512dd42373760" PRIMARY KEY (id),
	CONSTRAINT "UQ_a95e949168be7b7ece1a2382fed" UNIQUE (uuid)
);

ALTER TABLE public."user" ADD CONSTRAINT "FK_c28e52f758e7bbc53828db92194" FOREIGN KEY ("roleId") REFERENCES public."role"(id);

INSERT INTO public."role" ("id") VALUES('ADMIN'), ('EDITOR');
INSERT INTO public."user" ("uuid", "name", surname, email, "password", "roleId") VALUES
(
	'6aaf0d1e-d382-4dc7-80d1-5a9d44a1a902',
	'Robert',
	'Patiño',
	'mock@gmail.com',
	'$2b$10$peYkA/EX6CzvTRn/wOMab.RDr4j7KbmWn6tHBorL.o6lGrZUvo.a6',
	'ADMIN'
);

COMMIT;