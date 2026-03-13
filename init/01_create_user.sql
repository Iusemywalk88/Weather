-- init/01_create_user.sql
CREATE ROLE weather_user WITH LOGIN PASSWORD 'qwerty';
CREATE DATABASE weather_service_db OWNER weather_user;
GRANT ALL PRIVILEGES ON DATABASE weather_service_db TO weather_user;
GRANT USAGE ON SCHEMA public TO weather_user;
GRANT CREATE ON SCHEMA public TO weather_user;
