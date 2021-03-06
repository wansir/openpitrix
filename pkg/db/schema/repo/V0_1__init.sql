CREATE TABLE repo (
	repo_id     VARCHAR(50) PRIMARY KEY             NOT NULL,
	name        VARCHAR(255)                        NOT NULL,
	description TEXT                                NOT NULL,
	url         TEXT                                NOT NULL,
	credential  JSON                                NOT NULL,
	visibility  VARCHAR(50)                         NOT NULL,
	owner       VARCHAR(50)                         NOT NULL,
	status      VARCHAR(50)                         NOT NULL,
	create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	status_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE repo_label (
	repo_label_id  VARCHAR(50) PRIMARY KEY             NOT NULL,
	repo_id        VARCHAR(50)                         NOT NULL,
	label_key      VARCHAR(50)                         NOT NULL,
	label_value    VARCHAR(50)                         NOT NULL,
	status         VARCHAR(50)                         NOT NULL,
	create_time    TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	status_time    TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE repo_selector (
	repo_selector_id  VARCHAR(50) PRIMARY KEY             NOT NULL,
	repo_id           VARCHAR(50)                         NOT NULL,
	selector_key      VARCHAR(50)                         NOT NULL,
	selector_value    VARCHAR(50)                         NOT NULL,
	status            VARCHAR(50)                         NOT NULL,
	create_time       TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	status_time       TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);


CREATE TABLE repo_task (
	repo_task_id VARCHAR(50) PRIMARY KEY             NOT NULL,
	repo_id      VARCHAR(50)                         NOT NULL,
	owner        VARCHAR(50)                         NOT NULL,
	status       VARCHAR(50)                         NOT NULL,
	result       TEXT                                NOT NULL,
	create_time  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	status_time  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
