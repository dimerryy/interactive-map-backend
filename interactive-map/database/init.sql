-- Create user
CREATE ROLE dimerryy WITH LOGIN PASSWORD 'Soo577dis';

-- Allow the user to create DBs (optional)
ALTER ROLE dimerryy CREATEDB;

-- Create database owned by the new user
CREATE DATABASE interactive_map OWNER dimerryy;
