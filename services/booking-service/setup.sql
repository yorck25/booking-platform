CREATE TABLE barbers
(
    id            uuid PRIMARY KEY,
    internal_name varchar(255) NOT NULL,
    display_name  varchar(255) NOT NULL,
    email         varchar(255),
    phone         varchar(50),
    mobile_phone  varchar(50),
    postal_code   varchar(20),
    street        varchar(128),
    city          varchar(128),
    active        boolean      NOT NULL DEFAULT true,
    deleted       boolean      NOT NULL DEFAULT false,
    created_at    timestamp    NOT NULL DEFAULT now(),
    updated_at    timestamp    NOT NULL DEFAULT now()
);

CREATE TABLE barber_user_login
(
    id              uuid PRIMARY KEY,
    barber_id       uuid         NOT NULL REFERENCES barbers (id) ON DELETE CASCADE,
    username        varchar(128) NOT NULL,
    firs_name       varchar(128),
    last_name       varchar(128),
    email           varchar(255),
    password_hash   varchar(255) NOT NULL,
    role            varchar(32)  NOT NULL DEFAULT 'user',
    last_login      timestamp,
    failed_attempts smallint              DEFAULT 0,
    active          boolean      NOT NULL DEFAULT true,
)


CREATE TABLE barber_opening_hours
(
    id         uuid PRIMARY KEY,
    barber_id  uuid     NOT NULL REFERENCES barbers (id) ON DELETE CASCADE,
    weekday    smallint NOT NULL CHECK (weekday BETWEEN 0 AND 6),
    start_time time     NOT NULL,
    end_time   time     NOT NULL,
    is_closed  boolean  NOT NULL DEFAULT false
);

CREATE TABLE barber_breaks
(
    id         uuid PRIMARY KEY,
    barber_id  uuid     NOT NULL REFERENCES barbers (id) ON DELETE CASCADE,
    weekday    smallint NOT NULL CHECK (weekday BETWEEN 0 AND 6),
    start_time time     NOT NULL,
    end_time   time     NOT NULL
);

CREATE TABLE bookings
(
    id                    uuid PRIMARY KEY,
    barber_id             uuid         NOT NULL REFERENCES barbers (id),

    customer_first_name   varchar(128) NOT NULL,
    customer_last_name    varchar(255) NOT NULL,
    customer_phone_number varchar(50)  NOT NULL,
    customer_email        varchar(255),

    booking_date          date         NOT NULL,
    start_time            time         NOT NULL,
    end_time              time         NOT NULL,

    status                varchar(32)  NOT NULL DEFAULT 'pending',
    cancel_reason         varchar(255),

    notes                 text,

    confirmed_at          timestamp,
    canceled_at           timestamp,

    created_at            timestamp    NOT NULL DEFAULT now(),
    updated_at            timestamp    NOT NULL DEFAULT now()
);