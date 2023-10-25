-- +goose Up
CREATE TYPE git_repo_access_t AS ENUM ('public', 'private', 'archived');

CREATE TABLE git_repo (
    id UUID PRIMARY KEY,
    organization TEXT NOT NULL,
    name TEXT NOT NULL,
    access git_repo_access_t NOT NULL,
    code JSONB NOT NULL,
    UNIQUE (organization, name)
);


CREATE TABLE ci_cd_pipeline (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    git_repo_id UUID NOT NULL,
    FOREIGN KEY (git_repo_id) REFERENCES git_repo (id) 
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    UNIQUE (name, git_repo_id),
    command TEXT NOT NULL,
    
    parent_id UUID, 
    FOREIGN KEY (parent_id) REFERENCES ci_cd_pipeline (id) 
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

-- CREATE TABLE ci_cd_pipeline_chain (
--    child_id UUID PRIMARY KEY, 
--    FOREIGN KEY (child_id) REFERENCES ci_cd_pipeline (id) 
--         ON DELETE CASCADE
--         ON UPDATE CASCADE,

--    parent_id UUID NOT NULL, 
--    FOREIGN KEY (parent_id) REFERENCES ci_cd_pipeline (id) 
--         ON DELETE CASCADE
--         ON UPDATE CASCADE
-- );


CREATE TABLE app (
    name TEXT PRIMARY KEY,
    git_repo_id UUID NOT NULL,
    FOREIGN KEY (git_repo_id) REFERENCES git_repo (id) 
        ON DELETE RESTRICT
        ON UPDATE CASCADE,
    UNIQUE(git_repo_id)
);

CREATE TYPE oltp_store_type_t AS ENUM ('queue', 'db');
CREATE TABLE oltp_store (
    id UUID PRIMARY KEY,
    provider TEXT NOT NULL,
    auth JSONB NOT NULL,
    type oltp_store_type_t NOT NULL,

    service_name TEXT NOT NULL,
    FOREIGN KEY (service_name) REFERENCES service (name) 
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    UNIQUE(service_name)
);

CREATE TABLE data_warehouse (
    id UUID PRIMARY KEY,
    provider TEXT NOT NULL,
    auth JSONB NOT NULL
);


CREATE TABLE service (
    name TEXT PRIMARY KEY,
    FOREIGN KEY (name) REFERENCES app (name) 
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

-- CREATE TABLE service_oltp_store_operation (
--     oltp_store_id UUID PRIMARY KEY,
--     FOREIGN KEY (oltp_store_id) REFERENCES oltp_store (id) 
--         ON DELETE CASCADE
--         ON UPDATE CASCADE,

--     service_name TEXT NOT NULL,
--     FOREIGN KEY (service_name) REFERENCES service (name) 
--         ON DELETE CASCADE
--         ON UPDATE CASCADE
-- );

CREATE TABLE etl_job (
    name TEXT PRIMARY KEY,
    FOREIGN KEY (name) REFERENCES app (name) 
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    oltp_store_id UUID NOT NULL,
    FOREIGN KEY (oltp_store_id) REFERENCES oltp_store (id) 
        ON DELETE RESTRICT
        ON UPDATE CASCADE,
    UNIQUE(oltp_store_id),

    data_warehouse_id UUID NOT NULL,
    FOREIGN KEY (data_warehouse_id) REFERENCES data_warehouse (id) 
        ON DELETE RESTRICT
        ON UPDATE CASCADE
);


CREATE TABLE image (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    version TEXT NOT NULL,
    UNIQUE(name, version),

    built_by_ci_id UUID NOT NULL,
    FOREIGN KEY (built_by_ci_id) REFERENCES ci_cd_pipeline (id) 
        ON DELETE RESTRICT
        ON UPDATE CASCADE,
    UNIQUE (built_by_ci_id)
);
CREATE TABLE app_container (
    name TEXT PRIMARY KEY,
    FOREIGN KEY (name) REFERENCES app (name) 
        ON DELETE CASCADE
        ON UPDATE CASCADE,

    image_id UUID NOT NULL,
    FOREIGN KEY (image_id) REFERENCES image (id) 
        ON DELETE RESTRICT
        ON UPDATE CASCADE
);
CREATE TABLE container_env (
    name TEXT PRIMARY KEY
);
CREATE TABLE trace (
    id TEXT PRIMARY KEY,
    approx_time TIMESTAMP NOT NULL
);

CREATE TABLE trace_spans (
    trace_id TEXT NOT NULL,
    FOREIGN KEY (trace_id) REFERENCES trace (id) 
        ON DELETE CASCADE
        ON UPDATE CASCADE,

    container_env_name TEXT NOT NULL,
    FOREIGN KEY (container_env_name) REFERENCES container_env (name) 
        ON DELETE RESTRICT
        ON UPDATE CASCADE,

    app_container_name TEXT NOT NULL,
    FOREIGN KEY (app_container_name) REFERENCES app_container (name) 
        ON DELETE RESTRICT
        ON UPDATE CASCADE,

    payload JSONB NOT NULL
);



-- +goose Down
DROP TABLE trace_spans;
DROP TABLE trace;
DROP TABLE container_env;
DROP TABLE app_container;
DROP TABLE image;

DROP TABLE service_oltp_store_operation;
DROP TABLE service;
DROP TABLE etl_job;
DROP TABLE app;

DROP TABLE oltp_store;
DROP TYPE oltp_store_type_t;
DROP TABLE data_warehouse;

DROP TABLE ci_cd_pipeline_chain;
DROP TABLE ci_cd_pipeline;
DROP TABLE git_repo;
DROP TYPE git_repo_access_t;
