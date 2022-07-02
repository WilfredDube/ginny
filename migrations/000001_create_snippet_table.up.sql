-- Create a `snippets` table.
CREATE TABLE public.snippets (
  id serial NOT NULL,
  guid uuid NOT NULL,
  title character varying(100) NOT NULL,
  content text NOT NULL,
  created timestamp with time zone NOT NULL,
  expires timestamp with time zone NOT NULL
);
ALTER TABLE
  public.snippets
ADD
  CONSTRAINT snippets_pkey PRIMARY KEY (id)