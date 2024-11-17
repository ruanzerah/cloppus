CREATE TABLE messages (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  owner VARCHAR(80) NOT NULL,
  subject VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  likes int8 DEFAULT 0 NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL,
  deleted_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE users (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  username VARCHAR(80) NOT NULL,
  email VARCHAR(90) NOT NULL,
  auth BOOLEAN NOT NULL,
  hash VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL,
  deleted_at TIMESTAMPTZ NOT NULL
);